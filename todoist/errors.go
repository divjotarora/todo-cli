package todoist

import "fmt"

func newError(wrapped error) error {
	return fmt.Errorf("todoist API error: %w", wrapped)
}
