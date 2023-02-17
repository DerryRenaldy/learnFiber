package customercache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/go-redis/redis/v8"
	"github.com/valyala/fasthttp"
	"time"
)

type CustomerCache struct {
	client    *redis.Client
	namespace string
	l         logger.ILogger
	timeout   time.Duration
}

func (o CustomerCache) getCtx(ctx *fasthttp.RequestCtx) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 10*time.Second)
}

func (o CustomerCache) redisKey(CustomerID int64) string {
	return fmt.Sprintf("%v:Customer:%v", o.namespace, CustomerID)
}

func (o CustomerCache) GetByCode(ctx *fasthttp.RequestCtx, CustomerID int64) (*entity.Customer, error) {
	var cancel context.CancelFunc
	cont := context.Background()
	cont, cancel = o.getCtx(ctx)
	defer cancel()

	rs, err := o.client.Get(cont, o.redisKey(CustomerID)).Result()
	if err != nil {
		return nil, err
	}
	out := &entity.Customer{}
	err = json.Unmarshal([]byte(rs), &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (o CustomerCache) Set(ctx *fasthttp.RequestCtx, obj *entity.Customer) error {
	var cancel context.CancelFunc
	cont := context.Background()
	cont, cancel = o.getCtx(ctx)
	defer cancel()

	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	o.l.Infof("[redis key] : %s", o.redisKey(obj.ID))
	return o.client.Set(cont, o.redisKey(obj.ID), string(data), o.timeout).Err()
}

func (o CustomerCache) Remove(ctx *fasthttp.RequestCtx, CustomerID int64) error {
	var cancel context.CancelFunc
	cont := context.Background()
	cont, cancel = o.getCtx(ctx)
	defer cancel()
	return o.client.Del(cont, o.redisKey(CustomerID)).Err()
}

func (o CustomerCache) GetList(ctx *fasthttp.RequestCtx) ([]*entity.Customer, error) {
	var cancel context.CancelFunc
	cont := context.Background()
	cont, cancel = o.getCtx(ctx)
	defer cancel()

	listKeys, err := o.client.Keys(cont, fmt.Sprintf("%v:Customer:*", o.namespace)).Result()
	if err != nil {
		return nil, err
	}

	var results []*entity.Customer

	var str []*redis.StringCmd
	pipe := o.client.Pipeline()
	for _, item := range listKeys {
		str = append(str, pipe.Get(cont, item))
	}
	_, err = pipe.Exec(cont)
	if err != nil {
		return nil, err
	}

	for _, item := range str {
		rs, err := item.Result()
		if err != nil {
			return nil, err
		}

		out := &entity.Customer{}
		err = json.Unmarshal([]byte(rs), &out)
		if err != nil {
			return nil, err
		}
		results = append(results, out)
	}
	return results, nil
}

func NewRedis(client *redis.Client, namespace string, timeout time.Duration, log logger.ILogger) *CustomerCache {
	return &CustomerCache{
		l:         log,
		client:    client,
		namespace: namespace,
		timeout:   timeout,
	}
}
