package xstrings

import (
    "testing"
    "log"
)

func TestSort(t *testing.T) {
    strs := []string {"apple", "pear", "cherry"}

    log.Println("排序前：", strs)

    strs = Sort(strs, "asc")

    log.Println("升序排序后：", strs)

    strs = Sort(strs, "desc")

    log.Println("降序排序后：", strs)
}
