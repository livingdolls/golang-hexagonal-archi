package config

type DBConfig struct {
	Driver                  string
	Url                     string
	ConnMaxLifetimeInMinute int
	MaxOpenConns            int
	MaxIdleConns            int
}
