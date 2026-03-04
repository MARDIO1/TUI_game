# Bubbles v2 - AI友好文档

## 概述

Bubbles v2 (`charm.land/bubbles/v2`) 是Charmbracelet的TUI组件库，专为Bubble Tea v2设计。提供了一系列可复用的UI组件，用于构建终端用户界面。

### 核心特性
- **现代化API**：使用Getter/Setter方法替代直接字段访问
- **类型安全**：使用`image/color.Color`替代字符串颜色
- **明确样式**：需要显式指定Light/Dark模式
- **函数式选项**：使用`WithXxx()`选项模式

### 导入路径
```go
import (
    "charm.land/bubbles/v2/textarea"
    "charm.land/bubbles/v2/viewport"
    "charm.land/bubbles/v2/textinput"
    "charm.land/bubbles/v2/list"
    "charm.land/bubbles/v2/table"
    "charm.land/bubbles/v2/progress"
    "charm.land/bubbles/v2/spinner"
    "charm.land/bubbles/v2/help"
    "charm.land/bubbles/v2/key"
    tea "charm.land/bubbletea/v2"
    "charm.land/lipgloss/v2"
)
```

---

## Textarea - 多行文本输入

### 概述
多行文本输入组件，支持滚动、粘贴、单词导航等高级功能。

### 构造函数
```go
// 基本创建
ta := textarea.New()

// 带选项创建
ta := textarea.New(
    textarea.WithWidth(80),
    textarea.WithHeight(10),
    textarea.WithPlaceholder("输入内容..."),
    textarea.WithMaxLength(1000),
)
```

### 关键方法
```go
// 焦点控制
ta.Focus()           // 获取焦点
ta.Blur()            // 失去焦点
ta.Focused() bool    // 检查是否有焦点

// 内容操作
ta.SetValue("文本")           // 设置内容
ta.Value() string            // 获取内容
ta.Reset()                   // 重置内容
ta.InsertString("插入")       // 在光标处插入

// 光标控制
ta.SetCursorColumn(5)        // 设置光标列位置
ta.Column() int              // 获取当前列
ta.SetCursorLine(2)          // 设置光标行
ta.Line() int                // 获取当前行

// 滚动控制
ta.SetYOffset(10)            // 设置垂直偏移
ta.YOffset() int             // 获取垂直偏移
ta.GotoBottom()              // 滚动到底部
ta.GotoTop()                 // 滚动到顶部

// 尺寸控制
ta.SetWidth(80)              // 设置宽度
ta.Width() int               // 获取宽度
ta.SetHeight(20)             // 设置高度
ta.Height() int              // 获取高度
```

### 样式配置
```go
// 获取默认样式（需要指定Light/Dark模式）
isDark := true // 或通过tea.BackgroundColorMsg检测
styles := textarea.DefaultStyles(isDark)

// 自定义样式
styles.Focused.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
styles.Focused.Text = lipgloss.NewStyle().Bold(true)
styles.Blurred.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

ta.SetStyles(styles)
```

### 键绑定
```go
// 获取默认键绑定
km := textarea.DefaultKeyMap()

// 自定义键绑定
km.InsertNewline = key.NewBinding(
    key.WithKeys("enter", "ctrl+m"),
    key.WithHelp("enter", "插入新行"),
)

ta.SetKeyMap(km)
```

### 示例：基本使用
```go
type Model struct {
    textarea textarea.Model
}

func NewModel() Model {
    ta := textarea.New()
    ta.Placeholder = "输入指令..."
    ta.Focus()
    ta.SetHeight(3)
    return Model{textarea: ta}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    m.textarea, cmd = m.textarea.Update(msg)
    return m, cmd
}

func (m Model) View() string {
    return m.textarea.View()
}
```

---

## Viewport - 视口组件

### 概述
用于垂直滚动内容的视口组件，支持高亮、行号、软换行等高级功能。

