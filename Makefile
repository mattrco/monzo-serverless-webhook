.PHONY: check-env deploy

build: main.go
	@go build -o monzo-webhook

zip: build
	zip deployment.zip monzo-webhook

check-env:
	test -n "$(ACCOUNT_ID)" # $$ACCOUNT_ID env var must be set
	test -n "$(ROLE)"       # $$ROLE env var must be set

deploy: zip check-env
	aws lambda create-function \
	--cli-connect-timeout 300 \
	--region eu-west-2 \
	--function-name monzo-webhook \
	--zip-file fileb://./deployment.zip \
	--runtime go1.x \
	--tracing-config Mode=Active \
	--role arn:aws:iam::$(ACCOUNT_ID):role/$(ROLE) \
	--handler monzo-webhook
