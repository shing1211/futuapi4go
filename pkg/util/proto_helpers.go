package util

func ProtoStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ProtoInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func ProtoBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func ProtoFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func ProtoFloat32(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

func ProtoInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func ProtoUint64(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

func ProtoInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
