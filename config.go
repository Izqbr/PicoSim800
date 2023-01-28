package main

import (
	"machine"
)

var (
	led = machine.LED
	led2 = machine.GP16
	
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