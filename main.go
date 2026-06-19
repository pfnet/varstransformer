package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
)

func BuildCmd() *cobra.Command {
	cmd := command.Build(buildProcessor(), command.StandaloneEnabled, false)
	return cmd
}

func main() {
	if err := BuildCmd().Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
