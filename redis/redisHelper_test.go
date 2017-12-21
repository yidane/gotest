package redis

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestRedisHelper_Dail(t *testing.T) {
	type fields struct {
		Server string
		con    *redis.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "localhost", fields: fields{Server: "127.0.0.1:6379"}, wantErr: false},
		{name: "localhostErrorPort", fields: fields{Server: "127.0.0.1:6380"}, wantErr: true},
		{name: "emptyServer", fields: fields{Server: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisHelper{
				Server: tt.fields.Server,
				con:    tt.fields.con,
			}
			if err := r.Dail(); (err != nil) != tt.wantErr {
				t.Errorf("RedisHelper.Dail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
