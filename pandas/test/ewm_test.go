package test

import (
	"fmt"
	"testing"
)

// https://blog.csdn.net/huangguohui_123/article/details/105066646?spm=1001.2101.3001.6650.3&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-3-105066646-blog-99670246.pc_relevant_multi_platform_whitelistv4&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-3-105066646-blog-99670246.pc_relevant_multi_platform_whitelistv4&utm_relevant_index=5

func TestEWM(t *testing.T) {
	samples := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	fmt.Println("---1---")
	t1 := EWM(samples, 5)
	fmt.Println(t1)
	//t2 := formula.MA(t1, 5)
	//fmt.Println(t2)
	fmt.Println("---2---")

	//fmt.Println("---1---")
	//t1 := EWM(samples, 5)
	//fmt.Println(t1)
	//fmt.Println("---2---")

	// pervious = beta*pervious + float64(1-beta)*v
	//x := 1.00

	//df=pd.DataFrame({'x':[1,2,3,4,5,6,7,8,9]})
	//df['y2']=df['x'].ewm(span=5,adjust=False).mean()
	//print(df)
	//
	// span = 5
	//   x        y2
	//0  1  1.000000
	//1  2  1.250000
	//2  3  1.687500
	//3  4  2.265625
	//4  5  2.949219
	//5  6  3.711914
	//6  7  4.533936
	//7  8  5.400452
	//8  9  6.300339
	//t2 := x * 6.117055 + (1-x) * 9
	//x * 6.117055 + 9 - x *9
	//x * (6.117055 - 9) + 9 = 7.078037
	/*
		x8 := (6.117055 - 8) / (5.175583 - 8)
		fmt.Println("x8=", x8)

		x9 := (7.078037 - 9) / (6.117055 - 9)
		fmt.Println("x9=", x9)
	*/

}
