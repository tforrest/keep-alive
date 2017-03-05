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
	sendSuccess() error
	sendFailure() error
}

type slackAlert struct {
	userName string
	channel  string
	mentions []string
	apiKey   string
}

type twilioAlert struct {
	sid     string
	numbers []string
	apiKey  string
}

type emailAlert struct {
	emailAddress string
}
