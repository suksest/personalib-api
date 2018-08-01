package main

import (
	"time"
)

type (
	//Book type
	Book struct {
		ID          int       `json:"id" gorm:"primary_key;auto_increment"`
		ISBN        string    `json:"isbn" gorm:"not null;unique;type:varchar(13)"`
		Title       string    `json:"title" gorm:"not null;type:varchar(255)"`
		Page        int       `json:"page" gorm:"not null"`
		Synopsis    string    `json:"synopsis" gorm:"type:text"`
		PublishDate time.Time `json:"published_at"`

		PublisherID int        `json:"publisher_id" gorm:"index"`
		AuthorID    int        `json:"author_id" gorm:"index"`
		Categories  []Category `gorm:"many2many:book_categories;"`
	}

	//Publisher type
	Publisher struct {
		ID      int    `json:"id" gorm:"primary_key;auto_increment"`
		Name    string `json:"name" gorm:"not null;type:varchar(255)"`
		Address string `json:"address" gorm:"type:text"`
		Books   []Book `json:"books"`
	}

	//Author type
	Author struct {
		ID    int    `json:"id" gorm:"primary_key;auto_increment"`
		Name  string `json:"name" gorm:"not null;type:varchar(255)"`
		Books []Book `json:"books"`
	}

	//Category type
	Category struct {
		ID   int    `json:"id" gorm:"primary_key;auto_increment"`
		Name string `json:"name" gorm:"not null;type:varchar(50)"`
	}

	//BookCategory type
	BookCategory struct {
		BookID     int `json:"book_id" gorm:"index"`
		CategoryID int `json:"publisher_id" gorm:"index"`
	}

	//Message type
	Message struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)
