package email

import (
	"testing"
)

func TestSendText(t *testing.T) {
	err := SendText("mail.domain.com", 25, "", "", []string{"canuxcheng@gmail.com"}, "go-devops@domain.com", "", "subject", "body")
	if err != nil {
		t.Error("failed")
	}
}
