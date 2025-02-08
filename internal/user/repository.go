package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Respository interface {
	Create(user *User) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Respository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(user *User) error {

	user.ID = uuid.New().String()

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	repo.log.Println("User creatd withc id: ", user.ID)
	return nil
}
