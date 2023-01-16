package main

import "machine"


var (
	led = machine.LED
	led2 = machine.GP7
	btn = machine.GP2
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
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	btn.SetInterrupt(machine.PinFalling|machine.PinRising,
		func(p machine.Pin) {
			led2.Set(!p.Get())
			Serial.Write([]byte("Button is pressed \n\r"))
		})
}