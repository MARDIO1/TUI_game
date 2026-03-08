package main

/*启动文件，理论不应该包含任何库*/
import (
	"TUI_game/core"
	"TUI_game/output"

	tea "charm.land/bubbletea/v2"
)

func main() {
	//播放固定音乐的调用方式
	bgmPlayer := output.NewPlayer()
	err := bgmPlayer.Play("assets/audio/pathos/AnimenzCallOfSilence.mp3")
	if err != nil {
		panic(err)
	}

	// // 播放随机音乐调用方式
	// bgmPlayer := output.NewPlayer()
	// err := bgmPlayer.PlaySceneBGM(output.SceneCandy)
	// if err != nil {
	// 	panic(err)
	// }

	m := core.Initial() //自定义结构体初始化，看似正常

	if _, err := tea.NewProgram(m).Run(); err != nil {
		panic(err)
	}
}
