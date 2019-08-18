package devicetype

import "fmt"

type NotRecognised struct{}

func (e *NotRecognised) Error() string {
	return fmt.Sprintf("The device type provided was not recognised.")
}
