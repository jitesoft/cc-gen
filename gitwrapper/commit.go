package gitwrapper

import (
	"time"
)

type Commit struct {
	Hash string
	Message string
	Time time.Time
	Author string
	Email string
}
