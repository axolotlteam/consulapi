package cmd

import (
	"consulapi/consul"
	"consulapi/env"

	"github.com/spf13/cobra"
)

// consulCmd represents the server command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "service cmd",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.Flags().StringVarP(&env.ConsulHost, "consul", "c", "127.0.0.1:8500", "consul address")
	serviceCmd.Flags().StringVarP(&env.DeRegister, "deregister", "d", "", "deregister service name")
}

func start() {
	switch env.DeRegister {
	case "all":
		consul.DeregisterAll(env.ConsulHost)
		return
	case "":
		panic("deregister name required")
	default:
		consul.Deregister(env.ConsulHost, env.DeRegister)
	}
}
