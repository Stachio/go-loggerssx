package logger

type NoLog struct {
}

func Void() *NoLog {
	return &NoLog{}
}

// Write - Void Write function
func (no *NoLog) Write(p []byte) (n int, err error) {
	return 0, nil
}
