package factoryMethod

type AbstractFactory interface {
	GetInstanceFactory(name string) Handler
}

type ConcreteFactory struct {
}

func NewConcreteFactory() *ConcreteFactory {
	return &ConcreteFactory{}
}

func (s *ConcreteFactory) GetInstanceFactory(name string) Handler {
	switch name {
	case "cache":
		return NewCacheHandler()
	case "db":
		return NewDBHandler()
	default:
		return nil
	}
}

type Handler interface {
	process()
}

type CacheHandler struct {
}

func NewCacheHandler() *CacheHandler {
	return &CacheHandler{}
}

func (s *CacheHandler) process() {
	panic("implement me")
}

type DBHandler struct {
}

func NewDBHandler() *DBHandler {
	return &DBHandler{}
}

func (s *DBHandler) process() {
	panic("implement me")
}
