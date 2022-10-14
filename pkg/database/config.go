package database

import "fmt"

type Config struct {
	Username    string
	Password    string
	Host        string
	Port        uint16
	DbName      string
	Timeout     int
	MaxAttempts int
}

func (c Config) DsnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%v/%s?connect_timeout=%v", c.Username, c.Password, c.Host, c.Port, c.DbName, c.Timeout)
}
