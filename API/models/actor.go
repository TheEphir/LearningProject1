package models

import (
	"github.com/TheEphir/LearningProject1/API/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Actor struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type DbActor struct {
	id uint
	Actor
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Actor{})
}

func GetActors() []DbActor {
	var DbActors []DbActor
	db.Raw("SELECT actor_id, first_name, last_name FROM actor").Scan(&DbActors)
	return DbActors
}

func (a *Actor) AddActor() *Actor {
	db.Create(&a)
	return a
}
