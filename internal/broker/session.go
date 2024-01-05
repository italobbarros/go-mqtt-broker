package broker

import (
	"sync"
	"time"
)

// NewSessionManager creates a new SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		partitionMap: make(map[int]*SessionPartition),
		sessionMap:   make(map[string]*Session),
	}
}

// AddSession adds a new session to the top of the list
func (sm *SessionManager) AddSession(sessionCfg *SessionConfig) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	session := &Session{
		config:    sessionCfg,
		Timestamp: time.Now(),
	}
	sessionPartition, ok := sm.partitionMap[sessionCfg.Timeout]
	if !ok {
		sessionPartition = &SessionPartition{}
		sm.partitionMap[sessionCfg.Timeout] = sessionPartition
	}

	sm.sessionMap[sessionCfg.Id] = session
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
func (sm *SessionManager) UpdateSession(id string, timeout int) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if SessionPartition, ok := sm.partitionMap[timeout]; ok {
		// If already at the top, do nothing
		session, ok := sm.sessionMap[id]
		if !ok {
			return
		}
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

// RemoveSession removes sessions from the map and updates pointers
func (sm *SessionManager) RemoveSession(id string, timeout int) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	if SessionPartition, ok := sm.partitionMap[timeout]; ok {
		// If the session is the head of the list
		session, ok := sm.sessionMap[id]
		if !ok {
			return
		}
		if session == SessionPartition.head {
			SessionPartition.head = session.bottom
		}
		if session == SessionPartition.tail {
			SessionPartition.tail = session.top
		}

		// Update pointers to remove the session from the list
		if session.bottom != nil {
			session.bottom.top = session.top
		}
		if session.top != nil {
			session.top.bottom = session.bottom
		}

		// Remove the session from the map
		delete(sm.sessionMap, id)
	}
}

// CheckSessionTimeouts verifica e remove sessões cujo tempo limite expirou
func (sm *SessionManager) CheckSessionTimeouts() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	currentTimestamp := time.Now()

	// Use WaitGroup para esperar até que todas as goroutines sejam concluídas
	var wg sync.WaitGroup

	// Itera sobre cada SessionPartition
	for _, sessionPartition := range sm.partitionMap {
		wg.Add(1) // Incrementa o WaitGroup para cada goroutine iniciada

		// Processa cada SessionPartition em uma goroutine separada
		go func(partition *SessionPartition) {
			defer wg.Done() // Sinaliza que a goroutine terminou

			currentSession := partition.tail
			for currentSession != nil {
				var elapsed float64
				if !currentSession.config.Clean {
					elapsed = currentTimestamp.Sub(currentSession.Timestamp).Seconds()
					if elapsed > 3600 {
						sm.RemoveSession(currentSession.config.Id, currentSession.config.Timeout)
					} else {
						// Se a sessão ainda está dentro do intervalo, pode sair do loop
						break
					}
				} else {
					elapsed = currentTimestamp.Sub(currentSession.Timestamp).Seconds()
					if elapsed > float64(currentSession.config.Timeout) {
						sm.RemoveSession(currentSession.config.Id, currentSession.config.Timeout)
					} else {
						// Se a sessão ainda está dentro do intervalo, pode sair do loop
						break
					}
				}
				currentSession = currentSession.top
			}
		}(sessionPartition)
	}

	// Aguarde até que todas as goroutines sejam concluídas
	wg.Wait()
}
