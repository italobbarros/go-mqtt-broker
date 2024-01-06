package protocol

import "bytes"

func mqttVersionCompatible(b []byte) bool {
	if bytes.Equal(b, []byte{0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x0}) { //3.1 MQTT 0
		return true
	}
	if bytes.Equal(b, []byte{0x00, 0x04, 0x4D, 0x51, 0x54, 0x54, 0x4}) { //3.1 MQTT 1
		return true
	}
	if bytes.Equal(b, []byte{0x00, 0x05, 0x4D, 0x51, 0x54, 0x54, 0x0}) { //5.1 MQTT 0
		return true
	}
	return false
}

func calcLength(b byte) int {
	valor := 0
	multiplicador := 1

	for b&128 != 0 {
		valor += int(b&127) * multiplicador
		multiplicador *= 128
		b >>= 7
	}

	return valor + int(b)
}
