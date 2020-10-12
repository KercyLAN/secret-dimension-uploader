package upload

type Node interface {
	// 获取节点id，
	GetId() string
	// 获取节点host。
	GetHost() string
}

// 默认的单机节点
type StandaloneNode struct {

}

func (slf StandaloneNode) GetId() string {
	return "standalone-node"
}

func (slf StandaloneNode) GetHost() string {
	return "localhost"
}
