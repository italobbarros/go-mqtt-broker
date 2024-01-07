package broker

import (
	"fmt"
	"sync"
	"time"
)

// NewSessionManager creates a new SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		partitionMap: make(map[int16]*SessionPartition),
		sessionMap:   make(map[string]*Session),
	}
}

func (sm *SessionManager) Exist(id string) bool {
	sm.lockSession.Lock()
	_, ok := sm.sessionMap[id]
	sm.lockSession.Unlock()
	return ok
}

// AddSession adds a new session to the top of the list
func (sm *SessionManager) AddSession(sessionCfg *SessionConfig) {
	sm.lockPartition.Lock()
	defer sm.lockPartition.Unlock()
	session := &Session{
		config:    sessionCfg,
		Timestamp: time.Now(),
	}
	sessionPartition, ok := sm.partitionMap[sessionCfg.KeepAlive]
	if !ok {
		sessionPartition = &SessionPartition{}
		sm.partitionMap[sessionCfg.KeepAlive] = sessionPartition
	}
	sm.lockSession.Lock()
	sm.sessionMap[sessionCfg.Id] = session
	sm.lockSession.Unlock()
	if sessionPartition.head == nil {
		sessionPartition.head = session
		sessionPartition.tail = session
		return
	}
	session.bottom = sessionPartition.head
	sessionPartition.head.top = session
	sessionPartition.head = session

}

// UpdateSession moves the updated session to the top of the list
func (sm *SessionManager) UpdateSession(sessionCfg *SessionConfig) {
	sm.lockPartition.Lock()
	defer sm.lockPartition.Unlock()
	if SessionPartition, ok := sm.partitionMap[sessionCfg.KeepAlive]; ok {
		// If already at the top, do nothing
		sm.lockSession.Lock()
		session, ok := sm.sessionMap[sessionCfg.Id]
		sm.lockSession.Unlock()
		if !ok {
			return
		}
		if session.config.KeepAlive != sessionCfg.KeepAlive {
			//TODO caso tenha mudado o keepalive
		}
		session.config = sessionCfg
		session.Timestamp = time.Now()
		if session == SessionPartition.head {
			return
		}
		if session == SessionPartition.tail {
			SessionPartition.tail = session.top
		}

		// Remove the node from its current position
		if session.bottom != nil {
			session.bottom.top = session.top
		}
		if session.top != nil {
			session.top.bottom = session.bottom
		}

		// Move the node to the top
		session.bottom = SessionPartition.head
		session.top = nil
		SessionPartition.head.top = session
		SessionPartition.head = session
	}
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

	sm.lockSession.Lock()
	delete(sm.sessionMap, session.config.Id)
	sm.lockSession.Unlock()
	// Remove the session from the map
}

// RemoveSession removes sessions from the map and updates pointers
func (sm *SessionManager) RemoveSession(id string, keepAlive int16) {
	sm.lockPartition.Lock()
	defer sm.lockPartition.Unlock()
	if sessionPartition, ok := sm.partitionMap[keepAlive]; ok {
		// If the session is the head of the list

		sm.lockSession.Lock()
		session, ok := sm.sessionMap[id]
		sm.lockSession.Unlock()
		if !ok {
			return
		}
		sm.onlyRemoveSession(sessionPartition, session)
	}
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
	sm.lockPartition.Lock()
	defer sm.lockPartition.Unlock()

	if len(sm.sessionMap) == 0 {
		return fmt.Errorf("sessionMap is empty")
	}
	if len(sm.partitionMap) == 0 {
		return fmt.Errorf("partitionMap is empty")
	}

	currentTimestamp := time.Now()
	var wg sync.WaitGroup
	// Itera sobre cada SessionPartition
	for _, sessionPartition := range sm.partitionMap {
		wg.Add(1) // Incrementa o WaitGroup para cada goroutine iniciada
		go sm.CheckSession(sessionPartition, &wg, &currentTimestamp)
	}
	wg.Wait()

	return nil
}

func (sm *SessionManager) DebugPrint() {
	sm.lockPartition.Lock()
	defer sm.lockPartition.Unlock()

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

	}
	fmt.Print("-------------------------------------------------")
}
