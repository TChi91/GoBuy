package data

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/TChi91/GoBuy/db"
	"github.com/jinzhu/gorm"
)

//User model
type User struct {
	gorm.Model
	UserName  string `json:"username" gorm:"size:15;unique_index" validate:"required"`
	FirstName string `json:"first_name" gorm:"size:15" validate:"required"`
	LastName  string `json:"last_name" gorm:"size:15" validate:"required"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index" validate:"required"`
	Password  string `json:"password" gorm:"size:255" validate:"required"`
}

//Clean the user fields
func (u *User) Clean() *User {
	u.UserName = strings.ToLower(u.UserName)
	u.FirstName = strings.Title(strings.ToLower(u.FirstName))
	u.LastName = strings.Title(strings.ToLower(u.LastName))
	u.Email = strings.ToLower(u.Email)
	return u
}

//Create new user
func Create(u *User) error {
	err := db.Db.Create(u).Error
	return err
}

//FromJSON for marshaling product
func (u *User) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(u)
	return err
}

//ToJSON for marshaling product
func (u *User) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(u)
	return err
}
