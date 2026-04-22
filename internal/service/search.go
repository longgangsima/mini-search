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

// 为什么要 return 回来一个 service 的地址呢？在 questions.md 有解释。
// * + & 是标准：通过 * 要 return 地址，然后 & return 给你地址。
//
// NewSearchService 返回指针，让方法接收者共享同一份实例状态。
func NewSearchService() *SearchService {
	return &SearchService{}
}

// 详细在 questions.md 有解释。
// (1) Receiver  (2) Name  (3) Input Parameters  (4) Return Values
func (s *SearchService) Search(ctx context.Context, req SearchRequest) (SearchResponse, error) {
	_ = ctx // Day 1 占位；接下游后在这里尊重 ctx 取消
	// Day 1：先返回 mock 数据；return 的是 req.Query 拼出来的串，不是真实数据源。
	return SearchResponse{
		Results: []string{"result for " + req.Query},
		Total:   1,
	}, nil
}
