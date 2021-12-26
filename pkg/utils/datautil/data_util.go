package datautil

func CalcSize(val int) int {
	if val > 10 {
		val = 8
	}
	if val > 20 {
		val = 16
	}
	return val
}

func CollectDistictStrings(ln int, callback func(i int) string) (res []string) {
	var size = CalcSize(ln)
	var mp = make(map[string]bool, size)
	for i := 0; i < ln; i++ {
		var out = callback(i)
		if mp[out] {
			continue
		}
		mp[out] = true
		res = append(res, out)
	}
	return res

}
func MapString(ln int,u func(i int) string) (res []string) {
	for i:=0;i<ln;i++ {
		res = append(res,u(i))
	}
	return res
}