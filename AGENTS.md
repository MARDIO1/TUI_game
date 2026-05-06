# go_interact 协作说明

这是一个 Go 语言的 Bubble Tea + Bubbles v2 TUI 游戏项目。默认从仓库根目录启动，音频资源和相对路径都依赖当前工作目录。

## 启动方式

在仓库根目录执行：

```powershell
go run .
```

也可以先编译再运行：

```powershell
go build -o bin\go_interact.exe .
.\bin\go_interact.exe
```

## 部署方式

这个项目没有独立后端服务，部署本质上是把可执行文件和资源一起发出去。

需要保留的内容：

- 可执行文件
- `assets/audio/` 下的音频资源
- 仓库根目录作为运行目录，或者保证程序读取资源的相对路径不变

建议的发布方式：

1. 在仓库根目录运行 `go build -o bin\go_interact.exe .`
2. 把 `bin\go_interact.exe` 和 `assets/` 一起打包
3. 在目标机器上从项目根目录或等价目录启动

## 代码入口

- `main.go`：程序入口，只负责初始化音乐播放器并启动 Bubble Tea 程序
- `core/master.go`：主控制器，负责 100Hz tick、输入分发、视图拼接
- `input/input.go`：输入框组件
- `output/TUI.go`：上半屏输出区域
- `output/sound.go`：音频播放与资源扫描

## 协作约定

- 先看现有代码和 `doc/`，再动手改动
- 保持改动最小化，不要顺手重构无关模块
- 不要删除用户已有文件
- 不要自动提交 git
- 维持现有命名风格和模块边界

## 关键注意点

- 音频初始化会扫描 `assets/audio/peace`、`assets/audio/candy`、`assets/audio/pathos` 等目录
- 程序依赖相对路径，运行时不要随意切换工作目录
- 当前 `go.mod` 使用 Go 1.25.0，最低应使用兼容版本的 Go 环境

## 常用检查命令

```powershell
go build ./...
```

如果构建失败，优先检查：

- Go 版本是否兼容
- `assets/audio/` 是否存在
- 工作目录是否在仓库根目录