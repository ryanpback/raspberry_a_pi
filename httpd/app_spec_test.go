package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type queryTest struct {
	query    query
	expected []string
}

var queryTests = []queryTest{
	{
		query{
			"test":  "{test}",
			"test2": "{test2}",
		},
		[]string{
			"test",
			"{test}",
			"test2",
			"{test2}",
		},
	},
	{
		query{
			"test3": "{test3}",
			"test4": "{test4}",
		},
		[]string{
			"test3",
			"{test3}",
			"test4",
			"{test4}",
		},
	},
}

func TestConvertQueries(t *testing.T) {
	assert := assert.New(t)

	for _, qt := range queryTests {
		convertedQueries := convertQueries(qt.query)

		assert.Equal(qt.expected, convertedQueries)
	}
}
