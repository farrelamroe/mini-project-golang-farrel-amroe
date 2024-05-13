package controller

import (
	// "database/sql"
	// "log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"project/models"
)

var users []models.User

// Register handles user registration
func Register(c echo.Context) error {
	user := new(models.User)
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
		if u.Email == email && u.Password == password && u.UserName == username{
			return c.JSON(http.StatusOK, "Login successful")
		}
	}

	return c.JSON(http.StatusUnauthorized, "Invalid credentials")
}

// Get a book by ID
// Fetch books from database
func GetBook(c echo.Context) error {
	// Implement your SQL query here
	// rows, err := DB.Query("SELECT id, title, tags FROM books")
	// if err != nil {
	// 	log.Println("Error fetching books:", err)
	// 	return c.JSON(http.StatusInternalServerError, "Internal server error")
	// }
	// defer rows.Close()

	// var books []models.Book
	// for rows.Next() {
	// 	var book models.Book
	// 	if err := rows.Scan(&book.ID, &book.Title, &book.Tags); err != nil {
	// 		log.Println("Error scanning row:", err)
	// 		continue
	// 	}
	// 	books = append(books, book)
	// }

	// if err := rows.Err(); err != nil {
	// 	log.Println("Error iterating rows:", err)
	// 	return c.JSON(http.StatusInternalServerError, "Internal server error")
	// }
	book := []models.Book{
		{ID: 1, Title: "Clean Code", Tags: "programming, software"},
		{ID: 2, Title: "Design Patterns", Tags: "programming, architecture"},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": book,
	})

}

func GetBookById(c echo.Context) error {
	book := models.Book{ID: int(uuid.New().ID()), Title: "Clean Code", Tags: "programming, software"}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": book,
	})

}

// Create a new book
func CreateBook(c echo.Context) error {
	// Parse request body
	// var book Book
	// if err := c.Bind(&book); err != nil {
	// 	return err
	// }

	// // Insert book into database
	// result, err := DB.Exec("INSERT INTO books (title, tags) VALUES (?, ?)", book.Title, book.Tags)
	// if err != nil {
	// 	log.Println("Error creating book:", err)
	// 	return c.JSON(http.StatusInternalServerError, "Internal server error")
	// }

	// // Get the auto-generated ID of the newly inserted book
	// bookID, _ := result.LastInsertId()
	// book.ID = int(bookID)
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, book)
}

// Update a book by ID
func UpdateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))
	bookCatalog := models.Book{}
	c.Bind(&bookCatalog)

	book := []models.Book{
		{ID: int(uuid.New().ID()), Title: "Clean Code", Tags: "programming, software"},
		{ID: id, Title: "Design Patterns", Tags: "programming, architecture"},
	}

	for i, books := range book {
		if id == books.ID {
			book[i].ID = bookCatalog.ID
			book[i].Title = bookCatalog.Title
			book[i].Tags = bookCatalog.Tags
		}
	}

	// Update book in database
	// _, err := DB.Exec("UPDATE books SET title = ?, tags = ? WHERE id = ?", book.Title, book.Tags, id)
	// if err != nil {
	// 	log.Println("Error updating book:", err)
	// 	return c.JSON(http.StatusInternalServerError, "Internal server error")
	// }

	return c.JSON(http.StatusOK, book)
}

// Delete a book by ID
func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))

	// Delete book from database
	// _, err := DB.Exec("DELETE FROM books WHERE id = ?", id)
	// if err != nil {
	// 	log.Println("Error deleting book:", err)
	// 	return c.JSON(http.StatusInternalServerError, "Internal server error")
	// }

	book := []models.Book{
		{ID: int(uuid.New().ID()), Title: "Clean Code", Tags: "programming, software"},
		{ID: id, Title: "Design Patterns", Tags: "programming, architecture"},
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