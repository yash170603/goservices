package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogUrl string `envconfig:"CATALOG_SERVICE_URL"`
	OrderUrl   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println("Error processing env config:", err)
		log.Fatal(err)
	}

	// Create a new GraphQL server
	s, err := NewGraphqlServer(cfg.AccountUrl, cfg.CatalogUrl, cfg.OrderUrl)
	if err != nil {
		log.Fatalf("Failed to create GraphQL server: %v", err)
	}
	server := handler.New(s.ToExecutableSchema())

	// Set up HTTP handlers
	http.Handle("/graphql", server)
	http.Handle("/playground", playground.Handler("yeder", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
