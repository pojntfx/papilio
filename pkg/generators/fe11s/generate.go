package fe11s

import (
	"github.com/pojntfx/papilio/pkg/utils"
)

const (
	// See http://www.linux-usb.org/usb.ids
	DefaultIdVendor                = uint16(0x1a40)
	DefaultIdProduct               = uint16(0x0101)
	DefaultBcdDevice               = uint16(0x0001)
	DefaultNumberOfDownstreamPorts = uint8(4)
)

// | Address     | Contents              | Note                                                                                                             |
// | ----------- | --------------------- | ---------------------------------------------------------------------------------------------------------------- |
// | `0x00`      | `0x40`                | Constant, low byte of check code                                                                                 |
// | `0x01`      | `0x1A`                | Constant, high byte of check code                                                                                |
// | `0x02`      | Vendor ID (Low)       | Low byte of vendor ID, `idVendor` field of standard device descriptor                                            |
// | `0x03`      | Vendor ID (High)      | High byte of vendor ID                                                                                           |
// | `0x04`      | Product ID (Low)      | Low byte of product ID, `idProduct` field of standard device descriptor                                          |
// | `0x05`      | Product ID (High)     | High byte of product ID                                                                                          |
// | `0x06`      | Device Release (Low)  | Low byte of device release number, must be binary coded decimal, `bcdDevice` field of standard device descriptor |
// | `0x07`      | Device Release (High) | High byte of device release number, must be binary coded decimal                                                 |
// | `0x08-0x19` | Filling All           | `0x00`                                                                                                           |
// | `0x1A`      | Port Number           | Number of downstream ports, `bNbrPorts` field of hub descriptor.                                                 |
// | `0x1B-0x1E` | Filling All           | `0x00`                                                                                                           |
// | `0x1F`      | Check Sum             | The 8-bit sum of all values from `0x00` to `0x1E`.                                                               |

func GenerateEEPROM(
	idVendor uint16, // i.e. 0x046d for Logitech
	idProduct uint16, // i.e. 0x082d for HD Pro Webcam C920
	bcdDevice uint16, // i.e. 0x0001 for release 1
	numberOfDownstreamPorts uint8, // i.e. 4 for 4 ports
) ([]byte, error) {
	buf := make([]byte, 0x1F+1) // Filling is `0x00`

	buf[0x00] = 0x40 // Constant, low byte of check code
	buf[0x01] = 0x1A // Constant, high byte of check code

	buf[0x02], buf[0x03] = utils.GetLowAndHighByte(idVendor) // Low and high byte of vendor ID, `idVendor` field of standard device descriptor

	buf[0x04], buf[0x05] = utils.GetLowAndHighByte(idProduct) // Low and high byte of product ID, `idProduct` field of standard device descriptor

	buf[0x06], buf[0x07] = utils.GetLowAndHighByte(utils.GetBCD(bcdDevice)) // Low and high byte of device release number, must be binary coded decimal, `bcdDevice` field of standard device descriptor

	buf[0x1A] = numberOfDownstreamPorts // Number of downstream ports, `bNbrPorts` field of hub descriptor.

	buf[0x1F] = utils.GetChecksum(buf[:0x1E]) // The 8-bit sum of all values from `0x00` to `0x1E`.

	return buf, nil
}
