/*
Package mdc Map Diagnostic Context, useing for append contextual data to an event.
*/
package mdc

// IMDC Map diagnostic data interface
type IMDC interface {
	//Get return MDC as map[string]string
	Get() map[string]string
	//Add adds new key value to an existent map
	Add(key, val string)
}

//Context IMDC interface implementation
type Context struct {
	data map[string]string
}

// NewContext return a new Context
// TODO: Need another constractor from existen map.
func NewContext() *Context {
	m := Context{
		data: make(map[string]string),
	}
	return &m
}

// Get interface implementation
func (m *Context) Get() map[string]string {
	return m.data
}

// Add interface implementation
func (m *Context) Add(key, val string) {
	m.data[key] = val
}
