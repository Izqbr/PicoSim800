package main

import (
	 
	"PicoSim800/buzzer"
	"PicoSim800/ds18b20"
	"machine"
	"time"
	"device/arm"
	
)

 func blinck(){
	    led.High()
		time.Sleep(time.Millisecond * 750)
		led.Low()
		time.Sleep(time.Millisecond * 750)
}

func BLINCK2(){
	led2.High()
	time.Sleep(time.Millisecond * 300)
	led2.Low()
	time.Sleep(time.Millisecond * 300)
}


func main() {
	Configure()
	machine.InitADC()
	sensor := machine.ADC{Pin: machine.GP26}
	sensor.Get()
	// timer fires 10 times per second
	arm.SetupSystemTimer(machine.CPUFrequency() / 10)
	
	dallas := ds18b20.NewDevice(machine.GP2)
	buzzer := buzzer.NewBuzzer(machine.GP15)

	
// go blinck()
// go BLINCK2()
	for {


		if Connect_status != true && comconnect == false {
		 	SIM800.Write([]byte("AT\r\n"))  //
			println("посылаем в GSM модуль AT")
		}
		if SIM800.Buffered() > 0 {
			Resp_modem()
		} // если что-то пришло от SIM800 отправляем в Raspberry для разбора
		if Serial.Buffered() > 0 {
			Resp_serial()
		}
		time.Sleep(time.Millisecond * 1000)
		buzzer.Beep(200,2)

		if Connect_status==true {
			T := dallas.GetTemp()
			println(T)
		} 
		
		
		
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