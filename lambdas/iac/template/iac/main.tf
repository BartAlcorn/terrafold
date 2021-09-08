module "constants" {
  source = "../../../../lambdas/iac/constants"
}

### Invoke lambda ================================================================
module "lambdaContiainerInvoke" {
  source               = "github.com/turnercode/cp-awfm-iac/modules/aws_lambda_container"
  constants            = module.constants
  app                  = var.app
  environment          = var.environment
  lambda_bucket_name   = module.constants.lambda_bucket_name
  lambda_function_name = "${var.environment}-${var.app}-invoke"
  image_uri            = "${module.constants.aws_account_id}.dkr.ecr.${module.constants.aws_region}.amazonaws.com/${var.app}:0.0.1"
  handler              = "/var/task/${var.app}"
  lambda_policy        = data.aws_iam_policy_document.lambda_INVOKE_permissions.json
  lambda_timeout       = 30
  description          = "example '${var.app}' container deployment"
  tags                 = merge(module.constants.tags, { environment = var.environment })
}

### Invoke Lambda Permissions =====================================================
data "aws_iam_policy_document" "lambda_INVOKE_permissions" {
  statement {
    # sid       = "AllowInvokingLambdas"
    effect    = "Allow"
    resources = ["arn:aws:lambda:${module.constants.aws_region}:${module.constants.aws_account_id}:function:${var.environment}-${var.app}-invoke:*"]
    actions   = ["lambda:*"]
  }

  statement {
    # sid       = "AllowSecretsManager"
    effect    = "Allow"
    resources = ["arn:aws:secretsmanager:${module.constants.aws_region}:${module.constants.aws_account_id}:secret:zephyr/configs/${var.environment}/*"]
    actions = [
      "secretsmanager:GetRandomPassword",
      "secretsmanager:ListSecrets",
      "secretsmanager:GetResourcePolicy",
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
      "secretsmanager:ListSecretVersionIds",
    ]
  }
}

### SQS lambda ================================================================
module "lambdaContiainerSQS" {
  source = "github.com/turnercode/cp-awfm-iac/modules/aws_lambda_container"
  # version              = "v1.0.2"
  constants            = module.constants
  app                  = var.app
  environment          = var.environment
  lambda_bucket_name   = module.constants.lambda_bucket_name
  lambda_function_name = "${var.environment}-${var.app}-sqs"
  image_uri            = "${module.constants.aws_account_id}.dkr.ecr.${module.constants.aws_region}.amazonaws.com/${var.app}:0.0.1"
  handler              = "/var/task/${var.app}"
  lambda_policy        = data.aws_iam_policy_document.lambda_SQS_permissions.json
  lambda_timeout       = 30
  description          = "example '${var.app}' container deployment"
  tags                 = merge(module.constants.tags, { environment = var.environment })
}

### Invoke Lambda Permissions =====================================================
data "aws_iam_policy_document" "lambda_SQS_permissions" {
  statement {
    # sid       = "AllowInvokingLambdas"
    effect    = "Allow"
    resources = ["arn:aws:lambda:${module.constants.aws_region}:${module.constants.aws_account_id}:function:${var.environment}-${var.app}-sqs:*"]
    actions   = ["lambda:*"]
  }

  statement {
    # sid       = "AllowSecretsManager"
    effect    = "Allow"
    resources = ["arn:aws:secretsmanager:${module.constants.aws_region}:${module.constants.aws_account_id}:secret:zephyr/configs/${var.environment}/*"]
    actions = [
      "secretsmanager:GetRandomPassword",
      "secretsmanager:ListSecrets",
      "secretsmanager:GetResourcePolicy",
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
      "secretsmanager:ListSecretVersionIds",
    ]
  }
}
