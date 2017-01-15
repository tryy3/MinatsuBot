package permissionmanager

import (
	"encoding/json"
	"time"
)

type User struct {
	*description
	realGroups map[string]*Group
	groups     []string
}

func (u *User) GetGroups() map[string]*Group {
	return u.realGroups
}

func (u *User) GetGroup(name string) (group *Group, ok bool) {
	group, ok = u.realGroups[name]
	return
}

func (u *User) HasPermission(perm string) bool {
	p, ok := u.permissions[perm]
	if !ok {
		for _, g := range u.realGroups {
			if g.HasPermission(perm) {
				return true
			}
		}
	}
	return p
}

func (u *User) AddPermission(perm string, positive bool) bool {
	if u.HasPermission(perm) {
		return false
	}
	u.permissions[perm] = positive
	return true
}

func (u *User) DelPermission(perm string) bool {
	if !u.HasPermission(perm) {
		return false
	}
	delete(u.permissions, perm)
	return false
}

func (u *User) AddGroup(group *Group) bool {
	if _, ok := u.GetGroup(group.GetName()); ok {
		return false
	}
	u.realGroups[group.GetName()] = group
	return true
}

func (u *User) DelGroup(name string) bool {
	if _, ok := u.GetGroup(name); !ok {
		return false
	}
	delete(u.realGroups, name)
	return false
}

func (u *User) delete() {
	for _, group := range u.realGroups {
		group.DelUser(u.GetName())
	}

	u.name = ""
	u.permissions = nil
	u.realGroups = nil
	u.groups = nil
	u.created = time.Time{}
}

func (u *User) MarshalJSON() ([]byte, error) {
	var groups []string

	for _, g := range u.realGroups {
		groups = append(groups, g.GetName())
	}

	return json.Marshal(&struct {
		Name        string
		Permissions map[string]bool
		Groups      []string
		Created     int64
	}{
		Name:        u.name,
		Permissions: u.permissions,
		Groups:      groups,
		Created:     u.created.Unix(),
	})
}
