package singleton

var lazyInstance *LazySingleton

// 单例模式-懒汉式
type LazySingleton struct {
}

func GetLazyInstance() *LazySingleton {
	if lazyInstance == nil {
		lazyInstance = &LazySingleton{}
	}

	return lazyInstance
}
