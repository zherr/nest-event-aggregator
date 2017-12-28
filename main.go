package main

// Redirect handling code courtesy of Nest GoLang example: https://developers.nest.com/documentation/cloud/how-to-read-data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	token, present := os.LookupEnv("NEST_TOKEN")
	if !present {
		log.Fatalln("NEST_TOKEN not set. Please see README for more details.")
	}
	camID, present := os.LookupEnv("NEST_CAMERA_ID")
	if !present {
		log.Fatalln("NEST_CAMERA_ID not set. Please see README for more details.")
	}

	url := fmt.Sprintf("https://developer-api.nest.com/devices/cameras/%s", camID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Unable to create http.NewRequest: %s\n", err)
	}

	req.Header.Add(
		"Authorization",
		fmt.Sprintf("Bearer %s", token),
	)

	customClient := http.Client{
		CheckRedirect: func(redirRequest *http.Request, via []*http.Request) error {
			// Go's http.DefaultClient does not forward headers when a redirect 3xx
			// response is recieved. Thus, the header (which in this case contains the
			// Authorization token) needs to be passed forward to the redirect
			// destinations.
			redirRequest.Header = req.Header

			// Go's http.DefaultClient allows 10 redirects before returning an
			// an error. We have mimicked this default behavior.
			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		},
	}
	queueHTTPRequest(customClient, *req)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt)
	<-interrupt
}

func queueHTTPRequest(client http.Client, request http.Request) {
	ticker := time.NewTicker(time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				response, err := client.Do(&request)
				if err != nil {
					log.Fatalf("HTTP GET request failed: %s\n", err)
				}

				if response.StatusCode != 200 {
					log.Fatalf("Expected a 200 status code; got a %d\n", response.StatusCode)
				}

				defer response.Body.Close()

				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Unable to read body: %s\n", err)
				}

				var nestCamResp NestCameraResponse
				err = json.Unmarshal([]byte(body), &nestCamResp)
				if err != nil {
					log.Fatalf("Unabled to unmarshal: %s\n", err)
				}

				logNestCamEvent(nestCamResp)

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
