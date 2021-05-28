package filescan

import "fmt"

type HTTPResponseStatusError struct {
	code int
	msg  string
}

func (e *HTTPResponseStatusError) Error() string {
	return fmt.Sprintf("%d %s", e.code, e.msg)
}

type ChecksumGenerationError struct{}

func (e *ChecksumGenerationError) Error() string {
	return "Error computing the checksum"
}
