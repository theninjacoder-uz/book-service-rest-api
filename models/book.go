package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint64    `json:"id" gorm:"primary_key:auto_increment;"`
	Isbn      string    `json:"isbn" gorm:"type:varchar(255);not null;uniqueIndex;unique"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Author    string    `json:"author" gorm:"type:varchar(255)"`
	Published uint32    `json:"published"`
	Pages     uint32    `json:"pages"`
	Status    uint32    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type OenLibData struct {
	Title   string `json:"title"`
	Authors []struct {
		Name string `json:"name"`
	}
	PablishDate   string `json:"publish_date"`
	NumberOfPages uint32 `json:"number_of_pages"`
}

type Info map[string]OenLibData

type BookDto struct {
	Book   Book   `json:"book"`
	Status uint32 `json:"status"`
}

func (b *Book) Prepare(data OenLibData) {
	b.ID = 0
	b.Isbn = strings.TrimSpace(b.Isbn)
	b.Author = data.Authors[0].Name
	b.Title = data.Title
	year, _ := strconv.ParseInt(data.PablishDate[len(data.PablishDate)-4:], 10, 32)
	b.Published = uint32(year)
	b.Pages = data.NumberOfPages
	b.CreatedAt = time.Now()
	b.UpdatedAt = b.CreatedAt
	b.Status = 0
}

func (b *Book) Validate(action string) error {

	switch strings.ToLower(action) {
	case "create":
		if b.Isbn == "" {
			return errors.New("Required Isbn")
		}
	case "update":
		if b.Isbn == "" {
			return errors.New("Required Isbn")
		} else if b.Title == "" {
			return errors.New("Required Title")
		} else if b.Author == "" {
			return errors.New("Required Author")
		} else if b.Status < 0 || b.Status > 2 {
			return errors.New("Book status is not valid")
		}
	}
	return nil
}

func (b *Book) SaveABook(db *gorm.DB) (*BookDto, error) {

	data, err := b.FindABookByIsbn(db, b.Isbn)
	if err == nil {
		return &BookDto{Book: *data, Status: data.Status}, nil
	}

	openLibraryUri := fmt.Sprintf(os.Getenv("OPEN_LIBRARY_API"), b.Isbn)
	res, err := http.Get(openLibraryUri)

	if err != nil {
		return &BookDto{}, err
	}
	defer res.Body.Close()
	var info Info

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error when read body")
	}

	err = json.Unmarshal(body, &info)

	if err != nil {
		fmt.Println("Error when unmarshal")
	}

	// fmt.Printf("Fetched data %s", info)

	b.Prepare(info["ISBN:"+b.Isbn])

	db = db.Debug().Model(&Book{}).Create(&b)
	if db.Error != nil {
		return &BookDto{}, db.Error
	}

	return &BookDto{Book: *b, Status: b.Status}, nil
}

func (b *Book) FindAllBooks(db *gorm.DB) (*[]BookDto, error) {
	books := []Book{}
	db = db.Debug().Model(&Book{}).Limit(100).Find(&books)
	if db.Error != nil {
		return &[]BookDto{}, db.Error
	}
	bookDtos := []BookDto{}

	for _, b := range books {
		bookDtos = append(bookDtos, BookDto{Book: b, Status: b.Status})
	}

	return &bookDtos, nil
}

func (b *Book) UpdateABook(db *gorm.DB, id uint64) (*BookDto, error) {

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
		return &BookDto{}, db.Error
	}

	return &BookDto{Book: *b, Status: b.Status}, nil
}

func (b *Book) DeleteABook(db *gorm.DB, id uint64) (int64, error) {
	db = db.Debug().Model(&Book{}).Where("id = ?", id).Take(&Book{}).Delete(&Book{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (b *Book) FindABookByIsbn(db *gorm.DB, isbn string) (*Book, error) {
	err := db.Debug().Model(&Book{}).Where("isbn = ?", isbn).Take(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}
