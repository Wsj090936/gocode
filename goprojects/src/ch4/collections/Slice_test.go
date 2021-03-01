package collections

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 切片的初始化 与数组稍微有一点点不同
	var s0 []int
	s0 = append(s0, 1)
	// 打印切片 s0 的可访问数组长度以及容量
	fmt.Println(len(s0),cap(s0))// 1,1
	// 切片初始化的第二种方式
	s1 := []int{1,2,3,4}
	fmt.Println(len(s1),cap(s1))// 4,4
	// 切片初始化的第三种方式 指定可访问容量为3，数组容量为5
	s2 := make([]int,2,5)
	//fmt.Println(s2[0],s2[1],s2[2])  这里s2[2] 不能通过编译，因为切片的容量只有2
	fmt.Println(len(s2),cap(s2))// 2,5
	s2 = append(s2,1)
	fmt.Println(s2[0],s2[1],s2[2])

}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0;i < 10;i++{
		s = append(s,i)
		fmt.Println(len(s),cap(s))
	}
	// 从这里可以看出，切片类似于java中的list一样，每次扩容的大小是前一次的2倍
	/**
	1 1
	2 2
	3 4
	4 4
	5 8
	6 8
	7 8
	8 8
	9 16
	10 16
	 */
}

// 切片的共享空间
func TestSliceShare(t *testing.T){
	var s1 []int
	s1 = append(s1, 1,2,3,4,5,6,7,8)
	fmt.Println(s1,len(s1),cap(s1)) // [1 2 3 4 5 6 7 8] 8 8
	// 进行截取 这里截取出来，数组元素只有三个，但是容量却有7个，因为是共享的是和 s1一样的空间
	s2 := s1[1:4]
	fmt.Println(s2,len(s2),cap(s2)) // [2 3 4] 3 7
	// 进行截取 这里截取出来，数组元素只有两个，但是容量却有7个，因为是共享的是和 s1一样的空间
	s3 := s1[2:4]
	fmt.Println(s3,len(s3),cap(s3)) // [3 4] 2 6
	// 这里改下标为2的元素，s2 和 s3中的值也会变，因此是共享空间的
	s3[0] = 0
	fmt.Println(s2,s3) // [2 0 4] [0 4]

}

func TestSliceComparing(t *testing.T) {
	//s1 := []int{1,2,3,4}
	//s2 := []int{1,2,3,4}
	////fmt.Println(s1 == s2) 这里不能运行
}