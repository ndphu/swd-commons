package service

import (
	"github.com/eclipse/paho.mqtt.golang"
	"time"
)

func GetDefaultOps(broker string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetConnectTimeout(30 * time.Second)
	return opts
}
