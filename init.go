package logger

import "gitlab.appsflyer.com/Architecture/af-logger-go/impl"

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
