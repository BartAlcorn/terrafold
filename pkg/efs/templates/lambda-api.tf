### API Lambda ====================================================================
module "lambdaAPI" {
  source               = "github.com/turnercode/cp-awfm-iac/modules/aws_lambda_container"
  constants            = module.constants
  app                  = var.app
  environment          = var.environment
  lambda_bucket_name   = module.constants.lambda_bucket_name
  lambda_function_name = "${var.environment}-${var.app}-api"
  image_uri            = "${module.constants.aws_account_id}.dkr.ecr.${module.constants.aws_region}.amazonaws.com/hello-world-api:0.0.1"
  handler              = "/var/task/${var.app}"
  lambda_policy        = data.aws_iam_policy_document.lambda_api_permissions.json
  lambda_timeout       = 30
  description          = "${var.description} - API Trigger"
  tags                 = merge(module.constants.tags, { environment = var.environment })
}

### APIGateway =======================================================================
#this module has references to existing resources needed by the aws_apigateway module
module "api" {
  source                   = "github.com/turnercode/cp-awfm-iac/modules/aws_apigateway"
  app                      = var.app
  environment              = var.environment
  api_lambda_arn           = module.lambdaAPI.live_alias.arn
  api_lambda_invoke_arn    = module.lambdaAPI.live_alias.invoke_arn
  api_request_parameters   = {}
  api_parent_resource_path = "v1"
  api_endpoint_path        = var.app
  api_request_method       = "POST"
  tags                     = merge(module.constants.tags, { environment = var.environment })
  api_key_required         = true
}

data "aws_iam_policy_document" "lambda_api_permissions" {
  statement {
    # sid       = "AllowInvokingLambdas"
    effect    = "Allow"
    resources = ["arn:aws:lambda:us-east-1:026155191598:function:${var.environment}-${var.app}-api:*"]
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
