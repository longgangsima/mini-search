package service

import "context"

// SearchRequest 表示一次搜索的输入；字段名大写才能被 json 包对外序列化。
type SearchRequest struct {
	Query string `json:"query"`
	Store string `json:"store,omitempty"`
	Page  int    `json:"page,omitempty"`
}

// SearchResponse 表示返回给客户端的 JSON 形状（Day 1 为 mock）。
type SearchResponse struct {
	Results []string `json:"results"`
	Total   int      `json:"total"`
}

// SearchService 业务服务；后续会注入 client、repo 等依赖（Day 4+）。
type SearchService struct {
	// Day 4 会加 clients 字段
}

// NewSearchService 返回指针，让方法接收者共享同一份实例状态（见 questions.md）。
func NewSearchService() *SearchService {
	return &SearchService{}
}

// Search 核心业务入口；ctx 用于后续超时/取消，req 来自 handler 解析结果。
func (s *SearchService) Search(ctx context.Context, req SearchRequest) (SearchResponse, error) {
	_ = ctx // Day 1 占位；接下游后在这里尊重 ctx 取消
	// Day 1：不查真实数据源，只拼接 query，验证整条调用链打通。
	return SearchResponse{
		Results: []string{"result for " + req.Query},
		Total:   1,
	}, nil
}
