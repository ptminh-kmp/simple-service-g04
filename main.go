package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id      int    `json:id,omitempty gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// insert new note
	// newNote := Note{Title: "Demo note", Content: "This is Content"}
	// if err := db.Create(&newNote); err != nil {
	// 	fmt.Println(err)
	// }

	var notes []Note
	db.Where("status = ?", 1).Find(&notes)

	var note Note
	if err := db.Where("id = 4").First(&note); err != nil {
		log.Println(err)
	}

	//db.Table(Note{}.TableName()).Where("id = 3").Delete(nil)

	db.Table(Note{}.TableName()).Where("id = 2").Updates(map[string]interface{}{
		"title": "Demo 2",
	})

	fmt.Println(note)
}
