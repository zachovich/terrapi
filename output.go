package terrapi

type OutputAction struct {
	name    string
	NoColor bool   `cli:"-no-color"`
	State   string `cli:"-state="`
	Json    bool   `cli:"-json"`

	// outBuff, and errBuff holds the action output and error when executed
	OutBytes []byte
	ErrBytes []byte
}

func NewOutputAction() *OutputAction {
	return &OutputAction{
		name: "output",
	}
}

func (a *OutputAction) setOutErr(outBytes, errBytes []byte) {
	a.OutBytes = outBytes
	a.ErrBytes = errBytes
}

func (a *OutputAction) GetOutErr() ([]byte, []byte) {
	return a.OutBytes, a.ErrBytes
}

func (a *OutputAction) unmarshal() ([]string, error) {
	//TODO: struct validation
	return unmarshalAction(a)
}

// Set output to mono-color
// Terraform CLI: -no-color
func (a *OutputAction) DisableOutputColor() {
	a.NoColor = true
}

// Path to read and write state. Defaults to terraform.tfstate
// Terraform CLI: -state=path
func (a *OutputAction) SetExistingStateFile(path string) {
	a.State = path
}

func (a *OutputAction) SetJSONOutput() {
	a.Json = true
}
