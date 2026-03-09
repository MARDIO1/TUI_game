package core

import (
	"fmt"
	"strings"
)

// Printer 打字机组件
type Printer struct {
	queue   []Text  // 文本队列
	idx     int     // 当前文本索引
	pos     int     // 当前字符索引
	timer   float64 // 计时器（秒）
	active  bool    // 是否激活
}

// colorToANSI 将RGB颜色转换为ANSI颜色代码
func colorToANSI(c [3]uint8) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", c[0], c[1], c[2])
}

// NewPrinter 创建打字机
func NewPrinter() *Printer {
	return &Printer{}
}

// Set 设置文本队列
func (p *Printer) Set(queue []Text) {
	p.queue = queue
	p.idx = 0
	p.pos = 0
	p.timer = 0
	p.active = true
}

// Update 更新打字机状态
func (p *Printer) Update(dt float64) {
	if !p.active || p.idx >= len(p.queue) {
		p.active = false
		return
	}

	t := p.queue[p.idx]
	p.timer += dt

	if p.timer >= t.Time {
		p.timer = 0
		if p.pos < len(t.Text) {
			p.pos++
		} else {
			p.idx++
			p.pos = 0
		}
	}
}

// Text 获取已显示的文本（带颜色）
func (p *Printer) Text() string {
	var result strings.Builder
	
	for i := 0; i < p.idx; i++ {
		if i < len(p.queue) {
			// 完整显示已完成的文本
			color := colorToANSI(p.queue[i].Color)
			result.WriteString(color)
			result.WriteString(p.queue[i].Text)
			result.WriteString("\033[0m\n")
		}
	}
	
	// 显示当前正在输入的文本
	if p.idx < len(p.queue) && p.pos > 0 {
		color := colorToANSI(p.queue[p.idx].Color)
		result.WriteString(color)
		result.WriteString(p.queue[p.idx].Text[:p.pos])
		result.WriteString("\033[0m")
	}
	
	return result.String()
}

// Active 检查是否激活
func (p *Printer) Active() bool {
	return p.active
}