package gitwrapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCommitRangeBetween(t *testing.T) {
	var list = []*Commit{
		{Message: "1", Hash: "aaa"},
		{Message: "2", Hash: "bbb"},
		{Message: "3", Hash: "ccc"},
		{Message: "4", Hash: "ddd"},
		{Message: "5", Hash: "eee"},
		{Message: "6", Hash: "fff"},
	}

	out := GetCommitRangeBetween(list, "bbb", "eee")

	assert.Equal(t, out[0].Message, "2")
	assert.Equal(t, out[1].Message, "3")
	assert.Equal(t, out[2].Message, "4")
	assert.Equal(t, out[3].Message, "5")
}

func TestGetCommitRangeFrom(t *testing.T) {
	var list = []*Commit{
		{Message: "1", Hash: "aaa"},
		{Message: "2", Hash: "bbb"},
		{Message: "3", Hash: "ccc"},
		{Message: "4", Hash: "ddd"},
		{Message: "5", Hash: "eee"},
		{Message: "6", Hash: "fff"},
	}

	out := GetCommitRangeFrom(list, "ddd")

	assert.Equal(t, out[0].Message, "4")
	assert.Equal(t, out[1].Message, "5")
	assert.Equal(t, out[2].Message, "6")
}

func TestGetCommitRangeTo(t *testing.T) {
	var list = []*Commit{
		{Message: "1", Hash: "aaa"},
		{Message: "2", Hash: "bbb"},
		{Message: "3", Hash: "ccc"},
		{Message: "4", Hash: "ddd"},
		{Message: "5", Hash: "eee"},
		{Message: "6", Hash: "fff"},
	}

	out := GetCommitRangeTo(list, "ccc")

	assert.Equal(t, out[0].Message, "1")
	assert.Equal(t, out[1].Message, "2")
	assert.Equal(t, out[2].Message, "3")
}
