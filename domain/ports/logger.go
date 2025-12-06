package ports

type Logger interface {
	Info(msg string, field map[string]interface{})
	Debug(msg string, field map[string]interface{})
	Warn(msg string, field map[string]interface{})
	Error(msg string, err error)
	ErrorF(msg string, err error, fields map[string]interface{})
}
