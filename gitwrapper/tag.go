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

