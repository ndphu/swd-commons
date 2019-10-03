package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

var (
	SitStatusPresent = "PRESENT"
	SitStatusMissing = "MISSING"
)

type SitTracking struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	UserId       bson.ObjectId `json:"userId" bson:"userId"`
	DeviceId     string        `json:"deviceId" bson:"deviceId"`
	Status       string        `json:"status" bson:"status"`
	TrackingTime time.Time     `json:"trackingTime" bson:"trackingTime"`
}
