terraform {
  required_version = ">= 0.14"

  backend "s3" {
    region  = "us-east-1"
    profile = "aws-zephyr-main:aws-zephyr-main-devops"
    bucket  = "tf-state-zephyr-lambdas"
    key     = "endq3_ecr/dev_tfstate"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }

}
