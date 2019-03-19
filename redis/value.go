package redis

import "github.com/garyburd/redigo/redis"

type Value struct {
	Val interface{}
	Err error
}

func NewValue(val interface{}, err error) Value {
	return Value{
		Val: val,
		Err: err,
	}
}

func (val Value) HasValue() bool {
	return val.Err == nil
}

func (val Value) Int() (int, error) {
	return redis.Int(val.Val, val.Err)
}
