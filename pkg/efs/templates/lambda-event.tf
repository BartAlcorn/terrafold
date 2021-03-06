### EVENT Lambda ====================================================================
module "lambdaEVENT" {
  source               = "github.com/turnercode/cp-awfm-iac/modules/aws_lambda_container"
  constants            = module.constants
  app                  = var.app
  environment          = var.environment
  lambda_bucket_name   = module.constants.lambda_bucket_name
  lambda_function_name = "${var.environment}-${var.app}-event"
  image_uri            = "${module.constants.aws_account_id}.dkr.ecr.${module.constants.aws_region}.amazonaws.com/hello-world-api:0.0.1"
  handler              = "/var/task/${var.app}"
  lambda_policy        = data.aws_iam_policy_document.lambda_event_permissions.json
  lambda_timeout       = 30
  description          = var.description
  tags                 = merge(module.constants.tags, { environment = var.environment })
}

data "aws_iam_policy_document" "lambda_event_permissions" {
  statement {
    # sid       = "AllowInvokingLambdas"
    effect    = "Allow"
    resources = ["arn:aws:lambda:us-east-1:026155191598:function:${var.environment}-${var.app}-event:*"]
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
    sid    = "EB-execute-lambda"
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["events.amazonaws.com"]
    }
    actions   = ["lambda:InvokeFunction"]
    resources = ["arn:aws:lambda:us-east-1:026155191598:function:${var.environment}-${var.app}-event"]
    condition {
      test     = "ArnLike"
      variable = "aws:SourceArn"
      values = [
        "arn:aws:events:us-east-1:026155191598:rule/zephyr-events-${var.environment}/resume-step-function",
      ]
    }
  }
}
