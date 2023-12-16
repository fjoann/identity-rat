package api

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	apigateway "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	s3assets "github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Api struct {
	constructs.Construct
}

func NewApi(scope constructs.Construct, id *string) Api {
	lambda := lambda.NewFunction(scope, jsii.String("LambdaProxy"), &lambda.FunctionProps{
		Runtime: lambda.Runtime_PROVIDED_AL2023(),
		Code: lambda.Code_FromAsset(jsii.String("backend/api/lambdaproxy"), &s3assets.AssetOptions{
			Bundling: &cdk.BundlingOptions{
				Image:      cdk.DockerImage_FromRegistry(jsii.String("golang:1.21")),
				Entrypoint: jsii.Strings("/bin/bash", "-c"),
				Command:    jsii.Strings("go build -tags lambda.norpc -o /asset-output/bootstrap"),
				Environment: &map[string]*string{
					"GOCACHE":    jsii.String("/tmp/go-build"),
					"GOMODCACHE": jsii.String("/tmp/go/pkg/mod"),
				},
			},
		}),
		Handler: jsii.String("bootstrap"),
	})

	apigateway.NewLambdaRestApi(scope, jsii.String("RestApi"), &apigateway.LambdaRestApiProps{
		Handler: lambda,
	})

	return Api{
		Construct: scope,
	}
}
