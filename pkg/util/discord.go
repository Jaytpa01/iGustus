package util

import "fmt"

func FormatUserMention(userID string) string {
	return fmt.Sprintf("<@%s>", userID)
}
