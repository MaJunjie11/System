package common_logic

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var LocalCache = cache.New(time.Second*60, time.Second*10)
