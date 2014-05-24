package roomba

import (
	"time"
)

/**
 *
 */
var senderFunc func(datas []uint8) = nil

/**
 *
 */
func SetSenderFunc(funcsion func(datas []uint8)) {
	senderFunc = funcsion
}

/**
 *
 */
func Start() {
	sender(START, []uint8{})
}

/**
 *
 */
func Baud(baudRate uint8) {
	sender(BAUD, []uint8{baudRate})
}

/**
 *
 */
func Safe() {
	sender(SAFE, []uint8{})
}

/**
 *
 */
func Full() {
	sender(FULL, []uint8{})
}

/**
 *
 */
func Clean() {
	sender(CLEAN, []uint8{})
}

/**
 *
 */
func Max() {
	sender(MAX, []uint8{})
}

/**
 *
 */
func Spot() {
	sender(SPOT, []uint8{})
}

/**
 *
 */
func SeekDock() {
	sender(SEEK_DOCK, []uint8{})
}

/**
 * days[0] = sun ... days[6] = sat
 */
func Schedule(days [7]time.Time) {
	var daybits uint8 = 0x00
	daybytes := []uint8{}

	for i, day := range days {
		if !day.IsZero() {
			daybits |= 0x01 << (uint8)(i)
			daybytes = append(daybytes, (uint8)(day.Hour()), (uint8)(day.Minute()))
		}
	}

	sender(SCHEDULE, append([]uint8{daybits}, daybytes...))
}

/**
 *
 */
func DisableSchedule() {
	sender(SCHEDULE, []uint8{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00})
}

/**
 *
 */
func SetDayTime(time time.Time) {
	sender(SET_DAY_TIME, []uint8{
		(uint8)(time.Weekday()),
		(uint8)(time.Hour()),
		(uint8)(time.Minute()),
	})
}

/**
 *
 */
func Power() {
	sender(POWER, []uint8{})
}

/**
 *
 */
func Drive(velocity int16, radius int16) {
	sender(DRIVE, []uint8{
		(uint8)((velocity >> 8) & 0xFF),
		(uint8)((velocity >> 0) & 0xFF),
		(uint8)((radius >> 8) & 0xFF),
		(uint8)((radius >> 0) & 0xFF),
	})
}

/**
 *
 */
func DriveDirect(rightVelocity int16, leftVelocity int16) {
	sender(DRIVE_DIRECT, []uint8{
		(uint8)((rightVelocity >> 8) & 0xFF),
		(uint8)((rightVelocity >> 0) & 0xFF),
		(uint8)((leftVelocity >> 8) & 0xFF),
		(uint8)((leftVelocity >> 0) & 0xFF),
	})
}

/**
 *
 */
func DrivePwm(rightPwm int16, leftPwm int16) {
	sender(DRIVE_PWM, []uint8{
		(uint8)((rightPwm >> 8) & 0xFF),
		(uint8)((rightPwm >> 0) & 0xFF),
		(uint8)((leftPwm >> 8) & 0xFF),
		(uint8)((leftPwm >> 0) & 0xFF),
	})
}

/**
 *
 */
func Motors(mainBrushDirection bool, sideBrushClockwise bool, mainBrush bool, vacuum bool, sideBrush bool) {
	var motorbits uint8 = 0x00

	args := []bool{sideBrush, vacuum, mainBrush, sideBrushClockwise, mainBrushDirection}
	for i, v := range args {
		if v {
			motorbits |= 0x01 << (uint)(i)
		}
	}

	sender(MOTORS, []uint8{motorbits})
}

/**
 *
 */
func PwmMotors(mainBrushPwm int8, sideBrushPwm int8, vacuumPwm uint8) {
	sender(PWM_MOTORS, []uint8{
		(uint8)(mainBrushPwm),
		(uint8)(sideBrushPwm),
		(uint8)(vacuumPwm)})
}

/**
 *
 */
func Leds(checkRobot bool, dock bool, spot bool, debris bool, cleanPowerColor uint8, cleanPowerIntensity uint8) {
	var ledbits uint8 = 0x00

	args := []bool{debris, spot, dock, checkRobot}
	for i, v := range args {
		if v {
			ledbits |= 0x01 << (uint)(i)
		}
	}

	sender(LEDS, []uint8{ledbits, cleanPowerColor, cleanPowerIntensity})
}

/**
 *
 */
func SchedulingLeds(weekday uint8, schedule bool, clock bool, am bool, pm bool, colon bool) {
	var ledbits uint8 = 0x00

	args := []bool{colon, pm, am, clock, schedule}
	for i, v := range args {
		if v {
			ledbits |= 0x01 << (uint)(i)
		}
	}

	sender(SCHEDULING_LEDS, []uint8{weekday, ledbits})
}

/**
 *
 */
func DigitLedsRaw(digit [4]uint8) {
	sender(DIGIT_LEDS_RAW, digit[:4])
}

/**
 *
 */
func DigitLedsAscii(message string) {
	sender(DIGIT_LEDS_ASCII, []uint8(message[:4]))
}

/**
 *
 */
func Buttons(clock bool, schedule bool, day bool, hour bool, dock bool, spot bool, clean bool) {
	var buttonbits uint8 = 0x00

	args := []bool{clean, spot, dock, hour, day, schedule, clock}
	for i, v := range args {
		if v {
			buttonbits |= 0x01 << (uint)(i)
		}
	}

	sender(BUTTONS, []uint8{buttonbits})
}

/**
 *
 */
func Song(songNumber uint8, notes []Note) {
	datas := []uint8{songNumber, (uint8)(len(notes))}

	for _, note := range notes {
		datas = append(datas, note.Number, note.Duration)
	}

	sender(SONG, datas)
}

/**
 *
 */
func Play(songNumber uint8) {
	sender(PLAY, []uint8{songNumber})
}

/**
 *
 */
func Sensors(packetId PacketId) {
	sender(SENSORS, []uint8{(uint8)(packetId)})
}

/**
 *
 */
func QueryList(packetIds []PacketId) {
	datas := []uint8{(uint8)(len(packetIds))}

	for _, packetId := range packetIds {
		datas = append(datas, (uint8)(packetId))
	}

	sender(QUERY_LIST, datas)
}

/**
 *
 */
func Stream(packetIds []PacketId) {
	datas := []uint8{(uint8)(len(packetIds))}

	for _, packetId := range packetIds {
		datas = append(datas, (uint8)(packetId))
	}

	sender(STREAM, datas)
}

/**
 *
 */
func PauseStream() {
	sender(PAUSE_RESUME_STREAM, []uint8{0x00})
}

/**
 *
 */
func ResumeStream() {
	sender(PAUSE_RESUME_STREAM, []uint8{0x01})
}

/**
 *
 */
func receive(data uint8) {

}

/**
 *
 */
func sender(command uint8, datas []uint8) {
	if senderFunc != nil {
		senderFunc(append([]uint8{command}, datas...))
	}
}
