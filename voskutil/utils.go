package voskutil

import (
	"fmt"
)

func HandleVoskMessage(msg any) *string {
	// currPartial := []string{}
	// Try to parse as JSON object
	if m, ok := msg.(map[string]any); ok {
		if text, ok := m["text"].(string); ok && text != "" {
			fmt.Printf("\nFinal: %s\n", text)
			return &text

		} else if partial, ok := m["partial"].(string); ok && partial != "" {
			fmt.Printf("\rPartial: %s", partial)
		}
	} else if str, ok := msg.(string); ok {
		fmt.Printf("Message: %s\n", str)
	}
	return nil
}
