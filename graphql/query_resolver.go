package main

import "context"

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Account(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {

	var accounts []*Account

	return accounts, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string, id *string) ([]*Product, error) {

	var products []*Product

	return products, nil
}
