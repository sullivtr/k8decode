package loader

import (
	"time"

	"github.com/briandowns/spinner"
)

// NewSpinner creates a new instance of a spinner
func NewSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Color("green")
	s.Prefix = "[ DECODING SECRETS ] [ "
	s.Suffix = " ]"

	return s
}
