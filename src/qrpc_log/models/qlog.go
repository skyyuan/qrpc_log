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
	TraceId   string       `bson:"trace_id", json:"trace_id"`
	Content    string        `bson:"content", json:"content"`
	CommonModel `bson:",inline"`
}

func GetQlogs(db *mgo.Database)(qlogs []QLog, err error){
	collection := db.C("qlogs")
	err = collection.Find(nil).Limit(10).Sort("-created_at").All(&qlogs)
	return
	return
}

func GetQlogsByParams(db *mgo.Database, level, bType, traceId, content string)(qlogs []QLog, err error){
	collection := db.C("qlogs")
	query := bson.M{}
	if level != "" {
	    query["level"] = level
	}
	if bType != "" {
	    query["b_type"] = bType
	}
	if traceId != "" {
		query["trace_id"] = traceId
	}
	if content != "" {
		query["content"] = bson.M{"$regex": content}
	}
	err = collection.Find(query).Limit(10).Sort("-created_at").All(&qlogs)
	return
}

func GetQlogsByTime(db *mgo.Database, time time.Time, level, bType, traceId, content string)(qlogs []QLog, err error){
	collection := db.C("qlogs")
	query := bson.M{"created_at": bson.M{"$gt": time}}
	if level != "" {
		query["level"] = level
	}
	if bType != "" {
		query["b_type"] = bType
	}
	if traceId != "" {
		query["trace_id"] = traceId
	}
	if content != "" {
		query["content"] = bson.M{"$regex": content}
	}
	err = collection.Find(query).Sort("-created_at").All(&qlogs)
	return
}