package utils

func GetLowAndHighByte(word uint16) (byte, byte) {
	// See https://www.mikrocontroller.net/topic/231566
	return uint8(word & 0xff), uint8(word >> 8)
}

func GetBCD(word uint16) uint16 {
	// See https://github.com/masayukioguni/bcd/blob/master/bcd.go#L4
	return (((word / 10) % 10) << 4) | (word % 10)
}

func GetChecksum(subject []byte) (sum uint8) {
	for _, el := range subject {
		sum += el
	}

	return sum
}
