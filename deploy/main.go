package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"

	"github.com/jtfm/vpc/stacks"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stacks.NewVPCStack(
		app,
		"VPCStack",
		&stacks.VPCStackProps{
			StackProps: awscdk.StackProps{
				Env: nil,
			},
		})

	app.Synth(nil)
}
