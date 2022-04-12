package conf

import "maoim/pkg/mysql"

type Config struct {
	Mysql *mysql.Mysql

}

type RPCServer struct {
	Network string
	Addr string
}

type Bucket struct {
	Size int
}

type Channel struct {
	Size int
}