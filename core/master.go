package core

/*master.go
游戏主干文件，仍然和bubble库耦合
Master类，目前是输入输出的数据*/
import (
	"TUI_game/input"
	"TUI_game/output"
	"time"

	tea "charm.land/bubbletea/v2"
)

type TickMsg time.Time

type Master struct {
	In      *input.Ctrl  // 输入组件
	Out     *output.Ctrl // 输出组件
	Printer *Printer     // 打字机组件
	Tick_ms int          // 心跳间隔 10ms (100Hz)
}

func Initial() *Master {
	m := &Master{
		In:      input.New(),
		Out:     output.New(),
		Printer: NewPrinter(),
		Tick_ms: 10,
	}
	
	// 测试数据 - 使用更慢的速度和更短的文本
	m.Printer.Set([]Text{
		{Time: 0.5, Text: "Test."},      // 每个字符0.5秒
		{Time: 0.3, Text: "Hello."},     // 每个字符0.3秒
		{Time: 0.4, Text: "World!"},     // 每个字符0.4秒
	})
	
	return m
}

func (m *Master) Init() tea.Cmd {
	return tea.Batch(m.In.Focus(), m.next_tick())
}

func (m *Master) run() {
	if m.Printer.Active() {
		// 更新打字机
		m.Printer.Update(float64(m.Tick_ms) / 1000.0)
		// 显示打字机文本
		m.Out.Set(m.Printer.Text())
	} else {
		// 正常输入模式
		val := m.In.Get()
		m.Out.Set("V2 Engine 100Hz | Input: " + val)
	}
}

func (m *Master) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		m.In.Update(msg)
	case tea.WindowSizeMsg:
		m.Out.Resize(msg.Width, msg.Height)
		m.In.SetWidth(msg.Width)
	case TickMsg:
		m.run()
		return m, m.next_tick()
	}
	return m, nil
}

// AIchange: V2 规范 - 使用 tea.NewView 构造
func (m *Master) View() tea.View {
	// 1. 获取拼接后的 UI 字符串
	ui_str := m.Out.Render(m.In.View())

	// 2. 创建 View 结构体
	v := tea.NewView(ui_str)

	// 3. 在这里声明“全屏模式” (取代了之前的 WithAltScreen)
	v.AltScreen = true

	return v
}

func (m *Master) next_tick() tea.Cmd {
	return tea.Tick(time.Duration(m.Tick_ms)*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
