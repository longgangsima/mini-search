package client

import "context"

// UpstreamClient 定义「下游怎么被调用」的契约；service 只依赖接口，不依赖具体实现（解耦）。
// Day 1 仅占位；Day 4 起会有 Catalog/Inventory/Pricing 等具体类型实现该接口。
type UpstreamClient interface {
	// Fetch 表示一次上游拉取；ctx 用于超时/取消，query 为搜索词。
	Fetch(ctx context.Context, query string) (string, error)
}
