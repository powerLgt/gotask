package main

import "fmt"

/*
	编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	考察点 ：指针的使用、值传递与引用传递的区别。
*/
func addNum(data *int, num int) {
	*data += num
}

/*
	实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func arrItemDouble(data *[]int) {
	for i := 0; i < len(*data); i++ {
		(*data)[i] *= 2
	}
}

/*
   对比切片的值传递以及指针传递
   1、修改影响原值
   2、传递切片时，实际传递的是切片头的副本（包含指针、长度、容量）。函数内可通过副本的指针修改‌底层数组的元素（影响原切片数据），但修改切片头的元数据（长度/容量）‌不会影响原切片
   3、传递切片的指针（*[]int），函数内直接操作‌ 原切片的元数据。修改长度、容量或底层数组指针时，‌原切片同步更新
*/
func slicePass(slicePointer *[]int, sliceVar []int) {
	(*slicePointer)[0] = 100
	sliceVar[0] = 100
	*slicePointer = append(*slicePointer, []int{1, 2, 3, 4, 5}...)
	sliceVar = append(sliceVar, []int{1, 2, 3, 4, 5}...)
}

func main() {
	data := 10
	fmt.Println("init, data: ", data)

	addNum(&data, 10)
	fmt.Println("after addNum, data: ", data)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("init arr: ", arr)
	arrItemDouble(&arr)
	fmt.Println("after arrItemDouble, arr: ", arr)

	slice1 := make([]int, 1)
	slice2 := make([]int, 1)
	fmt.Println("slice init", slice1, slice2)
	slicePass(&slice1, slice2)
	fmt.Println("slice after", slice1, slice2)

}
