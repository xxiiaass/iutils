package iutils

// 获取一个将字符串转为int64的算法句柄
func HashCode() hashCode {
	return hashCode{}
}

type hashCode struct {
}

func (hashCode) BKDRHash(str string) int64 {
	const seed int64 = 131
	r := []rune(str)
	count, size := len(r), len(r)
	var hash int64
	for count > 0 {
		hash = hash*seed + int64(r[size-count])
		count--
	}
	return hash & 0x7FFFFFFF
}

func (hashCode) APHash(str string) int64 {
	r := []rune(str)
	count := len(r)
	var hash int64
	for i := 0; i < count; i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ int64(r[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ int64(r[i]) ^ (hash >> 5))
		}
		count--
	}
	return hash & 0x7FFFFFFF
}

func (hashCode) SDBMHash(str string) int64 {
	r := []rune(str)
	count, size := len(r), len(r)
	var hash int64
	for count > 0 {
		hash = int64(r[size-count]) + (hash << 6) + (hash << 16) - hash
		count--
	}
	return hash & 0x7FFFFFFF
}

func (hashCode) RSHash(str string) int64 {
	var a int64 = 63689
	const b int64 = 378551
	r := []rune(str)
	count, size := len(r), len(r)
	var hash int64
	for count > 0 {
		hash = hash*a + int64(r[size-count])
		a *= b
		count--
	}
	return hash & 0x7FFFFFFF
}

func (hashCode) JSHash(str string) int64 {
	var hash int64 = 1315423911
	r := []rune(str)
	count, size := len(r), len(r)
	for count > 0 {
		hash ^= (hash << 5) + int64(r[size-count]) + (hash >> 2)
		count--
	}
	return hash & 0x7FFFFFFF
}

func (hashCode) ELFHash(str string) int64 {
	var x, hash int64
	r := []rune(str)
	count, size := len(r), len(r)
	for count > 0 {
		hash = (hash << 4) + int64(r[size-count])
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
			hash &= ^x
		}
		count--
	}
	return hash & 0x7FFFFFFF
}

func (hashCode) DJBHash(str string) int64 {
	r := []rune(str)
	count, size := len(r), len(r)
	var hash int64
	for count > 0 {
		hash += (hash << 5) + int64(r[size-count])
		count--
	}
	return hash & 0x7FFFFFFF
}
