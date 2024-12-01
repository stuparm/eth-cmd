package states

import "fmt"

type Summarizer interface {
	Summarize(states *BlockStates) string
}

type summarizer struct{}

func NewSummarizer() Summarizer {
	return &summarizer{}
}

func (s *summarizer) Summarize(states *BlockStates) string {
	uniqueStates := len(states.balanceCounter) +
		len(states.nonceCounter) +
		len(states.codesCounter)
	for _, storages := range states.storageCounter {
		uniqueStates += len(storages)
	}

	totalStates := 0
	for _, v := range states.balanceCounter {
		totalStates += v
	}

	for _, v := range states.nonceCounter {
		totalStates += v
	}

	for _, v := range states.codesCounter {
		totalStates += v
	}

	for _, storage := range states.storageCounter {
		for _, v := range storage {
			totalStates += v
		}
	}

	ratio := 1 - float64(uniqueStates)/float64(totalStates)
	return fmt.Sprintf("Unique states: %d, Total states: %d, Ratio: %.2f", uniqueStates, totalStates, ratio)
}
