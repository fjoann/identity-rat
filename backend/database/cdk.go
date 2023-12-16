package database

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Database struct {
	constructs.Construct
	UsersTable dynamodb.TableV2
}

func NewDatabase(scope constructs.Construct, id *string) Database {
	usersTableProps := dynamodb.TablePropsV2{
		RemovalPolicy: cdk.RemovalPolicy_DESTROY,
		PartitionKey: &dynamodb.Attribute{
			Name: jsii.String("id"),
			Type: dynamodb.AttributeType_STRING,
		},
	}

	usersTable := dynamodb.NewTableV2(scope, jsii.String("Users"), &usersTableProps)

	return Database{
		Construct:  scope,
		UsersTable: usersTable,
	}
}
