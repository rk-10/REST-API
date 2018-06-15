package dao

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)

type Movie struct {
	ID bson.ObjectId	`bson:"id" json:"id"`
	Name string			`bson:"name" json:"name"`
	CoverImage string	`bson:"cover_image" json:"cover_image"`
	Description string	`bson:"description" json:"description"`
}

type MoviesDAO struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MoviesDAO) Connect()  {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}