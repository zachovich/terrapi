package terrapi

import (
	"strconv"
	"time"
)

type DestroyAction struct {
	name            string
	AutoApprove     bool     `cli:"-auto-approve"`
	Backup          string   `cli:"-backup="`
	Lock            string   `cli:"-lock="`
	LockTimeout     string   `cli:"-lock-timeout="`
	NoColor         bool     `cli:"-no-color"`
	Parallelism     string   `cli:"-parallelism="`
	Refresh         string   `cli:"-refresh="`
	State           string   `cli:"-state="`
	StateOut        string   `cli:"-state-out="`
	Targets         []string `cli:"-target="`
	Vars            []string `cli:"-var "`
	VarFile         string   `cli:"-var-file="`

	// outBuff, and errBuff holds the action output and error when executed
	OutBytes []byte
	ErrBytes []byte
}

func NewDestroyAction() *DestroyAction {
	return &DestroyAction{
		name: "destroy",
	}
}

func (a *DestroyAction) setOutErr(outBytes, errBytes []byte) {
	a.OutBytes = outBytes
	a.ErrBytes = errBytes
}

func (a *DestroyAction) GetOutErr() ([]byte, []byte) {
	return a.OutBytes, a.ErrBytes
}

func (a *DestroyAction) unmarshal() ([]string, error) {
	//TODO: struct validation
	return unmarshalAction(a)
}

// Skip interactive plan approval before applying.
// Terraform CLI: -auto-approve
func (a *DestroyAction) EnableAutoApprove() {
	a.AutoApprove = true
}

// Path to backup the existing state file before modifying.
// Defaults to the "-state-out" path with ".backup"
// Terraform CLI: -backup=path
func (a *DestroyAction) SetBackup(path string) {
	a.Backup = path
}

// Disable automatic backup of state file before Terraform updating it.
func (a *DestroyAction) DisableBackup() {
	a.Backup = "-"
}

// Disable locking of state files during state-related operations.
// Terraform CLI: -lock=true
func (a *DestroyAction) DisableLock() {
	a.Lock = "false"
}

// Override the time Terraform will wait to acquire a state lock.
// The default is 0s (zero seconds), which causes immediate failure
// if the lock is already held by another process.
// Terraform CLI: -lock-timeout=0s
func (a *DestroyAction) SetLockTimeout(duration time.Duration) {
	a.LockTimeout = duration.String()
}

// Set output to mono-color
// Terraform CLI: -no-color
func (a *DestroyAction) DisableOutputColor() {
	a.NoColor = true
}

// Set the number of parallel resource operations. Defaults to 10.
// Terraform CLI: -parallelism=n
func (a *DestroyAction) SetParallelism(n int) {
	a.Parallelism = strconv.Itoa(n)
}

// Disables auto update to state prior to checking for differences.
// This has no effect if a plan file is given to apply.
// Terraform CLI: -refresh=true
func (a *DestroyAction) DisableAutoRefresh() {
	a.Refresh = "false"
}

// Path to read and write state. Defaults to terraform.tfstate
// Terraform CLI: -state=path
func (a *DestroyAction) SetExistingStateFile(path string) {
	a.State = path
}

// Set a new file to write new terraform state. This can be used
// as a way to preserve old state file.
// Terraform CLI: -state-out=path
func (a *DestroyAction) SetNewStateFile(path string) {
	a.StateOut = path
}

// Limit operations to a set of resources and its dependencies.
// Terraform CLI: -target=resource -target=resource ...
func (a *DestroyAction) RegisterResourcesForApply(resources ...string) {
	a.Targets = append(a.Targets, resources...)
}

// Set variables in Terraform configuration.
// Terraform CLI: -var 'foo=bar' -var 'x=ray' ...
func (a *DestroyAction) SetVariables(vars ...string) {
	a.Vars = append(a.Vars, vars...)
}

// Set variables in Terraform configuration from this file.
// NOTE: if terrafrom.tfvars or any *.auto.tfvars files are present,
// they will be automatically loaded.
// Terraform CLI: -var-file=path
func (a *DestroyAction) SetVarFile(path string) {
	// TODO: Check if file exists and maybe is valid
	a.VarFile = path
}
