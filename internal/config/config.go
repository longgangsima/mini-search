package config

import (
	"os"
	"time"
)

// Config 保存进程级配置；main 启动时读一次，避免端口/超时写死在各处。
type Config struct {
	// Port 传给 ListenAndServe，例如 ":8080"；可用环境变量 PORT 覆盖。
	Port string
	// UpstreamTimeout 留给后面调下游时用（Day 1 可先不用）。
	UpstreamTimeout time.Duration
}

// Load 组装默认配置并从环境变量覆盖；集中入口，方便以后加字段。
func Load() Config {
	return Config{
		Port:            getEnv("PORT", ":8080"),
		UpstreamTimeout: 500 * time.Millisecond,
	}
}

// getEnv 读环境变量；空则用 fallback，避免调用方重复写 if v != "" 逻辑。
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
