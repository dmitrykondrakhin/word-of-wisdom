package hashcash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	bits         uint
	date         string
	randomString string
	token        string
	result       string
	checkResult  bool
}{
	{20, "231114", "abcd", "token", "1:20:231114:token::abcd:b5fe8", true},
	{22, "231114", "abcd", "token", "1:22:231114:token::abcd:6254f7", true},
	{20, "231101", "abcd", "token", "1:20:231101:token::abcd:1bba66", false},
}

func TestHashCash(t *testing.T) {
	for i, testCase := range testCases {
		hashcash := Hash{
			bits:         testCase.bits,
			date:         testCase.date,
			randomString: testCase.randomString,
		}
		result, err := hashcash.GetHeader(testCase.token)

		assert.Nil(t, err, fmt.Sprintf("test case %d", i))
		assert.Equal(t, testCase.result, result, fmt.Sprintf("test case %d", i))
		assert.Equal(t, testCase.checkResult, hashcash.Check(result), fmt.Sprintf("test case %d", i))
	}
}
