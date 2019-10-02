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

func CallBulkRecognizeWithProvidedFacesData(opts *mqtt.ClientOptions, deskId string, frames [][]byte, faces []model.Face) (*model.BulkRecognizeResponse, error) {
	log.Println("[RECOGNIZE]", "CallBulkRecognize for desk", deskId, "with provided faces data")
	reqId := uuid.New().String()
	rpcReq := newRPCRequest(reqId, frames, deskId)
	rpcReq.FacesData = faces
	return callRPC(opts, rpcReq)
}

func CallBulkRecognize(opts *mqtt.ClientOptions, deskId string, frames [][]byte, accessToken string) (*model.BulkRecognizeResponse, error) {
	log.Println("[RECOGNIZE]", "CallBulkRecognize for desk", deskId, "with accessToken")
	reqId := uuid.New().String()
	rpcReq := newRPCRequest(reqId, frames, deskId)
	rpcReq.AccessToken = accessToken
	return callRPC(opts, rpcReq)
}

func callRPC(opts *mqtt.ClientOptions, rpcReq model.BulkRecognizeRequest) (*model.BulkRecognizeResponse, error) {
	rpcRequestTopic := "/3ml/recognize/request"
	rpcResponseTopic := "/3ml/worker/response/generated/" + uuid.New().String()
	rpcReq.ResponseTo = rpcResponseTopic

	clientId := uuid.New().String()
	opts.ClientID = clientId

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error().Error())
	}
	defer c.Disconnect(500)

	rpcResponse := make(chan model.BulkRecognizeResponse)

	c.Subscribe(rpcResponseTopic, 0, func(c mqtt.Client, m mqtt.Message) {
		resp := model.BulkRecognizeResponse{}
		if err := json.Unmarshal(m.Payload(), &resp); err != nil {
			log.Println("[RPC]", rpcReq.RequestId, "fail to unmarshal response")
			rpcResponse <- model.BulkRecognizeResponse{
				Error: err,
			}
		} else {
			log.Println("[RPC]", rpcReq.RequestId, "received response")
			rpcResponse <- resp
		}
	}).Wait()

	rpcReqPayload, _ := json.Marshal(rpcReq)
	c.Publish(rpcRequestTopic, 0, false, rpcReqPayload).Wait()

	rpcTimeout := time.NewTimer(15 * time.Second)
	select {
	case resp := <-rpcResponse:
		return &resp, nil
	case <-rpcTimeout.C:
		log.Println("[RPC]", rpcReq.RequestId, "timeout occurred.")
		return nil, errors.New("timeout")
	}
}

func newRPCRequest(reqId string, frames [][]byte, deskId string) model.BulkRecognizeRequest {
	return model.BulkRecognizeRequest{
		RequestId:           reqId,
		Images:              frames,
		IncludeFacesDetails: true,
		DeskId:              deskId,
	}
}
