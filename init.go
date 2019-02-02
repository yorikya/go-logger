package logger

import "github.com/yorikya/go-logger/impl"

var (
	defaultLoggerFactory LoggerFactory
)

func init() {
	defaultLoggerFactory = newLoggerFactory()
}

func newLoggerFactory() *nativeLoggerFactory {
	factory := &nativeLoggerFactory{}
	return factory
}

type nativeLoggerFactory struct{}

func (nativeLoggerFactory) GetLogger(name string) Logger {
	return impl.DefaultLogger(name)
}

func GetLogger(name string) Logger {
	return defaultLoggerFactory.GetLogger(name)
}
