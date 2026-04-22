package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/longgangsima/mini-search/internal/service"
)

// SearchHandler 实现 http.Handler；负责 HTTP 层，把请求转给 SearchService。
type SearchHandler struct {
	// svc 是业务核心；handler 只解析 HTTP、写响应，不塞业务规则。
	svc *service.SearchService
}

// NewSearchHandler 构造处理器并注入依赖；main 里用这一行完成「接线」。
// 1. 接线：h.svc 存下了 svc 的地址
func NewSearchHandler(svc *service.SearchService) *SearchHandler {
	return &SearchHandler{svc: svc}
}

// ServeHTTP 每来一个 /search 请求调用一次；由 ServeMux 自动调度。
// 2. 使用：当请求触发此方法时，h 已经带有了初始化好的 svc
func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[search] start: q=%s", r.URL.Query().Get("q"))
	defer log.Printf("[search] done")

	// 取出请求级 context；后面接超时/取消时从这里往下传（Day 7 起会用上）。
	ctx := r.Context()

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page")) // 转不了就是 0

	// 从 query string 取参数；GET /search?q=... 的标准读法。
	req := service.SearchRequest{
		Query: q.Get("q"),
		Store: q.Get("store"),
		Page:  page,
	}

	// 调用业务层；错误在这里映射成 HTTP（当前只简单 500）。
	// 这里的 h.svc 就是 main 传进来的那个零件：Search from service
	resp, err := h.svc.Search(ctx, req)
	if err != nil {
		// day 1: http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
		// Day 2:
		var verr *service.ValidationError
		if errors.As(err, &verr) {
			writeJSONError(w, http.StatusBadRequest, verr.Error())
			return
		}
		writeJSONError(w, http.StatusInternalServerError, "internal error")
		return
	}
	// 告诉客户端 body 是 JSON；浏览器和 curl 都依赖这个 header。
	w.Header().Set("Content-Type", "application/json")
	// 把结构体编码进 ResponseWriter；失败时通常打日志（此处略）。
	json.NewEncoder(w).Encode(resp)
}

func writeJSONError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// HealthHandler 探活专用；与业务路由分离，运维只打 /health。
type HealthHandler struct{}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
