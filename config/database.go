package config

import (
    "example.com/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var Database *gorm.DB
var DATABASE_URI string = "root:new_password@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
    var err error

    Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&models.Notes{},&models.User{})

    err1 := Database.Migrator().DropColumn(&models.User{}, "id")
    if err1 != nil {
        // Do whatever you want to do!
        log.Print("ERROR: We expect the description column to be  drop-able")
    }    
    
    return nil
}
func AddUser( Username, Password string ) error {
    var err error

    Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }
    User := models.User{
        Username:   Username,
        Passsword: Password,
    }
    Database.Create(&User) 

    return nil
}