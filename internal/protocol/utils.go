package protocol

import (
	"bytes"
	"fmt"
)

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

func (prot *MqttProtocol) IsValidMqttCmd() (*Command, []byte, error) {
	data, err := prot.conn.Read(2)
	var cmd Command
	if err != nil {
		return nil, make([]byte, 0), err
	}
	if len(data) > 2 {
		return nil, make([]byte, 0), fmt.Errorf("length data is < 2")
	}
	cmd = Command(data[0])
	length := calcLength(data[1])
	if length != 0 {
		data2, err := prot.conn.Read(length)
		if err != nil {
			return nil, make([]byte, 0), err
		}
		data = append(data, data2...)
	}
	return &cmd, data, nil
}

func (prot *MqttProtocol) isMqttCmd(Cmd Command) ([]byte, error) {
	data, err := prot.conn.Read(2)
	if err != nil {
		return make([]byte, 0), err
	}
	if len(data) > 2 {
		return make([]byte, 0), fmt.Errorf("length data is < 2")
	}
	if data[0]&byte(Cmd) != byte(Cmd) {
		return make([]byte, 0), fmt.Errorf("Isn't a first byte %d Mqtt protocol", Cmd)
	}
	length := calcLength(data[1])
	data2, err := prot.conn.Read(length)
	data = append(data, data2...)
	if !mqttVersionCompatible(data[2:9]) {
		return make([]byte, 0), fmt.Errorf("Invalid MQTT Protocol Name or Version")
	}
	return data, nil
}
