package database

import (
	"alterra-agmc/day-2/models"

	"gorm.io/gorm"
)

type LibUser struct {
	DB *gorm.DB
}

type UserRepository interface {
	GetUser() (*[]models.User, error)
	GetUserByID(int) (*models.User, error)
	CreateUser(models.User) error
	UpdateUser(int, models.User) error
	DeleteUser(int) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &LibUser{db}
}

func (l *LibUser) GetUser() (data *[]models.User, err error) {
	if err := l.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (l *LibUser) GetUserByID(id int) (data *models.User, err error) {
	db := l.DB
	if err = db.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (l *LibUser) CreateUser(user models.User) error {
	if err := l.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (l *LibUser) UpdateUser(id int, input models.User) error {
	db := l.DB
	if err := db.Where(`id=?`, id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

func (l *LibUser) DeleteUser(id int) (data *models.User, err error) {
	db := l.DB

	if err = db.Where(`id=?`, id).Delete(&data, id).Error; err != nil {
		return nil, err
	}

	return data, nil
}
