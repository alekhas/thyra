package subcommand

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/uruddarraju/thyra/pkg/gateway"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the api gateway service.",
	Long:  "Starts the api gateway service and listens on few default endpoints.",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	glog.Infof("Started the server.....")
	gateway := gateway.NewDefaultGateway()
	gateway.Start()
}
