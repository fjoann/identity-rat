package main

import (
	"identity-rat/backend"

	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := cdk.NewApp(nil)

	backend.NewIdentityRat(app, "IdentityRatDev", &backend.IdentityRatProps{
		StackProps: cdk.StackProps{
			Env: &cdk.Environment{
				Account: jsii.String("963674144719"),
				Region:  jsii.String("eu-central-1"),
			},
		},
	})

	backend.NewIdentityRat(app, "IdentityRat", &backend.IdentityRatProps{
		StackProps: cdk.StackProps{
			Env: &cdk.Environment{
				Account: jsii.String("999552020175"),
				Region:  jsii.String("eu-central-1"),
			},
		},
	})

	app.Synth(nil)
}
