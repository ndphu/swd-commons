package model

import (
	"github.com/globalsign/mgo/bson"
)

var (
	RuleTypeSittingMonitoring  = "SITTING_MONITORING"
	RuleTypeDrinkWaterReminder = "WATER_REMINDER"

	DefaultSittingRemindInterval = 60
	DefaultDrinkRemindInterval   = 15
)

type Rule struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	DeskId          string        `json:"deskId" bson:"deskId"`
	IntervalMinutes int           `json:"intervalMinutes" bson:"intervalMinutes"`
	Type            string        `json:"type" bson:"type"`
	UserId          bson.ObjectId `json:"userId" bson:"userId"`
}

type Action struct {
	Type string `json:"type" bson:"type"`
}
