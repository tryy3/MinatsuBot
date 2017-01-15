package permissionmanager

import (
	"encoding/json"
	"time"
)

type Group struct {
	*description
	users map[string]*User
}

func (g *Group) GetUsers() map[string]*User {
	return g.users
}

func (g *Group) GetUser(name string) (user *User, ok bool) {
	user, ok = g.users[name]
	return
}

func (g *Group) AddUser(user *User) bool {
	if _, ok := g.GetUser(user.GetName()); ok {
		return false
	}
	g.users[user.GetName()] = user
	return true
}

func (g *Group) DelUser(name string) bool {
	if _, ok := g.GetUser(name); !ok {
		return false
	}
	delete(g.users, name)
	return false
}

func (g *Group) HasPermission(perm string) bool {
	p, ok := g.permissions[perm]

	if !ok {
		return false
	}
	return p
}

func (g *Group) delete() {
	for _, user := range g.users {
		user.DelGroup(g.GetName())
	}
	g.permissions = nil
	g.users = nil
	g.name = ""
	g.created = time.Time{}
}

func (g *Group) MarshalJSON() ([]byte, error) {
	var users []string

	for _, u := range g.users {
		users = append(users, u.GetName())
	}

	return json.Marshal(&struct {
		Name        string
		Permissions map[string]bool
		Created     int64
	}{
		Name:        g.name,
		Permissions: g.permissions,
		Created:     g.created.Unix(),
	})
}
