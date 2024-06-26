package config

var port string

func SetPort(p string) {
	port = p
}

func GetPort() string {
	return port
}

func GetRoutesPort() string {
	return "2" + port
}
