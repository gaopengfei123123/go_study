package progress

import "fmt"

// Hello 测试用
func Hello() {
	fmt.Println("hello")
}

// Progress 进图结构体
type Progress struct {
	percent int64  // 百分比
	current int64  // 当前进度
	total   int64  // 总量
	rate    string // 进度条
	graph   string // 进度符号
}

// NewProgress 初始化方法
func NewProgress(start, total int64) *Progress {
	p := new(Progress)

	p.current = start
	p.total = total

	p.graph = "#"
	p.percent = p.GetPercent()

	return p
}

// GetPercent 获取百分比
func (p *Progress) GetPercent() int64 {
	return int64(float32(p.current) / float32(p.total) * 100)
}

// Add 打印
func (p *Progress) Add(i int64) {

	p.current += i

	if p.current > p.total {
		return
	}

	last := p.percent
	p.percent = p.GetPercent()

	if p.percent != last && p.percent%2 == 0 {
		p.rate += p.graph
	}

	fmt.Printf("\r[%-50s]%3d%% %8d/%d", p.rate, p.percent, p.current, p.total)
	// %-50s 左对齐, 占50个字符位置, 打印string
	// %3d   右对齐, 占3个字符位置 打印int

	if p.current == p.total {
		p.Done()
	}
}

// Done 完毕
func (p *Progress) Done() {
	fmt.Println()
}
