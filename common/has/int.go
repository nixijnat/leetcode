package has

// TODO(jixintan) : 待优化
func Int(c []int, e int) bool {
	for _, i := range c {
		if i == e {
			return true
		}
	}
	return false
}
