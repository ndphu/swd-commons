package model

import (
	"image"
	"time"
)

var (
	TopicRecognizeRequest        = "/3ml/recognize/request"
	TopicRecognizeResponsePrefix = "/3ml/worker/response/generated/"
	TopicEventBroadcast          = "/3ml/event/broadcast"
	TopicNotificationBroadcast   = "/3ml/notification/broadcast"
	DefaultRecognizeRPCTimeout   = 30 * time.Second
)

type RecognizeRequest struct {
	Images              [][]byte `json:"payload"`
	FacesData           []Face   `json:"facesData"`
	IncludeFacesDetails bool     `json:"includeFacesDetails"`
	ClassifyFaces       bool     `json:"classifyFaces"`
	RequestId           string   `json:"requestId"`
	ResponseTo          string   `json:"responseTo"`
	DeskId              string   `json:"deskId"`
	AccessToken         string   `json:"accessToken"`
	TimeoutSeconds      int      `json:"timeoutSeconds"`
}

type RecognizeResponse struct {
	Labels          []string        `json:"labels"`
	FaceDetailsList [][]FaceDetails `json:"faceDetailsList"`
	Error           error           `json:"error"`
}

type FaceDetails struct {
	Rect       image.Rectangle `json:"rect"`
	Descriptor [128]float32    `json:"descriptor"`
	Image      []byte          `json:"image"`
}
