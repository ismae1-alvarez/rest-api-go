package user

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Respository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
	Delete(id string) error
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

func (repo *repo) GetAll() ([]User, error) {
	var u []User

	result := repo.db.Model(&u).Order("created_at desc").Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil

}
func (repo *repo) Get(id string) (*User, error) {
	user := User{ID: id}

	result := repo.db.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func (r *repo) Delete(id string) error {
	user := User{ID: id}

	result := r.db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	// El otro no me funcionaba
	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}
