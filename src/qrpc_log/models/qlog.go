package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type QLog struct {
	Id_       bson.ObjectId `bson:"_id"`
	BType string           `bson:"b_type", json:"b_type"`
	BFlag    string        `bson:"b_flag", json:"b_flag"`
	Level    string        `bson:"level", json:"level"`
	Content    string        `bson:"content", json:"content"`
	CommonModel `bson:",inline"`
}

func GetQlogs(db *mgo.Database)(qlogs []QLog, err error){
	collection := db.C("qlogs")
	err = collection.Find(nil).Sort("-created_at").All(&qlogs)
	return
	return
}