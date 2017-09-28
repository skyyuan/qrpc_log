package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
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
	err = collection.Find(nil).Limit(10).Sort("-created_at").All(&qlogs)
	return
	return
}

func GetQlogsByTime(db *mgo.Database, time time.Time)(qlogs []QLog, err error){
	collection := db.C("qlogs")
	err = collection.Find(bson.M{"created_at": bson.M{"$gt": time}}).Sort("-created_at").All(&qlogs)
	return
	return
}