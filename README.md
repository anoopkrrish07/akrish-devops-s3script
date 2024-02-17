# Bash script for the search and filtering

This is a simple bash script that uses AWS CLI commands to efficiently search for .txt files within an S3 bucket, enabling you to download and filter those files based on a specific string.

## Prerequisites 

- AWS CLI Installed: Ensure that the AWS Command Line Interface (CLI) is installed on your system. Or you can install it using this [link](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- Create AWS User: Create an AWS IAM user with programmatic access and s3 access permission
- Configure AWS Profile: Configure AWS local user profile using aws cli.

```
aws configure --profile profile_name
```
- Clone this git repository 

## How to run?

```
cd akrish-simple-bash/bash
chown +x s3_file_string_filter.sh
./s3_file_string_filter.sh
``` 
It will ask for the inputs.

# How to use go for the searching

This is a simple go code that helps to search the objects according to a specific string. In this I have used AWS SDK latest go module, that is v2

## Prerequisites

- Configure AWS credentials as default profile
- Clone the repo
- Install go-lang latest binary
- Install go-lang AWS SDK dependencies/packages

## How to run?
Switch to the directory where the code resides and run the command listed below to use the latest Go dependencies for the AWS SDK.

```
cd akrish-simple-bash/golang
go mod init github.com/git_username/golang
go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3
```
Run the code using this command:
```
go run s3_file_filter.go
```

Please provide your inputs accordingly.

> Note: The AWS region must be entered in this format "us-east-1".