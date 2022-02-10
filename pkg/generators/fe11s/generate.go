package fe11s

// | Address     | Contents              | Note                                                                                                             |
// | ----------- | --------------------- | ---------------------------------------------------------------------------------------------------------------- |
// | `0x00`      | `0x40`                | Constant, low byte of check code                                                                                 |
// | `0x01`      | `0x1A`                | Constant, high byte of check code                                                                                |
// | `0x02`      | Vendor ID (Low)       | Low byte of Vendor ID, `idVendor` field of Standard Device Descriptor                                            |
// | `0x03`      | Vendor ID (High)      | High byte of Vendor ID                                                                                           |
// | `0x04`      | Product ID (Low)      | Low byte of Product ID, `idProduct` field of Standard Device Descriptor                                          |
// | `0x05`      | Product ID (High)     | High Byte of Product ID                                                                                          |
// | `0x06`      | Device Release (Low)  | Low byte of Device Release Number, must be Binary Coded Decimal, `bcdDevice` field of Standard Device Descriptor |
// | `0x07`      | Device Release (High) | High byte of Device Release Number, must be Binary Coded Decimal                                                 |
// | `0x08-0x19` | Filling All           | `0x00`                                                                                                           |
// | `0x1A`      | Port Number           | Number of Downstream Ports, `bNbrPorts` field of Hub Descriptor.                                                 |
// | `0x1B-0x1E` | Filling All           | `0x00`                                                                                                           |
// | `0x1F`      | Check Sum             | The 8-bit sum of all value from `0x00` to `0x1E`.                                                                |

func GenerateEEPROM(
	idVendor uint16, // i.e. 0x046d for Logitech
	idProduct uint16, // i.e. 0x082d for HD Pro Webcam C920
	bcdDevice uint16, // i.e. 0x0001 for release 1
	numberOfDownstreamPorts uint8, // i.e. 4 for 4 ports
) ([]byte, error) {
	buf := make([]byte, 0x1F) // Filling is `0x00`

	buf[0x00] = 0x40 // Constant, low byte of check code
	buf[0x01] = 0x1A // Constant, high byte of check code

	buf[0x02], buf[0x03] = getLowAndHighByte(idVendor) // Low and high byte of Vendor ID, `idVendor` field of Standard Device Descriptor

	buf[0x04], buf[0x05] = getLowAndHighByte(idProduct) // Low and high byte of Product ID, `idProduct` field of Standard Device Descriptor

	buf[0x06], buf[0x07] = getLowAndHighByte(getBCD(bcdDevice)) // Low and high byte of Device Release Number, must be Binary Coded Decimal, `bcdDevice` field of Standard Device Descriptor

	buf[0x1F] = numberOfDownstreamPorts // The 8-bit sum of all value from `0x00` to `0x1E`.
	for _, el := range buf[:0x1E] {
		buf[0x1A] += el
	}

	return buf, nil
}

func getLowAndHighByte(word uint16) (byte, byte) {
	// See https://www.mikrocontroller.net/topic/231566
	return uint8(word & 0xff), uint8(word >> 8)
}

func getBCD(word uint16) uint16 {
	// See https://github.com/masayukioguni/bcd/blob/master/bcd.go#L4
	return (((word / 10) % 10) << 4) | (word % 10)
}
