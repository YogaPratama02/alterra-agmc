package database

import (
	"alterra-agmc/day-3/models"

	"gorm.io/gorm"
)

type LibBook struct {
	DB *gorm.DB
}

type BookRepository interface {
	GetBook() (*[]models.Book, error)
	GetBookByID(int) (*models.Book, error)
	CreateBook(models.Book) error
	UpdateBook(int, models.Book) error
	DeleteBook(int) (*models.Book, error)
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &LibBook{db}
}

func (l *LibBook) GetBook() (book *[]models.Book, err error) {
	db := l.DB
	if err := db.Find(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (l *LibBook) GetBookByID(id int) (book *models.Book, err error) {
	db := l.DB
	if err = db.Where(`id=?`, id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (l *LibBook) CreateBook(book models.Book) error {
	db := l.DB
	if err := db.Create(&book).Error; err != nil {
		return err
	}

	return nil
}

func (l *LibBook) UpdateBook(id int, input models.Book) error {
	db := l.DB
	if err := db.Where(`id=?`, id).Updates(input).Error; err != nil {
		return err
	}

	return nil
}

func (l *LibBook) DeleteBook(id int) (data *models.Book, err error) {
	db := l.DB

	if err = db.Where(`id=?`, id).Delete(&data, id).Error; err != nil {
		return nil, err
	}

	return data, nil
}
