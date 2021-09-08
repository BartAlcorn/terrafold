module "constants" {
  source = "../../../../lambdas/iac/constants"
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-api" {
  name                 = "${var.app}-api"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-invoke" {
  name                 = "${var.app}-invoke"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-sns" {
  name                 = "${var.app}-sns"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-sqs" {
  name                 = "${var.app}-sqs"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}
