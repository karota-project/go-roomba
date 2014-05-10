package roomba

const (
	/**
	 * Starts the SCI. The Start command must be sent before any other SCI commands. This command puts the SCI in passive mode.
	 */
	START uint8 = 128

	/**
	  * Sets the baud rate in bits per second (bps) at which SCI commands and data are sent according to the baud code sent in the data byte. The default baud rate at power up is 57600 bps. (See Serial Port Settings, above.) Once the baud rate is changed, it will persist until Roomba is power cycled by removing the battery (or until the battery voltage falls below the minimum required for processor operation). You must wait 100ms after sending this command before sending additional commands
	  at the new baud rate. The SCI must be in passive, safe, or full mode to accept this command. This command puts the SCI in passive mode.
	*/
	BAUD uint8 = 129

	/**
	 * Enables user control of Roomba. This command must be sent after the start command and before any control commands are sent to the SCI. The SCI must be in passive mode to accept this command. This command puts the SCI in safe mode.
	 */
	CONTROL uint8 = 130

	/**
	 * This command puts the SCI in safe mode. The SCI must be in full mode to accept this command.
	 */
	SAFE uint8 = 131

	/**
	 * Enables unrestricted control of Roomba through the SCI and turns off the safety features. The SCI must be in safe mode to accept this command. This command puts the SCI in full mode.
	 */
	FULL uint8 = 132

	/**
	 * Puts Roomba to sleep, the same as a normal “power” button press. The Device Detect line must be held low for 500 ms to wake up Roomba from sleep. The SCI must be in safe or full mode to accept this command. This command puts the SCI in passive mode.
	 */
	POWER uint8 = 133

	/**
	 * Starts a spot cleaning cycle, the same as a normal “spot” button press. The SCI must be in safe or full mode to accept this command. This command puts the SCI in passive mode.
	 */
	SPOT uint8 = 134

	/**
	 * Starts a normal cleaning cycle, the same as a normal “clean” button press. The SCI must be in safe or full mode to accept this command. This command puts the SCI in passive mode.
	 */
	CLEAN uint8 = 135

	/**
	 * Starts a maximum time cleaning cycle, the same as a normal “max” button press. The SCI must be in safe or full mode to accept this command. This command puts the SCI in passive mode.
	 */
	MAX uint8 = 136

	/**
	 * Controls Roomba’s drive wheels. The command takes four data bytes, which are interpreted as two 16 bit signed values using twos-complement. The first two bytes specify the average velocity of the drive wheels in millimeters per second (mm/s), with the high byte sent first. The next two bytes specify the radius, in millimeters, at which Roomba should turn. The longer radii make Roomba drive straighter; shorter radii make it turn more. A Drive command with a positive velocity and a positive radius will make Roomba drive forward while turning toward the left. A negative radius will make it turn toward the right. Special cases for the radius make Roomba turn in place or drive straight, as specified below. The SCI must be in safe or full mode to accept this command. This command does change the mode.
	 */
	DRIVE uint8 = 137

	/**
	 * Controls Roomba’s cleaning motors. The state of each motor is specified by one bit in the data byte. The SCI must be in safe or full mode to accept this command. This command does not change the mode.
	 */
	MOTORS uint8 = 138

	/**
	  * Controls Roomba’s LEDs. The state of each of the spot, clean, max, and dirt detect LEDs is specified by one bit in the first data byte. The color of the status LED is specified by two bits in the first data byte. The power LED is specified by two data bytes, one for the color and one for the intensity. The SCI must be in safe
	  or full mode to accept this command. This command does not change the mode.
	*/
	LEDS uint8 = 139

	/**
	  * Specifies a song to the SCI to be played later. Each song is associated with a song number which the Play command uses
	  to select the song to play. Users can specify up to 16 songs
	  with up to 16 notes per song. Each note is specified by a note number using MIDI note definitions and a duration specified
	  in fractions of a second. The number of data bytes varies depending on the length of the song specified. A one note song is specified by four data bytes. For each additional note, two data bytes must be added. The SCI must be in passive, safe, or full mode to accept this command. This command does not change the mode.
	*/
	SONG uint8 = 140

	/**
	 * Plays one of 16 songs, as specified by an earlier Song command. If the requested song has not been specified yet, the Play command does nothing. The SCI must be in safe or full mode to accept this command. This command does not change the mode.
	 */
	PLAY uint8 = 141

	/**
	 * Requests the SCI to send a packet of sensor data bytes. The user can select one of four different sensor packets. The sensor data packets are explained in more detail in the next section. The SCI must be in passive, safe, or full mode to accept this command. This command does not change the mode.
	 */
	SENSORS uint8 = 142

	/**
	  * Turns on force-seeking-dock mode, which causes the robot
	  to immediately attempt to dock during its cleaning cycle if it encounters the docking beams from the Home Base. (Note, however, that if the robot was not active in a clean, spot or max cycle it will not attempt to execute the docking.) Normally the robot attempts to dock only if the cleaning cycle has completed or the battery is nearing depletion. This command can be sent anytime, but the mode will be cancelled if the robot turns off, begins charging, or is commanded into SCI safe or full modes.
	*/
	FORCE_SEEKING_DOCK uint8 = 143
)
