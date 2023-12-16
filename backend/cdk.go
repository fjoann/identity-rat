package backend

import (
	"identity-rat/backend/api"
	"identity-rat/backend/database"

	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type IdentityRatProps struct {
	StackProps cdk.StackProps
}

func NewIdentityRat(scope constructs.Construct, id string, props *IdentityRatProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	database.NewDatabase(stack, jsii.String("Database"))
	api.NewApi(stack, jsii.String("Api"))

	return stack
}
