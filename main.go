package main

import (
	"TUI_game/core"

	tea "charm.land/bubbletea/v2"
)

func main() {
	m := core.Initial()
	// AIchange: V2 建议移除启动选项，全部在 View 中声明
	if _, err := tea.NewProgram(m).Run(); err != nil {
		panic(err)
	}
}
