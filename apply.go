package terrapi

import (
	"strconv"
	"time"
)

type ApplyAction struct {
	name            string
	AutoApprove     bool     `cli:"-auto-approve"`
	Backup          string   `cli:"-backup="`
	CompactWarnings bool     `cli:"-compact-warnings"`
	Lock            string   `cli:"-lock="`
	LockTimeout     string   `cli:"-lock-timeout="`
	Input           string   `cli:"-input="`
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

func NewApplyAction() *ApplyAction {
	return &ApplyAction{
		name: "apply",
	}
}

func (a *ApplyAction) setOutErr(outBytes, errBytes []byte) {
	a.OutBytes = outBytes
	a.ErrBytes = errBytes
}

func (a *ApplyAction) GetOutErr() ([]byte, []byte) {
	return a.OutBytes, a.ErrBytes
}

func (a *ApplyAction) unmarshal() ([]string, error) {
	//TODO: struct validation
	return unmarshalAction(a)
}

// Skip interactive plan approval before applying.
// Terraform CLI: -auto-approve
func (a *ApplyAction) EnableAutoApprove() {
	a.AutoApprove = true
}

// Path to backup the existing state file before modifying.
// Defaults to the "-state-out" path with ".backup"
// Terraform CLI: -backup=path
func (a *ApplyAction) SetBackup(path string) {
	a.Backup = path
}

// Disable automatic backup of state file before Terraform updating it.
func (a *ApplyAction) DisableBackup() {
	a.Backup = "-"
}

// Show only warning summary. Apply only to warnings that are not
// accompanied by errors.
// Terraform CLI: -compact-warnings
func (a *ApplyAction) EnableCompactWarnings() {
	a.CompactWarnings = true
}

// Disable locking of state files during state-related operations.
// Terraform CLI: -lock=true
func (a *ApplyAction) DisableLock() {
	a.Lock = "false"
}

// Override the time Terraform will wait to acquire a state lock.
// The default is 0s (zero seconds), which causes immediate failure
// if the lock is already held by another process.
// Terraform CLI: -lock-timeout=0s
func (a *ApplyAction) SetLockTimeout(duration time.Duration) {
	a.LockTimeout = duration.String()
}

// Disable ask for input for variables if not directly set. Will
// error if input was required
// Terraform CLI: -input=true
func (a *ApplyAction) DisableInteractiveVarsInput() {
	a.Input = "false"
}

// Set output to mono-color
// Terraform CLI: -no-color
func (a *ApplyAction) DisableOutputColor() {
	a.NoColor = true
}

// Set the number of parallel resource operations. Defaults to 10.
// Terraform CLI: -parallelism=n
func (a *ApplyAction) SetParallelism(n int) {
	a.Parallelism = strconv.Itoa(n)
}

// Disables auto update to state prior to checking for differences.
// This has no effect if a plan file is given to apply.
// Terraform CLI: -refresh=true
func (a *ApplyAction) DisableAutoRefresh() {
	a.Refresh = "false"
}

// Path to read and write state. Defaults to terraform.tfstate
// Terraform CLI: -state=path
func (a *ApplyAction) SetExistingStateFile(path string) {
	a.State = path
}

// Set a new file to write new terraform state. This can be used
// as a way to preserve old state file.
// Terraform CLI: -state-out=path
func (a *ApplyAction) SetNewStateFile(path string) {
	a.StateOut = path
}

// Limit operations to a set of resources and its dependencies.
// Terraform CLI: -target=resource -target=resource ...
func (a *ApplyAction) RegisterResourcesForApply(resources ...string) {
	a.Targets = append(a.Targets, resources...)
}

// Set variables in Terraform configuration.
// Terraform CLI: -var 'foo=bar' -var 'x=ray' ...
func (a *ApplyAction) SetVariables(vars ...string) {
	a.Vars = append(a.Vars, vars...)
}

// Set variables in Terraform configuration from this file.
// NOTE: if terrafrom.tfvars or any *.auto.tfvars files are present,
// they will be automatically loaded.
// Terraform CLI: -var-file=path
func (a *ApplyAction) SetVarFile(path string) {
	// TODO: Check if file exists and maybe is valid
	a.VarFile = path
}