### 构造函数
```go
// 基本创建
vp := viewport.New()

// 带选项创建
vp := viewport.New(
    viewport.WithWidth(80),
    viewport.WithHeight(24),
    viewport.WithContent("初始内容"),
)
```

### 关键方法
```go
// 尺寸控制
vp.SetWidth(80)              // 设置宽度
vp.Width() int               // 获取宽度
vp.SetHeight(24)             // 设置高度
vp.Height() int              // 获取高度

// 内容操作
vp.SetContent("新内容")      // 设置内容
vp.GetContent() string       // 获取内容
vp.SetContentLines([]string{"行1", "行2"}) // 设置行内容

// 滚动控制
vp.SetYOffset(10)            // 设置垂直偏移
vp.YOffset() int             // 获取垂直偏移
vp.LineUp(5)                 // 向上滚动5行
vp.LineDown(5)               // 向下滚动5行
vp.PageUp()                  // 向上翻页
vp.PageDown()                // 向下翻页
vp.GotoTop()                 // 滚动到顶部
vp.GotoBottom()              // 滚动到底部

// 高亮功能
vp.SetHighlights([][]int{{0, 5}, {10, 15}}) // 设置高亮范围
vp.HighlightNext()                          // 下一个高亮
vp.HighlightPrevious()                      // 上一个高亮
vp.ClearHighlights()                        // 清除高亮
```

### 高级功能
```go
// 软换行
vp.SoftWrap = true

// 行号显示
vp.LeftGutterFunc = func(info viewport.GutterContext) string {
    if info.Soft {
        return "     │ "  // 软换行行
    }
    if info.Index >= info.TotalLines {
        return "   ~ │ "  // 超出范围
    }
    return fmt.Sprintf("%4d │ ", info.Index+1) // 行号
}

// 每行样式
vp.StyleLineFunc = func(lineIndex int) lipgloss.Style {
    if lineIndex%2 == 0 {
        return lipgloss.NewStyle().Background(lipgloss.Color("#222"))
    }
    return lipgloss.NewStyle()
}

// 填充高度
vp.FillHeight = true  // 用空行填充视口
```

### 示例：与Textarea配合使用
```go
type Model struct {
    viewport viewport.Model
    textarea textarea.Model
    ready    bool
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        if !m.ready {
            m.viewport = viewport.New(viewport.WithWidth(msg.Width))
            m.viewport.SetHeight(msg.Height - 3) // 留出输入框空间
            m.ready = true
        } else {
            m.viewport.SetWidth(msg.Width)
            m.viewport.SetHeight(msg.Height - 3)
        }
        m.textarea.SetWidth(msg.Width)
        
    case tea.KeyPressMsg:
        if msg.String() == "enter" {
            content := m.viewport.GetContent() + "\n" + m.textarea.Value()
            m.viewport.SetContent(content)
            m.viewport.GotoBottom()
            m.textarea.Reset()
        }
    }
    
    var vpCmd, taCmd tea.Cmd
    m.viewport, vpCmd = m.viewport.Update(msg)
    m.textarea, taCmd = m.textarea.Update(msg)
    
    return m, tea.Batch(vpCmd, taCmd)
}
```

---

## Textinput - 单行文本输入

### 概述
单行文本输入组件，支持占位符、自动完成、验证等功能。

### 构造函数
```go
ti := textinput.New()
ti.Placeholder = "请输入..."
ti.Focus()
ti.SetWidth(40)
```

### 关键方法
```go
// 焦点控制
ti.Focus()           // 获取焦点
ti.Blur()            // 失去焦点
ti.Focused() bool    // 检查焦点状态

// 内容操作
ti.SetValue("文本")   // 设置值
ti.Value() string    // 获取值
ti.Reset()           // 重置

// 光标控制
ti.SetCursorColumn(5) // 设置光标位置
ti.Column() int       // 获取光标位置

// 尺寸控制
ti.SetWidth(80)       // 设置宽度
ti.Width() int        // 获取宽度

// 验证和提示
ti.SetValidate(func(s string) error {
    if len(s) < 3 {
        return errors.New("至少3个字符")
    }
    return nil
})

ti.SetSuggestions([]string{"选项1", "选项2", "选项3"})
```

