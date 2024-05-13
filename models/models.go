package models

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Tags  string `json:"tags"`
}

type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}