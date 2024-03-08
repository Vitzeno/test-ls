package handlers

import (
	"fmt"
	"log"
	"os"

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
