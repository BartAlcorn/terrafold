
.ONESHELL:

build: getvars
	@echo "Building..."
	@echo "Building ${PROJECT}-$(TRIGGER) v${VERSION}"
	GOARCH=amd64 GOOS=linux go build -v -ldflags="-X ${APPNAME} -X ${APPVERS} -X ${APPBUILDER}" -o ./bin ./main.go
	@echo "Dockerizing ${PROJECT} v${VERSION}"
	docker build -t ${PROJECT} .
	rm ./bin

getvars:
define GetJQ
$(shell jq '.$(1)' ./package.json)
endef

define bumpsemver
$(shell terrafold bumpsemver ./package.json)
endef

	ACCT := ${AWS_ACCOUNT}
	PROJECT := $(call GetJQ,project)
	REGION :="us-east-1"
	REPO := "${ACCT}.dkr.ecr.${REGION}.amazonaws.com"
	STAGE := "dev"
	TRIGGER := $(call GetJQ,trigger)
	USR := $$(id -u -n)
	VERSION := $(call bumpsemver)
	VPATH="github.com/turnercode/cp-awfm-common/pkg/version"
	APPNAME := "${VPATH}.AppName=${PROJECT}"
	APPVERS := "${VPATH}.Version=${VERSION}"
	APPBUILDER := "${VPATH}.BuiltBy=${USR}"

push: build
	@echo "Logging into ECR"
	aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ACCT}.dkr.ecr.${REGION}.amazonaws.com
	@echo "Tagging and Pushing ${PROJECT} v${VERSION}"
	docker tag ${PROJECT} ${REPO}/${PROJECT}-$(TRIGGER):${VERSION}
	docker push ${REPO}/${PROJECT}-$(TRIGGER):${VERSION}
	aws lambda update-function-code --function dev-${PROJECT}-$(TRIGGER) --image-uri ${REPO}/${PROJECT}-$(TRIGGER):${VERSION}
