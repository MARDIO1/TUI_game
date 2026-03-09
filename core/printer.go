package core

// Printer 打字机组件
type Printer struct {
	queue  []Text  // 文本队列
	idx    int     // 当前文本索引
	pos    int     // 当前字符索引
	text   string  // 已显示的文本
	timer  float64 // 计时器（秒）
	active bool    // 是否激活
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
	p.text = ""
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

	// 显示字符直到计时器不足
	for p.timer >= t.Time && p.active {
		p.timer -= t.Time
		
		if p.pos < len(t.Text) {
			p.text += string(t.Text[p.pos])
			p.pos++
		} else {
			// 当前文本显示完毕
			p.idx++
			if p.idx >= len(p.queue) {
				p.active = false
				return
			}
			p.pos = 0
			p.text += "\n"
			t = p.queue[p.idx] // 切换到下一个文本
		}
	}
}

// Text 获取已显示的文本
func (p *Printer) Text() string {
	return p.text
}

// Active 检查是否激活
func (p *Printer) Active() bool {
	return p.active
}