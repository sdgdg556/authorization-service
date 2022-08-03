package storage

import (
	"sync"
)

type Dao struct {
	Users     *sync.Map // userName => password
	Roles     *sync.Map // roleName => {}
	UserRoles *sync.Map // userName => [roleName => {}]
	tokens    *sync.Map // token => userName
}

func InitDao() *Dao {
	return &Dao{
		Users:     &sync.Map{},
		Roles:     &sync.Map{},
		UserRoles: &sync.Map{},
		tokens:    &sync.Map{},
	}
}
