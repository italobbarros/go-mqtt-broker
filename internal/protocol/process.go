package protocol

func (prot *MqttProtocol) ConnectProcess() (*ResponseConnect, error) {
	r, err := prot.connectUnPack()
	if err != nil {
		if r == nil {
			return nil, err
		}
		err = prot.connAck(r, CONNECT_UNACCEPCTABLE_PROTOCOL)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	if r == nil { //Clean Session False
		err = prot.connAck(r, CONNECT_ACCEPCTED)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	err = prot.connAck(r, CONNECT_ACCEPCTED)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (prot *MqttProtocol) PubRelProcess(data []byte, packetIdentifier *[]byte) error {
	err := prot.unpackPubRel(data, packetIdentifier)
	if err != nil {
		return err
	}
	err = prot.pubComp(packetIdentifier)
	if err != nil {
		return err
	}
	return nil
}

func (prot *MqttProtocol) PublishProcess(data []byte) (*ResponsePublish, error) {
	r, err := prot.publishUnPack(data)
	if err != nil {
		return nil, err
	}
	ackErr := make(chan error)
	if r.Qos == 1 {
		go prot.pubAck(r, ackErr)
		err := <-ackErr
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	if r.Qos == 2 {
		go prot.pubRec(r, ackErr)
		err := <-ackErr
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return r, nil
}

func (prot *MqttProtocol) SubscribeProcess(data []byte) (*ResponseSubscribe, error) {
	r, err := prot.subscribeUnPack(data)
	if err != nil {
		return nil, err
	}
	r.Conn = prot.conn
	return r, nil
}

func (prot *MqttProtocol) UnSubscribeProcess(data []byte) (*ResponseSubscribe, error) {
	r, err := prot.unSubscribeUnPack(data)
	if err != nil {
		return nil, err
	}
	err = prot.unSubAck(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (prot *MqttProtocol) PingProcess() error {
	response := []byte{byte(COMMAND_PINGRESP), 0}
	err := prot.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}
