package repository

import (
	"errors"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/google/uuid"
)
var dsn string = "root:admin@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local" 

type Album struct {
	//gorm.Model
	ID int64 `gorm:"primaryKey,column:id"`
	Title string `gorm:"column:title"`
	Artist string `gorm:"column:artist"`
	Price float32 `gorm:"column:price"`
}
type Tabler interface {
	TableName() string
  }
  
  // TableName overrides the table name used by Album to `album`
  func (Album) TableName() string {
	return "album"
  }

func conn() (gorm.DB,error){
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:admin@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return *db, errors.New("Connection_Error")
	}
	log.Println("Connection Open: " + db.Name())
	return *db,nil
  }

func Create(album Album)(error){
	db , err := conn()
	if err != nil {
		return err
	}
	result := db.Create(&album)
	return result.Error
} 
func (u *Album) BeforeCreate(tx *gorm.DB) (err error) {
	//u.Title = uuid.New().String()
	//u.Artist = uuid.New().String()
	return
  }

func AlbumsByArtist(artist string)([]Album,error){
	db , err := conn()
	if err != nil {
		return nil,err
	}
	var result []Album
	//cb := db.ClauseBuilders
	rows, err := db.Model(&Album{}).Where("artist LIKE ?",artist).Select("*").Rows()
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	var album Album
	for rows.Next() {
	// ScanRows scan a row into album
	errror := db.ScanRows(rows, &album)
	if errror != nil {
		return nil,errror
	}
	result = append(result, album)
	// do something
	}
	return result,nil
}