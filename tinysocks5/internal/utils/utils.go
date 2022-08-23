package utils

import (
	"bytes"
	"sync"
)

var (
	SPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 576)
		},
	}

	LPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 64*1024+262)
		},
	}
)
