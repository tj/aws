package logs

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// group implements log fetching and polling for
// a single CloudWatchLogs group.
type group struct {
	Config
	name string
	log  log.Interface
	err  error
}

// Start consuming logs.
func (g *group) Start() <-chan *Event {
	ch := make(chan *Event)
	go g.start(ch)
	return ch
}

// start consuming and exit after pagination if Follow is not enabled.
func (g *group) start(ch chan<- *Event) {
	defer close(ch)

	g.log.Debug("enter")
	defer g.log.Debug("exit")

	var start = g.StartTime.UnixNano() / int64(time.Millisecond)
	var nextToken *string
	var err error

	for {
		g.log.WithField("start", start).Debug("request")
		nextToken, start, err = g.fetch(nextToken, start, ch)

		if e, ok := err.(awserr.Error); ok {
			if e.Code() == "ThrottlingException" {
				g.log.Debug("throttled")
				time.Sleep(time.Second * 2)
				continue
			}
		}

		if err != nil {
			g.err = fmt.Errorf("log %q: %s", g.name, err)
			break
		}

		if nextToken == nil && g.Follow {
			time.Sleep(g.PollInterval)
			g.log.WithField("start", start).Debug("poll")
			continue
		}

		if nextToken == nil {
			break
		}
	}
}

// fetch logs relative to the given token and start time. We ignore when the log group is not found.
func (g *group) fetch(nextToken *string, start int64, ch chan<- *Event) (*string, int64, error) {
	latest := start

	res, err := g.Service.FilterLogEvents(&cloudwatchlogs.FilterLogEventsInput{
		LogGroupName:  &g.name,
		FilterPattern: &g.FilterPattern,
		StartTime:     &start,
		NextToken:     nextToken,
	})

	if e, ok := err.(awserr.Error); ok {
		if e.Code() == "ResourceNotFoundException" {
			g.log.Debug("not found")
			return nil, 0, nil
		}
	}

	if err != nil {
		return nil, 0, err
	}

	for _, event := range res.Events {
		ts := *event.Timestamp

		if ts > latest {
			latest = ts
		}

		sec := ts / 1000

		ch <- &Event{
			Timestamp: time.Unix(sec, 0),
			GroupName: g.name,
			Message:   *event.Message,
		}
	}

	if res.NextToken == nil {
		return nil, latest + 1, nil
	}

	return res.NextToken, start, nil
}

// Err returns the first error, if any, during processing.
func (g *group) Err() error {
	return g.err
}
