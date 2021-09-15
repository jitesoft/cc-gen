package gitwrapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSortTagsBy_Sort(t *testing.T) {
	var tags []*Tag
	time1, _ := time.Parse(time.RFC3339, "2021-09-13T10:09:00+00:00")
	time2, _ := time.Parse(time.RFC3339, "2021-05-13T10:09:00+00:00")
	time3, _ := time.Parse(time.RFC3339, "2021-01-13T10:09:00+00:00")
	time4, _ := time.Parse(time.RFC3339, "2021-12-13T10:09:00+00:00")

	tags = append(tags, &Tag{
		Name: "second",
		Time: time1,
	})
	tags = append(tags, &Tag{
		Name: "third",
		Time: time2,
	})
	tags = append(tags, &Tag{
		Name: "last",
		Time: time3,
	})
	tags = append(tags, &Tag{
		Name: "first",
		Time: time4,
	})

	sortTagsBy(func(t1 *Tag, t2 *Tag) bool {
		return t1.Time.After(t2.Time)
	}).Sort(tags)

	assert.Equal(t, tags[0].Name, "first")
	assert.Equal(t, tags[1].Name, "second")
	assert.Equal(t, tags[2].Name, "third")
	assert.Equal(t, tags[3].Name, "last")
}
