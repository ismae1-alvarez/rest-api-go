package user

import (
	"log"
)

type Service interface {
	Create(firsName, lastName, email, phone string) (*User, error)
}

type service struct {
	log  *log.Logger
	repo Respository
}

func NewService(log *log.Logger, repo Respository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firsName, lastName, email, phone string) (*User, error) {

	s.log.Println("Create user services")

	user := User{
		FirstName: firsName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
