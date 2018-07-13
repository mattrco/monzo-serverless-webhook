Monzo Webhook Lambda
====================

Serverless applications are a good candidate for receiving webhooks as they only consume resources when they run.

This application serves as an example of receiving events via Monzo's webhooks on AWS Lambda.

Setup
-----

The lambda function is invoked by a webhook request and logs the request body to stdout (automatically picked
up by CloudWatch).

Configuration is managed with [Terraform](https://www.terraform.io/intro/index.html).

Caveats
-------

By default, the API Gateway configured by Terraform will use HTTPS and redirect non-HTTPS requests.
You must register an HTTPS URL with Monzo, otherwise your transaction data will be sent across the
internet unencrypted. Do not use plain HTTP.
