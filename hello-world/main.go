package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	// fmt.Println(fields, "field")
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	// fmt.Println(rootQuery, "rootQuery")
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	// fmt.Println(schemaConfig, "schemaConfig")
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := `{
		hello
	}`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		// +v will include the struct's field name
		log.Fatalf("failed to execute graphql operation, errors:%+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) //
}
