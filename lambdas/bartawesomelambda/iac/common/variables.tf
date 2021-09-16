/*
 * variables.tf
 * Common variables to use in various Terraform files (*.tf)
 */
# The application name
variable "app" {
  default = "bart-awesome-lambda"
}

variable "description" {
  default = "bart awesome lambda description"
}

# The environment
variable "environment" {
  default = ""
}


# the tags for accounting
variable "tags" {
  type = map(string)
  default = {
    application   = "zephyr"
    environment   = ""
    team          = "zephyr"
    customer      = "contentplatforms"
    contact-email = "bart.alcorn@warnermedia.com"
    provisoner    = "terraform"
  }
}
