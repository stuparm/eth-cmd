package states

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Reporter interface {
	WriteBlockStates(pre, post *BlockStates) error
	WriteSummary(summary string) error
}

type consoleReporter struct{}

func NewConsoleReporter() Reporter {
	return &consoleReporter{}
}

func (cr *consoleReporter) WriteBlockStates(pre, post *BlockStates) error {
	for address, i := range post.balanceCounter {
		fmt.Println("Address:", address.Hex(), "Balance:", i)
	}
	for address, i := range post.nonceCounter {
		fmt.Println("Address:", address.Hex(), "Nonce:", i)
	}
	for address, i := range post.codesCounter {
		fmt.Println("Address:", address.Hex(), "Code:", i)
	}
	for address, storage := range post.storageCounter {
		for key, i := range storage {
			fmt.Println("Address:", address.Hex(), "Storage:", key.Hex(), i)
		}
	}

	fmt.Println("Touched contracts")
	for address, _ := range pre.codesCounter {
		fmt.Println("Address:", address.Hex())
	}

	return nil
}

func (cr *consoleReporter) WriteSummary(summary string) error {
	fmt.Println(summary)

	return nil
}

type fileReporter struct {
	filePath string
}

func NewFileReporter(filePath string) Reporter {
	return &fileReporter{filePath: filePath}
}

func (fr *fileReporter) WriteBlockStates(pre, post *BlockStates) error {
	var sb strings.Builder

	for address, i := range post.balanceCounter {
		sb.WriteString(fmt.Sprintf("Address: %s Balance: %d\n", address.Hex(), i))
	}

	for address, i := range post.nonceCounter {
		sb.WriteString(fmt.Sprintf("Address: %s Nonce: %d\n", address.Hex(), i))
	}

	for address, i := range post.codesCounter {
		sb.WriteString(fmt.Sprintf("Address: %s Code: %d\n", address.Hex(), i))
	}

	for address, storage := range post.storageCounter {
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

func (fr *fileReporter) WriteSummary(summary string) error {
	file, err := os.Create(fr.filePath)
	if err != nil {
		return errors.Wrap(err, "failed to create file")
	}
	defer file.Close()

	_, err = file.WriteString(summary)
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
