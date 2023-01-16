package main

import "time"
import "machine"


var status bool 
var dt = machine.GP10
var bit uint8	
var tt uint16 = 0
var bt uint8

func dt_Reset() bool {
	dt.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dt.Set(false) //низкий уровень
	time.Sleep(time.Microsecond * 485)
	dt.Set(true) // высокий уровень
	time.Sleep(time.Microsecond * 65)
	dt.Configure(machine.PinConfig{Mode: machine.PinInput})
	status = dt.Get()
	time.Sleep(time.Microsecond * 500)
	return status
}

// func dt_init(mode uint8) uint8 {
// 	if dt_Reset() == true {
// 		return 1	
// 	}
// 	if mode == 0xCC{
// 		dt_WriteByte(0xCC)
// 		dt_WriteByte(0x4E)
// 		dt_WriteByte(0x64)
// 		dt_WriteByte(0x9E)
// 		dt_WriteByte(0x7F) //resolution 12bit
// 	}	
// }


func dt_ReadByte() uint8 {
	var data uint8
	for i:=0;i<=7;i++ {
		data += dt_ReadBit()<<i
	}
	return data
}

func dt_ReadBit() uint8 {
	dt.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dt.Set(false)
	time.Sleep(time.Microsecond * 2)
	dt.Set(true)
	time.Sleep(time.Microsecond * 13)
	dt.Configure(machine.PinConfig{Mode: machine.PinInput})

	if dt.Get() == false{
		bit = 0
	}else {
		bit = 1
	}
	time.Sleep(time.Microsecond * 45)
	return bit //вернем результат
}

//-----------------------------------------------
func dt_WriteByte(bt uint8 ) {
		for i:=0;i<8;i++ {
			dt_WriteBit(bt >> i & 1)
			time.Sleep(time.Microsecond * 5)
	}
}

//-----------------------------------------------
func dt_WriteBit(bit uint8) {
	dt.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dt.Set(false)
	if bit==1 {
		time.Sleep(time.Microsecond * 2)
	} else {time.Sleep(time.Microsecond * 65) } 
	dt.Set(true)	  
	if bit == 1 {
		time.Sleep(time.Microsecond * 65)
	} else {time.Sleep(time.Microsecond * 2) }	  
}

func dt_MeasureTemperCmd(mode uint8){
	dt_Reset()
	if mode == 0xCC {
		dt_WriteByte(0xCC)
	}
	dt_WriteByte(0x44)
}



func dt_GetSign(dt uint16)uint8{
	if dt&(1<<11) == 1{
		return 1
	}else {
		return 0
	}
}

// func dt_Convert(dt uint16) float32 {
// 	var t float32
// 	t = (float32((dt&0x07FF)>>4))
// 	t += (float32(dt&0x000F) / 16.0)

// }

func dt_check() uint16 {
	//unsigned char bt;//переменная для считывания байта
	//unsigned int tt=0;
	if status==true{  //если устройство нашлось
	 
		dt_WriteByte(0xCC) //пропустить идентификацию, тк у нас только одно устройство на шине
		dt_WriteByte(0x44) //измеряем температуру
		time.Sleep(time.Millisecond * 750) //в 12битном режиме преобразования - 750 милисекунд
		dt_Reset() //снова используем  те же манипуляции с шиной что и при проверке ее присутствия
		dt_WriteByte(0xCC) //пропустить идентификацию, тк у нас только одно устройство на шине
		dt_WriteByte(0xBE) //даем команду на чтение данных с устройства
		//bt = dt_ReadByte() //читаем младший бит
		//tt = tt<<8| dt_ReadByte() //читаем старший бит MS
		//tt = (tt<<8)|bt    //сдвигаем старший влево, младший пишем на его место, тем самым получаем общий результат
	}
	return tt
}

func converttemp (tt uint8) uint8 {
 t := tt>>4;//сдвиг и отсечение части старшего байта
	return t
}