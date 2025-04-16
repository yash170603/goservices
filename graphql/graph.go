package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/yash170603/goservices/account"
	"github.com/yash170603/goservices/catalog"
	"github.com/yash170603/goservices/order"
)

type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func newGraphqlServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {

	var _ = graphql.HandlerExtension(nil)
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	//catalog is dependent on accountClient
	catalogClient, err := catalog.NewClient(catalogUrl)
	if err != nil {
		accountClient.close()
		return nil, err
	}

	// the ordeClient is both dependetn on catalogClient and accountClient
	orderClient, err := order.NewClient(orderUrl)
	if err != nil {
		catalogClient.Close()
		accountClient.close()
		return nil, err
	}

	return &Server{
		accountClient: accountClient,
		catalogClient: catalogClient,
		orderClient:   orderClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (S *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: S,
	})
}
