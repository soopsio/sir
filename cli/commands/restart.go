package commands

import (
	"fmt"
	"net/http"
	"github.com/soopsio/sir/cli/config"
	"github.com/soopsio/sir/cli/utils"
	"github.com/soopsio/sir/lib/httpclient"

	cli "gopkg.in/urfave/cli.v1"
)

var CmdRestart = cli.Command{
	Name:      "restart",
	UsageText: "<task>",
	Category:  string(TaskCategory),
	Usage:     "restart a task",
	Action:    ActionRestart,
}

func ActionRestart(c *cli.Context) error {
	taskName := c.Args().First()

	var response map[string]interface{}
	httpclient.Client.DoJSON(http.MethodPost, config.ApiPath("/task/"+taskName+"/restart"), nil, &response)

	println()
	fmt.Println(utils.Style.Success("[INFO]"), "TASK", utils.Style.Title(taskName), "RESTARTED", "\n")

	return nil
}
