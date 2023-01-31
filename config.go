package main

import (
	"machine"
)

var (
	led = machine.LED
	led2 = machine.GP9

	MQTT_type string = "MQIsdp"
	MQTT_CID string = "Lada"
	MQTT_user string = "user_xxxxxxxx"
	MQTT_pass string = "pass_xxxxxxxx"
	
	broker bool = false
	Serial      = machine.UART0
	SIM800      = machine.UART1
	MQTT_SERVER = "srv.clusterfly.ru"
	PORT        = "9991"
	AT          = ""
	comconnect  = false
	timerCh = make(chan struct{}, 1)
)
func Configure(){
	machine.UART0.Configure(machine.UARTConfig{
		BaudRate: 115200,
		TX: machine.GP0,
		RX: machine.GP1,
	})
	
// sim800
	machine.UART1.Configure(machine.UARTConfig{
		BaudRate: 115200,
		TX: machine.GP4,
		RX: machine.GP5,
	})
	



	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	
}