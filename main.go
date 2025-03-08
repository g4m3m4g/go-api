package main

import (
	"log"
	//"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	//"github.com/gofiber/jwt/v2"
)

type Book struct {
	ID     int		`json:"id"`
	Title  string	`json:"title"`
	Author string	`json:"author"`
}

var books []Book

func main() {
	if err := godotenv.Load(); err != nil{
		log.Fatal("loading .env file failed")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Gamie", Author: "John Doe"})
	books = append(books, Book{ID: 2, Title: "Game of Thrones", Author: "George Martin"})
	books = append(books, Book{ID: 3, Title: "Harry Potter", Author: "J.K. Rowling"})

	//app.Post("/login", login)
	//app.Use(jwtware.New(jwtware.Config{
	//	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	//}))

	app.Use(middleware)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test", testHTML)

	app.Listen(":8080")

}




