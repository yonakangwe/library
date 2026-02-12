package helpers

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const aclSuffix = "-acl"
const ResponseMsgKey = "respmsg"

type ResponseMessage struct {
	Error   bool     `json:"error"`
	Message []string `json:"message"`
}

type Map map[string]interface{}

func StoreCache(key string, value interface{}) error {
	if Cache == nil {
		Cache = cache.New(12*60*time.Minute, 13*60*time.Minute)
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
