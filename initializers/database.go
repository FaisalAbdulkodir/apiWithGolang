package initializers

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var db *sql.DB
// var err error

// func ConnectDB() {
// 	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/music")

// if err != nil {
// 	log.Fatal(err)
// }

//		err = db.Ping()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := "root@tcp(127.0.0.1:3306)/music?charset=utf8&parseTime=True"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
}
