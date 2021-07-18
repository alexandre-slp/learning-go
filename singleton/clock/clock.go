package clock

import (
	"sync"
	"time"
)

type clockSingleton interface {
	Now() time.Time
}

type clock struct{
	clockSingleton
}

var clockInstance *clock
var once sync.Once

func New() *clock {
	once.Do(func() {
		clockInstance = &clock{}
	})
	return clockInstance
}
