package prototype

// cloneable接口
type Cloneable interface {
	Clone() Cloneable
}

// 原型类，实现cloneable接口
type Handler struct {
	Name string
}

func NewHandler(name string) *Handler {
	return &Handler{Name: name}
}

func (s *Handler) GetName() string {
	return s.Name
}

func (s *Handler) SetName(name string) {
	s.Name = name
}

func (s *Handler) Clone() Cloneable {
	cloneHandler := s
	return cloneHandler
}

// 原型管理类
type HandlerManager struct {
	list map[string]Handler
}

func NewManager() *HandlerManager {
	return &HandlerManager{
		list: make(map[string]Handler, 0),
	}
}

func (s *HandlerManager) Get(name string) Handler {
	return s.list[name]
}

func (s *HandlerManager) Set(name string, prototype Handler) {
	s.list[name] = prototype
}

