package log

import "fmt"

type Color uint8

// ANSI控制码:
// 	\x1b[0m      关闭所有属性
// 	\x1b[1m     设置高亮度
// 	\x1b[4m     下划线
// 	\x1b[5m     闪烁
// 	\x1b[7m     反显
// 	\x1b[8m     消隐
// 	\x1b[30m   --  \x1b[37m   设置前景色
// 	\x1b[40m   --  \x1b[47m   设置背景色
// 	\x1b[nA    光标上移n行
// 	\x1b[nB    光标下移n行
// 	\x1b[nC    光标右移n行
// 	\x1b[nD    光标左移n行
// 	\x1b[y;xH  设置光标位置
// 	\x1b[2J    清屏
// 	\x1b[K     清除从光标到行尾的内容
// 	\x1b[s     保存光标位置
// 	\x1b[u     恢复光标位置
// 	\x1b[?25l  隐藏光标
// 	\x1b[?25h  显示光标
// 字背景颜色范围: 40--49         	  字颜色: 30--39
// 40: 黑                          30: 黑
// 41: 红                          31: 红
// 42: 绿                          32: 绿
// 43: 黄                          33: 黄
// 44: 蓝                          34: 蓝
// 45: 紫                          35: 紫
// 46: 深绿                        36: 深绿
// 47: 白色                        37: 白色
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// 打印一条日志
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}
