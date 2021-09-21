package enums

import "fagin/pkg/enum"

type ByteUnit uint64

var _ enum.Enum = new(ByteUnit)

const (
	ByteUnitB  ByteUnit = 1 << (10 * iota)
	ByteUnitKB          // 1 << (10*1)
	ByteUnitMB          // 1 << (10*2)
	ByteUnitGB          // 1 << (10*3)
	ByteUnitTB          // 1 << (10*4)
	//PB // 1 << (10*5)
	//EB // 1 << (10*6)
	//ZB // 1 << (10*7)
	//YB // 1 << (10*8)
)

func (code ByteUnit) String() string {
	switch code {
	case ByteUnitB:
		return "B"
	case ByteUnitKB:
		return "KB"
	case ByteUnitMB:
		return "MB"
	case ByteUnitGB:
		return "GB"
	case ByteUnitTB:
		return "TB"
	}
	return ""
}

func (code ByteUnit) Get() uint64 {
	return uint64(code)
}

func AllByteUnit() map[string]uint64 {
	return map[string]uint64{
		ByteUnitB.String():  ByteUnitB.Get(),
		ByteUnitKB.String(): ByteUnitKB.Get(),
		ByteUnitMB.String(): ByteUnitMB.Get(),
		ByteUnitGB.String(): ByteUnitGB.Get(),
		ByteUnitTB.String(): ByteUnitTB.Get(),
		//b.PB: "PB",
		//b.EB: "EB",
		//b.ZB: "ZB",
		//b.YB: "YB",
	}
}
