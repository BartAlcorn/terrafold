### SQS Lambda ====================================================================
module "lambdaSQS" {
  source               = "github.com/turnercode/cp-awfm-iac/modules/aws_lambda_container"
  constants            = module.constants
  app                  = var.app
  environment          = var.environment
  lambda_bucket_name   = module.constants.lambda_bucket_name
  lambda_function_name = "${var.environment}-${var.app}-sqs"
  image_uri            = "${module.constants.aws_account_id}.dkr.ecr.${module.constants.aws_region}.amazonaws.com/hello-world-api:0.0.1"
  handler              = "/var/task/${var.app}"
  lambda_policy        = data.aws_iam_policy_document.lambda_sqs_permissions.json
  lambda_timeout       = 30
  description          = "${var.description} - SQS Trigger"
  tags                 = merge(module.constants.tags, { environment = var.environment })
}

### SQS =======================================================================
module "sqs" {
  source         = "github.com/turnercode/cp-awfm-iac/modules/aws_sqs"
  constants      = module.constants
  app            = var.app
  environment    = var.environment
  sqs_queue_name = "${var.environment}-${var.app}"
  sqs_lambda_arn = module.lambdaSQS.live_alias.arn
  description    = "'${var.app}' Airing Retriever SQS " #
  tags           = merge(module.constants.tags, { environment = var.environment })
}

data "aws_iam_policy_document" "lambda_sqs_permissions" {
  statement {
    # sid       = "AllowInvokingLambdas"
    effect    = "Allow"
    resources = ["arn:aws:lambda:us-east-1:026155191598:function:${var.environment}-${var.app}-sqs:*"]
    actions   = ["lambda:*"]
  }

  statement {
    # sid       = "AllowSecretsManager"
    effect    = "Allow"
    resources = ["arn:aws:secretsmanager:us-east-1:026155191598:secret:zephyr/configs/*"]
    actions = [
      "secretsmanager:GetRandomPassword",
      "secretsmanager:ListSecrets",
      "secretsmanager:GetResourcePolicy",
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
      "secretsmanager:ListSecretVersionIds",
    ]
  }

  statement {
    # sid = "AllowReceiveSQS"
    effect    = "Allow"
    resources = [
      module.sqs.arn
    ]
    actions = [
      "sqs:ReceiveMessage",
      "sqs:GetQueueAttributes",
      "sqs:DeleteMessage"
    ]
  }
}
