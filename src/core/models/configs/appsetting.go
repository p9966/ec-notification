package configs

type Appsetting struct {
	Name                string
	ConfigurationCenter ConfigurationCenter
}

type ConfigurationCenter struct {
	Local Local
	Files []File
}

type Local struct {
	Folder string
}

type File struct {
	Name   string
	Path   string
	Source string
}
