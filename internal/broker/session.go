package broker

import (
	"sync"
	"time"
)

type SessionConfig struct {
	Id      string
	Timeout int
	Clean   bool
}

// Session representa uma sessão MQTT
type Session struct {
	Timestamp time.Time
	config    *SessionConfig
	top       *Session // Ponteiro para o nó anterior na lista
	bottom    *Session // Ponteiro para o próximo nó na lista
}

// SessionManager gerencia sessões MQTT
type SessionManager struct {
	head       *Session            // Ponteiro para o primeiro nó da lista
	tail       *Session            // Ponteiro para o ultimo nó da lista
	sessionMap map[string]*Session // Mapa para acessar sessões por ID
	lock       sync.Mutex
}

// NewSessionManager cria um novo SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionMap: make(map[string]*Session),
	}
}

// AddSession adiciona uma nova sessão ao topo da lista
func (sm *SessionManager) AddSession(sessionCfg *SessionConfig) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	session := &Session{
		config:    sessionCfg,
		Timestamp: time.Now(),
	}

	sm.sessionMap[sessionCfg.Id] = session

	if sm.head == nil {
		sm.head = session
		sm.tail = session
		return
	}
	session.bottom = sm.head
	sm.head.top = session
	sm.head = session
}

// UpdateSession move a sessão atualizada para o topo da lista
func (sm *SessionManager) UpdateSession(id string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if session, ok := sm.sessionMap[id]; ok {
		// Se já estiver no topo, não faz nada
		session.Timestamp = time.Now()
		if session == sm.head {
			return
		}
		if session == sm.tail {
			sm.tail = session.top
		}

		// Remove o nó da sua posição atual
		if session.bottom != nil {
			session.bottom.top = session.top
		}
		if session.top != nil {
			session.top.bottom = session.bottom
		}

		// Move o nó para o topo
		session.bottom = sm.head
		session.top = nil
		sm.head.top = session
		sm.head = session
	}
}

func (sm *SessionManager) RemoveSession(id string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	// Verifica se a sessão existe no mapa
	if session, ok := sm.sessionMap[id]; ok {
		// Se a sessão for a cabeça da lista
		if session == sm.head {
			sm.head = session.bottom
		}
		if session == sm.tail {
			sm.tail = session.top
		}

		// Atualiza os ponteiros para remover a sessão da lista
		if session.bottom != nil {
			session.bottom.top = session.top
		}
		if session.top != nil {
			session.top.bottom = session.bottom
		}

		// Remove a sessão do mapa
		delete(sm.sessionMap, id)
	}
}

// CheckSessionTimeouts verifica e remove as sessões cujo timeout expirou
func (sm *SessionManager) CheckSessionTimeouts() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	currentTimestamp := time.Now()

	currentSession := sm.tail
	for currentSession != nil {
		if !currentSession.config.Clean {
			elapsed := currentTimestamp.Sub(currentSession.Timestamp)
			if elapsed.Seconds() > float64(3600) {
				sm.RemoveSession(currentSession.config.Id)
			}
		} else {
			elapsed := currentTimestamp.Sub(currentSession.Timestamp)
			if elapsed.Seconds() > float64(currentSession.config.Timeout) {
				sm.RemoveSession(currentSession.config.Id)
			}
		}
		currentSession = currentSession.top
	}
}

//func main() {
//	manager := NewSessionManager(&SessionConfig{})
//
// Adicionar algumas sessões
//	manager.AddSession("session1")
//	manager.AddSession("session2")
//	manager.AddSession("session3")
//
// Atualizar uma sessão
//	manager.UpdateSession("session2")
//
// Remover uma sessão
//	manager.RemoveSession("session2")
//
//	// Simplesmente imprimindo para demonstração
//	current := manager.head
//	for current != nil {
//		fmt.Println(current.ID)
//		current = current.bottom
//	}
//}
