package scheduler

import (
	"fmt"
	"net/url"
	"time"

	"github.com/nlopes/slack"
	"github.com/sfreiberg/gotwilio"
)

// Scheduler cycles through a job list and excutes them async
type Scheduler struct {
	jobs []Job
}

// AddJob adds an additional job to the scheduler to check
func (s *Scheduler) AddJob(job Job) error {
	return nil
}

// RemoveJob removes a current job in the scheduler or returns an error if not present
func (s *Scheduler) RemoveJob(jobName string) error {
	return nil
}

// Start starts the scheduler and sends requests based on each jobs' interval
func (s *Scheduler) Start() error {
	return nil
}

// Job holds the information needed to keep an link alive and alert
type Job struct {
	JobName  string
	Interval time.Time
	Alerts   []alert
	URL      url.URL
}

type alert interface {
	SendSuccess(string) error
	SendFailure(string, string) error
}

// SlackAlert sends alerts via the slack api
type SlackAlert struct {
	users    []string
	channels []string
	slackAPI *slack.Client
}

// NewSlackAlert returns a new slack alert struct with the proper fields
func NewSlackAlert(users, channels []string, apikey string) (*SlackAlert, error) {
	api := slack.New(apikey)
	if _, err := api.AuthTest(); err != nil {
		return nil, err
	}
	slackAlert := &SlackAlert{
		users:    users,
		channels: channels,
		slackAPI: api,
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
		if _, _, err := s.slackAPI.PostMessage(c, channelMessage, slackParams); err != nil {
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
func (s *SlackAlert) SendFailure(serviceName, failureReason string) error {
	message := fmt.Sprintf("Service: %v is down! Error Message: %v", serviceName, failureReason)
	if err := s.sendMessage(message); err != nil {
		return err
	}
	return nil
}

// TwilioAlert sends alerts via the twilio api
type TwilioAlert struct {
	sender    string
	numbers   []string
	twilioAPI *gotwilio.Twilio
}

// NewTwilioAlert create a TwilioAlert struct with the proper fields
func NewTwilioAlert(sid, sender string, numbers []string, apikey string) *TwilioAlert {
	api := gotwilio.NewTwilioClient(sid, apikey)

	twilioAlert := &TwilioAlert{
		sender:    sender,
		numbers:   numbers,
		twilioAPI: api,
	}

	return twilioAlert
}

func (t *TwilioAlert) sendTextMessages(message string) error {
	for _, number := range t.numbers {
		_, _, err := t.twilioAPI.SendSMS(t.sender, number, message, "", "")
		if err != nil {
			return err
		}
	}
	return nil
}

// SendSuccess lets twilio know of a success
func (t *TwilioAlert) SendSuccess(serviceName string) error {
	message := fmt.Sprintf("Service: %v is up!", serviceName)
	return t.sendTextMessages(message)
}

// SendFailure lets twilio know of a failure
func (t *TwilioAlert) SendFailure(serviceName, failureReason string) error {
	message := fmt.Sprintf("Service: %v is down! Error message: %v", serviceName, failureReason)
	return t.sendTextMessages(message)
}
