package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "xutils",
	Short: "Monitoring & Automation CLI",
	Long: `
██╗  ██╗██╗   ██╗████████╗██╗██╗     ███████╗
╚██╗██╔╝██║   ██║╚══██╔══╝██║██║     ██╔════╝
 ╚███╔╝ ██║   ██║   ██║   ██║██║     ███████╗
 ██╔██╗ ██║   ██║   ██║   ██║██║     ╚════██║
██╔╝ ██╗╚██████╔╝   ██║   ██║███████╗███████║
╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝╚══════╝╚══════╝

System Monitoring Automation Tool
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
