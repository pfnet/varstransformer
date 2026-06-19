package main

import (
	"testing"

	"sigs.k8s.io/kustomize/kyaml/fn/framework/frameworktestutil"
)

func TestCommand(t *testing.T) {
	prc := frameworktestutil.CommandResultsChecker{
		Command: BuildCmd,

		//UpdateExpectedFromActual: true,
	}
	prc.Assert(t)
}
