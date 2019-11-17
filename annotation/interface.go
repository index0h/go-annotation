package annotation

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Generator interface {
	Process(storage *Storage)
}

type AnnotationParserInterface interface {
	Parse(string) []interface{}
}
