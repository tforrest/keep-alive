package scheduler

import (
	"fmt"
	"github.com/nlopes/slack"
	"net/url"
)

// Scheduler cycles through a job list and excutes them async
type Scheduler struct {
	jobs []Job
}

// Job holds the information needed to keep an link alive and alert
type Job struct {
	jobType string
	alert   interface{}
	url     url.URL
}

type alert interface {
	SendSuccess(string) error
	SendFailure(string, string) error
}

// SlackAlert sends alerts via the slack api
type SlackAlert struct {
	users    []string
	channels []string
	api      *slack.Client
}

// NewSlackAlert returns a new slack alert struct with the proper fields
func NewSlackAlert(users []string, channels []string, apikey string) (*SlackAlert, error) {
	api := slack.New(apikey)
	if _, err := api.AuthTest(); err != nil {
		return nil, err
	}
	slackAlert := &SlackAlert{
		users:    users,
		channels: channels,
		api:      api,
	}
	return slackAlert, nil
}

// MakeAlertMessage creates the tag portion of the slack alert
func (s *SlackAlert) MakeAlertMessage(content string) string {
	var userTags string
	for _, u := range s.users {
		userTags += fmt.Sprintf("@%v", u)
	}
	return fmt.Sprintf("%v\n\n%v", userTags, content)
}

func (s *SlackAlert) sendMessage(message string) error {
	// iterate through each channel and post message
	channelMessage := s.MakeAlertMessage(message)
	slackParams := slack.NewPostMessageParameters()
	for _, c := range s.channels {
		if _, _, err := s.api.PostMessage(c, channelMessage, slackParams); err != nil {
			return err
		}
	}
	return nil
}

// SendSuccess lets slack know of a success
func (s *SlackAlert) SendSuccess(serviceName string) error {
	message := fmt.Sprintf("Service: %v is up and running!", serviceName)
	if err := s.sendMessage(message); err != nil {
		return err
	}
	return nil
}

// SendFailure lets slack know of a failure
func (s *SlackAlert) SendFailure(serviceName string, failureReason string) error {
	message := fmt.Sprintf("Service: %v is down! Error Message: %v", serviceName, failureReason)
	if err := s.sendMessage(message); err != nil {
		return err
	}
	return nil
}

// TwilioAlert sends alerts via the twilio api
type TwilioAlert struct {
	sid     string
	numbers []string
	apiKey  string
}

// SendSuccess lets twilio know of a success
func (t *TwilioAlert) SendSuccess(message string) error {
	return nil
}

// SendFailure lets twilio know of a failure
func (t *TwilioAlert) SendFailure(message string) error {
	return nil
}
