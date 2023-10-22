package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type mainWindow struct {
	app     fyne.App
	window  fyne.Window
	started bool

	// metamask状态的label
	metaMaskStatusLabel *widget.Label
	// wallet状态的label
	walletStatusLabel *widget.Label

	// 硬件钱包地址输入框
	walletAddressEntry *widget.Entry
	// 连接硬件钱包按钮
	connectWalletButton *widget.Button

	// 日志信息显示框
	logsText *widget.Entry
	// 清空日志按钮
	clearLogsButton *widget.Button
}

var _ RelayServerUI = (*mainWindow)(nil)

func NewMainWindow(title string) RelayServerUI {
	mainWindow := &mainWindow{}

	mainWindow.init(title)

	return mainWindow
}

// 初始化
func (mw *mainWindow) init(title string) {
	// 创建UI主app
	a := app.New()
	w := a.NewWindow(title)

	mw.app = a
	mw.window = w
	mw.started = false

	mw.makeUI()
}

// 创建UI界面
func (mw *mainWindow) makeUI() {
	// 最外层的 vbox
	outerVbox := container.NewVBox()

	// 第1行， 两个显示状态的label
	{
		// metamask状态的label
		metaMaskStatusLabel := widget.NewLabel("MetaMask 未连接")
		metaMaskStatusLabel.Alignment = fyne.TextAlignCenter
		metaMaskStatusLabel.Importance = widget.LowImportance

		// wallet状态的label
		walletStatusLabel := widget.NewLabel("HardWallet 未连接")
		walletStatusLabel.Alignment = fyne.TextAlignCenter
		walletStatusLabel.Importance = widget.LowImportance

		line1 := container.New(layout.NewGridLayout(2), metaMaskStatusLabel, walletStatusLabel)
		outerVbox.Add(line1)

		mw.metaMaskStatusLabel = metaMaskStatusLabel
		mw.walletStatusLabel = walletStatusLabel
	}

	// 第2行， 1个输入框，1个按钮
	// 因为 Hbox 会将输入框的长度缩减为最小，所以这里使用 FormLayout（它会将右侧元素展示为最大）
	{ // 硬件钱包地址输入框
		walletAddressEntry := widget.NewEntry()
		walletAddressEntry.SetPlaceHolder("请输入硬件钱包地址")
		walletAddressEntry.SetMinRowsVisible(1)

		// 连接硬件钱包按钮
		connectWalletButton := widget.NewButton("连接硬件钱包", nil)

		line2 := container.New(layout.NewFormLayout(), connectWalletButton, walletAddressEntry)
		outerVbox.Add(line2)

		mw.walletAddressEntry = walletAddressEntry
		mw.connectWalletButton = connectWalletButton
	}

	// 第3行，1个日志信息显示框
	{
		logsText := widget.NewMultiLineEntry()
		logsText.SetPlaceHolder("日志信息")
		logsText.SetMinRowsVisible(20)

		line3 := container.New(layout.NewGridLayout(1), logsText)
		outerVbox.Add(line3)

		mw.logsText = logsText
	}

	// 第4行，1个按钮, 清空日志按钮
	{
		clearLogsButton := widget.NewButton("清空日志", nil)

		line4 := container.New(layout.NewGridLayout(1), clearLogsButton)
		outerVbox.Add(line4)

		mw.clearLogsButton = clearLogsButton
	}

	mw.window.SetContent(outerVbox)
}

// 启动UI， 只能调用1次
func (mw *mainWindow) StartUI() {
	if !mw.started {
		mw.started = true
		mw.window.ShowAndRun()
	}
}

// 设置MetaMask状态,
// @param connected: true: 连接成功, false: 断开连接
func (mw *mainWindow) SetMetaMaskStatus(connected bool) {
	if connected {
		mw.metaMaskStatusLabel.SetText("MetaMask 已连接")
		mw.metaMaskStatusLabel.Importance = widget.SuccessImportance
	} else {
		mw.metaMaskStatusLabel.SetText("MetaMask 未连接")
		mw.metaMaskStatusLabel.Importance = widget.LowImportance
	}
}

// 设置硬件钱包连接状态
// @param connected: true: 连接成功, false: 断开连接
// @param walletAddress: 硬件钱包地址, 在断开连接时, 该参数无效
func (mw *mainWindow) SetWalletStatus(connected bool, walletAddress string) {
	if connected {
		mw.walletStatusLabel.SetText("HardWallet 已连接 (地址：" + walletAddress + ")")
		mw.walletStatusLabel.Importance = widget.SuccessImportance
	} else {
		mw.walletStatusLabel.SetText("HardWallet 未连接")
		mw.walletStatusLabel.Importance = widget.LowImportance
	}
}

// 添加日志信息
// @param log: 日志信息
func (mw *mainWindow) AddLog(log string) {
	mw.logsText.SetText(mw.logsText.Text + log + "\n")
}

// 设置连接硬件钱包按钮的回调函数
// @param callback: 回调函数
func (mw *mainWindow) SetConnectWalletButtonCallback(callback func()) {
	mw.connectWalletButton.OnTapped = callback
}
