package log

import "testing"

func TestLog(t *testing.T) {
	logger := NewLogger()
	logger.Println("a", "b")
}
