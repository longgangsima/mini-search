package config

import(
	"os"
	"time"
)

type Config struct {
	Port string
	UpstreamTimeout time.Duration
}

func Load() Config{
	return Config{
		Port: getEnv("PORT", ":8080"), // if no env return :8080
		UpstreamTimeout: 500* time.Millisecond,
	}
}

func getEnv(key, fallback string) string {
	if v:= os.Getenv(key); v!=""{
		return v
	}
	return fallback
}
/**
v := os.Getenv(key)：尝试从操作系统中获取名为 key 的环境(环境变量.MD)。这里使用了短变量声明。
if 初始化语句：注意这里的 v := ... 是写在 if 关键字后面的。
这在 Go 中很常见，意思是：先执行声明并赋值，然后紧接着判断 v != ""。
作用域限制：变量 v 的生命周期仅限于这个 if 代码块内部。

*/