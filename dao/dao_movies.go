package dao

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
	. "github.com/rk-10/REST-API/models"
)


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
		log.Fatal("Could not connect to mongo", err)
	}

	db = session.DB(m.Database)
}

func (m *MoviesDAO) FindAll() ([]Movie, error)  {
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