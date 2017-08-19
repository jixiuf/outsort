package main

// while read i;do echo "$i $RANDOM";done<file|sort -k2n|cut -d" " -f1
import (
	"./outsort"
)

const (
	mem_limit = 1024 // 10k ,内存最大值为10k
	// mem_limit = 10 * 1024 // 10k ,内存最大值为10k
)

func main() {
	sorter := outsort.NewSort(mem_limit)
	sorter.Sort("in.txt", "out.txt")

}
