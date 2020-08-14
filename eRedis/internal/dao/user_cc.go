package dao

import (
	"context"
	"fmt"
	"log"
	"strconv"

	m "github.com/aivuca/goms/eRedis/internal/model"

	"github.com/gomodule/redigo/redis"
)

//
func getRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 10)
}

func (d *dao) existUserCC(c context.Context, uid int64) (bool, error) {
	cc := d.redis
	key := getRedisKey(uid)
	exist, err := redis.Bool(cc.Do("EXISTS", key))
	if err != nil {
		err = fmt.Errorf("cc do EXISTS: %w", err)
		return exist, err
	}
	log.Printf("cc %v exist user, uid = %v", exist, uid)
	return exist, nil
}

func (d *dao) setUserCC(c context.Context, user *m.User) error {
	cc := d.redis
	key := getRedisKey(user.Uid)
	if _, err := cc.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...); err != nil {
		err = fmt.Errorf("cc do HMSET: %w", err)
		return err
	}
	log.Printf("cc set user = %v", *user)
	return nil
}

func (d *dao) getUserCC(c context.Context, uid int64) (*m.User, error) {
	cc := d.redis
	user := &m.User{}
	key := getRedisKey(uid)
	value, err := redis.Values(cc.Do("HGETALL", key))
	if err != nil {
		err = fmt.Errorf("cc do HGETALL: %w", err)
		return user, err
	}
	if err = redis.ScanStruct(value, user); err != nil {
		err = fmt.Errorf("cc ScanStruct: %w", err)
		return user, err
	}
	log.Printf("cc get user = %v", *user)
	return user, nil
}

func (d *dao) delUserCC(c context.Context, uid int64) error {
	cc := d.redis
	key := getRedisKey(uid)
	if _, err := cc.Do("DEL", key); err != nil {
		err = fmt.Errorf("cc do DEL: %w", err)
		return err
	}
	log.Printf("cc delete user, uid = %v", uid)
	return nil
}
