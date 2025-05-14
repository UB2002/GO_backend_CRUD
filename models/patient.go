package models

import "github.com/jinzhu/gorm"

type Patient struct {
    gorm.Model
    Name   string
    Age    int
    Gender string
    Notes  string
}
