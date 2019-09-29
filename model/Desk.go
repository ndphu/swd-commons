package model

import "github.com/globalsign/mgo/bson"

type Desk struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	DeskId string        `json:"deskId" bson:"deskId"`
	Name   string        `json:"name" bson:"name"`
}
