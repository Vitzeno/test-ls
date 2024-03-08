package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Vitzeno/test-ls/types"
)

const (
	prePrompt = `
	You are a large language model designed to debug code. 
	You are given a code snippet and an error. 
	Your task is to provide a one sentense response to the error. 
	
	The error is: %s
	The code snippet is: %s

	What is a potential solution?
	`
)

type Ollama struct {
	Endpoint  string `json:"endpoint"`
	Model     string `json:"model"`
	PrePrompt string `json:"pre_prompt"`
}

func New(opts ...OllamaOption) *Ollama {
	o := Ollama{
		Endpoint:  "http://localhost:11434/api/generate",
		Model:     "mistral",
		PrePrompt: prePrompt,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return &o
}

type OllamaOption func(o *Ollama)

func WithEndpoint(endpoint string) OllamaOption {
	return func(o *Ollama) {
		o.Endpoint = endpoint
	}
}

func WithModel(model string) OllamaOption {
	return func(o *Ollama) {
		o.Model = model
	}
}

func WithPrePrompt(prePrompt string) OllamaOption {
	return func(o *Ollama) {
		o.PrePrompt = prePrompt
	}
}

func (o *Ollama) Prompt(errorText, codeText string) (string, error) {
	req := Request{
		Model:  o.Model,
		Prompt: fmt.Sprintf(o.PrePrompt, errorText, codeText),
		Stream: types.P(false),
	}

	j, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %w", err)
	}

	resp, err := http.Post(
		o.Endpoint,
		"application/json",
		bytes.NewReader(j),
	)
	if err != nil {
		return "", fmt.Errorf("error sending prompt: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	err = json.Unmarshal(buf, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	log.Printf("response: %+v", response.Response)

	return response.Response, nil
}
