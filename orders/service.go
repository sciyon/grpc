package main

import "context"

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) OrdersService {
	return &service{store}
}

func (s *service) Create(ctx context.Context) error {
	return nil
}
