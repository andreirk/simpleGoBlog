package models

import "time"

// Here we will create structs that model
// our DB tables
type User struct {
	Name        string
	Email       string
	Password    string
	AcctCreated time.Time
	LastLogin   time.Time
	UserType    int
	ID          int
}

type Post struct {
	Title   string
	Content string
	UserID  int
	ID      int
}
