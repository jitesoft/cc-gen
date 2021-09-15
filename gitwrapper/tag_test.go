package gitwrapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTagRangeBetween(t *testing.T) {
	var list = []*Tag{
		{Message: "1", Hash: "aaa"},
		{Message: "2", Hash: "bbb"},
		{Message: "3", Hash: "ccc"},
		{Message: "4", Hash: "ddd"},
		{Message: "5", Hash: "eee"},
		{Message: "6", Hash: "fff"},
	}

	out := GetTagRangeBetween(list, "bbb", "eee")

	assert.Equal(t, out[0].Message, "2")
	assert.Equal(t, out[1].Message, "3")
	assert.Equal(t, out[2].Message, "4")
	assert.Equal(t, out[3].Message, "5")
}

func TestGetTagRangeFrom(t *testing.T) {
	var list = []*Tag{
		{Message: "1", Hash: "aaa"},
		{Message: "2", Hash: "bbb"},
		{Message: "3", Hash: "ccc"},
		{Message: "4", Hash: "ddd"},
		{Message: "5", Hash: "eee"},
		{Message: "6", Hash: "fff"},
	}

	out := GetTagRangeFrom(list, "ddd")

	assert.Equal(t, out[0].Message, "4")
	assert.Equal(t, out[1].Message, "5")
	assert.Equal(t, out[2].Message, "6")
}

func TestGetTagRangeTo(t *testing.T) {
	var list = []*Tag{
		{Message: "1", Hash: "aaa"},
		{Message: "2", Hash: "bbb"},
		{Message: "3", Hash: "ccc"},
		{Message: "4", Hash: "ddd"},
		{Message: "5", Hash: "eee"},
		{Message: "6", Hash: "fff"},
	}

	out := GetTagRangeTo(list, "ccc")

	assert.Equal(t, out[0].Message, "1")
	assert.Equal(t, out[1].Message, "2")
	assert.Equal(t, out[2].Message, "3")
}
