package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

var (
	EventRecognizeSuccess = "RECOGNIZE_SUCCESS"
	EventRecognizeFail    = "RECOGNIZE_FAIL"
	EventCaptureFail      = "CAPTURE_FAIL"

	ResultPresent = "PRESENT"
	ResultMissing = "MISSING"
)

type Event struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	DeviceId  string        `json:"deviceId" bson:"deviceId"`
	Labels    []string      `json:"labels"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	Type      string        `json:"type" bson:"type"`
	Error     string        `json:"error" bson:"error"`
	Result    string        `json:"result" bson:"result"`
	UserId    bson.ObjectId `json:"owner" bson:"owner"`
}
