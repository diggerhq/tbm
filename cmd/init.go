/*
Copyright © 2023 Digger.dev
*/
package cmd

import (
	"context"
	"errors"
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
				// creating metadata table since it doesn't exist
				// TODO create table here
			} else {
				log.Fatal(err)
			}
		} else {
			log.Printf("TBM was already initialized in this AWS region. Table %v already exists.\n", metadataTableName)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
