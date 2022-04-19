package id

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var node *snowflake.Node
var once sync.Once

func init() {
	once.Do(func() {
		var err error
		node, err = snowflake.NewNode(1)
		if err != nil {
			panic(err)
		}
	})
}

func Generate() string {
	return node.Generate().String()
}
