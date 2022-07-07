package _map

import "sync"

func NewMap() *sync.Map {
	var storage sync.Map
	return &storage
}
