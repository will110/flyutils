package flyutils

/**
描述：删除slice中的某个元素；
参数：
	val: 要删除的元素
	sliceList： 切片列表
 */
func SliceDeleteString(val string, sliceList []string) []string {
	if len(sliceList) == 0 {
		return []string{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]string, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteInt(val int, sliceList []int) []int {
	if len(sliceList) == 0 {
		return []int{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]int, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteInt8(val int8, sliceList []int8) []int8 {
	if len(sliceList) == 0 {
		return []int8{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]int8, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteInt16(val int16, sliceList []int16) []int16 {
	if len(sliceList) == 0 {
		return []int16{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]int16, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteInt32(val int32, sliceList []int32) []int32 {
	if len(sliceList) == 0 {
		return []int32{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]int32, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteInt64(val int64, sliceList []int64) []int64 {
	if len(sliceList) == 0 {
		return []int64{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]int64, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteUint8(val uint8, sliceList []uint8) []uint8 {
	if len(sliceList) == 0 {
		return []uint8{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]uint8, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteUint16(val uint16, sliceList []uint16) []uint16 {
	if len(sliceList) == 0 {
		return []uint16{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]uint16, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteUint32(val uint32, sliceList []uint32) []uint32 {
	if len(sliceList) == 0 {
		return []uint32{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]uint32, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteUint64(val uint64, sliceList []uint64) []uint64 {
	if len(sliceList) == 0 {
		return []uint64{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]uint64, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteFloat32(val float32, sliceList []float32) []float32 {
	if len(sliceList) == 0 {
		return []float32{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]float32, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

func SliceDeleteFloat64(val float64, sliceList []float64) []float64 {
	if len(sliceList) == 0 {
		return []float64{}
	}

	index := 0
	endIndex := len(sliceList) - 1
	result := make([]float64, 0, len(sliceList))
	for k := range sliceList {
		if sliceList[k] == val {
			result = append(result, sliceList[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, sliceList[index:endIndex+1]...)
		}
	}

	return result
}

/**
描述：判断slice中是否存在某个元素；
参数：
	v: 要查找的元素
	sliceList： 切片列表
*/
func InSliceString(v string, sliceList []string) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceInt8(v int8, sliceList []int8) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceInt16(v int16, sliceList []int16) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceInt32(v int32, sliceList []int32) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceInt64(v int64, sliceList []int64) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceInt(v int, sliceList []int) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceUint8(v uint8, sliceList []uint8) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceUint16(v uint16, sliceList []uint16) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceUint32(v uint32, sliceList []uint32) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceUint64(v uint64, sliceList []uint64) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceUint(v uint, sliceList []uint) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceFloat32(v float32, sliceList []float32) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

func InSliceFloat64(v float64, sliceList []float64) bool {
	for k := range sliceList {
		if sliceList[k] == v {
			return true
		}
	}

	return false
}

/**
描述：获取slice1在slice2差值，即返回差集；
参数：
	slice1: 第一个切片列表
	slice2： 第二个切片列表
*/
func SliceDiffString(slice1, slice2 []string) []string {
	diffSlice := make([]string, 0, len(slice2))
	for k := range slice1 {
		if !InSliceString(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffInt8(slice1, slice2 []int8) []int8 {
	diffSlice := make([]int8, 0, len(slice2))
	for k := range slice1 {
		if !InSliceInt8(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffInt16(slice1, slice2 []int16) []int16 {
	diffSlice := make([]int16, 0, len(slice2))
	for k := range slice1 {
		if !InSliceInt16(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffInt32(slice1, slice2 []int32) []int32 {
	diffSlice := make([]int32, 0, len(slice2))
	for k := range slice1 {
		if !InSliceInt32(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffInt64(slice1, slice2 []int64) []int64 {
	diffSlice := make([]int64, 0, len(slice2))
	for k := range slice1 {
		if !InSliceInt64(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffInt(slice1, slice2 []int) []int {
	diffSlice := make([]int, 0, len(slice2))
	for k := range slice1 {
		if !InSliceInt(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffUint8(slice1, slice2 []uint8) []uint8 {
	diffSlice := make([]uint8, 0, len(slice2))
	for k := range slice1 {
		if !InSliceUint8(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffUint16(slice1, slice2 []uint16) []uint16 {
	diffSlice := make([]uint16, 0, len(slice2))
	for k := range slice1 {
		if !InSliceUint16(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffUint32(slice1, slice2 []uint32) []uint32 {
	diffSlice := make([]uint32, 0, len(slice2))
	for k := range slice1 {
		if !InSliceUint32(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffUint64(slice1, slice2 []uint64) []uint64 {
	diffSlice := make([]uint64, 0, len(slice2))
	for k := range slice1 {
		if !InSliceUint64(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffUint(slice1, slice2 []uint) []uint {
	diffSlice := make([]uint, 0, len(slice2))
	for k := range slice1 {
		if !InSliceUint(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffFloat32(slice1, slice2 []float32) []float32 {
	diffSlice := make([]float32, 0, len(slice2))
	for k := range slice1 {
		if !InSliceFloat32(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

func SliceDiffFloat64(slice1, slice2 []float64) []float64 {
	diffSlice := make([]float64, 0, len(slice2))
	for k := range slice1 {
		if !InSliceFloat64(slice1[k], slice2) {
			diffSlice = append(diffSlice, slice1[k])
		}
	}

	return diffSlice
}

/**
描述：获取slice1与slice2共同存在的值，即返回交集；
参数：
	slice1: 第一个切片列表
	slice2： 第二个切片列表
*/
func SliceIntersectString(slice1, slice2 []string) []string {
	intersectSlice := make([]string, 0, 8)
	for k := range slice1 {
		if InSliceString(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectInt(slice1, slice2 []int) []int {
	intersectSlice := make([]int, 0, 8)
	for k := range slice1 {
		if InSliceInt(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectInt8(slice1, slice2 []int8) []int8 {
	intersectSlice := make([]int8, 0, 8)
	for k := range slice1 {
		if InSliceInt8(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectInt16(slice1, slice2 []int16) []int16 {
	intersectSlice := make([]int16, 0, 8)
	for k := range slice1 {
		if InSliceInt16(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectInt32(slice1, slice2 []int32) []int32 {
	intersectSlice := make([]int32, 0, 8)
	for k := range slice1 {
		if InSliceInt32(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectInt64(slice1, slice2 []int64) []int64 {
	intersectSlice := make([]int64, 0, 8)
	for k := range slice1 {
		if InSliceInt64(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectUint8(slice1, slice2 []uint8) []uint8 {
	intersectSlice := make([]uint8, 0, 8)
	for k := range slice1 {
		if InSliceUint8(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectUint16(slice1, slice2 []uint16) []uint16 {
	intersectSlice := make([]uint16, 0, 8)
	for k := range slice1 {
		if InSliceUint16(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectUint32(slice1, slice2 []uint32) []uint32 {
	intersectSlice := make([]uint32, 0, 8)
	for k := range slice1 {
		if InSliceUint32(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectUint64(slice1, slice2 []uint64) []uint64 {
	intersectSlice := make([]uint64, 0, 8)
	for k := range slice1 {
		if InSliceUint64(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectUint(slice1, slice2 []uint) []uint {
	intersectSlice := make([]uint, 0, 8)
	for k := range slice1 {
		if InSliceUint(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectFloat32(slice1, slice2 []float32) []float32 {
	intersectSlice := make([]float32, 0, 8)
	for k := range slice1 {
		if InSliceFloat32(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

func SliceIntersectFloat64(slice1, slice2 []float64) []float64 {
	intersectSlice := make([]float64, 0, 8)
	for k := range slice1 {
		if InSliceFloat64(slice1[k], slice2) {
			intersectSlice = append(intersectSlice, slice1[k])
		}
	}

	return intersectSlice
}

/**
描述：返回sliceList元素所有元素之和；
参数：
	sliceList: 切片列表
*/
func SliceSumInt8(sliceList []int8) int64 {
	var sum int64
	for k := range sliceList {
		sum += int64(sliceList[k])
	}

	return sum
}

func SliceSumInt16(sliceList []int16) int64 {
	var sum int64
	for k := range sliceList {
		sum += int64(sliceList[k])
	}

	return sum
}

func SliceSumInt32(sliceList []int32) int64 {
	var sum int64
	for k := range sliceList {
		sum += int64(sliceList[k])
	}

	return sum
}


func SliceSumInt64(sliceList []int64) int64 {
	var sum int64
	for k := range sliceList {
		sum += sliceList[k]
	}

	return sum
}

func SliceSumUint8(sliceList []uint8) uint64 {
	var sum uint64
	for k := range sliceList {
		sum += uint64(sliceList[k])
	}

	return sum
}

func SliceSumUint16(sliceList []uint16) uint64 {
	var sum uint64
	for k := range sliceList {
		sum += uint64(sliceList[k])
	}

	return sum
}

func SliceSumUint32(sliceList []uint32) uint64 {
	var sum uint64
	for k := range sliceList {
		sum += uint64(sliceList[k])
	}

	return sum
}

func SliceSumUint64(sliceList []uint64) uint64 {
	var sum uint64
	for k := range sliceList {
		sum += sliceList[k]
	}

	return sum
}

func SliceSumUint(sliceList []uint) uint64 {
	var sum uint64
	for k := range sliceList {
		sum += uint64(sliceList[k])
	}

	return sum
}

func SliceSumInt(sliceList []int) int64 {
	var sum int64
	for k := range sliceList {
		sum += int64(sliceList[k])
	}

	return sum
}

func SliceSumFloat32(sliceList []float32) float64 {
	var sum float64
	for k := range sliceList {
		sum += float64(sliceList[k])
	}

	return sum
}

func SliceSumFloat64(sliceList []float64) float64 {
	var sum float64
	for k := range sliceList {
		sum += sliceList[k]
	}

	return sum
}

/**
描述：去掉重复的元素，并返回一个新的slice；
参数：
	sliceList：要处理的切片
 */
func SliceUniqueString(sliceList []string) []string {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceString(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueInt(sliceList []int) []int {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceInt(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueInt8(sliceList []int8) []int8 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceInt8(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueInt16(sliceList []int16) []int16 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceInt16(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueInt32(sliceList []int32) []int32 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceInt32(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueInt64(sliceList []int64) []int64 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceInt64(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueUint(sliceList []uint) []uint {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceUint(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueUint8(sliceList []uint8) []uint8 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceUint8(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueUint16(sliceList []uint16) []uint16 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceUint16(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueUint32(sliceList []uint32) []uint32 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceUint32(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueUint64(sliceList []uint64) []uint64 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceUint64(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueFloat32(sliceList []float32) []float32 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceFloat32(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}

func SliceUniqueFloat64(sliceList []float64) []float64 {
	result := sliceList[:0]
	for k := range sliceList {
		if !InSliceFloat64(sliceList[k], result) {
			result = append(result, sliceList[k])
		}
	}

	return result
}