package states

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"sort"
	"strings"
)

type kv struct {
	key   any
	value int
}

func SortedAddressCounter(counter map[common.Address]int) string {
	kvs := make([]kv, 0, len(counter))
	for k, v := range counter {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value > kvs[j].value
	})

	sb := strings.Builder{}
	for _, kv := range kvs {
		sb.WriteString(fmt.Sprintf("Address: %s, Count: %d\n", kv.key, kv.value))
	}
	return sb.String()
}
