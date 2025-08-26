package models

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Role     string `json:"role"` // admin or employee
}