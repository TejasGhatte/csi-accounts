package schemas

import "time"

type User struct {
	Email       string    `json:"email" validate:"required,email"`
	Did         Did       `json:"did" validate:"required"`
	Name        string    `json:"name" validate:"required,max=25"`
	Role        string    `json:"role" validate:"required,oneof='jrcore' 'srcore' 'board'"`
	Domain      string    `json:"domain" validate:"required,oneof='management' 'design' 'tech'"`
	DateCreated time.Time `json:"dateCreated" validate:"required"`
}
