package services

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

const aclSuffix = "-acl"
const ResponseMsgKey = "respmsg"

type ResponseMessage struct {
	Error   bool     `json:"error"`
	Message []string `json:"message"`
}

func Init() {
	//var err error
	Cache = cache.New(5*time.Minute, 15*time.Minute)
}

func StoreCache(key string, value interface{}) error {
	if Cache == nil {
		Cache = cache.New(30*time.Minute, 90*time.Minute)
	}
	_, ok := Cache.Get(key)
	if ok {
		return nil
	}
	return Cache.Add(key, value, 30*time.Minute)
}

func GetCache(key string) (interface{}, bool) {
	return Cache.Get(key)
}

func ClearCache(key string) {
	Cache.Set(key, nil, 1*time.Second)
}

func GetACLKey(email string) string {
	return email + aclSuffix
}

func SetResponseMessage(isError bool, message ...string) error {
	err := &ResponseMessage{
		Error:   isError,
		Message: message,
	}
	return Cache.Add(ResponseMsgKey, err, 500*time.Millisecond)
}
