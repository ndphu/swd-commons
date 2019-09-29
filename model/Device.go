package model

import "github.com/globalsign/mgo/bson"

type Device struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	DeviceId string        `json:"deviceId" bson:"deviceId"`
	Name     string        `json:"name" bson:"name"`
	Status   string        `json:"status" bson:"status"`
	DeskId   string        `json:"deskId" bson:"deskId"`
	Owner    bson.ObjectId `json:"owner" bson:"owner"`
}
