package service

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"time"
)

func GetDefaultOps() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetConnectTimeout(30 * time.Second)
	return opts
}

func NewClientOpts(broker string) *mqtt.ClientOptions {
	opts := GetDefaultOps()
	opts.AddBroker(broker)
	opts.ClientID = uuid.New().String()
	return opts
}

func NewClientOptsWithId(broker string, clientId string) *mqtt.ClientOptions {
	opts := NewClientOpts(broker)
	opts.ClientID = clientId
	return opts
}
