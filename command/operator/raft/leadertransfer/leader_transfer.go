package leadertransfer

import (
	"flag"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
)

func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui}
	c.init()
	return c
}

type cmd struct {
	UI    cli.Ui
	flags *flag.FlagSet
	http  *flags.HTTPFlags
	help  string
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.http = &flags.HTTPFlags{}
	flags.Merge(c.flags, c.http.ClientFlags())
	flags.Merge(c.flags, c.http.ServerFlags())
	c.help = flags.Usage(help, c.flags)
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		c.UI.Error(fmt.Sprintf("Failed to parse args: %v", err))
		return 1
	}

	// Set up a client.
	client, err := c.http.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error initializing client: %s", err))
		return 1
	}

	// Fetch the current configuration.
	result, err := raftTransferLeader(client, c.http.Stale())
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error transfering leadership: %v", err))
		return 1
	}

	c.UI.Output(result)
	return 0
}

func raftTransferLeader(client *api.Client, stale bool) (string, error) {
	q := &api.QueryOptions{
		AllowStale: stale,
	}
	reply, err := client.Operator().RaftLeaderTransfer(q)
	if err != nil {
		return "", fmt.Errorf("Failed to transfer leadership %w", err)
	}
	if !reply.Success {
		return "", fmt.Errorf("Failed to transfer leadership")
	}
	return "Success", nil
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return c.help
}

const synopsis = "Display the current Raft peer configuration"
const help = `
Usage: consul operator raft list-peers [options]

  Displays the current Raft peer configuration.
`