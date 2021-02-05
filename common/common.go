package common


var Semaphore = make(chan int8,1)

/**
合并两个map,相同key以第二个为主
 */
func CombineMap(m1 map[string]string, m2 map[string]string) map[string]string{
	var m = make(map[string]string)
	for k,v := range m1{
		m[k] = v
	}
	for k,v := range m2{
		m[k] = v
	}
	return m

}