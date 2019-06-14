package models

import (
	"sort"
)

//	=====================
//	=== Posts 的排序 ===
//	=====================

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

//	=====================
//	=== Tags 的排序 ===
//	=====================

type byName Tags

func (b byName) Len() int {
	return len(b)
}

func (b byName) Less(i, j int) bool {
	if b[i].Name != b[j].Name {
		return b[i].Name < b[j].Name
	}
	if b[i].Count != b[j].Count {
		return b[i].Count < b[j].Count
	}
	return false
}

func (b byName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byName) sort() {
	sort.Sort(b)
}
