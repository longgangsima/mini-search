package main

import (
	"github.com/longgangsima/mini-search/internal/config"
	"github.com/longgangsima/mini-search/internal/handler"
	"github.com/longgangsima/mini-search/internal/service"
	"log"
	"net/http"
)

func main() {
	// 从环境变量等读出端口、超时等配置；集中在一处，避免魔法数字散落在代码里。
	cfg := config.Load()

	// 构造业务核心：负责 /search 背后的逻辑（当前是 mock，后面会接 client、校验等）。
	svc := service.NewSearchService()

	// 把 HTTP 请求转交给 service：解析 query、写 JSON、映射错误码都在这一层。
	searchH := handler.NewSearchHandler(svc)

	// 单独的健康检查处理器：负载均衡 / K8s 探活只打 /health，不碰业务路径。
	healthH := &handler.HealthHandler{}

	// 标准库路由表：把 URL 前缀映射到实现了 http.Handler 的对象上。
	mux := http.NewServeMux()

	// 注册搜索路由；没有这行，浏览器/curl 访问 /search 会 404。
	mux.Handle("/search", searchH)

	// 注册存活探针；与 /search 解耦，运维脚本只依赖这一路径。
	mux.Handle("/health", healthH)

	// 启动前打一行日志，方便确认进程真的在监听、端口读自配置。
	log.Printf("listening on %s", cfg.Port)
	// 阻塞监听 TCP；出错（例如端口被占用）时 log.Fatal 打印错误并退出进程。
	log.Fatal(http.ListenAndServe(cfg.Port, mux))

}