### 样式配置
```go
isDark := true
styles := textinput.DefaultStyles(isDark)

// 自定义样式
styles.Focused.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
styles.Focused.Text = lipgloss.NewStyle().Bold(true)
styles.Blurred.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
styles.Focused.Suggestion = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))

ti.SetStyles(styles)
```

### 示例：表单输入
```go
type FormModel struct {
    inputs []textinput.Model
    focusIndex int
}

func NewFormModel() FormModel {
    m := FormModel{
        inputs: make([]textinput.Model, 3),
    }
    
    // 用户名输入
    m.inputs[0] = textinput.New()
    m.inputs[0].Placeholder = "用户名"
    m.inputs[0].Focus()
    
    // 邮箱输入
    m.inputs[1] = textinput.New()
    m.inputs[1].Placeholder = "邮箱"
    m.inputs[1].SetValidate(func(s string) error {
        if !strings.Contains(s, "@") {
            return errors.New("请输入有效的邮箱")
        }
        return nil
    })
    
    // 密码输入
    m.inputs[2] = textinput.New()
    m.inputs[2].Placeholder = "密码"
    m.inputs[2].EchoMode = textinput.EchoPassword
    m.inputs[2].EchoCharacter = '•'
    
    return m
}
```

---

## List - 列表组件

### 概述
功能丰富的列表组件，支持分页、过滤、帮助、状态消息等。

### 构造函数
```go
// 创建列表项
type item string
func (i item) FilterValue() string { return string(i) }

items := []list.Item{
    item("选项1"),
    item("选项2"),
    item("选项3"),
}

// 创建列表
l := list.New(items, list.NewDefaultDelegate(), 0, 0)
l.Title = "选择项目"
l.SetShowStatusBar(true)
l.SetFilteringEnabled(true)
```

### 关键方法
```go
// 尺寸控制
l.SetWidth(80)       // 设置宽度
l.Width() int        // 获取宽度
l.SetHeight(20)      // 设置高度
l.Height() int       // 获取高度

// 项目操作
l.SetItems(items)    // 设置项目
l.Items() []list.Item // 获取项目
l.SelectedItem() list.Item // 获取选中项
l.Index() int        // 获取选中索引

// 功能控制
l.SetFilteringEnabled(true)     // 启用过滤
l.SetShowTitle(true)            // 显示标题
l.SetShowStatusBar(true)        // 显示状态栏
l.SetShowHelp(true)             // 显示帮助
l.SetShowPagination(true)       // 显示分页

// 过滤
l.SetFilterValue("搜索词")      // 设置过滤值
l.FilterValue() string          // 获取过滤值
```

### 样式配置
```go
isDark := true
styles := list.DefaultStyles(isDark)

// 自定义样式
styles.Title = lipgloss.NewStyle().
    Foreground(lipgloss.Color("205")).
    Background(lipgloss.Color("0")).
    Padding(0, 1)

styles.Filter.Cursor = lipgloss.NewStyle().
    Foreground(lipgloss.Color("212"))

l.SetStyles(styles)
```

### 委托（Delegate）配置
```go
delegate := list.NewDefaultDelegate()
delegate.Styles = list.NewDefaultItemStyles(isDark)

// 自定义委托
delegate.ShowDescription = true
delegate.SetHeight(2)  // 每项高度
delegate.SetSpacing(1) // 项间距

l.SetDelegate(delegate)
```

### 示例：交互式列表
```go
type Model struct {
    list list.Model
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch msg.String() {
        case "enter":
            selected := m.list.SelectedItem().(item)
            // 处理选中项
            return m, tea.Println("选择了:", selected)
        }
    }
    
    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m Model) View() string {
    return m.list.View()
}
```

---

## Table - 表格组件

### 概述
用于显示和导航表格数据的组件，支持列排序、行选择、自定义渲染。

