package broker

import (
	"testing"
	"time"
)

func TestNewSessionManager(t *testing.T) {
	sm := NewSessionManager()
	if sm == nil {
		t.Error("Expected a non-nil SessionManager")
	}

}

func TestAddSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:        "testID",
		KeepAlive: 10,
		Clean:     false,
	}
	chSession := make(chan *Session)
	sm.AddSession(cfg, chSession)

}

func TestUpdateSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:        "testID",
		KeepAlive: 10,
		Clean:     false,
	}
	chSession := make(chan *Session)
	sm.AddSession(cfg, chSession)
	//session := <-chSession
	time.Sleep(1 * time.Second) // Espera 1 segundo para simular uma atualização
	//sm.UpdateSession(cfg, chSession)
	// Você pode adicionar mais verificações aqui para validar a funcionalidade.
}

func TestRemoveSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:        "testID",
		KeepAlive: 10,
		Clean:     false,
	}
	chSession := make(chan *Session)
	sm.AddSession(cfg, chSession)
	sm.RemoveSession("testID", 10)

}

func TestCheckSessionTimeouts(t *testing.T) {
	sm := NewSessionManager()
	chSession := make(chan *Session)

	sm.AddSession(&SessionConfig{
		Id:        "testID1",
		KeepAlive: 1,
		Clean:     true,
	}, chSession)
	sm.AddSession(&SessionConfig{
		Id:        "testID2",
		KeepAlive: 0,
		Clean:     true,
	}, chSession)
	sm.AddSession(&SessionConfig{
		Id:        "testID3",
		KeepAlive: 0,
		Clean:     false,
	}, chSession)
	sm.AddSession(&SessionConfig{
		Id:        "testID3",
		KeepAlive: 10,
		Clean:     true,
	}, chSession)
	//t.Log("SessionMap length:", sm.sessionCount)
	//t.Log("partitionMap length:", sm.partitionCount)
	time.Sleep(3 * time.Second) // Espera 2 segundos para que a sessão expire

	err := sm.CheckSessionTimeouts()

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	//if sm.sessionCount != 2 {
	//	t.Errorf("Expected SessionMap to be 2 after timeout check, but is %d", sm.sessionCount)
	//}
}

// Certifique-se de executar esses testes usando "go test" no diretório do seu pacote.
