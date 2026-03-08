package output

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

// Player 是播放器类
type Player struct {
	ctrl       *beep.Ctrl
	isInit     bool
	sampleRate beep.SampleRate
}

// NewPlayer 是实例化播放器的函数
func NewPlayer() *Player {
	return &Player{}
}

// Play 播放指定的 MP3 文件
func (p *Player) Play(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Open audio file failed: %v", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		return fmt.Errorf("Decode MP3 failed: %v", err)
	}

	// 扬声器全局初始化
	if !p.isInit || p.sampleRate != format.SampleRate {
		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			f.Close()
			return fmt.Errorf("Initialize speaker failed: %v", err)
		}
		p.isInit = true
		p.sampleRate = format.SampleRate
	}

	// 包装控制流并设置播放完毕后自动关闭文件
	p.ctrl = &beep.Ctrl{
		Streamer: beep.Seq(streamer, beep.Callback(func() {
			f.Close()
		})),
		Paused: false,
	}

	// 异步播放
	speaker.Play(p.ctrl)
	return nil
}

// Stop 停止播放并清空控制流
func (p *Player) Stop() {
	if p.ctrl != nil {
		speaker.Lock()
		p.ctrl.Streamer = nil // 切断音频流
		speaker.Unlock()
	}
}

// 动态音频相关实现
// SceneType 定义游戏中不同情景类型
type SceneType string

const (
	ScenePeace  SceneType = "peace"
	SceneFire   SceneType = "fire"
	SceneCandy  SceneType = "candy"
	ScenePathos SceneType = "pathos"
)

var playlist = map[SceneType][]string{
	ScenePeace: {
		"assets/audio/peace/OuterWilds.mp3",
		"assets/audio/peace/WetHands.mp3",
	},
	SceneFire: {},
	SceneCandy: {
		"assets/audio/candy/AnonCallOfSilence.mp3",
		"assets/audio/candy/AnonWanderer.mp3",
	},
	ScenePathos: {
		"assets/audio/pathos/AnimenzCallOfSilence.mp3",
	},
}

// 根据情景随机播放
func (p *Player) PlaySceneBGM(scene SceneType) error {
	// 获取该情景下的所有音乐列表
	tracks, exists := playlist[scene]

	if !exists || len(tracks) == 0 {
		return fmt.Errorf("can't find music for scene [%v]", scene)
	}

	// 随机挑一首歌的索引
	randomIndex := rand.Intn(len(tracks))
	selectedTrack := tracks[randomIndex]

	// 先把之前正在播放的音乐停掉
	p.Stop()

	// 播放选中的新音乐
	return p.Play(selectedTrack)
}
