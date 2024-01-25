package models

type Config struct {
	DB     DBConf
	Server ServerConf
}

type DBConf struct {
	UserName string
	Pass     string
	Database string
	Host     string
	Port     string
}

type ServerConf struct {
	Port string
}
