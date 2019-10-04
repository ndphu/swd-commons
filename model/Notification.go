package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

var (
	NotificationSeverityGood    = "good"
	NotificationSeverityWarning = "warning"
	NotificationSeverityDanger  = "danger"

	NotificationTypeSlack = "SLACK"
)

type Notification struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Type            string        `json:"type" bson:"type"`
	DeskId          string        `json:"deskId" bson:"deskId"`
	DeviceId        string        `json:"deviceId" bson:"deviceId"`
	Timestamp       time.Time     `json:"timestamp" bson:"timestamp"`
	UserId          bson.ObjectId `json:"owner" bson:"owner"`
	SitDuration     time.Duration `json:"sitDuration" bson:"sitDuration"`
	Rule            Rule          `json:"rule" bson:"rule"`
}
