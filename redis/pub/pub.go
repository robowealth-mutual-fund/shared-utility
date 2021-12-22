package pub

import (
	"encoding/json"

	"github.com/adjust/redismq"
)

// RedisPub ...
type RedisPub struct {
	host, port, password string
}

// NewRedisPub ...
func NewRedisPub(host, port, password string) *RedisPub {
	return &RedisPub{
		host:     host,
		port:     port,
		password: password,
	}
}

// Put ...
func (r *RedisPub) Put(queueName string, req interface{}, redisDB ...int64) error {
	var db int64 = 9
	if len(redisDB) > 0 {
		db = redisDB[0]
	}
	queue := redismq.CreateQueue(r.host, r.port, r.password, db, queueName)
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return queue.Put(string(payload))
}
