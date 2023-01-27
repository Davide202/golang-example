package gorm

import (
	"fmt"
	"log"
	"math/rand"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	"root", "admin", "127.0.0.1:3306", "db_users",
)

type User struct {
	Id          int64  `gorm:"column:id"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	Email       string `gorm:"column:email"`
	DateCreated string `gorm:"column:date_created"`
	Status      string `gorm:"column:status"`
	Password    string `gorm:"column:password"`
}

type UserDTO struct {
	//Id          int64  `gorm:"column:id"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	Email       string `gorm:"column:email"`
	DateCreated string `gorm:"column:date_created"`
	Status      string `gorm:"column:status"`
	//Password    string `gorm:"column:password"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Album to `album`
func (User) TableName() string {
	return "users"
}
func conn() (gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return *db, err
	}
	log.Println("Connection Open: " + db.Name())
	return *db, nil
}

func GetUserByEmail(email *string) ([]User, error) {
	db, err := conn()
	if err != nil {
		return nil, err
	}
	const per string = "%"
	param := per + *email + per
	rows, err := db.Model(&User{}).Where("email LIKE ?", param).Find(&UserDTO{}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]User, 0)
	var user User
	for rows.Next() {
		errror := db.ScanRows(rows, &user)
		if errror != nil {
			return nil, errror
		}
		users = append(users, user)
	}
	return users, nil
}

func Create(user *User) error {
	user.Id = rand.Int63()
	db, err := conn()
	if err != nil {
		return err
	}
	result := db.Create(&user)
	return result.Error
}

func Update(user *User) error {
	user.Id = rand.Int63()
	db, err := conn()
	if err != nil {
		return err
	}
	result := db.Save(&user)
	return result.Error
}
