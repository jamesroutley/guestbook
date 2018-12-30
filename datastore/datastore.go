package datastore

import "github.com/jamesroutley/guestbook/domain"

type Datastore interface {
	Store(visit domain.Visit) error
}
