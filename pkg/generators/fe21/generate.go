package fe21s

import (
	"errors"

	"github.com/pojntfx/papilio/pkg/utils"
	"golang.org/x/exp/utf8string"
)

const (
	// See http://www.linux-usb.org/usb.ids
	DefaultIdVendor                = uint16(0x1a40)
	DefaultIdProduct               = uint16(0201)
	DefaultBcdDevice               = uint16(0x0001)
	DefaultNumberOfDownstreamPorts = uint8(7)
)

var (
	ErrSerialToLong   = errors.New("serial number is too long")
	ErrSerialNotASCII = errors.New("serial number is not valid ASCII")
)

// See https://yourbasic.org/golang/bitmask-flag-set-clear/
type portMask byte

const (
	_ portMask = 1 << iota // Bit 0 is reserved and should be 0
	port1
	port2
	port3
	port4
	port5
	port6
	port7
)

func setPortRemovable(b, flag portMask) portMask { return b | flag }

// | Address     | Contents                | Note                                                                                                              |
// | ----------- | ----------------------- | ----------------------------------------------------------------------------------------------------------------- |
// | `0x00`      | `0x40`                  | Constant, low byte of check code                                                                                  |
// | `0x01`      | `0x1A`                  | Constant, high byte of check code                                                                                 |
// | `0x02`      | Vendo ID (Low)          | Low byte of vendor ID, `idVendor` field of standard device descriptor                                             |
// | `0x03`      | Vendor ID (High)        | High byte of vendor ID, `idVendor` field of standard device descriptor                                            |
// | `0x04`      | Product ID (Low)        | Low byte of product ID, `idProduct` field of standard device descriptor                                           |
// | `0x05`      | Product ID (High)       | High Byte of product ID, `idProduct` field of standard device descriptor                                          |
// | `0x06`      | Device Release (Low)    | Low byte of device release number, must be binary coded decimal, `bcdDevice` field of standard device descriptor  |
// | `0x07`      | Device Release (High)   | High byte of device release number, must be binary coded decimal, `bcdDevice` field of standard device descriptor |
// | `0x08-0x17` | Device Serial           | Number device's serial number - the contents of string descriptor describing the device's serial number.          |
// | `0x18`      | Length of Serial Number | Length of effective device serial number stored in `0x08-0x17`.                                                   |
// | `0x19`      | Filling                 | `0x00`                                                                                                            |
// | `0x1A`      | Port Number             | Number of downstream ports, `bNbrPorts` field of hub descriptor.                                                  |
// | `0x1B`      | Filling                 | `0x00`                                                                                                            |
// | `0x1C`      | Device Removable Device | See below                                                                                                         |
// | `0x1D`      | Filling                 | `0x00`                                                                                                            |
// | `0x1E`      | Device Attributes       | See below                                                                                                         |
// | `0x1F`      | Check Sum               | The 8-bit sum of all values from `0x00` to `0x1E`                                                                 |
//
// **Device Removable Device**: Removable field of hub descriptor - indicates if a port has a removable device attached. If bit `N` is set to 1, then the device on downstream facing port `N` is non-removable. Otherwise, it is removable. Bit 0 is reserved and should be 0.
//
// **Device Attributes**:
//     **Bit 0**: Port indicators support, bit 7 of `wHubCharacteristics` field of hub descriptor
//         0: Port indicators are not supported on its downstream facing ports and `PORT_INDICATOR` request has no effect.
//         1: Port indicators are supported on its downstream facing ports and `PORT_INDICATOR` request controls the indicators.
//     **Bit 1**: Identifies a compound device, bit 2 of `wHubCharacteristics` field of hub descriptor
//         0: Hub is not part of a compound device.
//         1: Hub is part of a compound device.
//     **Bit 2**: Maximum current requirements of the hub controller electronics, `bHubContrCurrent` field of hub descriptor
//         0: 200mA.
//         1: 500mA.
//         Bit 3 to 7, reserved, must be 0s.

func GenerateEEPROM(
	idVendor uint16, // i.e. 0x046d for Logitech
	idProduct uint16, // i.e. 0x082d for HD Pro Webcam C920
	bcdDevice uint16, // i.e. 0x0001 for release 1
	numberOfDownstreamPorts uint8, // i.e. 4 for 4 ports
	serial string, // ASCII serial number, max. 15 chars (i.e. `sadfasdfasdi3ds`)
	portsWithRemovableDevices [7]bool, // Which ports have removable devices (true = removable, false = non-removable)
) ([]byte, error) {
	buf := make([]byte, 0x1F+1) // Filling is `0x00`

	buf[0x00] = 0x40 // Constant, low byte of check code
	buf[0x01] = 0x1A // Constant, high byte of check code

	buf[0x02], buf[0x03] = utils.GetLowAndHighByte(idVendor) // Low and high byte of vendor ID, `idVendor` field of standard device descriptor

	buf[0x04], buf[0x05] = utils.GetLowAndHighByte(idProduct) // Low and high byte of product ID, `idProduct` field of standard device descriptor

	buf[0x06], buf[0x07] = utils.GetLowAndHighByte(utils.GetBCD(bcdDevice)) // Low and high byte of device release number, must be binary coded decimal, `bcdDevice` field of standard device descriptor

	serialBuf := utf8string.NewString(serial)
	if !serialBuf.IsASCII() {
		return []byte{}, ErrSerialNotASCII
	}
	serial = serialBuf.String()

	if len(serial) > 15 { // 0x17-0x08
		return []byte{}, ErrSerialToLong
	}

	for i, c := range serial {
		buf[0x08+i] = byte(c) // Number device's serial number - the contents of string descriptor describing the device's serial number.
	}

	buf[0x18] = byte(len(serial)) // Length of effective device serial number stored in `0x08-0x17`.

	var removableDevices portMask
	for i, port := range portsWithRemovableDevices {
		if port {
			// hic sunt dracones
			switch i {
			case 0:
				removableDevices = setPortRemovable(removableDevices, port1)
			case 1:
				removableDevices = setPortRemovable(removableDevices, port2)
			case 2:
				removableDevices = setPortRemovable(removableDevices, port3)
			case 3:
				removableDevices = setPortRemovable(removableDevices, port4)
			case 4:
				removableDevices = setPortRemovable(removableDevices, port5)
			case 5:
				removableDevices = setPortRemovable(removableDevices, port6)
			case 6:
				removableDevices = setPortRemovable(removableDevices, port7)
			}
		}
	}
	buf[0x1C] = byte(removableDevices) // Removable field of hub descriptor - indicates if a port has a removable device attached

	// TODO: Add Device Attributes

	buf[0x1A] = numberOfDownstreamPorts // Number of downstream ports, `bNbrPorts` field of hub descriptor.

	buf[0x1F] = utils.GetChecksum(buf[:0x1E]) // The 8-bit sum of all values from `0x00` to `0x1E`.

	return buf, nil
}
