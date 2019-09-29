package model

import "github.com/globalsign/mgo/bson"

type Rule struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	DeskId string        `json:"deskId" bson:"deskId"`
	DeviceId  string        `json:"deviceId" bson:"deviceId"`
	Action    Action        `json:"action" bson:"action"`
	Interval  int        `json:"interval" bson:"interval"`
}

type Action struct {
	Type string `json:"type" bson:"type"`
}
