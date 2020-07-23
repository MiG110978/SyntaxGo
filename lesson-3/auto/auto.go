package main

import (
	"fmt"
)

type Auto struct {
	Marka        string
	Year         int
	VolumeTrunk  int
	EngineOnOff  string
	WindowOpenYN string
	TrunkPercent int
}

func EngineСondition(car Auto) {
	fmt.Printf("Желаете ли Вы изменить состояние двиагателя автомобиля %+v \n ? Y/N \n", car.Marka)
	var x string
	fmt.Scanln(&x)
	if x == "Y" {
		if car.EngineOnOff == "On" {
			car.EngineOnOff = "Off"
		} else {
			car.EngineOnOff = "On"
		}
	}
	fmt.Printf("Новые данные автомобиля \n%+v \n", car)
}

func WindowСondition(car Auto) {
	fmt.Printf("Желаете ли Вы изменить состояние окон автомобиля %+v \n ? Y/N \n", car.Marka)
	var x string
	fmt.Scanln(&x)
	if x == "Y" {
		if car.WindowOpenYN == "On" {
			car.WindowOpenYN = "Off"
		} else {
			car.WindowOpenYN = "On"
		}
	}
	fmt.Printf("Новые данные автомобиля \n%+v \n", car)
}

func main() {
	ford := Auto{
		Marka:        "Ford",
		Year:         2015,
		VolumeTrunk:  800,
		EngineOnOff:  "Off",
		WindowOpenYN: "No",
		TrunkPercent: 60,
	}
	renault := Auto{
		Marka:        "Renault",
		Year:         2017,
		VolumeTrunk:  650,
		EngineOnOff:  "On",
		WindowOpenYN: "No",
		TrunkPercent: 65,
	}
	kia := Auto{
		Marka:        "Kia",
		Year:         2019,
		VolumeTrunk:  720,
		EngineOnOff:  "Off",
		WindowOpenYN: "No",
		TrunkPercent: 70,
	}
	fmt.Printf("%+v \n", ford)
	fmt.Printf("%+v \n", renault)
	fmt.Printf("%+v \n", kia)

	EngineСondition(ford)
	EngineСondition(renault)
	EngineСondition(kia)
	WindowСondition(ford)
	WindowСondition(renault)
	WindowСondition(kia)
}
