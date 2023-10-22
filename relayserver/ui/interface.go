package ui


// UI 对外暴漏的操作接口
type RelayServerUI interface {
	// 启动UI， 只能调用1次
	// 启动后会挂起程序， 直到程序退出
	StartUI()

	// 设置MetaMask状态, 
	// @param connected: true: 连接成功, false: 断开连接
	SetMetaMaskStatus(connected bool)

	// 设置硬件钱包连接状态
	// @param connected: true: 连接成功, false: 断开连接
	// @param walletAddress: 硬件钱包地址, 在断开连接时, 该参数无效
	SetWalletStatus(connected bool, walletAddress string)

	// 添加日志信息
	// @param log: 日志信息
	AddLog(log string)

	// 设置连接硬件钱包按钮的回调函数
	// @param callback: 回调函数
	SetConnectWalletButtonCallback(callback func())
}