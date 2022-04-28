package logging

type Logger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

var _ Logger = (*ConsoleLogger)(nil) //insuring at compile time that ConsoleLogger implements Logger

type ConsoleLogger struct {
}

func (c *ConsoleLogger) Info(args ...interface{}) {

}
func (c *ConsoleLogger) Debug(args ...interface{}) {

}
func (c *ConsoleLogger) Error(args ...interface{}) {

}
func (c *ConsoleLogger) Fatal(args ...interface{}) {

}
