package central

import "time"

var (
	livenessCheckDuration = time.Duration(10 * time.Second)
	staleDuration         = time.Duration(1 * time.Minute)
	expireDuration        = time.Duration(5 * time.Minute)
)

var (
	DefaultPort = 9100
)
