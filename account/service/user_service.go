package service

import (
	"context"
	"github.com/mbakin/custom-dictionary/account/model/apperrors"
	"log"

	"github.com/google/uuid"
	"github.com/mbakin/custom-dictionary/account/model"
)

// userService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type userService struct {
	UserRepository model.UserRepository
}

// USConfig will hold repositories that will eventually be injected into this
//  service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *userService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}

// Signup reaches our two a UserRepository to verify the
// email address is available and signs up the user if this is the case
func (s *userService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)

	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", u.Email)
		return apperrors.NewInternal()
	}

	// now I realize why I originally used Signup(ctx, email, password)
	// then created a user. It's somewhat un-natural to mutate the user here
	u.Password = pw

	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}

	// If we get around to adding events, we'd Publish it here
	// err := s.EventsBroker.PublishUserUpdated(u, true)

	// if err != nil {
	// 	return nil, apperrors.NewInternal()
	// }

	return nil
}
