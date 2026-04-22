package service

import "context"

type SearchRequest struct {
	Query string `json:"query"`
	Store string `json:"store,omitempty"`
	Page int `json: "page,omitempty"`
}

type SearchResponse struct {
	Results [] string `json:"results"`
	Total int `json:"total"`
}

type SearchService struct {
	// Day 4 add client code
}
// 为什么要return 回来一个service 的地址呢？ 在questions.md 有解释
// * + & 是标准 我通过 * 要 return 地址，然后& return给你地址
func NewSearchService() *SearchService{
	return &SearchService{}
}


// 详细在questions.md 有解释
//   (1) Receiver      (2) Name  (3) Input Parameters                   (4) Return Values
func (s *SearchService) Search(ctx context.Context, req SearchRequest)(SearchResponse, error){
	// Day 1 先返回 mock 数据
	return SearchResponse{
		Results: []string{"result for " + req.Query},
		Total: 1,
	},nil
}