### 构造函数
```go
// 定义列
columns := []table.Column{
    {Title: "ID", Width: 10},
    {Title: "名称", Width: 20},
    {Title: "价格", Width: 15},
}

// 定义行
rows := []table.Row{
    {"1", "商品A", "$10.00"},
    {"2", "商品B", "$20.00"},
    {"3", "商品C", "$30.00"},
}

// 创建表格
t := table.New(
    table.WithColumns(columns),
    table.WithRows(rows),
    table.WithFocused(true),
    table.WithHeight(10),
)
```

### 关键方法
```go
// 尺寸控制
t.SetWidth(100)      // 设置宽度
t.Width() int        // 获取宽度
t.SetHeight(20)      // 设置高度
t.Height() int       // 获取高度

// 数据操作
t.SetRows(rows)              // 设置行数据
t.Rows() []table.Row         // 获取行数据
t.SetColumns(columns)        // 设置列
t.Columns() []table.Column   // 获取列

// 选择控制
t.Select(2)                  // 选择第3行（0-based）
t.SelectedRow() table.Row    // 获取选中行
t.Cursor() int               // 获取光标位置（选中行索引）

// 排序
t.SetSort([]table.SortColumn{
    {Column: 0, Ascending: true}, // 按第1列升序
})
```

### 样式配置
```go
isDark := true
styles := table.DefaultStyles(isDark)

// 自定义样式
styles.Header = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("212")).
    Background(lipgloss.Color("0")).
    Padding(0, 1)

styles.Selected = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("229")).
    Background(lipgloss.Color("57")).
    Padding(0, 1)

t.SetStyles(styles)
```

### 示例：可排序表格
```go
type Model struct {
    table table.Model
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch msg.String() {
        case "s":
            // 切换排序
            m.table.SetSort([]table.SortColumn{
                {Column: 2, Ascending: true}, // 按价格排序
            })
        }
    }
    
    var cmd tea.Cmd
    m.table, cmd = m.table.Update(msg)
    return m, cmd
}
```

---

## Progress - 进度条

### 概述
可定制的进度条组件，支持渐变、动画、自定义字符。

### 构造函数
```go
// 基本进度条
p := progress.New()

// 带选项创建
p := progress.New(
    progress.WithWidth(50),
    progress.WithColors(lipgloss.Color("#5A56E0"), lipgloss.Color("#EE6FF8")),
    progress.WithScaled(true), // 渐变缩放
    progress.WithPercentStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("255"))),
)
```

### 关键方法
```go
// 尺寸控制
p.SetWidth(80)       // 设置宽度
p.Width() int        // 获取宽度

// 进度控制
p.SetPercent(0.75)   // 设置进度（0.0-1.0）
p.Percent() float64  // 获取进度

// 颜色配置
p.SetFullColor(lipgloss.Color("#FF0000"))
p.SetEmptyColor(lipgloss.Color("#333333"))

// 字符配置
p.SetFull("█")      // 填充字符
p.SetEmpty("░")     // 空字符
p.SetShowPercent(true) // 显示百分比
```

### 动画进度条
```go
type Model struct {
    progress progress.Model
    percent  float64
}

func (m Model) Init() tea.Cmd {
    // 每100ms更新一次进度
    return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
        m.percent += 0.01
        if m.percent > 1.0 {
            m.percent = 1.0
        }
        return progress.PercentMsg(m.percent)
    })
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case progress.PercentMsg:
        m.percent = float64(msg)
        m.progress.SetPercent(m.percent)
        if m.percent < 1.0 {
            return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
                return progress.PercentMsg(m.percent + 0.01)
            })
        }
    }
    return m, nil
}
```

---

## Spinner - 加载指示器

### 概述
加载动画组件，支持自定义帧和样式。

### 构造函数
```go
// 默认spinner
s := spinner.New()

// 自定义spinner
s := spinner.New(
    spinner.WithStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205"))),
    spinner.WithFrames([]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}),
)
```

