package configs

type ServerList struct {
	Servers []Server
}

type Server struct {
	Name     string
	HostList []Host
}

type Host struct {
	Channel string
	Host    string
}
