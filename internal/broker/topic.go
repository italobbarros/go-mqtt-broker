package broker

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/italobbarros/go-mqtt-broker/internal/api/models"
	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
	"github.com/italobbarros/go-mqtt-broker/pkg/client"
)

func (b *Broker) GetTopicNode(topic string) (*models.TopicResponse, error) {
	queryParams := map[string]string{
		"Name": topic,
	}

	// Cabeçalhos
	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := client.Get(os.Getenv("API_GET_TOPIC"), client.RequestOptions{
		Params:  queryParams,
		Headers: headers,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Decodifica o corpo da resposta em um valor do tipo TopicResponse
		var topicResponse models.TopicResponse
		if err := json.NewDecoder(resp.Body).Decode(&topicResponse); err != nil {
			return nil, err
		}

		return &topicResponse, nil
	}

	// Se a resposta não foi bem-sucedida, retorna um erro
	return nil, fmt.Errorf("Erro na resposta. Código de status: %d", resp.StatusCode)
}

func (b *Broker) AddPublish(publishRequest models.PublishRequest) error {
	// Cabeçalhos
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	// Realiza a requisição POST usando a função do pacote client
	resp, err := client.Post(os.Getenv("API_POST_PUBLISH"), client.RequestOptions{
		Headers:    headers,
		Body:       publishRequest,
		JSONEncode: true,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verifica se a resposta foi bem-sucedida (código 2xx)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	// Se a resposta não foi bem-sucedida, retorna um erro
	return fmt.Errorf("Erro na resposta. Código de status: %d, body:%s", resp.StatusCode, resp.Body)
}

func (b *Broker) AddSubscribeTopicNode(topic string, id string, subs *SubscriberConfig, topicNode chan bool) error {

	return nil
}

func (b *Broker) notifyNewSubscriber(topic string, sub *SubscriberConfig) error {
	b.logger.Debug("NotifyNewSubscriber: %s", sub.Identifier)
	return nil
}

func (b *Broker) NotifyAllSubscribers(topic string, topicReady chan bool) {
	b.logger.Debug("NotifyAllSubscribers")
	<-topicReady
	topicNode, err := b.GetTopicNode(topic)
	if err != nil {
		b.logger.Warning("topicNode empty: %v")
		return
	}
	if topicNode == nil {
		b.logger.Warning("topicNode empty: %v")
		return
	}
	/*
		topicNode.MessageCount += 1
		for _, sub := range topicNode.Subscribers {
			currentProt := sub.session.prot
			go b.publishSubscribe(topicNode.Topic, currentProt, topicNode.TopicConfig.Retained, topicNode.TopicConfig.Payload, sub)
		}*/
}

func (b *Broker) publishSubscribe(topic string, currentProt *protocol.MqttProtocol, retained bool, payload string, sub *SubscriberConfig) {
	currentProt.Start()
	err := currentProt.PublishSend(sub.Qos, true, retained, payload, topic, sub.Identifier)
	if err != nil {
		currentProt.End()
	}
	b.logger.Debug("publishSubscribe Identifier: %v", sub.Identifier)
	currentProt.End()
}

func (b *Broker) IncMsgCount(topic string) error {
	return nil
}
