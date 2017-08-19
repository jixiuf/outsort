package main

import (
	"fmt"

	"github.com/jixiuf/outsort/outsort"

	"flag"

	"os"
)

const (
	// 这个常量的含义，假如内存限制为1G,则arr_size的大小为1G/4 ,
	// 即 排序过程中会用到三个数组，其大小分别是为 1G/4,1G/4,1G/2
	// 即将前两个已排序的数组 归并到 第3个数组中
	arr_size = 10 * 1024 // 10k ,
	// arr_size = 10 * 1024 // 10k ,内存最大值为10k
)

var inFile string
var outFile string

func main() {
	fmt.Printf("用法: %s -i in.txt -o out.txt\n", os.Args[0])

	flag.StringVar(&inFile, "i", "in.txt", "未排序的文件")
	flag.StringVar(&outFile, "o", "out.txt", "已排序的数据存放在此文件")
	flag.Parse()
	sorter := outsort.NewSort(arr_size)
	err := sorter.Sort(inFile, outFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("生成的已排序的数据存放在%s文件中", outFile)

}
