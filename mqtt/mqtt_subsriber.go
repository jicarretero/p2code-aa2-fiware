package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jicarretero/p2code-aa2-fiware/config"
	_ "github.com/jicarretero/p2code-aa2-fiware/models"
)

var client mqtt.Client
var callback func(string, []byte)

func OnConnect(c mqtt.Client) {
	fmt.Println("Client connected ")
}

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	callback(msg.Topic(), msg.Payload())
}

func Subscribe(cfg *config.Config, cb func(string, []byte)) error {
	opts := mqtt.NewClientOptions().AddBroker(cfg.MQTT.MqttUrl)

	// Set Auto reconnection to true
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetMaxReconnectInterval(30 * time.Second)
	opts.SetKeepAlive(60 * time.Second)

	if cfg.MQTT.MqttUser != "" && cfg.MQTT.MqttPassword != "" {
		opts.SetUsername(cfg.MQTT.MqttUser)
		opts.SetPassword(cfg.MQTT.MqttPassword)
	}

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	fmt.Printf("Subscribing to topic: %s\n", cfg.MQTT.MqttTopic)
	if token := client.Subscribe(cfg.MQTT.MqttTopic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	callback = cb
	return nil
}
