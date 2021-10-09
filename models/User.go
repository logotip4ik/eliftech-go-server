package models

import (
	"errors"
	"home/work/utils/token"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Posts    []Post `json:"posts"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
type CreateUserInput struct {
	Username string `json:"username" binding:"required,alphanum,min=3"`
	Email    string `json:"email" binding:"required,email,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}
type UpdateUserInput struct {
	Username string `json:"username" binding:"alphanum,min=3"`
	Email    string `json:"email" binding:"email,min=5"`
}

// func GetUserByUsername(username string) (User, error) {
// 	var user User

// 	if err := DB.Where("username = ?", username).Select("username", "email", "posts").First(&user).Error; err != nil {
// 		return user, errors.New("user not found")
// 	}

// 	var posts []Post
	
// 	if err := DB.Table("posts").Where("user_id = ?", user.ID).Order("created_at desc").Select("title", "content", "created_at", "updated_at").Find(&posts).Error; err != nil {
// 		return user, err
// 	}

// 	user.Password = ""
// 	user.Posts = posts

// 	return user, nil
// }

func GetUserByID(userID uint) (User, error) {
	var user User

	if err := DB.First(&user, userID).Error; err != nil {
		return user, errors.New("user not found")
	}

	var posts []Post
	
	if err := DB.Table("posts").Where("user_id = ?", user.ID).Select("username", "email").Order("created_at desc").Find(&posts).Error; err != nil {
		return user, err
	}

	user.Password = ""
	user.Posts = posts

	return user, nil 
}

func LoginCheck(email, password string) (string, error) {
	var err error
	u := User{}

	if err = DB.Model(User{}).Where("email = ?", email).Take(&u).Error; err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.Generate(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) SaveUser() (*User, error) {
	if err := DB.Create(&u).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
