package models

import "github.com/jinzhu/gorm"

type Role string

const (
    RoleReceptionist Role = "receptionist"
    RoleDoctor       Role = "doctor"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Role     Role   `gorm:"not null"`
}
