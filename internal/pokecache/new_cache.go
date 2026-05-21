package pokecache

import (
	"time"
)

type Cache struct {
	createdAt time.Time
	val       []byte
}
