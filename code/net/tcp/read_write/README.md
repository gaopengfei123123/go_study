组合1 模拟正常的读写请求
组合2 是对一次性读取不完请求数据时的处理
组合3 模拟当 server 读取数据时 client 端主动关闭链接
组合4 模拟 server 读取超时的情况
组合5 模拟 client 向 server 发送大量数据模拟两边OS 协议栈堆满的情景
组合6 是组合5的升级版, client 写入时加上了超时判定

代码来源: https://tonybai.com/2015/11/17/tcp-programming-in-golang/


综上, tcp 因为是双工链接, client 和 server 的读&写都应该加上超时判定, 保持程序的高速处理请求的状态