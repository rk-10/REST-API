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

func (m *MoviesDAO) FindAll(id string) ([]Movie, error)  {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesDAO) FindbyId(id string)  (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MoviesDAO) Insert(movie Movie) error  {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesDAO) Remove(movie Movie) error  {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

func (m *MoviesDAO) Update(movie Movie) error  {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}