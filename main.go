package main

import (
	"PicoSim800/buzzer"
	"PicoSim800/ds18b20"
	"machine"
	"time"
	"device/arm"
	
)

var (
	Serial      = machine.UART0
	SIM800      = machine.UART1
	MQTT_SERVER = "srv.clusterfly.ru"
	PORT        = "9991"
	AT          = ""
	comconnect  = false
	timerCh = make(chan struct{}, 1)
)

func main() {
	Configure()

	// timer fires 10 times per second
	arm.SetupSystemTimer(machine.CPUFrequency() / 10)
	
	dallas := ds18b20.NewDevice(machine.GP2)
	buzzer := buzzer.NewBuzzer(machine.GP15)
	
	for {


		if Connect_status != true && comconnect == false {
		 	SIM800.Write([]byte("AT\r\n"))  //посылаем в GSM модуль
		}
		if SIM800.Buffered() > 0 {
			Resp_modem()
		} // если что-то пришло от SIM800 отправляем в Raspberry для разбора
		if Serial.Buffered() > 0 {
			Resp_serial()
		}
		time.Sleep(time.Millisecond * 1000)
		buzzer.Beep(200,2)
		T := dallas.GetTemp()
		println(T)
		
		led.High()
		time.Sleep(time.Millisecond * 750)
		led.Low()
		time.Sleep(time.Millisecond * 750)
	}

}


//export SysTick_Handler
func timer_isr() {
	select {
	case timerCh <- struct{}{}:
	default:
		// The consumer is running behind.
	}
}