// Package scheduler provides the core interface that Empire uses when
// interacting with a cluster of machines to run tasks.
package scheduler

import (
	"encoding/json"
	"io"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/remind101/empire/pkg/image"
	"github.com/remind101/pkg/logger"
)

type App struct {
	// The id of the app.
	ID string

	// An identifier that represents the version of this release.
	Release string

	// The name of the app.
	Name string

	// The application environment.
	Env map[string]string

	// The application labels.
	Labels map[string]string

	// Process that belong to this app.
	Processes []*Process
}

type Process struct {
	// The type of process.
	Type string

	// The Image to run.
	Image image.Image

	// The Command to run.
	Command []string

	// Environment variables to set.
	Env map[string]string

	// Labels to set on the container.
	Labels map[string]string

	// Exposure is the level of exposure for this process.
	Exposure *Exposure

	// Instances is the desired instances of this service to run.
	Instances uint

	// The amount of RAM to allocate to this process in bytes.
	MemoryLimit uint

	// The amount of CPU to allocate to this process, out of 1024. Maps to
	// the --cpu-shares flag for docker.
	CPUShares uint

	// ulimit -u
	Nproc uint
}

// Exposure controls the exposure settings for a process.
type Exposure struct {
	// External means that this process will be exposed to internet facing
	// traffic, as opposed to being internal. How this is used is
	// implementation specific. For ECS, this means that the attached ELB
	// will be "internet-facing".
	External bool

	// The exposure type (e.g. HTTPExposure, HTTPSExposure, TCPExposure).
	Type ExposureType
}

// Exposure represents a service that a process exposes, like HTTP/HTTPS/TCP or
// SSL.
type ExposureType interface {
	Protocol() string
}

// HTTPExposure represents an HTTP exposure.
type HTTPExposure struct{}

func (e *HTTPExposure) Protocol() string { return "http" }

// HTTPSExposure represents an HTTPS exposure
type HTTPSExposure struct {
	// The certificate to attach to the process.
	Cert string
}

func (e *HTTPSExposure) Protocol() string { return "https" }

// Instance represents an Instance of a Process.
type Instance struct {
	Process *Process

	// The instance ID.
	ID string

	// The State that this Instance is in.
	State string

	// The time that this instance was last updated.
	UpdatedAt time.Time
}

type Runner interface {
	// Run runs a process.
	Run(ctx context.Context, app *App, process *Process, in io.Reader, out io.Writer) error
}

// Scheduler is an interface for interfacing with Services.
type Scheduler interface {
	Runner

	// Submit submits an app, creating it or updating it as necessary.
	Submit(context.Context, *App, StatusStream) error

	// Remove removes the App.
	Remove(ctx context.Context, app string) error

	// Instance lists the instances of a Process for an app.
	Instances(ctx context.Context, app string) ([]*Instance, error)

	// Stop stops an instance. The scheduler will automatically start a new
	// instance.
	Stop(ctx context.Context, instanceID string) error
}

// Env merges the App environment with any environment variables provided
// in the process.
func Env(app *App, process *Process) map[string]string {
	return merge(app.Env, process.Env)
}

// Labels merges the App labels with any labels provided in the process.
func Labels(app *App, process *Process) map[string]string {
	return merge(app.Labels, process.Labels)
}

// merges the maps together, favoring keys from the right to the left.
func merge(envs ...map[string]string) map[string]string {
	merged := make(map[string]string)
	for _, env := range envs {
		for k, v := range env {
			merged[k] = v
		}
	}
	return merged
}

type Status struct {
	// A friendly human readable message about the status change.
	Message string
}

// String implements the fmt.Stringer interface.
func (s *Status) String() string {
	return s.Message
}

// StatusStream is an interface for publishing status updates while a scheduler
// is executing.
type StatusStream interface {
	// Publish publishes an update to the status stream
	Publish(context.Context, Status) error

	// Done finalizes the status stream
	Done(error)

	// Wait returns a channel that receives once Done() is called. Consumers
	// should call the Err() method to determine if an error occurred.
	Wait() <-chan struct{}

	// Returns the error from calling Done().
	Err() error
}

// jsonmessageStatusStream implements the StatusStream interface with support
// for writing jsonmessages to the provided io.Writer
type jsonmessageStatusStream struct {
	sync.Mutex
	done chan struct{}
	err  error
	w    io.Writer
}

// NewJSONMessageStream returns a new instance of the default status stream.
func NewJSONMessageStream(w io.Writer) StatusStream {
	return &jsonmessageStatusStream{w: w, done: make(chan struct{}, 1)}
}

func (s *jsonmessageStatusStream) Publish(ctx context.Context, status Status) error {
	select {
	case <-s.done:
		logger.Warn(ctx, "Publish called on a finalized stream")
		return nil
	default:
	}
	return json.NewEncoder(s.w).Encode(jsonmessage.JSONMessage{Status: status.Message})
}

func (s *jsonmessageStatusStream) Done(err error) {
	s.Lock()
	defer s.Unlock()
	close(s.done)
	s.err = err
}

func (s *jsonmessageStatusStream) Err() error {
	return s.err
}

func (s *jsonmessageStatusStream) Wait() <-chan struct{} {
	return s.done
}
