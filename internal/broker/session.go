package broker

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/italobbarros/go-mqtt-broker/internal/api/models"
	"github.com/italobbarros/go-mqtt-broker/pkg/client"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

// NewSessionManager creates a new SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionMap: &sync.Map{},
	}
}

func (sm *SessionManager) Exist(id string) bool {
	_, ok := sm.sessionMap.Load(id)
	return ok
}

func (sm *SessionManager) GetSessionCount() int {
	return 1
}

func addSession(sessionRequest models.SessionRequest) (*models.SessionResponse, error) {
	// Cabeçalhos
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	// Realiza a requisição POST usando a função do pacote client
	resp, err := client.Post(os.Getenv("API_POST_SESSION"), client.RequestOptions{
		Headers:    headers,
		Body:       sessionRequest,
		JSONEncode: true,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica se a resposta foi bem-sucedida (código 2xx)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Session adicionado com sucesso!")
		var sessionResponse models.SessionResponse
		if err := json.NewDecoder(resp.Body).Decode(&sessionResponse); err != nil {
			return nil, err
		}
		return &sessionResponse, nil
	}

	// Se a resposta não foi bem-sucedida, retorna um erro
	return nil, fmt.Errorf("Erro na resposta. Código de status: %d, body:%s", resp.StatusCode, resp.Body)
}

// AddSession adds a new session to the top of the list
func (sm *SessionManager) AddSession(sessionCfg *SessionConfig, chSession chan *Session) {

	intNumber, _ := strconv.Atoi(os.Getenv("CONTAINER_ID"))
	r, _ := addSession(models.SessionRequest{
		IdContainer: uint64(intNumber),
		ClientId:    sessionCfg.Id,
		KeepAlive:   sessionCfg.KeepAlive,
		Clean:       sessionCfg.Clean,
		Username:    sessionCfg.username,
		Password:    sessionCfg.password,
	})
	session := &Session{
		Id:        r.ClientId,
		KeepAlive: r.KeepAlive,
		Clean:     r.Clean,
		username:  r.Username,
		password:  r.Password,
		Timestamp: r.Updated,
		logger:    logger.NewLogger(r.ClientId),
	}
	chSession <- session
}

// UpdateSession moves the updated session to the top of the list
func (sm *SessionManager) UpdateSession(sessionCfg *SessionConfig, chSession chan *Session) {

	sessionVar, ok := sm.sessionMap.Load(sessionCfg.Id)
	if !ok {
		return
	}
	session := sessionVar.(*Session)

	session.Id = sessionCfg.Id
	session.KeepAlive = sessionCfg.KeepAlive
	session.Clean = sessionCfg.Clean
	session.Timestamp = time.Now()

	chSession <- session
}

func (sm *SessionManager) onlyRemoveSession(session *Session) {
	sm.sessionMap.Delete(session.Id)
	// Remove the session from the map
}

// RemoveSession removes sessions from the map and updates pointers
func (sm *SessionManager) RemoveSession(id string, keepAlive int16) {
	sessionVar, ok := sm.sessionMap.Load(id)
	if !ok {
		return
	}
	sm.onlyRemoveSession(sessionVar.(*Session))
}

func (sm *SessionManager) CheckSession(partition *SessionPartition, wg *sync.WaitGroup, currentTimestamp *time.Time) {
	defer wg.Done() // Sinaliza que a goroutine terminou
	currentSession := partition.tail
	for currentSession != nil {
		var elapsed float64
		if !currentSession.Clean {
			elapsed = currentTimestamp.Sub(currentSession.Timestamp).Seconds()
			if elapsed > 3600 {
				sm.onlyRemoveSession(currentSession)
			} else {
				break
			}
		} else {
			elapsed = currentTimestamp.Sub(currentSession.Timestamp).Seconds()
			if elapsed > float64(currentSession.KeepAlive) {
				sm.onlyRemoveSession(currentSession)
			} else {
				break
			}
		}
	}
	//return nil
}

// CheckSessionTimeouts verifica e remove sessões cujo tempo limite expirou
func (sm *SessionManager) CheckSessionTimeouts() error {

	/*
		currentTimestamp := time.Now()
		var wg sync.WaitGroup
		// Itera sobre cada SessionPartition
		for _, sessionPartition := range sm.partitionMap {
			wg.Add(1) // Incrementa o WaitGroup para cada goroutine iniciada
			go sm.CheckSession(sessionPartition, &wg, &currentTimestamp)
		}
		wg.Wait()*/

	return nil
}

func (sm *SessionManager) DebugPrint() {
	/*
		for k, partition := range sm.partitionMap {
			fmt.Println("------------------------------------------------")
			currentSession := partition.head
			fmt.Println("Partition Key:", k) // Supondo que você tenha um ID para cada partição

			for currentSession != nil {
				fmt.Println("                             ^")
				fmt.Printf("-> Session ID: %s, KeepAlive: %d, Clean: %v ,Timestamp: %s\n",
					currentSession.config.Id,
					currentSession.config.KeepAlive,
					currentSession.config.Clean,
					currentSession.Timestamp.Format("2006-01-02 15:04:05"))
				currentSession = currentSession.bottom
				//fmt.Println("                    v")
			}

		}*/
	fmt.Println("-------------------------------------------------")
}
