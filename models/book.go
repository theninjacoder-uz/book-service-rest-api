package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint64    `json:"id" gorm:"primary_key:auto_increment;"`
	Isbn      string    `json:"isbn" gorm:"type:varchar(255);not null"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Author    string    `json:"author" gorm:"type:varchar(255)"`
	Published uint32    `json:"published"`
	Pages     uint32    `json:"pages"`
	Status    uint8     `json: "status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type BookDto struct {
	Book   Book  `json:"book"`
	Status uint8 `json:"status"`
}

func (b *Book) Prepare() {
	b.ID = 0
	b.Title = strings.TrimSpace(b.Title)
	b.Isbn = strings.TrimSpace(b.Isbn)
	b.Author = strings.TrimSpace(b.Author)
	b.CreatedAt = time.Now()
	b.UpdatedAt = b.CreatedAt
}

func (b *Book) Validate() error {

	if b.Isbn == "" {
		return errors.New("Required Isbn")
	} else if b.Title == "" {
		return errors.New("Required Title")
	} else if b.Author == "" {
		return errors.New("Required Author")
	} else if b.Status < 0 || b.Status > 2 {
		return errors.New("Book status is not valid")
	}
	return nil
}

func (b *Book) SaveABook(db *gorm.DB) (*Book, error) {

	db = db.Debug().Model(&Book{}).Create(&b)
	if db.Error != nil {
		return &Book{}, db.Error
	}

	return b, nil
}

func (b *Book) FindAllBooks(db *gorm.DB) (*[]Book, error) {
	books := []Book{}
	db = db.Debug().Model(&Book{}).Limit(100).Find(&books)
	if db.Error != nil {
		return &[]Book{}, db.Error
	}

	return &books, nil
}

func (b *Book) UpdateABook(db *gorm.DB, id uint64) (*Book, error) {

	db = db.Debug().Model(&Book{}).Where("id = ?", id).Updates(
		Book{
			Isbn:      b.Isbn,
			Title:     b.Title,
			Author:    b.Author,
			Published: b.Published,
			Pages:     b.Pages,
			UpdatedAt: time.Now(),
		},
	)

	if db.Error != nil {
		return &Book{}, db.Error
	}

	return b, nil
}

func (b *Book) DeleteABook(db *gorm.DB, id uint64) (int64, error) {
	db = db.Debug().Model(&Book{}).Where("id = ?", id).Take(&Book{}).Delete(&Book{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
