/*
Copyright © 2023 Digger.dev
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/spf13/cobra"
)

const metadataTableName = "digger-tbm-metadata"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize TBM metadata on AWS",
	Long:  `TBM creates a DynamoDB table to store S3 bucket IDs of backends and other metadata`,
	Run: func(cmd *cobra.Command, args []string) {
		awsConfig, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("Failed to load AWS configuration, %v", err)
		}
		dynamodbClient := dynamodb.NewFromConfig(awsConfig)
		_, err = dynamodbClient.DescribeTable(
			context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(metadataTableName)},
		)
		if err != nil {
			var notFoundEx *types.ResourceNotFoundException
			if errors.As(err, &notFoundEx) {
				err = nil
				_ = createDynamodbTable(metadataTableName, dynamodbClient)
				fmt.Printf("TBM successfully initialized. Table %v created", metadataTableName)
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("TBM was already initialized in %v. Table %v already exists.\n", awsConfig.Region, metadataTableName)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createDynamodbTable(tableName string, dynamodbClient *dynamodb.Client) *dynamodb.CreateTableOutput {
	table, err := dynamodbClient.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("backendName"),
			AttributeType: types.ScalarAttributeTypeN,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("backendName"),
			KeyType:       types.KeyTypeHash,
		}},
		TableName: aws.String(tableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return table
}
