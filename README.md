# TerraFold

The goal of Terrafold is to provide a convenience tool to help auto the scaffolding of AWS Lambdas, and their associated Terraform files.

A new lambda can be achieved by creating a simple json file in the repo root or, preferably, in the lambda directory root:

```json
{
  "name": "new-lambda",
  "description": "new lambda description",
  "triggers": ["api","invoke","sns","sqs"],
  "overwrite": false
}
```

This tool is extremely opinionated, by design. It was written for a specific group of developers I work with and as such follows and enforces their particular application structure pattern.

The name TerraFold is the playful combination of `Terra`form and Scaf`fold`.

## Commands

- newprofile - creates a new TerraFold profile json file.
- scaffold   - takes the named profile json and creates the handlers and iac files. Will not overwrite unless `overwrite` flag in profile is set to true.
- bumpsemver - will search a named json file for `version`, and increment the patch number.
- dumptemplates - writes the embedded template files to a folder named `TerrafoldTemplates`.


## Installation

Download this repo, and then run:

`go build -o terrafold ./cmd/cli/main.go`

You need to move or toerhwise assure `terrafold` is in your execution path.
