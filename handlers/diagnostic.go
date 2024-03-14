package handlers

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Vitzeno/test-ls/llm"
	"github.com/Vitzeno/test-ls/types"
	"github.com/vitzeno/llvm-test/parser"
)

func GetDiagnostics(fileURI string) ([]Diagnostic, error) {
	var diagnostics []Diagnostic

	// Remove the "file:///" prefix
	filePath := fileURI[7:]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return diagnostics, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure the file is closed

	parser.YYDebug = 0
	lexer := parser.NewLexer(file)
	errCount := parser.YYParse(lexer)
	if errCount != 0 {
		log.Println("parsing failed found error(s) in source file")
		//return diagnostics, fmt.Errorf("parsing failed found error(s) in source file")
	}

	log.Println("parsing file")
	for _, lexErr := range lexer.Errors() {
		log.Println("Error:", lexErr)
		diagnostics = append(diagnostics, Diagnostic{
			Range: Range{
				Start: Position{
					Line:      lexErr.Line - 1,
					Character: lexErr.Col,
				},
			},
			Severity: types.P(DiagError),
			Message:  lexErr.Message,
		})
	}

	return diagnostics, nil
}

func GetHover(fileURI string, pos Position) (HoverResponse, error) {
	var hoverResponse HoverResponse

	// Remove the "file:///" prefix
	filePath := fileURI[7:]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file for hover:", err)
		return hoverResponse, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure the file is closed

	lineNumber := pos.Line + 1

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		return hoverResponse, fmt.Errorf("error reading file: %w", err)
	}

	var hoverLine string

	// Split the file content into lines
	lines := strings.Split(string(fileContent), "\n")

	// Retrieve the requested line
	if lineNumber >= 1 && lineNumber <= len(lines) {
		hoverLine = lines[lineNumber-1]
	} else {
		log.Println("Line number out of range")
		return hoverResponse, fmt.Errorf("line number out of range")
	}

	ollama := llm.New(llm.WithExplainPrompt())
	resp, err := ollama.Prompt(hoverLine, string(fileContent))
	if err != nil {
		log.Printf("error getting llm response: %v", err)
		return hoverResponse, fmt.Errorf("error getting llm response: %w", err)
	}

	hoverResponse = HoverResponse{
		Contents: MarkupContent{
			Kind: MarkupKind{
				PlainText: "plaintext",
			},
			Value: resp,
		},
	}

	return hoverResponse, nil
}
