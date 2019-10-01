package model

import "github.com/globalsign/mgo/bson"

type SlackConfig struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	UserId         bson.ObjectId `json:"userId" bson:"userId"`
	SlackUserId    string        `json:"slackUserId" bson:"slackUserId"`
	SentInvitation bool          `json:"sendInvitation" bson:"sendInvitation"`
}
