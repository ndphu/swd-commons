package model

import (
	"github.com/globalsign/mgo/bson"
)

type Face struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Label      string        `json:"label" bson:"label"`
	Descriptor [128]float32  `json:"descriptor" bson:"descriptor"`
	MD5        string        `json:"md5" bson:"md5"`
	DeviceId   string        `json:"deviceId" bson:"deviceId"`
	ProjectId  string        `json:"projectId" bson:"projectId"`
}