package configs

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var MySQLConfig = &DBConfig{
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "",
	DBName:   "golomba_db",
}
