.PHONY: check-env deploy

build: main.go
	GOOS=linux go build -o monzo-webhook

test:
	go test ./...

zip: build test
	zip deployment.zip monzo-webhook

check-env:
	test -n "$(VERSION)" # $$VERSION env var must be set

deploy: zip check-env
	aws s3 cp deployment.zip s3://monzo-webhook-lambda/versions/$(VERSION)/deployment.zip
	terraform apply -var="app_version=$(VERSION)"
