package enums

import "fagin/pkg/enum"

type ByteUnit uint64

var _ enum.Enum = new(ByteUnit)

const (
	base                = 10
	ByteUnitB  ByteUnit = 1 << (base * iota)
	ByteUnitKB          // 1 << (10*1)
	ByteUnitMB          // 1 << (10*2)
	ByteUnitGB          // 1 << (10*3)
	ByteUnitTB          // 1 << (10*4)

	// PB // 1 << (10*5)
	// EB // 1 << (10*6)
	// ZB // 1 << (10*7)
	// YB // 1 << (10*8)
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

func AllByteUnit() map[uint64]string {
	return map[uint64]string{
		ByteUnitB.Get():  ByteUnitB.String(),
		ByteUnitKB.Get(): ByteUnitKB.String(),
		ByteUnitMB.Get(): ByteUnitMB.String(),
		ByteUnitGB.Get(): ByteUnitGB.String(),
		ByteUnitTB.Get(): ByteUnitTB.String(),
		// b.PB: "PB",
		// b.EB: "EB",
		// b.ZB: "ZB",
		// b.YB: "YB",
	}
}
