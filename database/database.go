package database

import (
    "fmt"
    "github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/database/migrations"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/gallery-db?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }

    fmt.Println("Connected to the database")

}

func MigrateDB() {
    migrations.Migrate(DB)
    fmt.Println("Database migrated")
}
