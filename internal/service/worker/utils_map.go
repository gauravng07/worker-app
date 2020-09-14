package worker

import "sort"

func rankMapStringInt(values map[string]int) (string, int) {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	var postCode string
	var count int
	for i, kv := range ss {
		if i == 0 {
			postCode = kv.Key
			count = kv.Value
		}
	}
	return postCode, count
}
