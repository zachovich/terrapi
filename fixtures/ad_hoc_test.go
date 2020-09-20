package fixtures

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zachovich/terrapi"
	"testing"
)

func TestExec(t *testing.T) {
	t.Run("fixtures", func(t *testing.T) {
		logrus.SetLevel(logrus.DebugLevel)

		init := terrapi.NewInitAction()
		init.SetPluginDirs("/home/zach/tmp/terraform-plugins/")
		init.DisableInteractiveVarsInput()
		init.SetBackendConfiguration(
			"bucket=zach-terraform-state-files",
			"key=terrapi-fixtures",
			"region=eu-west-1",
			"profile=zach",
		)
		//init.DisablePluginsVerification()

		apply := terrapi.NewApplyAction()
		apply.DisableInteractiveVarsInput()
		apply.EnableAutoApprove()

		output := terrapi.NewOutputAction()
		output.SetJSONOutput()

		destroy := terrapi.NewDestroyAction()
		destroy.EnableAutoApprove()

		terraform, err := terrapi.NewTerraform(terrapi.WithActions(init, apply, output, destroy))
		if err != nil {
			fmt.Println(err)
		}

		terraform.CodePath = "/home/zach/Devel/go/github.com/zachovich/terrapi/fixtures"

		_, se, e := terraform.Exec()
		if e != nil {
			fmt.Println(e.Error())
			fmt.Println(string(se))
		}

		oo, _ := output.GetOutErr()
		fmt.Println(string(oo))
	})
}
