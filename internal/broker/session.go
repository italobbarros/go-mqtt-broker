package broker

import (
	"fmt"
	"sync"
	"time"

	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

// NewSessionManager creates a new SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionMap:   &sync.Map{},
		partitionMap: &sync.Map{},
	}
}

func (sm *SessionManager) Exist(id string) bool {
	_, ok := sm.sessionMap.Load(id)
	return ok
}

func (sm *SessionManager) GetSessionCount() int {
	return sm.sessionCount
}

// AddSession adds a new session to the top of the list
func (sm *SessionManager) AddSession(sessionCfg *SessionConfig) *Session {
	session := &Session{
		config:    sessionCfg,
		Timestamp: time.Now(),
		logger:    logger.NewLogger(sessionCfg.Id),
	}
	sessionPartitionVar, ok := sm.partitionMap.Load(sessionCfg.KeepAlive)
	var sessionPartition *SessionPartition
	if !ok {
		sessionPartition = &SessionPartition{}
		sm.partitionMap.Store(sessionCfg.KeepAlive, sessionPartition)
		sm.partitionCount++
	} else {
		sessionPartition = sessionPartitionVar.(*SessionPartition)

	}
	sm.sessionMap.Store(sessionCfg.Id, session)
	if sessionPartition.head == nil {
		sessionPartition.head = session
		sessionPartition.tail = session
		return session
	}
	session.bottom = sessionPartition.head
	sessionPartition.head.top = session
	sessionPartition.head = session
	sm.sessionCount++
	return session
}

// UpdateSession moves the updated session to the top of the list
func (sm *SessionManager) UpdateSession(sessionCfg *SessionConfig) *Session {
	sessionPartitionVar, ok := sm.partitionMap.Load(sessionCfg.KeepAlive)
	if !ok {
		return nil
	}
	sessionPartition := sessionPartitionVar.(*SessionPartition)

	sessionVar, ok := sm.sessionMap.Load(sessionCfg.Id)
	if !ok {
		return nil
	}
	session := sessionVar.(*Session)

	if session.config.KeepAlive != sessionCfg.KeepAlive {
		//TODO caso tenha mudado o keepalive
	}
	session.config = sessionCfg
	session.Timestamp = time.Now()
	if session == sessionPartition.head {
		return session
	}
	if session == sessionPartition.tail {
		sessionPartition.tail = session.top
	}

	// Remove the node from its current position
	if session.bottom != nil {
		session.bottom.top = session.top
	}
	if session.top != nil {
		session.top.bottom = session.bottom
	}

	// Move the node to the top
	session.bottom = sessionPartition.head
	session.top = nil
	sessionPartition.head.top = session
	sessionPartition.head = session
	return session
}

func (sm *SessionManager) onlyRemoveSession(sessionPartition *SessionPartition, session *Session) {
	if session == sessionPartition.head {
		sessionPartition.head = session.bottom
	}
	if session == sessionPartition.tail {
		sessionPartition.tail = session.top
	}

	// Update pointers to remove the session from the list
	if session.bottom != nil {
		session.bottom.top = session.top
	}
	if session.top != nil {
		session.top.bottom = session.bottom
	}
	sm.sessionMap.Delete(session.config.Id)
	sm.sessionCount--
	// Remove the session from the map
}

// RemoveSession removes sessions from the map and updates pointers
func (sm *SessionManager) RemoveSession(id string, keepAlive int16) {
	sessionPartitionVar, ok := sm.partitionMap.Load(keepAlive)
	if !ok {
		return
	}
	sessionVar, ok := sm.sessionMap.Load(id)
	if !ok {
		return
	}
	sm.onlyRemoveSession(sessionPartitionVar.(*SessionPartition), sessionVar.(*Session))
}

func (sm *SessionManager) CheckSession(partition *SessionPartition, wg *sync.WaitGroup, currentTimestamp *time.Time) {
	defer wg.Done() // Sinaliza que a goroutine terminou
	currentSession := partition.tail
	for currentSession != nil {
		var elapsed float64
		if !currentSession.config.Clean {
			elapsed = currentTimestamp.Sub(currentSession.Timestamp).Seconds()
			if elapsed > 3600 {
				sm.onlyRemoveSession(partition, currentSession)
			} else {
				break
			}
		} else {
			elapsed = currentTimestamp.Sub(currentSession.Timestamp).Seconds()
			if elapsed > float64(currentSession.config.KeepAlive) {
				sm.onlyRemoveSession(partition, currentSession)
			} else {
				break
			}
		}
		currentSession = currentSession.top
	}
	//return nil
}

// CheckSessionTimeouts verifica e remove sessões cujo tempo limite expirou
func (sm *SessionManager) CheckSessionTimeouts() error {

	if sm.sessionCount == 0 {
		return fmt.Errorf("SessionMap is empty")
	}
	if sm.partitionCount == 0 {
		return fmt.Errorf("partitionMap is empty")
	}
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
