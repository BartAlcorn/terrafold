/*
 * variables.tf
 * Common variables to use in various Terraform files (*.tf)
 */
# The application name
variable "app" {
  default = "APP"
}

variable "description" {
  description = "the description of this lambda"
  default     = ""
}

# The environment
variable "environment" {
  default = "dev"
}
