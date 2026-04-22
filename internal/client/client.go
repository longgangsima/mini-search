package client
import "context"

// UpstreamClient 是下游服务的通接口
// 用通俗的话说，这段代码是在**“定规矩”**。
// inteface -> 接口: qustions.md 有解释
// 用这种方式 属于decouple.
// UpstreamClient -> is a contract
// Day 1 先定义这个contract，Day 4 会有真实现

type UpstreamClient interface{
	func(ctx context.Context, query string)(string, error)
}