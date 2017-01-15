package minatsubot

type Settings struct {
	Token   string
	Prefix  string
	Logging string
}

type Description struct {
	Version *Version
	Author  string
	Website string
	Info    string
}

type PluginDescription struct {
	Name string
	Description
}

// Version holds version data.
// Based on Semantic versioning 2.0.0 http://semver.org/
// MAJOR version when you make incompatible API changes.
// MINOR version when you add functionality in a backwards-compatible manner.
// PATCH version when you make backwards-compatible bug fixes.
type Version struct {
	Major string
	Minor string
	Patch string
}

func (v *Version) Get() string {
	return v.Major + "." + v.Minor + "." + v.Patch
}

type CommandDescription struct {
	Name        string
	Aliases     []string
	Description string
	Usage       string
}