### 关键方法
```go
// 控制
s.Start()           // 开始动画
s.Stop()            // 停止动画
s.Tick() tea.Cmd    // 获取tick命令（用于Update）

// 样式
s.SetStyle(lipgloss.NewStyle().Bold(true))
s.SetFrames([]string{".", "..", "...", "...."})
```

### 示例：加载状态
```go
type Model struct {
    spinner spinner.Model
    loading bool
}

func (m Model) Init() tea.Cmd {
    m.spinner.Start()
    m.loading = true
    return m.spinner.Tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    if m.loading {
        var cmd tea.Cmd
        m.spinner, cmd = m.spinner.Update(msg)
        return m, cmd
    }
    return m, nil
}

func (m Model) View() string {
    if m.loading {
        return m.spinner.View() + " 加载中..."
    }
    return "加载完成"
}
```

---

## Help - 帮助系统

### 概述
自动生成帮助视图的组件，基于键绑定显示帮助信息。

### 构造函数
```go
h := help.New()
h.Styles = help.DefaultStyles(isDark)
h.SetWidth(80)
```

### 关键方法
```go
// 尺寸控制
h.SetWidth(80)      // 设置宽度
h.Width() int       // 获取宽度

// 显示模式
h.SetShowAll(true)           // 显示所有帮助
h.ShortSeparator = " • "     // 短格式分隔符
h.FullSeparator = " | "      // 完整格式分隔符

// 键绑定
h.SetKeyMap(myKeyMap)        // 设置键绑定映射
```

### 示例：集成帮助
```go
type KeyMap struct {
    Up    key.Binding
    Down  key.Binding
    Enter key.Binding
    Help  key.Binding
    Quit  key.Binding
}

var DefaultKeyMap = KeyMap{
    Up: key.NewBinding(
        key.WithKeys("up", "k"),
        key.WithHelp("↑/k", "上移"),
    ),
    Down: key.NewBinding(
        key.WithKeys("down", "j"),
        key.WithHelp("↓/j", "下移"),
    ),
    Enter: key.NewBinding(
        key.WithKeys("enter"),
        key.WithHelp("enter", "选择"),
    ),
    Help: key.NewBinding(
        key.WithKeys("?"),
        key.WithHelp("?", "切换帮助"),
    ),
    Quit: key.NewBinding(
        key.WithKeys("ctrl+c", "q"),
        key.WithHelp("ctrl+c/q", "退出"),
    ),
}

type Model struct {
    help   help.Model
    showHelp bool
    keys   KeyMap
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch {
        case key.Matches(msg, m.keys.Help):
            m.showHelp = !m.showHelp
        }
    }
    return m, nil
}

func (m Model) View() string {
    if m.showHelp {
        return m.help.View(m.keys)
    }
    return "主界面..."
}
```

---

## Key - 键绑定管理

### 概述
非可视化组件，用于管理键绑定和生成帮助文本。

### 创建键绑定
```go
// 基本键绑定
binding := key.NewBinding(
    key.WithKeys("enter", "ctrl+m"),
    key.WithHelp("enter", "确认选择"),
    key.WithDisabled(false),
)

// 多个键的绑定
multiKey := key.NewBinding(
    key.WithKeys("up", "k", "ctrl+p"),
    key.WithHelp("↑/k", "向上移动"),
)

// 禁用绑定
disabledKey := key.NewBinding(
    key.WithKeys("tab"),
    key.WithHelp("tab", "切换（已禁用）"),
    key.WithDisabled(true),
)
```

### 键绑定匹配
```go
type KeyMap struct {
    Up    key.Binding
    Down  key.Binding
    Enter key.Binding
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch {
        case key.Matches(msg, m.keys.Up):
            // 处理上键
        case key.Matches(msg, m.keys.Down):
            // 处理下键
        case key.Matches(msg, m.keys.Enter):
            // 处理回车
        }
    }
    return m, nil
}
```

