package core

/*最小渲染单元结构体*/
type Text struct {
    Time  float64     // 打字机等待时间（秒）
    Color [3]uint16   // RGB颜色，0-65535
    Text  string      // 文字内容
}
