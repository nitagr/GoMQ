package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

// understand the private and public vars
const (
	defaultNs      = "rsmq"
	defaultVt      = 30
	defaultDelay   = 0
	defaultMaxsize = 65536
)

type RedisSMQ struct {
	client *redis.Client
	ns     string
}

type QueueAttributes struct {
	Vt         uint
	Delay      uint
	Maxsize    int
	TotalRecv  uint64
	TotalSent  uint64
	Created    uint64
	Modified   uint64
	Msgs       uint64
	HiddenMsgs uint64
}

type QueueMessage struct {
	ID      string
	Message string
	Rc      uint64
	Fr      time.Time
	Sent    time.Time
}

// NewRedisSMQ creates and returns new rsmq client
func NewRedisSMQ(client *redis.Client, ns string) *RedisSMQ {
	if client == nil {
		panic("")
	}

	if ns == "" {
		ns = defaultNs
	}
	if !strings.HasSuffix(ns, ":") {
		ns += ":"
	}

	rsmq := &RedisSMQ{
		client: client,
		ns:     ns,
	}

	return rsmq
}
func main() {
	fmt.Println("Hi! Welcome")
	fmt.Println(scriptPopMessage)

}
