package ds18b20

import (
	
	"machine"
	 "time"
	
)

type Device interface{
	
	Init() bool
	SendCommand(command uint8)
	SendBit(bit uint8)
	
}

type dallas struct {
	load machine.Pin
}

func NewDevice(load machine.Pin) Device {
	return &dallas {
		load: load,
	}
}


func (dallas *dallas) Init() bool{
	dallas.load.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dallas.load.Set(false) //низкий уровень
	time.Sleep(time.Microsecond * 485)
	dallas.load.Set(true) // высокий уровень
	time.Sleep(time.Microsecond * 65)
	dallas.load.Configure(machine.PinConfig{Mode: machine.PinInput})
	status := dallas.load.Get()
	time.Sleep(time.Microsecond * 500)
	return status
}

func (dallas *dallas)SendCommand(command uint8) {
	for i:=0;i<8;i++ {
		dallas.SendBit(command >> i & 1)
		time.Sleep(time.Microsecond * 5)
} 
	}
func(dallas *dallas) SendBit(bit uint8) {
	dallas.load.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dallas.load.Set(false)
	if bit==1 {
		time.Sleep(time.Microsecond * 2)
	} else {time.Sleep(time.Microsecond * 65) } 
	dallas.load.Set(true)	  
	if bit == 1 {
		time.Sleep(time.Microsecond * 65)
	} else {time.Sleep(time.Microsecond * 2) }	
}


