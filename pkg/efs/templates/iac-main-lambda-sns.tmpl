### SNS Lambda ====================================================================
module "lambdaSNS" {
  source               = "github.com/turnercode/cp-awfm-iac/modules/aws_lambda_container"
  constants            = module.constants
  app                  = var.app
  environment          = var.environment
  lambda_bucket_name   = module.constants.lambda_bucket_name
  lambda_function_name = "${var.environment}-${var.app}-sns"
  image_uri            = "${module.constants.aws_account_id}.dkr.ecr.${module.constants.aws_region}.amazonaws.com/hello-world-api:0.0.1"
  handler              = "/var/task/${var.app}"
  lambda_policy        = data.aws_iam_policy_document.lambda_sns_permissions.json
  lambda_timeout       = 30
  description          = "example '${var.app}' container deployment"
  tags                 = merge(module.constants.tags, { environment = var.environment })
}

module "sns" {
  source         = "github.com/turnercode/cp-awfm-iac/modules/aws_sns"
  constants      = module.constants
  app            = var.app
  environment    = var.environment
  sns_topic_name = "${var.environment}-${var.app}"
  description    = "{{.Description}}"
  tags           = var.tags
}

resource "aws_sns_topic_subscription" "{{.Name}}_lambda_target" {
  topic_arn = module.sns.sns_arn
  protocol  = "lambda"
  endpoint  = module.lambdaSNS.lambda_arn
}

resource "aws_lambda_permission" "allow_sns_{{.Name}}" {
  statement_id  = "AllowExecutionFromSNS"
  action        = "lambda:InvokeFunction"
  function_name = module.lambdaSNS.lambda_func.function_name
  principal     = "sns.amazonaws.com"
  source_arn    = module.sns.sns_arn
}

data "aws_iam_policy_document" "lambda_sns_permissions" {
  statement {
    # sid       = "AllowInvokingLambdas"
    effect    = "Allow"
    resources = ["arn:aws:lambda:us-east-1:026155191598:function:${var.environment}-${var.app}-sns:*"]
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
    # sid  = "AllowSNS"
    effect = "Allow"
    actions = [
      "sns:*",
    ]
    resources = [
      module.sns.sns_arn,
    ]
  }
}
