package command

import "fmt"

func newValidationError(cmd Command, err error) error {
	return fmt.Errorf("validation error for command %s: %w", cmd.Name(), err)
}

func invalidNumberOfArguments(cmd Command, expectedNumArgs, actualNumArgs int) error {
	wrapped := fmt.Errorf("expected %d arguments, got %d", expectedNumArgs, actualNumArgs)
	return newValidationError(cmd, wrapped)
}

func newExecutionError(cmd Command, err error) error {
	return fmt.Errorf("execution error for command %s: %w", cmd.Name(), err)
}
