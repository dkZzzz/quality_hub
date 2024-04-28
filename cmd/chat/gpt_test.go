package chat

import (
	"testing"
)

func TestChat(t *testing.T) {
	code := "some code here"
	message := "some problem description here"

	reply, err := Chat(code, message)
	if err != nil {
		t.Errorf("Chat() returned an error: %v", err)
	}

	if reply == "" {
		t.Error("Chat() returned an empty reply")
	}

	t.Logf("Chat() test passed successfully")
}
