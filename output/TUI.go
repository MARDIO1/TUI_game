package output

import (
	"charm.land/bubbles/v2/viewport"
	"charm.land/lipgloss/v2"
)

type Ctrl struct {
	VP viewport.Model
}

func New() *Ctrl { return &Ctrl{} }

func (c *Ctrl) Set(s string) {
	c.VP.SetContent(s) // 手册：设置内容
	c.VP.GotoBottom()  // 手册：滚动到底部
}

func (c *Ctrl) Resize(w, h int) {
	content_h := h - 6
	if c.VP.Width() == 0 {
		// 手册：使用选项模式构造
		c.VP = viewport.New(
			viewport.WithWidth(w),
			viewport.WithHeight(content_h),
		)
	} else {
		c.VP.SetWidth(w)
		c.VP.SetHeight(content_h)
	}
}

func (c *Ctrl) Render(inView string) string {
	line := "------------------------------------------"
	// 手册：JoinVertical 组合多个组件视图
	return lipgloss.JoinVertical(lipgloss.Left, c.VP.View(), line, inView)
}
