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
	vendorID [2]byte, // i.e. 2109 -> []byte{0x1a, 0x40} for VIA Labs, Inc.
	productID [2]byte, // i.e. 0817 -> []byte{0x01, 0x01} for USB3.0 Hub
	deviceReleaseNumber [2]byte, // i.e. 3.83 → []byte{0x3, 0x83}
	numberOfDownstreamPorts byte, // i.e. 4
) ([]byte, error) {
	buf := make([]byte, 0x1F) // Filling is `0x00`

	buf[0x00] = 0x40 // Constant, low byte of check code
	buf[0x01] = 0x1A // Constant, high byte of check code

	buf[0x02] = vendorID[0] // Low byte of Vendor ID, `idVendor` field of Standard Device Descriptor
	buf[0x03] = vendorID[1] // High byte of Vendor ID

	buf[0x04] = productID[0] // Low byte of Product ID, `idProduct` field of Standard Device Descriptor
	buf[0x05] = productID[1] // High Byte of Product ID

	buf[0x06] = deviceReleaseNumber[0] // Low byte of Device Release Number, must be Binary Coded Decimal, `bcdDevice` field of Standard Device Descriptor
	buf[0x07] = deviceReleaseNumber[1] // High byte of Device Release Number, must be Binary Coded Decimal

	buf[0x1F] = numberOfDownstreamPorts // The 8-bit sum of all value from `0x00` to `0x1E`.
	for _, el := range buf[:0x1E] {
		buf[0x1A] += el
	}

	return buf, nil
}
