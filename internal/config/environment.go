package config

import "fmt"

type ConnectionInfo struct {
	Host string
	Port string
	User string
	Pass string
	DB   string
}

func (c *ConnectionInfo) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Pass, c.Host, c.Port, c.DB)
}
