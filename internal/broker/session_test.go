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
	if sm.partitionCount != 0 || sm.sessionCount != 0 {
		t.Error("Expected empty maps for partitionMap and SessionMap")
	}
}

func TestAddSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:        "testID",
		KeepAlive: 10,
		Clean:     false,
	}
	sm.AddSession(cfg)
	if sm.partitionCount != 0 || sm.sessionCount != 0 {
		t.Error("Expected one entry in partitionMap and SessionMap")
	}
}

func TestUpdateSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:        "testID",
		KeepAlive: 10,
		Clean:     false,
	}
	sm.AddSession(cfg)
	time.Sleep(1 * time.Second) // Espera 1 segundo para simular uma atualização
	sm.UpdateSession(cfg)
	// Você pode adicionar mais verificações aqui para validar a funcionalidade.
}

func TestRemoveSession(t *testing.T) {
	sm := NewSessionManager()
	cfg := &SessionConfig{
		Id:        "testID",
		KeepAlive: 10,
		Clean:     false,
	}
	sm.AddSession(cfg)
	sm.RemoveSession("testID", 10)
	if sm.sessionCount != 0 {
		t.Error("Expected SessionMap to be empty after removal")
	}
}

func TestCheckSessionTimeouts(t *testing.T) {
	sm := NewSessionManager()
	sm.AddSession(&SessionConfig{
		Id:        "testID1",
		KeepAlive: 1,
		Clean:     true,
	})
	sm.AddSession(&SessionConfig{
		Id:        "testID2",
		KeepAlive: 0,
		Clean:     true,
	})
	sm.AddSession(&SessionConfig{
		Id:        "testID3",
		KeepAlive: 0,
		Clean:     false,
	})
	sm.AddSession(&SessionConfig{
		Id:        "testID3",
		KeepAlive: 10,
		Clean:     true,
	})
	t.Log("SessionMap length:", sm.sessionCount)
	t.Log("partitionMap length:", sm.partitionCount)
	time.Sleep(3 * time.Second) // Espera 2 segundos para que a sessão expire

	err := sm.CheckSessionTimeouts()

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if sm.sessionCount != 2 {
		t.Errorf("Expected SessionMap to be 2 after timeout check, but is %d", sm.sessionCount)
	}
}

// Certifique-se de executar esses testes usando "go test" no diretório do seu pacote.
