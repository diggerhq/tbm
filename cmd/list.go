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
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View existing backends in the current region",
	Long:  `Outputs name and S3 bucket id for every backend in the current region`,
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
				fmt.Printf("Table %v does not exist in %v. Run `tbm init` first.", metadataTableName, awsConfig.Region)
			} else {
				log.Fatal(err)
			}
		}
		backends := scanDynamodbTable(metadataTableName, dynamodbClient)
		if len(backends) > 0 {
			for _, backend := range backends {
				fmt.Printf("%v - Bucket: %v", backend.name, backend.bucket)
			}
		} else {
			fmt.Printf("No backends found in %v. Run `tbm new` to create one", awsConfig.Region)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func scanDynamodbTable(tableName string, dynamodbClient *dynamodb.Client) []Backend {
	var backends []Backend
	var err error
	response, err := dynamodbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		log.Printf("Failed to scan DynanamoDB table %v. Error: %v", metadataTableName, err)
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &backends)
		if err != nil {
			log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
		}
	}
	return backends
}
