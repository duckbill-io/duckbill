package models

import (
	"sort"
)

type byCreatedAt Posts

func (b byCreatedAt) Len() int {
	return len(b)
}

func (b byCreatedAt) Less(i, j int) bool {
	if b[i].CreatedAt != b[j].CreatedAt {
		return b[i].CreatedAt > b[j].CreatedAt
	}
	if b[i].Name != b[j].Name {
		return b[i].Name < b[j].Name
	}
	return false
}

func (b byCreatedAt) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byCreatedAt) sort() {
	sort.Sort(b)
}
