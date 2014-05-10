package roomba

import ()

var senderFunc func(command uint8, datas []uint8) = nil

func SetSenderFunc(_senderFunc func(command uint8, datas []uint8)) {
	senderFunc = _senderFunc
}

func Start() {
	sender(START, []uint8{})
}

func Baud(baudRate uint8) {
	sender(BAUD, []uint8{baudRate})
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

func Drive(velocity int16, radius int16) {
	sender(DRIVE, []uint8{
		(uint8)((velocity >> 8) & 0xFF),
		(uint8)((velocity >> 0) & 0xFF),
		(uint8)((radius >> 8) & 0xFF),
		(uint8)((radius >> 0) & 0xFF)})
}

func Motors(mainBrush bool, vacuum bool, sideBrush bool) {
	var bits uint8 = 0

	if mainBrush {
		bits |= 0x01 << 2
	}

	if vacuum {
		bits |= 0x01 << 1
	}

	if sideBrush {
		bits |= 0x01 << 0
	}

	sender(MOTORS, []uint8{bits})
}

func Leds(statusRed bool, statusGreen bool, spot bool, clean bool, max bool, dirtDetect bool, powerColor uint8, powerIntensity uint8) {
	var ledBits uint8 = 0

	if statusRed {
		ledBits |= 0x01 << 5
	}

	if statusGreen {
		ledBits |= 0x01 << 4
	}

	if spot {
		ledBits |= 0x01 << 3
	}

	if clean {
		ledBits |= 0x01 << 2
	}

	if max {
		ledBits |= 0x01 << 1
	}

	if dirtDetect {
		ledBits |= 0x01 << 0
	}

	sender(LEDS, []uint8{ledBits, powerColor, powerIntensity})
}

func Song(songNumber uint8, songs []SongData) {
	datas := []uint8{songNumber, (uint8)(len(songs))}

	for _, v := range songs {
		datas = append(datas, v.Note, v.Duration)
	}

	sender(SONG, datas)
}

func Play(songNumber uint8) {
	sender(PLAY, []uint8{songNumber})
}

func Sensors(packetCode uint8) {
	sender(SENSORS, []uint8{packetCode})
}

func ForceSeekingDock() {
	sender(FORCE_SEEKING_DOCK, []uint8{})
}

func sender(command uint8, datas []uint8) {
	if senderFunc != nil {
		senderFunc(command, datas)
	}
}
