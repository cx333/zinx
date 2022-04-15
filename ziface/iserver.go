package ziface

type IServer interface {

	// Start 启动
	Start()

	// Stop 停止
	Stop()

	// Serve 开启业务服务方法
	Serve()
}
