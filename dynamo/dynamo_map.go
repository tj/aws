package dynamo

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Map of attributes.
type Map map[string]*dynamodb.AttributeValue

// NewMap returns a new map.
func NewMap() Map {
	return make(Map)
}

// Struct value.
func (m Map) Struct(name string, value interface{}) (Map, error) {
	v, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		return nil, err
	}

	m[name] = &dynamodb.AttributeValue{
		M: v,
	}

	return m, nil
}

// MustStruct value.
func (m Map) MustStruct(name string, value interface{}) Map {
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
func (m Map) String(name, value string) Map {
	m[name] = &dynamodb.AttributeValue{
		S: &value,
	}
	return m
}

// StringSet value.
func (m Map) StringSet(name string, value []string) Map {
	m[name] = &dynamodb.AttributeValue{
		SS: aws.StringSlice(value),
	}
	return m
}

// Map value.
func (m Map) Map(name string) Map {
	v := NewMap()
	m[name] = &dynamodb.AttributeValue{
		M: v,
	}
	return v
}

// Null value.
func (m Map) Null(name string) Map {
	m[name] = &dynamodb.AttributeValue{
		NULL: aws.Bool(true),
	}
	return m
}

// Blob value.
func (m Map) Blob(name string, value []byte) Map {
	m[name] = &dynamodb.AttributeValue{
		B: value,
	}
	return m
}

// Bool value.
func (m Map) Bool(name string, value bool) Map {
	m[name] = &dynamodb.AttributeValue{
		BOOL: &value,
	}
	return m
}

// BinarySet value.
func (m Map) BinarySet(name string, value [][]byte) Map {
	m[name] = &dynamodb.AttributeValue{
		BS: value,
	}
	return m
}

// Int value.
func (m Map) Int(name string, value int) Map {
	s := strconv.Itoa(value)
	m[name] = &dynamodb.AttributeValue{
		N: &s,
	}
	return m
}

// IntSet value.
func (m Map) IntSet(name string, value []int) Map {
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
func (m Map) Float(name string, value float64) Map {
	s := strconv.FormatFloat(value, 'b', -1, 64)
	m[name] = &dynamodb.AttributeValue{
		N: &s,
	}
	return m
}

// FloatSet value.
func (m Map) FloatSet(name string, value []float64) Map {
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
