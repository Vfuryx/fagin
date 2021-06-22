package byte_unit

import "fagin/pkg/constant"

const (
	B  = 1 << (10 * iota)
	KB // 1 << (10*1)
	MB // 1 << (10*2)
	GB // 1 << (10*3)
	TB // 1 << (10*4)
	//PB // 1 << (10*5)
	//EB // 1 << (10*6)
	//ZB // 1 << (10*7)
	//YB // 1 << (10*8)
)

// 标准 64 位机器
type byteUnit struct {
	B  int
	KB int
	MB int
	GB int
	TB int
	//PB int
	//EB int
	//ZB uint64
	//YB uint64
	constant.Constant
}

func ByteUnit() *byteUnit {
	s := byteUnit{
		B:  B,
		KB: KB,
		MB: MB,
		GB: GB,
		TB: TB,
		//PB: PB,
		//EB: EB,
		//ZB: ZB,
		//YB: YB,
	}
	s.Constant.C = &s
	return &s
}

func (b *byteUnit) All() map[interface{}]string {
	return map[interface{}]string{
		b.B:  "B",
		b.KB: "KB",
		b.MB: "MB",
		b.GB: "GB",
		b.TB: "TB",
		//b.PB: "PB",
		//b.EB: "EB",
		//b.ZB: "ZB",
		//b.YB: "YB",
	}
}
