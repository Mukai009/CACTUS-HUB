package main

import (
	"fmt"
	"os"

	"github.com/elthworth/Cactus-HUB/node"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Launches the CTH node and its HTTP API.",
		Run: func(cmd *cobra.Command, args []string) {
			ip, _ := cmd.Flags().GetString(flagIP)
			port, _ := cmd.Flags().GetUint64(flagPort)

			fmt.Println("Launching CTH node and its HTTP API...")

			bootstrap := node.NewPeerNode(
				"127.0.0.1",
				7000,
				true,
				false,
			)

			n := node.New(getDataDirFromCmd(cmd), ip, port, bootstrap)
			err := n.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	addDefaultRequiredFlags(runCmd)
	runCmd.Flags().String(flagIP, node.DefaultIP, "exposed IP for communication with peers")
	runCmd.Flags().Uint64(flagPort, node.DefaultHTTPort, "exposed HTTP port for communication with peers")

	return runCmd
}
