package utils

import "github.com/google/uuid"

type DummyUser struct {
	ID uuid.UUID
}

var AnonymousUser = &DummyUser{}

func (u *DummyUser) IsAnonymous() bool {
	return u == AnonymousUser
}
