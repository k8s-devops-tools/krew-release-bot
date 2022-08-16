package main

import (
	"github.com/k8s-devops-tools/krew-release-bot/pkg/source/actions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(actionCmd)
}

// actionCmd is the github action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "github action for updating plugin manifests in krew-index repo",
	Run: func(cmd *cobra.Command, args []string) {
		err := actions.RunAction()
		if err != nil {
			logrus.Fatal(err)
		}
	},
}
