package core

import (
	"TUI_game/input"
	"TUI_game/output"
	"time"

	tea "charm.land/bubbletea/v2"
)

type TickMsg time.Time

type Master struct {
	In      *input.Ctrl
	Out     *output.Ctrl
	Tick_ms int
}

func Initial() *Master {
	return &Master{
		In:      input.New(),
		Out:     output.New(),
		Tick_ms: 10,
	}
}

func (m *Master) Init() tea.Cmd {
	return tea.Batch(m.In.Focus(), m.next_tick())
}

func (m *Master) run() {
	val := m.In.Get()
	m.Out.Set("V2 Engine 100Hz | Input: " + val)
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
