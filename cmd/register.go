package cmd

import (
	"github.com/agajdosi/uhlobot/parlamentni"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register all bots on given platform.",
	Long:  "Will register all bots on given platform, may need a manual cooperation during the process.",
	Run: func(cmd *cobra.Command, args []string) {
		parlamentni.RegisterAll()
	},
}
