package conf

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(config.Http.Host)
	fmt.Println(config.Http.Port)
	fmt.Println(config.Mysql.Host)
	fmt.Println(config.Mysql.Port)
	fmt.Println(config.Mysql.Username)
	fmt.Println(config.Mysql.Password)
	fmt.Println(config.Mysql.Database)

}
