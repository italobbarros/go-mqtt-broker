package main

import (
	"encoding/json"
	"net/http"
)

type TreeNode struct {
	Name     string      `json:"name"`
	Children []*TreeNode `json:"children,omitempty"`
}

func main() {
	http.HandleFunc("/mqtt-tree", corsHandler(mqttTreeHandler))
	http.ListenAndServe(":8080", nil)
}

func corsHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Configurar cabeçalhos CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permitir qualquer origem
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			return
		}

		fn(w, r)
	}
}

func mqttTreeHandler(w http.ResponseWriter, r *http.Request) {
	// Simulação de uma árvore de tópicos MQTT simples
	root := &TreeNode{
		Name: "MQTT",
		Children: []*TreeNode{
			{
				Name: "client1",
				Children: []*TreeNode{
					{
						Name: "teste",
						Children: []*TreeNode{
							{Name: "client1_io1"},
							{Name: "client1_io2"},
						},
					},
				},
			},
			{
				Name: "client2",
				Children: []*TreeNode{
					{Name: "client2_io1"},
					{Name: "client2_io2"},
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}
