package voskutil

import (
	"fmt"
)

func HandleVoskMessage(msg any) *string {
	// currPartial := []string{}
	// Try to parse as JSON object
	if m, ok := msg.(map[string]any); ok {
		if text, ok := m["text"].(string); ok && text != "" {
			fmt.Print("\r\033[2K") // \033[2K clears entire line

			fmt.Printf("Final: %s\n", text)
			return &text

		} else if partial, ok := m["partial"].(string); ok && partial != "" {
			fmt.Printf("\rListening: %s", partial)
		}
	} else if str, ok := msg.(string); ok {
		fmt.Printf("Message: %s\n", str)
	}
	return nil
}
