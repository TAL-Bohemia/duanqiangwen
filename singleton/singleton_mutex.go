package singleton

import (
	"sync"
)

var instance *Singleton
var once sync.Once

// 线程安全的单例模式
type Singleton struct {
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
