package database

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    "log"
	"github.com/lokesh2201013/Usermanagement/models"
)

var DB *gorm.DB
func ConnectDB(){
 dsn := "host=localhost user=postgres password= 9910994194lokesh dbname= usermanagement port= 5432 sslmode=disable"
 db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

 if err!=nil{
 	log.Fatalf("Failed to connect to db ,err = %v\n",err)								
 }

 db.AutoMigrate(&models.User{},&models.Admin{})
 log.Println("Connected to PostgreSQL using GORM")
 DB=db
}
