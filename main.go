package main

/*启动文件，理论不应该包含任何库*/
import (
	"TUI_game/core"

	tea "charm.land/bubbletea/v2"
)

func main() {
	m := core.Initial()//自定义结构体初始化，看似正常
	// AIchange: V2 建议移除启动选项，全部在 View 中声明
	/*
		tea.NewProgram(m): 创建一个TUI程序实例，接管终端
		.Run(): 这是阻塞调用。程序会在这里进入死循环
		_, err := ...: Run() 返回两个值：最终的状态 Model（我们不需要，所以用 _ 忽略）和可能发生的错误（err）
		注意go的特性就是返回值可以有多个
		if ...; err != nil: 这是一个组合语句。它先执行运行操作，紧接着判断：如果运行过程中出错了（比如终端环境不支持、权限不足等），就执行 panic(err) 强制退出并打印报错。*/
	if _, err := tea.NewProgram(m).Run(); err != nil {
		panic(err)
	}
}
