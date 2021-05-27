package ohdear

import (
	"fmt"
	"strings"
)

type (
	APIError struct {
		Message string              `json:"message,omitempty"`
		Errors  map[string][]string `json:"errors,omitempty"`
	}
)

func (e *APIError) Error() string {
	causes := []string{}
	for key, cause := range e.Errors {
		causes = append(causes, fmt.Sprintf("%s: %s", key, strings.Join(cause, ", ")))
	}

	return fmt.Sprintf("%s, Causes: %s", e.Message, strings.Join(causes, ", "))
}
