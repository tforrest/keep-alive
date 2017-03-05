package scheduler

import (
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
	sendSuccess(string) error
	sendFailure(string) error
}

// SlackAlert sends alerts via the slack api
type SlackAlert struct {
	userName string
	channel  string
	mentions []string
	apiKey   string
}

// SendSuccess lets slack know of a success
func (s *SlackAlert) SendSuccess(message string) error {
	return nil
}

// SendFailure lets slack know of a failure
func (s *SlackAlert) SendFailure(message string) error {
	return nil
}

// TwilioAlert sends alerts via the twilio api
type TwilioAlert struct {
	sid     string
	numbers []string
	apiKey  string
}

// SendSuccess lets twilio know of success
func (t *TwilioAlert) SendSuccess(message string) error {
	return nil
}

// SendFailure lets twilio know of a failure
func (t *TwilioAlert) SendFailure(message string) error {
	return nil
}
