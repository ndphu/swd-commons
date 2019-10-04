package service

import (
	"encoding/json"
	"errors"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/ndphu/swd-commons/model"
	"log"
	"time"
)

func CallRecognizeWithRequest(opts *mqtt.ClientOptions, req model.RecognizeRequest) (*model.RecognizeResponse, error) {
	log.Println("[RECOGNIZE]", "CallRecognizeWithRequest for desk", req.DeskId)
	return callRPC(opts, req)
}

func CallBulkRecognizeWithProvidedFacesData(opts *mqtt.ClientOptions, deskId string, frames [][]byte, faces []model.Face) (*model.RecognizeResponse, error) {
	log.Println("[RECOGNIZE]", "CallBulkRecognize for desk", deskId, "with provided faces data")
	reqId := uuid.New().String()
	rpcReq := NewRecognizeRequest(reqId, frames, deskId)
	rpcReq.FacesData = faces
	return callRPC(opts, rpcReq)
}

func CallBulkRecognize(opts *mqtt.ClientOptions, deskId string, frames [][]byte, accessToken string) (*model.RecognizeResponse, error) {
	log.Println("[RECOGNIZE]", "CallBulkRecognize for desk", deskId, "with accessToken")
	reqId := uuid.New().String()
	rpcReq := NewRecognizeRequest(reqId, frames, deskId)
	rpcReq.AccessToken = accessToken
	return callRPC(opts, rpcReq)
}

func callRPC(opts *mqtt.ClientOptions, rpcReq model.RecognizeRequest) (*model.RecognizeResponse, error) {
	if rpcReq.RequestId == "" {
		rpcReq.RequestId = uuid.New().String()
	}

	rpcRequestTopic := model.TopicRecognizeRequest
	rpcReq.ResponseTo = model.TopicRecognizeResponsePrefix + uuid.New().String()

	clientId := uuid.New().String()
	opts.ClientID = clientId

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error().Error())
	}
	defer c.Disconnect(500)

	rpcResponse := make(chan model.RecognizeResponse)

	c.Subscribe(rpcReq.ResponseTo, 0, func(c mqtt.Client, m mqtt.Message) {
		resp := model.RecognizeResponse{}
		if err := json.Unmarshal(m.Payload(), &resp); err != nil {
			log.Println("[RPC]", rpcReq.RequestId, "fail to unmarshal response")
			rpcResponse <- model.RecognizeResponse{
				Error: err,
			}
		} else {
			log.Println("[RPC]", rpcReq.RequestId, "received response")
			rpcResponse <- resp
		}
	}).Wait()

	rpcReqPayload, _ := json.Marshal(rpcReq)
	c.Publish(rpcRequestTopic, 0, false, rpcReqPayload).Wait()

	timeoutTimer := time.NewTimer(model.DefaultRecognizeRPCTimeout)
	if rpcReq.TimeoutSeconds > 0 {
		timeoutTimer = time.NewTimer(time.Duration(rpcReq.TimeoutSeconds) * time.Second)
	}

	select {
	case resp := <-rpcResponse:
		return &resp, nil
	case <-timeoutTimer.C:
		log.Println("[RPC]", rpcReq.RequestId, "timeout occurred.")
		return nil, errors.New("timeout")
	}
}

func NewRecognizeRequest(reqId string, frames [][]byte, deskId string) model.RecognizeRequest {
	return model.RecognizeRequest{
		RequestId:           reqId,
		Images:              frames,
		IncludeFacesDetails: false,
		ClassifyFaces:       false,
		DeskId:              deskId,
	}
}
