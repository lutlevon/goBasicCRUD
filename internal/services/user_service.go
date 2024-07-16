package services

import (
	"CRUD/ent"
	"CRUD/internal/models"
	"context"
)

type UserService struct {
	client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{client: client}
}

func (s *UserService) GetUsers(ctx context.Context) ([]*ent.User, error) {
	users, err := s.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserById(ctx context.Context, id int) (*ent.User, error) {
	u, err := s.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) CreateUser(ctx context.Context, input models.UserInput) (*ent.User, error) {
	u, err := s.client.User.Create().
		SetName(input.Name).
		SetEmail(input.Email).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, input models.UserInput) (*ent.User, error) {
	u, err := s.client.User.UpdateOneID(id).
		SetName(input.Name).
		SetEmail(input.Email).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return u, nil
}
