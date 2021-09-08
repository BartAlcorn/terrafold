# Deploy Lambda via Docker Image

## Initial Deployment Steps

* iac/base/ ```terraform init``` and ```terraform apply```
* lambda root ```gmake push``` This pushes initial version to ECR
* iac/dev ```terraform init``` and ```terraform apply```

## Redeployment steps

* lambda root ```gmake deploy```
