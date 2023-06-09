package connection

import (
	"fmt"
	"os"
	"github.com/james-kariuki-source/Timetable-Management-API/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	godotenv.Load()
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection failed!")
	}

	DB = db

	fmt.Println("The databse connection was successful!")

	AutoMigrate(db)


}

func AutoMigrate(connection *gorm.DB){
	connection.AutoMigrate(
		&models.Admin{},
		&models.Building{},
		&models.Halls{},
		&models.Rooms{},
		&models.Lecturer{},
		&models.PremisesManager{},
		&models.Registrar{},
		&models.Student{},
	)
}