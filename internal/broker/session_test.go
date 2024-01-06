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
	if len(sm.partitionMap) != 0 || len(sm.sessionMap) != 0 {
		t.Error("Expected empty maps for partitionMap and sessionMap")
	}
}

func TestAddSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:      "testID",
		Timeout: 10,
		Clean:   false,
	}
	sm.AddSession(cfg)
	if len(sm.partitionMap) != 1 || len(sm.sessionMap) != 1 {
		t.Error("Expected one entry in partitionMap and sessionMap")
	}
}

func TestUpdateSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:      "testID",
		Timeout: 10,
		Clean:   false,
	}
	sm.AddSession(cfg)
	time.Sleep(1 * time.Second) // Espera 1 segundo para simular uma atualização
	sm.UpdateSession("testID", 10)
	// Você pode adicionar mais verificações aqui para validar a funcionalidade.
}

func TestRemoveSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:      "testID",
		Timeout: 10,
		Clean:   false,
	}
	sm.AddSession(cfg)
	sm.RemoveSession("testID", 10)
	if len(sm.sessionMap) != 0 {
		t.Error("Expected sessionMap to be empty after removal")
	}
}

func TestCheckSessionTimeouts(t *testing.T) {
	sm := NewSessionManager()
	sm.AddSession(&SessionConfig{
		Id:      "testID1",
		Timeout: 1,
		Clean:   true,
	})
	sm.AddSession(&SessionConfig{
		Id:      "testID2",
		Timeout: 0,
		Clean:   true,
	})
	sm.AddSession(&SessionConfig{
		Id:      "testID3",
		Timeout: 0,
		Clean:   false,
	})
	sm.AddSession(&SessionConfig{
		Id:      "testID3",
		Timeout: 10,
		Clean:   true,
	})
	t.Log("sessionMap length:", len(sm.sessionMap))
	t.Log("partitionMap length:", len(sm.partitionMap))
	time.Sleep(3 * time.Second) // Espera 2 segundos para que a sessão expire

	err := sm.CheckSessionTimeouts()

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if len(sm.sessionMap) != 2 {
		t.Errorf("Expected sessionMap to be 2 after timeout check, but is %d", len(sm.sessionMap))
	}
}

// Certifique-se de executar esses testes usando "go test" no diretório do seu pacote.
