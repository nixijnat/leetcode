package main

func convert(s string, numRows int) string {
	length := len(s)
	if numRows == 1 || length <= numRows {
		return s
	}
	ret := make([]byte, 0, length)
	eleNum := numRows*2 - 2
	groupNum := length/eleNum + 1
	for i := 0; i < numRows; i++ {
		midRow := i != 0 && i != numRows-1
		for j := 0; j <= groupNum; j++ {
			index := eleNum*j + i
			if index >= length {
				break
			}
			ret = append(ret, s[index])
			if !midRow {
				continue
			}
			index = eleNum*(j+1) - i
			if index >= length {
				break
			}
			ret = append(ret, s[index])
		}
	}
	return string(ret)
}
