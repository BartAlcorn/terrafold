module "constants" {
  source = "../../../../lambdas/iac/constants"
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-api" {
  name                 = "bart-awesome-lambda-api"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-invoke" {
  name                 = "bart-awesome-lambda-invoke"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-sns" {
  name                 = "bart-awesome-lambda-sns"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-sqs" {
  name                 = "bart-awesome-lambda-sqs"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

