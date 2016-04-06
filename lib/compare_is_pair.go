package lib
 
type IsCompare struct {
    	s1 string   
    	i2  int  
}
 
// 按照 Person.Age 从大到小排序
type IsCompareSlice [] IsCompare
 
func (a IsCompareSlice) Len() int {    // 重写 Len() 方法
    	return len(a)
}
func (a IsCompareSlice) Swap(i, j int){     // 重写 Swap() 方法
    	a[i], a[j] = a[j], a[i]
}
func (a IsCompareSlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
    	if a[j].s1 > a[i].s1 {
		return true
	} else {
		if a[j].i2 > a[i].i2 {
			return true
		} else {
			return false
		}
	}
}
