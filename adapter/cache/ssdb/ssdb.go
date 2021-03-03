package ssdb

import (
	"github.com/Sonek-HoangBui/beego/v2/adapter/cache"
	ssdb2 "github.com/Sonek-HoangBui/beego/v2/client/cache/ssdb"
)

// NewSsdbCache create new ssdb adapter.
func NewSsdbCache() cache.Cache {
	return cache.CreateNewToOldCacheAdapter(ssdb2.NewSsdbCache())
}

func init() {
	cache.Register("ssdb", NewSsdbCache)
}
