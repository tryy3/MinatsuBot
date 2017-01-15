package permissionmanager

import "time"

func NewDescription(name string) *description {
	return &description{
		name:        name,
		permissions: map[string]bool{},
		created:     time.Now(),
	}
}

type description struct {
	name        string
	permissions map[string]bool
	created     time.Time
}

func (d *description) GetName() string {
	return d.name
}

func (d *description) GetPermissions() map[string]bool {
	return d.permissions
}

func (d *description) GetCreated() time.Time {
	return d.created
}

func (d *description) AddPermission(perm string, positive bool) bool {
	if _, ok := d.permissions[perm]; ok {
		return false
	}
	d.permissions[perm] = positive
	return true
}

func (d *description) DelPermission(perm string) bool {
	if _, ok := d.permissions[perm]; !ok {
		return false
	}
	delete(d.permissions, perm)
	return false
}
