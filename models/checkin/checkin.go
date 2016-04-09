package checkin

import (
	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2"
	"time"
)

type CheckIn struct {
	Id_         bson.ObjectId `bson:"_id"`
	UserId      bson.ObjectId `bson:"user_id"`
	Content     string        `bson:"content"`
	CreateAt    time.Time     `bson:"create_at"`
	CreateDay   int           `bson:"create_day"`
	CreateMonth int           `bson:"create_month"`
	CreateYear  int           `bson:"create_year"`
	CreateHour	int           `bson:"create_hour"`
	CreateMin   int           `bson:"create_min"`
	CreateSec	int           `bson:"create_sec"`
	Timestamp   int64         `bson:"timestamp"`
	Images      []string      `bson:"images"`
}
