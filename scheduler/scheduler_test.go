package scheduler

import (
	"testing"
)

// TestSlackMessage tests parsing a slackAlert struct into a message
func TestSlackMessage(t *testing.T) {

	slackAlert := &SlackAlert{
		users:    []string{"Thomas", "Bob"},
		channels: []string{"general", "random"},
	}
	expected := "@Thomas@Bob\n\nTest Message"

	actual := slackAlert.MakeAlertMessage("Test Message")

	if expected != actual {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, actual)
	}
}
