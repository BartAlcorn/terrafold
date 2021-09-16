/*
 * variables.tf
 * Common variables to use in various Terraform files (*.tf)
 */
# The application name
variable "app" {
  default = "{{.Name}}"
}

variable "description" {
  default = "{{.Description}}"
}

# The environment
variable "environment" {
  default = "{{.Stage}}"
}


# the tags for accounting
variable "tags" {
  type = map(string)
  default = {
    application   = "zephyr"
    environment   = "{{.Stage}}"
    team          = "zephyr"
    customer      = "contentplatforms"
    contact-email = "bart.alcorn@warnermedia.com"
    provisoner    = "terraform"
  }
}
