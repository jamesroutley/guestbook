package domain

import "time"

type Visit struct {
	URL      string
	Referrer string
	IP       string
	Created  time.Time
}
