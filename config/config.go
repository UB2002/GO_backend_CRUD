package config

import (
	"os"
	"patient_portal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Config struct {
    DBUrl     string
    JWTSecret string
}

func seedUsers(db *gorm.DB) {
  defaults := []models.User{
    {Username: "receptionist1", Password: "pass123", Role: models.RoleReceptionist},
    {Username: "doctor1",       Password: "pass123", Role: models.RoleDoctor},
  }
  for _, u := range defaults {
    var existing models.User
    if db.Where("username = ?", u.Username).First(&existing).RecordNotFound() {
      db.Create(&u)
    }
  }
}


func Load() *Config {
    return &Config{
        DBUrl:     os.Getenv("DATABASE_URL"),
        JWTSecret: os.Getenv("JWT_SECRET"),
    }
}

func ConnectDB(c *Config) *gorm.DB {
    db, err := gorm.Open("postgres", c.DBUrl)
    if err != nil {
        panic(err)
    }
	  seedUsers(db)
    db.AutoMigrate(&models.User{}, &models.Patient{})
    return db
}