package input

import (
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
)

type Ctrl struct {
	Model textarea.Model
}

func New() *Ctrl {
	ta := textarea.New()
	ta.Placeholder = "输入指令..." // 手册示例：直接赋值
	ta.Focus()                 // 手册：获取焦点
	ta.SetHeight(3)            // 手册：设置高度
	return &Ctrl{Model: ta}
}

func (c *Ctrl) Focus() tea.Cmd { return c.Model.Focus() }
func (c *Ctrl) Get() string    { return c.Model.Value() }
func (c *Ctrl) SetWidth(w int) { c.Model.SetWidth(w) }

// AIchange: 手册中 View() 返回 string，但 Master 需要 tea.View
func (c *Ctrl) View() string {
	return c.Model.View()
}

func (c *Ctrl) Update(msg tea.Msg) {
	var cmd tea.Cmd
	c.Model, cmd = c.Model.Update(msg)
	_ = cmd
}
