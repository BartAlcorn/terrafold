## Zephyr Specific constants
output "aws_region" {
  value = "us-east-1"
}
output "aws_account_id" {
  value = "026155191598"
}
output "aws_profile" {
  value = "aws-zephyr-main"
}
output "lambda_bucket_name" {
  value = "tf-state-zephyr-lambdas"
}
# The stash context for secrets ex odt, zephr,compass
output "stash_context" {
  value = "zephyr"
}
output "tags" {
  value = {
    application   = "zephyr"
    team          = "dawgs"
    customer      = "cdp"
    contact-email = "bart.alcorn@warnermedia.com"
    provisoner    = "terraform"
  }
}
# user roles... Capitalization matters!!
output "user_roles" {
  description = "users roles for access"
  value = [
    "aws-zephyr-main-admin/*",
    "aws-zephyr-main-devops/*",
  ]
}
output "sns_allowed_acct_ids" {
  value = [
    "277635488776", # content-platforms account
    "837769064668", # ondemandtools
    "298547466439", # compass
  ]
}
#============================================================================#
# Network and security configuration (common to everything in an AWS account)
#============================================================================#
# The VPC to use for the Fargate cluster
output "vpc" {
  value = "vpc-0cc54b10f90de5084"
}
# The private subnets, minimum of 2, that are a part of the VPC(s)
output "private_subnets" {
  value = "subnet-0560e5b578cb056d5,subnet-018f5ce706432c0e9"
}
# The public subnets, minimum of 2, that are a part of the VPC(s)
# This isn't currently used anywhere, so it's commented out
# output "public_subnets" {
#   //I'm not sure what these should be for ODT
#   //value = "subnet-0bcfb6e3c01ee3e39,subnet-02223e04dba270af2"
# }
# The security groups
output "security_group_ids" {
  value = "sg-0f675abdc521fa28b"
}
