package internal

import (
	"os"
)

// Stdrwc wraps os.Stdin to satisfy io.ReadWriteCloser
type Stdrwc struct{}

// var _ io.ReadWriteCloser = (*Stdrwc)(nil)

func (Stdrwc) Read(p []byte) (int, error) {
	//log.Println("Reading from stdin: ", string(p))
	return os.Stdin.Read(p)
}

func (Stdrwc) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

// func (Stdrwc) Close() error {
// 	if err := os.Stdin.Close(); err != nil {
// 		return err
// 	}
// 	return os.Stdout.Close()
// }

func (s Stdrwc) Close() error {
	return nil // os.Stdin cannot be closed
}
