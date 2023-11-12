# tbm
A CLI to manage S3 backends for terraform in your AWS account

# Usage

1. Set AWS keys either as `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` env vars, or in `~/.aws/credentials` file
2. Run `tbm init` to initialize TBM metadata in DynamoDB
3. Run `tbm new` to create a new S3 backend. Use the values printed in your terraform backend configuration. 

# Commands

- `tbm init` creates a dynamodb table for TBM to store S3 bucket IDs and other metadata
- `tbm new` creates a new S3 backend
- `tbm list` shows all backends in the current AWS account TBM knows of
- `tbm import` registers a pre-existing S3 backend
