package main

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

var testBody = `{
    "type": "transaction.created",
    "data": {
        "account_id": "acc_00008gju41AHyfLUzBUk8A",
        "amount": -350,
        "created": "2015-09-04T14:28:40Z",
        "currency": "GBP",
        "description": "Ozone Coffee Roasters",
        "id": "tx_00008zjky19HyFLAzlUk7t",
        "category": "eating_out",
        "is_load": false,
        "settled": true,
        "merchant": {
            "address": {
                "address": "98 Southgate Road",
                "city": "London",
                "country": "GB",
                "latitude": 51.54151,
                "longitude": -0.08482400000002599,
                "postcode": "N1 3JD",
                "region": "Greater London"
            },
            "created": "2015-08-22T12:20:18Z",
            "group_id": "grp_00008zIcpbBOaAr7TTP3sv",
            "id": "merch_00008zIcpbAKe8shBxXUtl",
            "logo": "https://pbs.twimg.com/profile_images/527043602623389696/68_SgUWJ.jpeg",
            "emoji": "🍞",
            "name": "The De Beauvoir Deli Co.",
            "category": "eating_out"
        }
    }
}
`

func TestHandler(t *testing.T) {
	req := events.APIGatewayProxyRequest{}
	req.HTTPMethod = "POST"
	req.Body = testBody

	resp, err := Handler(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("Expected %d, got %d", http.StatusAccepted, resp.StatusCode)
	}
}
