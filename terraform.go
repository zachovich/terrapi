package terrapi

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
)

const pkgName = "terrapi"

type Terraform struct {
	// If not set, then assume searching PATH
	TerraformBinPath string

	// If not set, then assume current directory
	CodePath string

	// Plan file
	PlanPath string

	// A series of actions in queue that will be executed in
	// order and sequentially by Exec.
	queue []Action

	// Next action for execution
	head int

	//output map[Action][]byte

	// Stdin, Stdout, and Stderr example to test docs
	// TODO:
	//Stdin  io.Reader
	//Stdout io.Writer
	//Stderr io.Writer

	// Env
	// TODO: corresponding to exec.Cmd Env field

	Ctx context.Context

	//
	// StateFile
}

type (
	OptionType string
	Option     func(*Terraform) error
)

func NewTerraform(opts ...Option) (*Terraform, error) {
	t := new(Terraform)

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if err := opt(t); err != nil {
			return nil, &FunctionalOptionError{Opt: reflect.TypeOf(opt).Kind().String()}
		}
	}

	// defaults
	t.TerraformBinPath = "terraform"
	t.Ctx = context.Background()

	return t, nil
}

func WithActions(actions ...Action) Option {
	return func(t *Terraform) error {
		t.queue = append(t.queue, actions...)
		return nil
	}
}

const terraformBin = "terraform"

func (t *Terraform) AmendToQueue(a Action) {
	t.queue = append(t.queue, a)
}

// Head returns the index of the next action to be executed
func (t *Terraform) Head() int {
	return t.head
}

// Reset resets the head to '0'
func (t *Terraform) Reset() {
	t.head = 0
}

// SetHead manually sets the head to the next action to be executed.
// Head starts count from '0', so, the first action to be executed in
// the action queue is indexed at '0'
func (t *Terraform) SetHead(i int) error {
	if i > len(t.queue) {
		return errors.New("index bigger than current action queue size")
	}

	t.head = i

	return nil
}

// Exec will run actions in Terraform queue sequentially and update
// action output and error fields. If an action fails, Exec returns
// stdout and stderr of the action execution.
//
// Running Exec again will continue from the last failed
// action or the first newly amended action.
func (t *Terraform) Exec() ([]byte, []byte, error) {
	if len(t.queue) == 0 {
		return nil, nil, ErrNoRegisteredActions
	}

	log.Debugf(logMsg("queue length is: %d"), len(t.queue))
	log.Infof(logMsg("start queue execution at index: %d"), t.head)

	var outBuff, errBuff bytes.Buffer

	for _, j := range t.queue[t.head:] {
		u, _ := j.unmarshal()

		log.Debugf(
			logMsg("Action index: %d Terraform operation: %s Args: %s"),
			t.head, u[0], strings.Join(u[1:], " "),
		)

		e := exec.CommandContext(t.Ctx, terraformBin, u...)
		e.Dir = t.CodePath
		e.Stdout = &outBuff
		e.Stderr = &errBuff

		err := e.Run()

		// recording action output before returning
		t.queue[t.head].setOutErr(outBuff.Bytes(), errBuff.Bytes())

		if err != nil {
			return outBuff.Bytes(), errBuff.Bytes(), err
		}

		// move the head to the next action in the queue
		t.head++
	}

	return outBuff.Bytes(), errBuff.Bytes(), nil
}
