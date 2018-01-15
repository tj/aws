// Package logs implements AWS CloudWatchLogs tailing.
package logs

import (
	"time"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
)

// Event is a single log event from a group.
type Event struct {
	Timestamp time.Time
	GroupName string
	Message   string
}

// Config options.
type Config struct {
	Service       cloudwatchlogsiface.CloudWatchLogsAPI
	FilterPattern string
	PollInterval  time.Duration
	StartTime     time.Time
	Follow        bool
	GroupNames    []string
}

// Logs tailer.
type Logs struct {
	Config
	err error
}

// New log tailer with config.
func New(c Config) *Logs {
	return &Logs{
		Config: c,
	}
}

// Start consuming logs.
func (l *Logs) Start() <-chan *Event {
	ch := make(chan *Event)
	done := make(chan error)

	for _, name := range l.GroupNames {
		go l.consume(name, ch, done)
	}

	go func() {
		l.wait(done)
		close(ch)
	}()

	return ch
}

// Err returns the error, if any, during processing.
func (l *Logs) Err() error {
	return l.err
}

// wait for each log group to complete.
func (l *Logs) wait(done <-chan error) {
	for range l.GroupNames {
		if err := <-done; err != nil {
			l.err = err
			return
		}
	}
}

// consume logs for group `name`.
func (l *Logs) consume(name string, ch chan *Event, done chan error) {
	log := group{
		Config: l.Config,
		name:   name,
		log:    log.WithField("group", name),
	}

	for event := range log.Start() {
		ch <- event
	}

	done <- log.Err()
}
