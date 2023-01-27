package buzzer

import (
	"machine"
	"time"
)

type Buzzer interface {
	Configure()
	Beep(highDuration time.Duration, amount uint8 )
}

type buzzer struct {
	pin machine.Pin
}

func NewBuzzer(pin machine.Pin) Buzzer {
	return buzzer {
		pin: pin,
	}
}

func (buzzer buzzer) Configure() {
	buzzer.pin.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
}

func (buzzer buzzer) Beep(highDuration time.Duration, amount uint8 ) {
	for i:= amount; i > 0; i-- {
		buzzer.pin.High()
		time.Sleep(highDuration)
		buzzer.pin.Low()
		time.Sleep(highDuration)
	}
}