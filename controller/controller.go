package controller

import (
	"net/http"
	"strconv"
	"time"

	"project/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var users []model.User

// Register handles user registration
func Register(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Save user to database (in-memory slice for demonstration)
	users = append(users, *user)

	return c.JSON(http.StatusCreated, "User registered successfully")
}

// Login handles user login
func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	username := c.FormValue("username")

	// Check if user exists in the database
	for _, u := range users {
		if u.Email == email && u.Password == password && u.UserName == username {
			return c.JSON(http.StatusOK, "Login successful")
		}
	}

	return c.JSON(http.StatusUnauthorized, "Invalid credentials")
}

// Get a book by ID
func GetBook(c echo.Context) error {
	book := []model.Book{
		{ID: int(uuid.New().ID()), Title: "Clean Code", Tags: "programming, software", CreatedAt: time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: int(uuid.New().ID()), Title: "Design Patterns", Tags: "programming, architecture", CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: int(uuid.New().ID()), Title: "Learning Python", Tags: "programming", CreatedAt: time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: int(uuid.New().ID()), Title: "Learning Javascript", Tags: "programming", CreatedAt: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: int(uuid.New().ID()), Title: "History of Java", Tags: "history, culture", CreatedAt: time.Date(2023, 12, 11, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: int(uuid.New().ID()), Title: "Learning Basic Web Programming", Tags: "programming, architecture", CreatedAt: time.Now(), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": book,
	})

}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))
	book := []model.Book{
		{ID: int(uuid.New().ID()), Title: "Clean Code", Tags: "programming, software", CreatedAt: time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: id, Title: "Design Patterns", Tags: "programming, architecture", CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: id + 1, Title: "Learning Python", Tags: "programming", CreatedAt: time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: id + 2, Title: "Learning Javascript", Tags: "programming", CreatedAt: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: id + 3, Title: "History of Java", Tags: "history, culture", CreatedAt: time.Date(2023, 12, 11, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: id + 4, Title: "Learning Basic Web Programming", Tags: "programming, architecture", CreatedAt: time.Now(), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": book,
	})

}

// Create a new book
func CreateBook(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"books": book,
	})
}

// Update a book by ID
func UpdateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))
	bookCatalog := model.Book{}
	c.Bind(&bookCatalog)

	book := []model.Book{
		{ID: int(uuid.New().ID()), Title: "Clean Code", Tags: "programming, software", CreatedAt: time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: id, Title: "Design Patterns", Tags: "programming, architecture", CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
	}

	for i, books := range book {
		if id == books.ID {
			book[i].ID = bookCatalog.ID
			book[i].Title = bookCatalog.Title
			book[i].Tags = bookCatalog.Tags
		}
	}

	return c.JSON(http.StatusOK, book)
}

// Delete a book by ID
func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))

	book := []model.Book{
		{ID: int(uuid.New().ID()), Title: "Clean Code", Tags: "programming, software", CreatedAt: time.Date(2022, 3, 21, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
		{ID: id, Title: "Design Patterns", Tags: "programming, architecture", CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC), Category: model.Categories{CategoryID: int(uuid.New().ID()), Category: "Book"}},
	}

	indexToDelete := -1
	for i, books := range book {
		if books.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		book = append(book[:indexToDelete], book[indexToDelete+1:]...)
	}

	return c.NoContent(http.StatusNoContent)
}