### 键绑定组
```go
// 创建键绑定组
bindings := []key.Binding{
    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "上")),
    key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "下")),
    key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "选择")),
}

// 生成帮助文本
helpText := key.Bindings(bindings).Help().FullHelp()
```

---

## 其他组件

### Stopwatch - 秒表
```go
sw := stopwatch.New(stopwatch.WithInterval(time.Second))
sw.Start()
// 在Update中：sw, cmd = sw.Update(msg)
// 获取时间：sw.Elapsed()
```

### Timer - 定时器
```go
t := timer.New(30*time.Second, timer.WithInterval(time.Second))
t.Start()
// 检查是否超时：t.Timedout()
```

### Filepicker - 文件选择器
```go
fp := filepicker.New()
fp.AllowedTypes = []string{".txt", ".go", ".md"}
fp.CurrentDirectory, _ = os.UserHomeDir()
```

### Paginator - 分页器
```go
pg := paginator.New()
pg.SetTotalPages(10)
pg.Page() // 当前页码
```

### Cursor - 光标控制
```go
// 在Textinput/Textarea中自动管理
// 手动控制：
c := cursor.New()
c.Blink() // 闪烁命令
```

---

## 最佳实践

### 1. Light/Dark模式处理
```go
// 在Model中存储isDark状态
type Model struct {
    isDark bool
    help   help.Model
    list   list.Model
}

// 在Init中请求背景色
func (m Model) Init() tea.Cmd {
    return tea.RequestBackgroundColor
}

// 在Update中处理背景色
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.BackgroundColorMsg:
        m.isDark = msg.IsDark()
        // 更新所有组件样式
        m.help.Styles = help.DefaultStyles(m.isDark)
        m.list.Styles = list.DefaultStyles(m.isDark)
    }
    return m, nil
}
```

### 2. 组件组合
```go
type GameModel struct {
    viewport viewport.Model  // 显示游戏内容
    textarea textarea.Model  // 玩家输入
    help     help.Model      // 帮助系统
    keys     GameKeyMap      // 键绑定
    showHelp bool
}

// 统一更新
func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var vpCmd, taCmd, helpCmd tea.Cmd
    
    m.viewport, vpCmd = m.viewport.Update(msg)
    m.textarea, taCmd = m.textarea.Update(msg)
    m.help, helpCmd = m.help.Update(msg)
    
    return m, tea.Batch(vpCmd, taCmd, helpCmd)
}
```

### 3. 性能优化
- **批量更新**：使用`tea.Batch()`组合多个命令
- **懒加载**：只在需要时创建组件
- **缓存渲染**：对于复杂视图考虑缓存
- **避免阻塞**：长时间操作使用goroutine

### 4. 错误处理
```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    defer func() {
        if r := recover(); r != nil {
            // 恢复panic，记录错误
            log.Printf("Recovered from panic: %v", r)
        }
    }()
    
    // 正常更新逻辑
    return m, nil
}
```

---

## 常见问题

### Q: 如何检测Light/Dark模式？
A: 使用`tea.RequestBackgroundColor`命令，在`Update`中处理`tea.BackgroundColorMsg`。

### Q: 组件不响应按键？
A: 检查：
1. 组件是否获得焦点（`Focus()`）
2. 键绑定是否正确配置
3. 是否在`Update`中正确处理`tea.KeyPressMsg`

### Q: 如何自定义组件样式？
A: 每个组件都有`DefaultStyles(isDark bool)`函数，返回可修改的样式结构体。

### Q: 如何组合多个组件？
A: 在Model中嵌入多个组件，在`Update`中分别调用各自的`Update`方法，使用`tea.Batch()`组合命令。

### Q: 如何实现组件间通信？
A: 通过自定义消息类型和Bubble Tea的消息系统。

---

## 资源

- **官方文档**: `charm.land/bubbles/v2`
- **示例代码**: GitHub仓库中的examples目录
- **社区组件**: `github.com/charm-and-friends/additional-bubbles`

---

*文档最后更新: 2026年3月4日*
*基于Bubbles v2.0.0*
