package ziface

//IServer 定义服务接口,服务模块抽象层
type IServer interface {
	//启动
	Start()

	//关闭
	Stop()

	//开启业务服务方法
	Serve()
}
