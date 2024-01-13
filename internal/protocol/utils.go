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

func decodeLength(encoded []byte) (int, error) {
	var length int
	multiplier := 1

	for i := 0; i < len(encoded); i++ {
		// Verifique se o bit de continuação está definido
		if encoded[i]&0x80 == 0 {
			length += int(encoded[i]) * multiplier
			return length, nil
		}

		// Remova o bit de continuação e some ao comprimento
		length += int(encoded[i]&0x7F) * multiplier
		multiplier *= 128

		// Verifique se atingiu o máximo de 4 bytes para o comprimento
		if multiplier > 128*128*128 {
			return 0, fmt.Errorf("malformed remaining length")
		}
	}

	return 0, fmt.Errorf("incomplete remaining length")
}

func encodeLength(length int) []byte {
	var encoded []byte

	for {
		digit := length % 128
		length = length / 128

		// Se ainda houver dígitos, defina o bit de continuação
		if length > 0 {
			digit = digit | 0x80
		}

		encoded = append(encoded, byte(digit))

		if length <= 0 {
			break
		}
	}

	return encoded
}

type MqttCmdResult struct {
	Command *Command
	Data    []byte
	Err     error
}

func (prot *MqttProtocol) IsValidMqttCmd() *MqttCmdResult {

	data, err := prot.conn.Read(2)
	result := &MqttCmdResult{}

	if err != nil {
		result.Err = err
		return result
	}
	if len(data) > 2 {
		result.Err = fmt.Errorf("length data is < 2")
		return result
	}
	cmd := Command(data[0])
	result.Command = &cmd
	var d1 []byte
	if data[1] > 127 {
		d1, err = prot.conn.Read(1)
		if err != nil {
			result.Err = err
			return result
		}
		data = append(data, d1...)
		if data[2] > 127 {
			d1, err = prot.conn.Read(1)
			if err != nil {
				result.Err = err
				return result
			}
			data = append(data, d1...)
			if data[3] > 127 {
				d1, err = prot.conn.Read(1)
				if err != nil {
					result.Err = err
					return result
				}
				data = append(data, d1...)
			}
		}
	}
	length, err := decodeLength(data[1:])
	if err != nil {
		result.Err = err
		return result
	}
	if length != 0 {
		data2, err := prot.conn.Read(length)
		if err != nil {
			result.Err = err
			return result
		}
		data = append(data, data2...)
	}
	result.Data = data
	return result
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
	var d1 []byte
	if data[1] > 127 {
		d1, err = prot.conn.Read(1)
		if err != nil {
			return make([]byte, 0), err
		}
		data = append(data, d1...)
		if data[2] > 127 {
			d1, err = prot.conn.Read(1)
			if err != nil {
				return make([]byte, 0), err
			}
			data = append(data, d1...)
			if data[3] > 127 {
				d1, err = prot.conn.Read(1)
				if err != nil {
					return make([]byte, 0), err
				}
				data = append(data, d1...)
			}
		}
	}
	length, err := decodeLength(data[1:])
	if err != nil {
		return make([]byte, 0), err
	}
	data2, err := prot.conn.Read(length)
	data = append(data, data2...)
	if !mqttVersionCompatible(data[2:9]) {
		return make([]byte, 0), fmt.Errorf("Invalid MQTT Protocol Name or Version")
	}
	return data, nil
}

func IsCmdEqual(currentCommand *Command, testCommand Command) bool {
	return *currentCommand&testCommand == testCommand
}
