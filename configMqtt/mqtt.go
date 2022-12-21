package configmqtt

import (
	"fmt"
	"os"
	"os/signal"

	"rishabhjha12-hub/constants"

	"rishabhjha12-hub/helper"

	"syscall"
	"time"

	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	//"github.com/getsentry/sentry-go"
)

//https://uat.fastag.ai/api/v2/icd2.5/fastag_status/

var HandleMessage MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	log.Printf("MSG: %s\n", msg.Payload())

	//helper function to parse
	helper.ParseHelper(msg.Payload())

}

// this function will be called at the time of connection lost
var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
	//250ms
	client.Disconnect(250)

	//after 5 seconds try to reconnect
	time.AfterFunc(5*time.Second, func() {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())

		} else {
			fmt.Printf("Connected to server\n")
		}
	})
}

func MqttConfig() {

	//this will create a channel through signal
	//reference = https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/stdoutsub/main.go
	channel := make(chan os.Signal, constants.No_channels)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	//Connect to mqtt
	//opts := MQTT.NewClientOptions().AddBroker(constants.Broker_Protocol + constants.BrokerIpp + ":" + constants.Brokerport)
	//for testing
	opts := MQTT.NewClientOptions().AddBroker("tcp://172.27.0.103:8883")

	//setting defualt publish handler
	opts.SetDefaultPublishHandler(HandleMessage)

	//when the client is connected to the server
	opts.OnConnect = func(client MQTT.Client) {
		//subscribe to the topic  and request messages to be delivered
		//at a maximum qos(quality of service) of zero, wait for the receipt to confirm the subscription

		//qos reference = https://www.hivemq.com/blog/mqtt-essentials-part-6-mqtt-quality-of-service-levels/
		if token := client.Subscribe(constants.Gate+constants.Jetson_ANPR, 0, HandleMessage); token.Wait() && token.Error() != nil {
			//panic(token.Error())
			log.Println(token.Error())
		}
		if token := client.Subscribe(constants.Gate+constants.Parkzap_epc, 0, HandleMessage); token.Wait() && token.Error() != nil {
			// panic(token.Error())
			log.Println(token.Error())
			//sentry.CaptureException(token.Error())
		}

		if token := client.Subscribe(constants.Gate+constants.Parkzap_loop, 0, HandleMessage); token.Wait() && token.Error() != nil {
			//panic(token.Error())
			log.Println(token.Error())
			//sentry.CaptureException(token.Error())
		}
	}

	//creating a mqtt client
	client := MQTT.NewClient(opts)

	//making connection
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		//panic(token.Error())
		log.Fatal("error is ", token.Error())
		//sentry.CaptureException(token.Error())

	} else {
		log.Printf("Connected to server\n")
	}

	//Error handling /retry
	//when the client is disconnecected
	opts.OnConnectionLost = connectLostHandler

	//channel closed
	<-channel

}
