package roomba

import (
	"fmt"
	//"log"
)

func Start() {
	sender(START, []uint8{})
}

func Baud() {
	sender(BAUD, []uint8{})
}

func Control() {
	sender(CONTROL, []uint8{})
}

func Safe() {
	sender(SAFE, []uint8{})
}

func Full() {
	sender(FULL, []uint8{})
}

func Power() {
	sender(POWER, []uint8{})
}

func Spot() {
	sender(SPOT, []uint8{})
}

func Clean() {
	sender(CLEAN, []uint8{})
}

func Max() {
	sender(MAX, []uint8{})
}

func Drive() {
	sender(DRIVE, []uint8{})
}

func Motors() {
	sender(MOTORS, []uint8{})
}

func Leds() {
	sender(LEDS, []uint8{})
}

func Song() {
	sender(SONG, []uint8{})
}

func Play() {
	sender(PLAY, []uint8{})
}

func Sensors() {
	sender(SENSORS, []uint8{})
}

func ForceSeekingDock() {
	sender(FORCE_SEEKING_DOCK, []uint8{})
}

func sender(command uint8, datas []uint8) {
	fmt.Print("command: ")
	fmt.Println(command)
	fmt.Println("datas: ")
	fmt.Println(datas)
	fmt.Println("")
}
