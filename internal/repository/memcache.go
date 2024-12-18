package repository

import "github.com/bradfitz/gomemcache/memcache"

var cache = memcache.New("localhost:11211")
