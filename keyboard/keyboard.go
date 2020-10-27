package keyboard

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
)

// KeyLogger wrapper around file descriptior
type KeyLogger struct {
	fd *os.File
}

type devices []string

func (d *devices) hasDevice(str string) bool {
	for _, device := range *d {
		if strings.Contains(str, device) {
			return true
		}
	}

	return false
}

// use lowercase names for devices, as we turn the device input name to lower case
var restrictedDevices = devices{"mouse"}
var allowedDevices = devices{"keyboard", "logitech mx keys"}

// New creates a new keylogger for a device path
func New(devPath string) (*KeyLogger, error) {
	k := &KeyLogger{}
	if !k.IsRoot() {
		return nil, errors.New("Must be run as root")
	}
	fd, err := os.Open(devPath)
	k.fd = fd
	return k, err
}

// FindKeyboardDevice by going through each device registered on OS
// Mostly it will contain keyword - keyboard
// Returns the file path which contains events
func FindKeyboardDevice() string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(path, i))
		if err != nil {
			fmt.Println(err)
		}

		deviceName := strings.ToLower(string(buff))

		if restrictedDevices.hasDevice(deviceName) {
			continue
		} else if allowedDevices.hasDevice(deviceName) {
			return fmt.Sprintf(resolved, i)
		}
	}

	return ""
}

// Like FindKeyboardDevice, but finds all devices which contain keyword 'keyboard'
// Returns an array of file paths which contain keyboard events
func FindAllKeyboardDevices() []string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	valid := make([]string, 0)

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(path, i))

		// prevent from checking non-existant files
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		deviceName := strings.ToLower(string(buff))

		if restrictedDevices.hasDevice(deviceName) {
			continue
		} else if allowedDevices.hasDevice(deviceName) {
			valid = append(valid, fmt.Sprintf(resolved, i))
		}
	}
	return valid
}

// IsRoot checks if the process is run with root permission
func (k *KeyLogger) IsRoot() bool {
	return syscall.Getuid() == 0 && syscall.Geteuid() == 0
}

// Read from file descriptor
// Blocking call, returns channel
// Make sure to close channel when finish
func (k *KeyLogger) Read() chan InputEvent {
	emit := make(chan InputEvent)
	go func(emit chan InputEvent) {
		for {
			keyEvent, err := k.read()
			if err != nil {
				fmt.Println("keyRead:", err)
				close(emit)
				break
			}

			if keyEvent != nil {
				emit <- *keyEvent
			}
		}
		//shutdown flag
		shutdown := InputEvent{Code: SHUTDOWN}
		emit <- shutdown
	}(emit)
	return emit
}

// read from file description and parse binary into go struct
func (k *KeyLogger) read() (*InputEvent, error) {
	buffer := make([]byte, eventsize)
	n, err := k.fd.Read(buffer)
	if err != nil {
		return nil, err
	}
	// no input, dont send error
	if n <= 0 {
		return nil, nil
	}
	return k.eventFromBuffer(buffer)
}

// eventFromBuffer parser bytes into InputEvent struct
func (k *KeyLogger) eventFromBuffer(buffer []byte) (*InputEvent, error) {
	event := &InputEvent{}
	err := binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, event)
	return event, err
}

// Close file descriptor
func (k *KeyLogger) Close() error {
	if k.fd == nil {
		return nil
	}
	return k.fd.Close()
}
