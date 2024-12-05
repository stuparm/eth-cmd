package states

import "fmt"

type Summarizer interface {
	Summarize(pre, post *BlockStates) string
}

type summarizer struct{}

func NewSummarizer() Summarizer {
	return &summarizer{}
}

func (s *summarizer) Summarize(pre, post *BlockStates) string {
	uniqueStates := len(post.balanceCounter) +
		len(post.nonceCounter) +
		len(post.codesCounter)
	for _, storages := range post.storageCounter {
		uniqueStates += len(storages)
	}

	totalStates := 0
	for _, v := range post.balanceCounter {
		totalStates += v
	}

	for _, v := range post.nonceCounter {
		totalStates += v
	}

	for _, v := range post.codesCounter {
		totalStates += v
	}

	for _, storage := range post.storageCounter {
		for _, v := range storage {
			totalStates += v
		}
	}

	uniqueLoadedContracts := SortedAddressCounter(pre.codesCounter)

	ratio := 1 - float64(uniqueStates)/float64(totalStates)
	return fmt.Sprintf("Unique states: %d, Total states: %d,\n"+
		"Ratio: %.2f\n"+
		"Unique loaded contracts: \n %s\n",
		uniqueStates, totalStates, ratio, uniqueLoadedContracts)
}
