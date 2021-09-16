### ECR Repository ================================================================
resource "aws_ecr_repository" "ecr-{{.Trigger}}" {
  name                 = "{{.Name}}-{{.Trigger}}"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(module.constants.tags, { environment = "all" })
}

