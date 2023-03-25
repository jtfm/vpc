package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"

	"github.com/aws/constructs-go/constructs/v10"
)

type VPCStackProps struct {
	awscdk.StackProps
}

func VpcStack(
	scope constructs.Construct,
	id string,
	props *VPCStackProps) awscdk.Stack {

	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	_ = awsec2.NewVpc(
		stack,
		jsii.String("VPC"),
		&awsec2.VpcProps{})

	return stack
}
