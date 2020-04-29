package cmd

import "fmt"

// Mistake is spelling mistake
type Mistake struct {
	Word string
	S    []string
}

func (m Mistake) String() string {
	return fmt.Sprintf("[word: %s, suggestions: %s]", m.Word, m.S)
}
