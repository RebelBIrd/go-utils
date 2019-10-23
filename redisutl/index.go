package redisutl

import (
	"github.com/go-redis/redis"
	"github.com/qinyuanmao/go-utils/logutl"
	"github.com/qinyuanmao/go-utils/strutl"
	"net/url"
	"time"
)

// redis配置引擎
type RedisConf struct {
	Url           string `yaml:"url"`      //redis地址
	Username      string `yaml:"username"` //redis用户名
	Password      string `yaml:"password"` //redis密码
	Database      int    `yaml:"database"` //redis数据库
	*redis.Client                          //redis客户端引擎
}

func (conf *RedisConf) InitRedis() {
	redisOpt := redis.Options{
		Password: conf.Password,
		DB:       conf.Database,
		Addr:     conf.Url,
	}
	if conf.Username != "" {
		redisUrl, _ := url.Parse(strutl.ConnString("redis://", conf.Username, ":", conf.Password, "@", conf.Url))
		redisOpt.Addr = redisUrl.Host
	}
	conf.Client = redis.NewClient(&redisOpt)
	if pong, err := conf.Client.Ping().Result(); err != nil {
		logutl.Error(err.Error())
	} else {
		logutl.Info(pong)
	}
}

func (this *RedisConf) Save(key string, value interface{}) {
	if this.IsExist(key) {
		this.Del(key)
	}
	cmd := this.Set(key, value, time.Hour*2)
	if cmd.Val() != "OK" {
		logutl.Error(cmd.Err().Error())
	}
}

func (this *RedisConf) IsExist(key string) bool {
	if intCmd := this.Exists(key); intCmd.Val() > 0 {
		return true
	} else {
		return false
	}
}

func (this *RedisConf) ExpireKey(key string) bool {
	isSuccess := this.Expire(key, time.Hour*2)
	if isSuccess.Err() != nil {
		logutl.Error(isSuccess)
	}
	return isSuccess.Val()
}

func (this *RedisConf) GetAllKeys() (ids []string) {
	var err error
	ids, err = this.Keys("*").Result()
	if err != nil {
		logutl.Error(err)
	}
	return
}
