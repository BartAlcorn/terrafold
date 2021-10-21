### INVOKE Lambda ====================================================================
module "lambdaINVOKE" {
  source               = "github.com/turnercode/cp-awfm-iac/modules/aws_lambda_container"
  constants            = module.constants
  app                  = var.app
  environment          = var.environment
  lambda_bucket_name   = module.constants.lambda_bucket_name
  lambda_function_name = "${var.environment}-${var.app}-invoke"
  image_uri            = "${module.constants.aws_account_id}.dkr.ecr.${module.constants.aws_region}.amazonaws.com/hello-world-api:0.0.1"
  handler              = "/var/task/${var.app}"
  lambda_policy        = data.aws_iam_policy_document.lambda_invoke_permissions.json
  lambda_timeout       = 30
  description          = "${var.description} - Invoke Trigger"
  tags                 = merge(module.constants.tags, { environment = var.environment })
}

data "aws_iam_policy_document" "lambda_invoke_permissions" {
  statement {
    # sid       = "AllowInvokingLambdas"
    effect    = "Allow"
    resources = ["arn:aws:lambda:us-east-1:026155191598:function:${var.environment}-${var.app}-invoke:*"]
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

}
