package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Notification struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	RuleId    bson.ObjectId `json:"ruleId" bson:"ruleId"`
	Type      string        `json:"type" bson:"type"`
	DeskId    string        `json:"deskId" bson:"deskId"`
	DeviceId  string        `json:"deviceId" bson:"deviceId"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	UserId    bson.ObjectId `json:"owner" bson:"owner"`
}
