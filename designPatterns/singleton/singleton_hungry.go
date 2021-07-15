package singleton

var hungryInstance = &HungrySingleton{}

// 单例模式-饿汉式
type HungrySingleton struct {
}

func GetHungryInstance() *HungrySingleton {
	return hungryInstance
}
