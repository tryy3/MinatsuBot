package permissionmanager

import "github.com/tryy3/minatsubot"

var (
	log        *minatsubot.Logger
	permission PermissionProvider
)

type Permission struct {
}

func (p Permission) Init() {
	log = minatsubot.PluginAPI.GetLogger("PermissionPlugin")
}

func (p Permission) Enable() {
	permission = PermissionProvider{
		groups:    map[string]*Group{},
		users:     map[string]*User{},
		userFile:  "users.json",
		groupFile: "groups.json",
	}

	minatsubot.PluginAPI.SetPermissionHandler(permission)

	permission.load()
}

func (p Permission) Disable() {
	permission.save()
}

type PermissionProvider struct {
	groups map[string]*Group
	users  map[string]*User

	userFile  string
	groupFile string

	userConfig  *config
	groupConfig *config
}

func (p *PermissionProvider) save() {
	p.userConfig.save(p.users)
	p.groupConfig.save(p.groups)
}

func (p *PermissionProvider) load() {
	p.userConfig = &config{file: p.userFile}
	p.groupConfig = &config{file: p.groupFile}

	p.users = map[string]*User{}
	p.groups = map[string]*Group{}

	p.groupConfig.load(&p.groups)
	p.userConfig.load(&p.users)

	for _, user := range p.users {
		for _, group := range user.groups {
			g, ok := p.GetGroup(group)
			if !ok {
				log.Errorf("The user %s contains an invalid group %s.", user.GetName(), group)
				continue
			}

			g.AddUser(user)
		}
	}
}

func (p *PermissionProvider) GetGroup(name string) (group *Group, ok bool) {
	group, ok = p.groups[name]
	return
}

func (p *PermissionProvider) CreateGroup(name string) bool {
	if _, ok := p.GetGroup(name); ok {
		return false
	}

	p.groups[name] = &Group{
		description: NewDescription(name),
		users:       map[string]*User{},
	}
	return true
}

func (p *PermissionProvider) DeleteGroup(name string) bool {
	group, ok := p.GetGroup(name)
	if !ok {
		return false
	}
	group.delete()
	delete(p.groups, name)
	return true
}

func (p *PermissionProvider) GetUser(name string) (user *User, ok bool) {
	user, ok = p.users[name]
	return
}

func (p *PermissionProvider) CreateUser(name string) bool {
	if _, ok := p.GetUser(name); ok {
		return false
	}
	p.users[name] = &User{
		description: NewDescription(name),
		realGroups:  map[string]*Group{},
	}
	return true
}

func (p *PermissionProvider) DeleteUser(name string) bool {
	user, ok := p.GetUser(name)
	if !ok {
		return false
	}
	user.delete()
	delete(p.users, name)
	return true
}

func (p PermissionProvider) HasUserPermission(name, perm string) bool {
	user, ok := p.GetUser(name)
	if !ok {
		return false
	}
	return user.HasPermission(perm)
}

func (p PermissionProvider) HasGroupPermission(name, perm string) bool {
	group, ok := p.GetGroup(name)
	if !ok {
		return false
	}
	return group.HasPermission(perm)
}
