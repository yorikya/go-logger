package mdc

type IMDC interface {
	Get() map[string]string
	Add(key, val string)
}

type Context struct {
	data map[string]string
}

// TODO: Need another constractor from existen map.
func NewContext() *Context {
	m := Context{
		data: make(map[string]string),
	}
	return &m
}

func (m *Context) Get() map[string]string {
	return m.data
}

func (m *Context) Add(key, val string) {
	m.data[key] = val
}
