package application

// A Logger belong to the application layer.
type Logger interface {
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
