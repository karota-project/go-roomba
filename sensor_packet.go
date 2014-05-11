package roomba

const (
	BUMPS_WHEELDROPS = iota
	WALL
	CLIFF_LEFT
	CLIFF_FRONT_LEFT
	CLIFF_FRONT_RIGHT
	CLIFF_RIGHT
	VIRTUAL_WALL
	MOTOR_OVERCURRENTS
	DIRT_DETECT_LEFT
	DIRT_DETECT_RIGHT

	REMOTE_OPCODE
	BUTTONS
	DISTANCE
	ANGLE

	CHARGING_STATE
	VOLTAGE
	CURRENT
	TEMPERATURE
	CHARGE
	CAPACITY
)

type BumpsWheeldrops struct {
  WheeldropCaster bool
  WheeldropLeft bool
  WheeldropRight bool
  BumpLeft bool
  BumpRight bool
}

type MotorOvercurrents struct {
  DriveLeft bool
  DriveRight bool
  MainBrush bool
  Vacuum bool
  SideBrush bool
}

type Buttons struct {
  Power bool
  Spot bool
  Clean bool
  Max bool
}

// Charging State
const (
  NOT_CHARGING = iota
  CHARGING_RECOVERY
  CHARGING
  TRICKLE_CHARGING
  WAITING
  CHARGING_ERROR
)

type SendorPacketGroup1 struct {
	BUMPS_WHEELDROPS BumpsWheeldrops
	WALL bool
	CLIFF_LEFT bool
	CLIFF_FRONT_LEFT bool
	CLIFF_FRONT_RIGHT bool
	CLIFF_RIGHT bool
	VIRTUAL_WALL bool
	MOTOR_OVERCURRENTS MotorOvercurrents
	DIRT_DETECT_LEFT uint8
	DIRT_DETECT_RIGHT uint8
}

type SendorPacketGroup2 struct {
	REMOTE_OPCODE uint8
	BUTTONS Buttons
	DISTANCE int16
	ANGLE int16
}

type SendorPacketGroup3 struct {
	CHARGING_STATE uint8
	VOLTAGE uint16
	CURRENT int16
	TEMPERATURE int8
	CHARGE uint16
	CAPACITY uint16
}

