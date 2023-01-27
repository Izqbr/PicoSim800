package ds18b20

import (
	"fmt"
	"machine"
	
	"time"
)

type Device interface {
	Init() bool
	SendCommand(command uint8)
	SendBit(bit uint8)
	GetTemp() string
	Readbyte() byte
}

type dallas struct {
	load machine.Pin
}

func NewDevice(load machine.Pin) Device {
	return &dallas{
		load: load,
	}
}

func (dallas *dallas) Init() bool {
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

func (dallas *dallas) SendCommand(command uint8) {
	dallas.load.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for i := 0; i < 8; i++ {
		dallas.SendBit(command >> i & 1)
		time.Sleep(time.Microsecond * 5)
	}
}
func (dallas *dallas) SendBit(bit uint8) {
	dallas.load.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dallas.load.Set(false)
	if bit == 1 {
		time.Sleep(time.Microsecond * 2)
	} else {
		time.Sleep(time.Microsecond * 65)
	}
	dallas.load.Set(true)
	if bit == 1 {
		time.Sleep(time.Microsecond * 65)
	} else {
		time.Sleep(time.Microsecond * 2)
	}
}

func (dallas *dallas) Readbyte() byte {
	var data byte
	for i := 0; i < 8; i++ {
		data += dallas.ReadBit() << i
	}

	return data

}

func (dallas *dallas) ReadBit() byte {

	var bit byte
	dallas.load.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dallas.load.Set(false)
	time.Sleep(time.Microsecond * 2)
	dallas.load.Set(true)
	time.Sleep(time.Microsecond * 13)
	dallas.load.Configure(machine.PinConfig{Mode: machine.PinInput})

	if dallas.load.Get() == false {
		bit = 0
	} else {
		bit = 1
	}
	time.Sleep(time.Microsecond * 45)
	return bit //вернем результат
}

func (dallas *dallas) GetTemp() string {

	dallas.Init()
	dallas.SendCommand(SKIP_ROM)
	dallas.SendCommand(CONVERT_T)
	time.Sleep(time.Millisecond * 750)
	dallas.Init()
	dallas.SendCommand(SKIP_ROM)
	dallas.SendCommand(READ_SCRETCHPAD)
	sign := ""
	lbt := uint16(dallas.Readbyte())
	dr:= (lbt & 0b1111)/16
	
	hbt := uint16(dallas.Readbyte())
	if hbt&128 == 0 {
		sign = "+"
	} else {
		sign = "-"
	}
	temp := hbt<<8 | lbt
	
	temperature := temp >> 4
	//temperature := 0x19
	return sign + fmt.Sprintf("%d",temperature)+"."+ fmt.Sprintf("%1d",dr)
	

}
