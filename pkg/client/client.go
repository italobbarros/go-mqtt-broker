// client/client.go
package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RequestOptions contém opções para personalizar uma requisição HTTP.
type RequestOptions struct {
	Headers    map[string]string
	Body       interface{}
	Params     map[string]string
	JSONEncode bool
}

func Request(method, url string, options RequestOptions) (*http.Response, error) {
	var bodyReader *bytes.Reader

	// Se houver um corpo na requisição e a opção JSONEncode estiver ativada, codificar para JSON.
	if options.Body != nil && options.JSONEncode {
		jsonBody, err := json.Marshal(options.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	} else if options.Body != nil {
		// Se houver um corpo na requisição, converte para string.
		bodyString, ok := options.Body.(string)
		if ok {
			bodyReader = bytes.NewReader([]byte(bodyString))
		}
	}

	// Criar uma requisição HTTP com o método, URL e corpo fornecidos.
	var req *http.Request
	var err error
	if bodyReader != nil {
		req, err = http.NewRequest(method, url, bodyReader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	// Adicionar cabeçalhos à requisição.
	for key, value := range options.Headers {
		req.Header.Add(key, value)
	}

	// Adicionar parâmetros de consulta à URL.
	query := req.URL.Query()
	for key, value := range options.Params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// Realizar a requisição HTTP.
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get realiza uma requisição GET para a URL fornecida com as opções especificadas.
func Get(url string, options RequestOptions) (*http.Response, error) {
	return Request(http.MethodGet, url, options)
}

// Post realiza uma requisição POST para a URL fornecida com as opções especificadas.
func Post(url string, options RequestOptions) (*http.Response, error) {
	return Request(http.MethodPost, url, options)
}

// Put realiza uma requisição PUT para a URL fornecida com as opções especificadas.
func Put(url string, options RequestOptions) (*http.Response, error) {
	return Request(http.MethodPut, url, options)
}

// Delete realiza uma requisição DELETE para a URL fornecida com as opções especificadas.
func Delete(url string, options RequestOptions) (*http.Response, error) {
	return Request(http.MethodDelete, url, options)
}
