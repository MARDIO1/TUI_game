package output

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
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
func (p *Player) Play(songName string) error {

	// 先把之前可能在播放的音乐停掉（防重叠）
	p.Stop()

	// 如果有,那么去掉 MP3 后缀
	songName = strings.TrimSuffix(songName, ".mp3")

	// 去字典里查全路径
	fullPath, exists := trackMap[songName]
	if !exists {
		return fmt.Errorf("Can't find audio [%s]", songName)
	}

	// 使用查到的全路径打开文件
	f, err := os.Open(fullPath)
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

// 情景类型字典,key是情景类型,value是这个情景下所有音乐的完整路径列表
var playlist = make(map[SceneType][]string)

// 音乐路径字典,key是音乐文件的名字,value是它的完整路径
var trackMap = make(map[string]string)

// 字典init
func init() {
	basePath := "assets/audio"

	folders, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Printf("Can't read audio root directory: %v\n", err)
		return
	}

	// 遍历这些情景文件夹
	for _, folder := range folders {
		if !folder.IsDir() {
			continue // 如果读到的不是文件夹直接跳过
		}

		// 文件夹的名字刚好就是我们的 SceneType
		scene := SceneType(folder.Name())
		folderPath := filepath.Join(basePath, folder.Name())

		// 读取这个情景文件夹里的所有文件
		files, err := os.ReadDir(folderPath)
		if err != nil {
			continue
		}

		// 遍历里面的每一首歌
		for _, file := range files {
			// 必须是普通文件，并且后缀名是.mp3
			if !file.IsDir() && filepath.Ext(file.Name()) == ".mp3" {

				fullPath := filepath.Join(folderPath, file.Name())

				songName := strings.TrimSuffix(file.Name(), ".mp3")

				playlist[scene] = append(playlist[scene], songName)
				// 所有歌曲路径的初始化
				trackMap[songName] = fullPath
			}
		}
	}
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

	// 播放选中的新音乐
	return p.Play(selectedTrack)
}
