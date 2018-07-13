provider "aws" {
  region = "eu-west-2"
}

variable "app_version" {
}

resource "aws_lambda_function" "monzo_webhook" {
  function_name = "MonzoWebhook"
  handler       = "monzo-webhook"
  runtime       = "go1.x"
  s3_bucket     = "monzo-webhook-lambda"
  s3_key        = "versions/${var.app_version}/deployment.zip"
  role          = "${aws_iam_role.lambda_exec.arn}"
}

# Version explained at https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
resource "aws_iam_role" "lambda_exec" {
  name = "monzo_webhook_lambda"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "api_gateway" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.monzo_webhook.arn}"
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any method on any resource
  # within the API Gateway "REST API".
  source_arn = "${aws_api_gateway_rest_api.monzo_webhook.execution_arn}/*/*"
}
