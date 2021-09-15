package conventional

import (
	"github.com/jitesoft/cc-gen/gitwrapper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testCommit = &gitwrapper.Commit{
	Hash:    "abc",
	Message: "feat: Test 123",
	Time:    time.Time{},
	Author:  "johannes",
	Email:   "johannes@jitesoft.com",
}

func TestIsConventional(t *testing.T) {
	testCommit.Message = "feat: Test 123"
	assert.True(t, IsConventional(testCommit))
	testCommit.Message = `test(git_wrapper): Added this test

and the test was good.
`
	assert.True(t, IsConventional(testCommit))
	testCommit.Message = `refactor(nothing): Not worth it

multi
row
body
`
	assert.True(t, IsConventional(testCommit))

	testCommit.Message = "test"
	assert.False(t, IsConventional(testCommit))
	testCommit.Message = `feat(test):abc` // Missing space.
	assert.False(t, IsConventional(testCommit))
	testCommit.Message = `feat(test) abc` // missing :.
	assert.False(t, IsConventional(testCommit))
	testCommit.Message = "feat"
	assert.False(t, IsConventional(testCommit))
	testCommit.Message = "just a message"
	assert.False(t, IsConventional(testCommit))
	testCommit.Message = "a() ¤3###!! standard ###!#%%# commit."
	assert.False(t, IsConventional(testCommit))
}

func TestParseConventional(t *testing.T) {
	testCommit.Message = "feat: Test 123"
	result, _ := ParseConventional(testCommit)
	assert.Equal(t, "feat", result.Type)
	assert.Equal(t, "", result.SubType)
	assert.Equal(t, "Test 123", result.Header)
	assert.Equal(t, "", result.Message)

	testCommit.Message = `test(git_wrapper): Added this test

and the test was good.
`
	result, _ = ParseConventional(testCommit)
	assert.Equal(t, "test", result.Type)
	assert.Equal(t, "git_wrapper", result.SubType)
	assert.Equal(t, "Added this test", result.Header)
	assert.Equal(t, "and the test was good.", result.Message)

	testCommit.Message = `refactor(nothing): Not worth it

multi
row
body
`
	result, _ = ParseConventional(testCommit)
	assert.Equal(t, "refactor", result.Type)
	assert.Equal(t, "nothing", result.SubType)
	assert.Equal(t, "Not worth it", result.Header)
	assert.Equal(t, `multi
row
body`, result.Message)
}

func TestGroupByType(t *testing.T) {
	var list = []*Commit{
		{
			Type:    "test1",
			SubType: "",
			Header:  "abc",
			Message: "",
			Hash:    "",
			Author:  "",
			Time:    time.Time{},
		},
		{
			Type:    "test2",
			SubType: "",
			Header:  "abc123",
			Message: "",
			Hash:    "",
			Author:  "",
			Time:    time.Time{},
		},
		{
			Type:    "test1",
			SubType: "",
			Header:  "abcdef",
			Message: "",
			Hash:    "",
			Author:  "",
			Time:    time.Time{},
		},
	}

	out := GroupByType(list)

	assert.Equal(t, out["test1"][0].Header, "abc")
	assert.Equal(t, out["test1"][1].Header, "abcdef")
	assert.Equal(t, out["test2"][0].Header, "abc123")
}
