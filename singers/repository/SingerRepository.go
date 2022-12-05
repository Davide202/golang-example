package repository

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	//"github.com/google/uuid"
)
var init_db string = "create database if not exists recordings; use recordings;  DROP TABLE IF EXISTS album; CREATE TABLE album (   id         INT AUTO_INCREMENT NOT NULL,   title      VARCHAR(128) NOT NULL,   artist     VARCHAR(255) NOT NULL,   price      DECIMAL(5,2) NOT NULL,   PRIMARY KEY (`id`) );  INSERT INTO album   (title, artist, price) VALUES   ('Blue Train', 'John Coltrane', 56.99),   ('Giant Steps', 'John Coltrane', 63.99),   ('Jeru', 'Gerry Mulligan', 17.99),   ('Sarah Vaughan', 'Sarah Vaughan', 34.98);"

//var dsn string = "root:admin@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
//var dsn string = "https://davide202-probable-rotary-phone-w47vv6wj5jp396xj-3306.preview.app.github.dev/"
var dsn string = "root:admin@tcp(172.18.0.2:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"

type Album struct {
	//gorm.Model
	ID int64 `gorm:"primaryKey,column:id"`
	//APIAlbum 
	Title  string  `gorm:"column:title"`
	Artist string  `gorm:"column:artist"`
	Price  float32 `gorm:"column:price"`
}
type APIAlbum struct {
	//gorm.Model
	//ID int64 `gorm:"primaryKey,column:id"`
	Title  string  `gorm:"column:title"`
	Artist string  `gorm:"column:artist"`
	Price  float32 `gorm:"column:price"`
}
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Album to `album`
func (Album) TableName() string {
	return "album"
}

func conn() (gorm.DB, error) {
	user := os.Getenv("DB_USER")
	log.Println("Env variable: " + user)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:admin@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	dtb := mysql.Open(dsn)
	db, err := gorm.Open(dtb, &gorm.Config{})
	if err != nil {
		return *db, errors.New("Connection_Error")
	}
	database , er := db.DB()
	if er != nil {
		return *db, errors.New("Connection_Error")
	}
	database.Query(init_db)
	log.Println("Connection Open: " + db.Name())
	return *db, nil
}

func Create(album Album) error {
	db, err := conn()
	if err != nil {
		return err
	}
	result := db.Create(&album)
	db.Commit()
	return result.Error
}
func (u *Album) BeforeCreate(tx *gorm.DB) (err error) {
	//u.Title = uuid.New().String()
	//u.Artist = uuid.New().String()
	return
}

func AlbumsByArtist(artist string) ([]Album, error) {
	db, err := conn()
	if err != nil {
		return nil, err
	}
	var result []Album
	//cb := db.ClauseBuilders
	const per string = "%"
	param := per + artist + per
	rows, err := db.Model(&Album{}).Where("artist LIKE ?", param).Find(&APIAlbum{}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var album Album
	for rows.Next() {
		// ScanRows scan a row into album
		errror := db.ScanRows(rows, &album)
		if errror != nil {
			return nil, errror
		}
		result = append(result, album)
		// do something
	}
	return result, nil
}
