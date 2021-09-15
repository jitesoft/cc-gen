package gitwrapper

import (
	"time"
)

type Tag struct {
	Name string
	Hash string
	TagHash string
	Tagger string
	Email string
	Message string
	Time time.Time
}

func FindTag(tags []*Tag, name string) *Tag {
	for _, t := range tags {
		if t.Name == name {
			return t
		}
	}
	return nil
}


func GetTagRangeFrom(tags []*Tag, from string) []*Tag {
	for i, t := range tags {
		if t.Hash == from {
			// We can sub-slice here as we won't modify the commit either way.
			return tags[i:]
		}
	}
	return tags
}

func GetTagRangeTo(tags []*Tag, from string) []*Tag  {
	for i, t := range tags {
		if t.Hash == from {
			// We can sub-slice here as we won't modify the commit either way.
			return tags[:i+1]
		}
	}
	return tags
}

func GetTagRangeBetween(tags []*Tag, fromHash string, toHash string) []*Tag {
	j, k := 0, 0

	for i, t := range tags {
		if t.Hash == fromHash {
			j = i
		}
		if t.Hash == toHash {
			k = i+1
			break
		}
	}

	return tags[j:k]
}

