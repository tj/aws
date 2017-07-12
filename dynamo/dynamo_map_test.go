package dynamo

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/tj/assert"
)

func Test(t *testing.T) {
	a := NewMap()

	a.Int("id", 12345).
		String("user", "Tobi")

	a.Map("meta").
		String("email", "tobi@apex.sh").
		String("species", "ferret")

	expected := map[string]*dynamodb.AttributeValue{
		"id": {
			N: aws.String("12345"),
		},
		"user": {
			S: aws.String("Tobi"),
		},
		"meta": {
			M: map[string]*dynamodb.AttributeValue{
				"email": {
					S: aws.String("tobi@apex.sh"),
				},
				"species": {
					S: aws.String("ferret"),
				},
			},
		},
	}

	assert.Equal(t, expected, map[string]*dynamodb.AttributeValue(a))
}
