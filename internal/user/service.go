package user

import (
	"log"
)

type Service interface {
	Create(firsName, lastName, email, phone string) (*User, error)
	Get(id string) (*User, error)
	GetAll() ([]User, error)
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
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

func (s service) GetAll() ([]User, error) {

	users, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil

}

func (s service) Get(id string) (*User, error) {
	user, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *service) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {

	return s.repo.Update(id, firstName, lastName, email, phone)
}
