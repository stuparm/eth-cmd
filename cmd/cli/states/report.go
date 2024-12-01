package states

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Reporter interface {
	WriteBlockStates(bs *BlockStates) error
}

type consoleReporter struct{}

func NewConsoleReporter() Reporter {
	return &consoleReporter{}
}

func (cr *consoleReporter) WriteBlockStates(bs *BlockStates) error {
	for address, i := range bs.balanceCounter {
		fmt.Println("Address:", address.Hex(), "Balance:", i)
	}
	for address, i := range bs.nonceCounter {
		fmt.Println("Address:", address.Hex(), "Nonce:", i)
	}
	for address, i := range bs.codesCounter {
		fmt.Println("Address:", address.Hex(), "Code:", i)
	}
	for address, storage := range bs.storageCounter {
		for key, i := range storage {
			fmt.Println("Address:", address.Hex(), "Storage:", key.Hex(), i)
		}
	}

	return nil
}

type fileReporter struct {
	filePath string
}

func NewFileReporter(filePath string) Reporter {
	return &fileReporter{filePath: filePath}
}

func (fr *fileReporter) WriteBlockStates(bs *BlockStates) error {
	var sb strings.Builder

	for address, i := range bs.balanceCounter {
		sb.WriteString(fmt.Sprintf("Address: %s Balance: %d\n", address.Hex(), i))
	}

	for address, i := range bs.nonceCounter {
		sb.WriteString(fmt.Sprintf("Address: %s Nonce: %d\n", address.Hex(), i))
	}

	for address, i := range bs.codesCounter {
		sb.WriteString(fmt.Sprintf("Address: %s Code: %d\n", address.Hex(), i))
	}

	for address, storage := range bs.storageCounter {
		for key, i := range storage {
			sb.WriteString(fmt.Sprintf("Address: %s Storage: %s %d\n", address.Hex(), key.Hex(), i))
		}
	}
	output := sb.String()

	file, err := os.Create(fr.filePath)
	if err != nil {
		return errors.Wrap(err, "failed to create file")
	}
	defer file.Close()

	_, err = file.WriteString(output)
	if err != nil {
		return errors.Wrap(err, "failed to write to file")
	}

	return nil

}

func NewReporter(output string) Reporter {
	if output == "console" || output == "" {
		return NewConsoleReporter()
	}

	return NewFileReporter(output)
}
