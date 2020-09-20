package terrapi

import "time"

type InitAction struct {
	name               string
	Backend            string   `cli:"-backend="`
	BackendConfig      []string   `cli:"-backend-config="`
	ForceCopy          bool     `cli:"-force-copy"`
	CopyModuleFromPath string   `cli:"-from-module="`
	DownloadModules    string   `cli:"-get="`
	DownloadPlugins    string   `cli:"-get-plugins="`
	Input              string   `cli:"-input="`
	Lock               string   `cli:"-lock="`
	LockTimeout        string   `cli:"-lock-timeout="`
	NoColor            bool     `cli:"-no-color"`
	PluginDirs         []string `cli:"-plugin-dir="`
	Reconfigure        bool     `cli:"-reconfigure"`
	Upgrade            string   `cli:"-upgrade="`
	VerifyPlugins      string   `cli:"-verify-plugins="`

	// outBuff, and errBuff holds the action output and error when executed
	OutBytes []byte
	ErrBytes []byte
}

func NewInitAction() *InitAction {
	return &InitAction{
		name: "init",
	}
}

func (a *InitAction) setOutErr(outBytes, errBytes []byte) {
	a.OutBytes = outBytes
	a.ErrBytes = errBytes
}

func (a *InitAction) GetOutErr() ([]byte, []byte) {
	return a.OutBytes, a.ErrBytes
}

func (a *InitAction) unmarshal() ([]string, error) {
	//TODO: struct validation
	return unmarshalAction(a)
}

// Skip configuring the backend for this configuration
// Terraform CLI: -backend=true
func (a *InitAction) SkipBackendConfiguration() {
	a.Backend = "false"
}

// Used for partial backend configuration, in situations where
// the backend settings are dynamic or sensitive and so cannot
// be statically specified in the configuration file. It can be
// specified multiple times.
//
// The following example contrasts the difference between having
// a complete backend configuration and a partial backend
// configuration:
//
// A complete `consul` backend configuration:
// terraform {
//  backend "consul" {
//    address = "demo.consul.io"
//    scheme  = "https"
//    path    = "example_app/terraform_state"
//  }
//}
//
// Partial `consul` backend configuration and the use of
// `-backend-config`:
// terraform {
//  backend "consul" {}
//}
//
// Option#1 using a backend config file:
//address = "demo.consul.io"
//path    = "example_app/terraform_state"
//scheme  = "https"
//
// -backend-config=/path/to/file
//
// Option#2 using inline variables:
// terraform init \
// -backend-config="address=demo.consul.io" \
// -backend-config="path=example_app/terraform_state" \
// -backend-config="scheme=https"
//
// Terraform CLI: -backend-config=path
// Terraform CLI: -backend-config="key=value"
func (a *InitAction) SetBackendConfiguration(backendConfigs ...string) {
	a.BackendConfig = append(a.BackendConfig, backendConfigs...)
}

// Suppress prompts about copying state data.
// Terraform CLI: -force-copy
func (a *InitAction) ForceCopyStateFile() {
	a.ForceCopy = true
}

// Terraform CLI: -from-module=SOURCE
func (a *InitAction) CopyModule(path string) {
	a.CopyModuleFromPath = path
}

// Terraform CLI: -get=true
func (a *InitAction) DisableModulesDownload() {
	a.DownloadModules = "false"
}

// Terraform CLI: -get-plugins=true
func (a *InitAction) DisablePluginsDownload() {
	a.DownloadPlugins = "false"
}

// Disable ask for input for variables if not directly set. Will
// error if input was required
// Terraform CLI: -input=true
func (a *InitAction) DisableInteractiveVarsInput() {
	a.Input = "false"
}

// Disable locking of state files during state-related operations.
// Terraform CLI: -lock=true
func (a *InitAction) DisableLock() {
	a.Lock = "false"
}

// Override the time Terraform will wait to acquire a state lock.
// The default is 0s (zero seconds), which causes immediate failure
// if the lock is already held by another process.
// Terraform CLI: -lock-timeout=0s
func (a *InitAction) SetLockTimeout(duration time.Duration) {
	a.LockTimeout = duration.String()
}

// Set output to mono-color
// Terraform CLI: -no-color
func (a *InitAction) DisableOutputColor() {
	a.NoColor = true
}

// Terraform CLI: -plugin-dir path -plugin-dir path ...
func (a *InitAction) SetPluginDirs(paths ...string) {
	a.PluginDirs = append(a.PluginDirs, paths...)
}

// This will disregards any existing configuration, preventing migration
// of any existing state.
// Terraform CLI: -reconfigure
func (a *InitAction) DisregardExistingConfiguration() {
	a.Reconfigure = true
}

// Terraform CLI: -upgrade=false
func (a *InitAction) EnableUpgrade() {
	a.Upgrade = "true"
}

// Terraform CLI: -verify-plugins=true
func (a *InitAction) DisablePluginsVerification() {
	a.VerifyPlugins = "false"
}