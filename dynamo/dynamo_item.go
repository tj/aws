package dynamo

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Item is a map of attributes.
type Item map[string]*dynamodb.AttributeValue

// NewItem returns a new item.
func NewItem() Item {
	return make(Item)
}

// Marshal returns a new item from struct.
func Marshal(value interface{}) (Item, error) {
	v, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		return nil, err
	}

	return Item(v), nil
}

// MustMarshal returns a new item from struct.
func MustMarshal(value interface{}) Item {
	v, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		panic(err)
	}

	return Item(v)
}

// Marshal value.
func (m Item) Marshal(name string, value interface{}) (Item, error) {
	v, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		return nil, err
	}

	m[name] = &dynamodb.AttributeValue{
		M: v,
	}

	return m, nil
}

// MustMarshal value.
func (m Item) MustMarshal(name string, value interface{}) Item {
	v, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		panic(err)
	}

	m[name] = &dynamodb.AttributeValue{
		M: v,
	}

	return m
}

// String value.
func (m Item) String(name, value string) Item {
	m[name] = &dynamodb.AttributeValue{
		S: &value,
	}
	return m
}

// StringSet value.
func (m Item) StringSet(name string, value []string) Item {
	m[name] = &dynamodb.AttributeValue{
		SS: aws.StringSlice(value),
	}
	return m
}

// Map value.
func (m Item) Map(name string) Item {
	v := make(Item)
	m[name] = &dynamodb.AttributeValue{
		M: v,
	}
	return v
}

// Null value.
func (m Item) Null(name string) Item {
	m[name] = &dynamodb.AttributeValue{
		NULL: aws.Bool(true),
	}
	return m
}

// Blob value.
func (m Item) Blob(name string, value []byte) Item {
	m[name] = &dynamodb.AttributeValue{
		B: value,
	}
	return m
}

// Bool value.
func (m Item) Bool(name string, value bool) Item {
	m[name] = &dynamodb.AttributeValue{
		BOOL: &value,
	}
	return m
}

// BinarySet value.
func (m Item) BinarySet(name string, value [][]byte) Item {
	m[name] = &dynamodb.AttributeValue{
		BS: value,
	}
	return m
}

// Int value.
func (m Item) Int(name string, value int) Item {
	s := strconv.Itoa(value)
	m[name] = &dynamodb.AttributeValue{
		N: &s,
	}
	return m
}

// IntSet value.
func (m Item) IntSet(name string, value []int) Item {
	var s []*string

	for _, n := range value {
		v := strconv.Itoa(n)
		s = append(s, &v)
	}

	m[name] = &dynamodb.AttributeValue{
		NS: s,
	}

	return m
}

// Float value.
func (m Item) Float(name string, value float64) Item {
	s := strconv.FormatFloat(value, 'b', -1, 64)
	m[name] = &dynamodb.AttributeValue{
		N: &s,
	}
	return m
}

// FloatSet value.
func (m Item) FloatSet(name string, value []float64) Item {
	var s []*string

	for _, n := range value {
		v := strconv.FormatFloat(n, 'b', -1, 64)
		s = append(s, &v)
	}

	m[name] = &dynamodb.AttributeValue{
		NS: s,
	}

	return m
}

// Attributes returns the AttributeValue map.
func (m Item) Attributes() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue(m)
}
