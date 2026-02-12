package helpers

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func Init() {
	Cache = cache.New(15*time.Minute, 15*time.Minute)
}
