package account

import (
	"common/mysql"
	"common/redis"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("Test")
	fmt.Println(mysql.Initialize("192.168.189.142", "3306", "stayreal", "123456", "bokerchain"))
	fmt.Println(redis.Initialize("127.0.0.1", "6379", "", 2, 10))
	fmt.Println(NewAccount("bctestappid", "userid111", "phone111", "123456", 1))
}
