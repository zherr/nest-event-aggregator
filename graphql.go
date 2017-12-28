package main

import (
	"github.com/graphql-go/graphql"
	"log"
)

var nestCameraEventType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NestCameraEvent",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"has_sound": &graphql.Field{
				Type: graphql.Boolean,
			},
			"has_motion": &graphql.Field{
				Type: graphql.Boolean,
			},
			"has_person": &graphql.Field{
				Type: graphql.Boolean,
			},
			"start_time": &graphql.Field{
				Type: graphql.DateTime,
			},
			"end_time": &graphql.Field{
				Type: graphql.DateTime,
			},
			"urls_expire_time": &graphql.Field{
				Type: graphql.DateTime,
			},
			"web_url": &graphql.Field{
				Type: graphql.String,
			},
			"app_url": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"animated_image_url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"eventById": &graphql.Field{
				Type: nestCameraEventType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						db, err := getDbConnection()
						if err != nil {
							log.Fatal(err)
						}
						defer db.Close()
						nestCameraEvent := NestCameraEvent{}
						db.First(&nestCameraEvent, idQuery)
						return nestCameraEvent, nil
					}
					return nil, nil
				},
			},
			"allEvents": &graphql.Field{
				Type: graphql.NewList(nestCameraEventType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db, err := getDbConnection()
					if err != nil {
						log.Fatal(err)
					}
					defer db.Close()
					nestCameraEvents := []NestCameraEvent{}
					db.Find(&nestCameraEvents)
					return nestCameraEvents, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
