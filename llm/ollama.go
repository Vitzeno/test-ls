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
	preText = `
	You are a large language model designed to debug code. 
	You are given a code snippet and an error. 
	Your task is to provide a one sentense response to the error. 
	The error is: %s
	The code snippet is: %s

	What is a potentail solution?
	`
)

func Prompt(errorText, codeText string) (string, error) {
	req := Request{
		Model:  "mistral",
		Prompt: fmt.Sprintf(preText, errorText, codeText),
		Stream: types.P(false),
	}

	j, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %w", err)
	}

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
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
