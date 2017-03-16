package sqlgo

import "fmt"

// Option defines
type Option interface {
	Apply(*db) error
}

// OptErr returns a new OptError
func OptErr(option string, err error, fmt ...string) *OptError {
	if len(fmt) == 0 {
		fmt[0] = "error with %s: %s"
	}

	return &OptError{
		fmt[0],
		option,
		err,
	}
}

// OptError wraps errors for options
type OptError struct {
	Fmt    string
	Option string
	Err    error
}

func (oe *OptError) Error() string {
	return fmt.Sprintf(oe.Fmt, oe.Option, oe.Err)
}
