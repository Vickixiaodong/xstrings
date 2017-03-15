// 字符串操作

package xstrings

// 字符串数组排序
func Sort(strs []string, sortType string) ([]string) {
    if len(strs) == 0 {
        return strs
    }

    // 排序
    switch sortType {
    case "asc":
        strs = ascSort(strs)
    case "desc":
        strs = descSort(strs)
    default:
        strs = ascSort(strs)
    }

    return strs
}

// 升序排序
func ascSort(strs []string) ([]string) {
    for i := 0; i < len(strs) - 1; i++ {
        minValueIndex := i
        for j := i + 1; j < len(strs); j++ {
            if strs[minValueIndex] > strs[j] {
                minValueIndex = j
            }
        }

        if minValueIndex != i {
            tmp := strs[i]
            strs[i] = strs[minValueIndex]
            strs[minValueIndex] = tmp
        }
    }

    return strs
}

// 降序排序
func descSort(strs []string) ([]string) {
    for i := 0; i < len(strs) - 1; i++ {
        minValueIndex := i
        for j := i + 1; j < len(strs); j++ {
            if strs[minValueIndex] < strs[j] {
                minValueIndex = j
            }
        }

        if minValueIndex != i {
            tmp := strs[i]
            strs[i] = strs[minValueIndex]
            strs[minValueIndex] = tmp
        }
    }

    return strs
}