package conf

import "time"

const (
	AuthCodeCacheKey  = "gs:auth:code:"
	AuthTokenCacheKey = "gs:auth:token:"
)

const (
	AuthCodeCacheKeyTimeout         = time.Minute * 2     // code 过期时间
	AuthRefreshTokenCacheKeyTimeout = time.Hour * 24 * 30 // RefreshToken 过期时间
	AuthAccessTokenCacheKeyTimeout  = time.Hour * 24      // AccessToken 过期时间
)
