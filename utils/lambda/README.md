# lambda

## overview

`lambda` is a lambda expression for go,lets you extract elements from array by lambda expression

## Installation

go get github.com/favar/lambda

## Getting Started

#### use LambdaArray returns Array interface

```go
sa := []int{1,2,3,4,5,6,7,8,9}
arr := LambdaArray(sa) // return Array<int>
```



#### interface Array

```go
type Array interface {
	IsSlice() bool
	Join(options JoinOptions) string
	Filter(express interface{}) Array
	Sort(express interface{}) Array
	SortMT(express interface{}) Array
	Map(express interface{}) Array
	Append(elements ...interface{}) Array
	Max(express interface{}) interface{}
	Min(express interface{}) interface{}
	Any(express interface{}) bool
	All(express interface{}) bool
	Count(express interface{}) int
	First(express interface{}) (interface{}, error)
	Last(express interface{}) (interface{}, error)
	index(i int) (interface{}, error)
	Take(skip, count int) Array
	Sum(express interface{}) interface{}
	Average(express interface{}) float64
	Contains(express interface{}) bool
	Pointer() interface{}
}
```

## Usage

***define test struct***

```go
type user struct {
	name string
	age  int
}
```



#### Join

array join into string

```go
type JoinOptions struct {
	Symbol  string // split string,default `,`
    express interface{} // express match func(ele TElement) string
}
Join(options JoinOptions) string
```

```
arr := []int{1,2,3,4,5}
str1 := LambdaArray(arr).Join(JoinOptions{
	express: func(e int) string { return strconv.Itoa(e) },
})
fmt.Println(str1) // 1,2,3,4,5 default `,`

str2 := LambdaArray(arr).Join(JoinOptions{
    express: func(e int) string { return strconv.Itoa(e) },
    Symbol:  "|",
})
fmt.Println(str2) // 1|2|3|4|5

 
```



#### Filter

array filter

```go
Filter(express interface{}) Array // express match func(ele TElement) bool
```

```go
arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
larr := LambdaArray(arr)
ret1 := larr.Filter(func(ele int) bool { return ele > 5 }).Pointer().([]int)
fmt.Println(ret1) // [6 7 8 9 10]

ret2 := larr.Filter(func(ele int) bool { return ele%2 == 0 }).Pointer().([]int)
fmt.Println(ret2) // [2 4 6 8 10]

ret3 := LambdaArray(users).Filter(func(u user) bool { return u.age < 30 }).Pointer().([]user)
fmt.Println(ret3) // [{Abraham 20} {Edith 25} {Anthony 26}]
```



#### Sort

quick sort

```go
Sort(express interface{}) Array // express match func(e1, e2 TElement) bool
```

```go
arr := []int{1, 3, 8, 6, 12, 5, 9}
// order by asc
ret1 := LambdaArray(arr).Sort(func(a, b int) bool { return a < b }).Pointer().([]int)
// order by desc
ret2 := LambdaArray(arr).Sort(func(a, b int) bool { return a > b }).Pointer().([]int)

fmt.Println(ret1) // [1 3 5 6 8 9 12]
fmt.Println(ret2) // [12 9 8 6 5 3 1]

users := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
// order by user.age asc
ret3 := LambdaArray(users).Sort(func(a, b user) bool { return a.age < b.age }).Pointer().([]user)
fmt.Println(ret3) // [{Abraham 20} {Edith 25} {Anthony 26} {Abel 33} {Charles 40}]
```



#### SortMT

sort by quick multithreading

usage like Sort

#### Map 

.map to new array 

```go
Map(express interface{}) Array // express match func(ele TElement) TOut
```

```go
arr := LambdaArray([]int{1, 2, 3, 4, 5})
users := arr.Map(func(i int) user {
    return user{name: "un:" + strconv.Itoa(i), age: i}
}).Pointer().([]user)
fmt.Println(users) // [{un:1 1} {un:2 2} {un:3 3} {un:4 4} {un:5 5}]

```



#### Append

.append element

```go
Append(elements ...interface{}) Array // each of elements type must be TElmenent
```

```go
arr := LambdaArray([]int{1, 2, 3})
arr.Append(4)
fmt.Println(arr.Pointer().([]int)) // [1 2 3 4]
arr.Append(5, 6)
fmt.Println(arr.Pointer().([]int)) // [1 2 3 4 5 6]
```

#### Max

.maximum element of array

```go
Max(express interface{}) interface{}
```

```go
users := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
eldest := LambdaArray(users).Max(func(u user) int { return u.age }).(user)
fmt.Println(eldest.name + " is the eldest") // Charles is the eldest

want := []int{1, 5, 6, 3, 8, 9, 3, 12, 56, 186, 4, 9, 14}
var iArr = LambdaArray(want)
ret := iArr.Max(nil).(int)
fmt.Println(ret) // 186
```



#### Min

.minimum element of array

```go
Min(express interface{}) interface{}
```

```go
users := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
eldest := LambdaArray(users).Min(func(u user) int { return u.age }).(user)
fmt.Println(eldest.name + " is the eldest") // Abraham is the Charles

want := []int{1, 5, 6, 3, 8, 9, 3, 12, 56, 186, 4, 9, 14}
var iArr = LambdaArray(want)
ret := iArr.Min(nil).(int)
fmt.Println(ret) // 1
```



#### Any

.Determines whether the Array contains any elements

```go
Any(express interface{}) bool
```

```go
us := []user{
   {"Abraham", 20},
   {"Edith", 25},
   {"Charles", 40},
   {"Anthony", 26},
   {"Abel", 33},
}
ret1 := LambdaArray(us).Any(func(u user) bool { return u.age > 30 })
fmt.Println(ret1) // true
ret2 := LambdaArray(us).Any(func(u user) bool { return u.age < 0 })
fmt.Println(ret2) // false
```

#### All

Determines whether the condition is satisfied for all elements in the Array

```go
All(express interface{}) bool
```

```go
us := []user{
   {"Abraham", 20},
   {"Edith", 25},
   {"Charles", 40},
   {"Anthony", 26},
   {"Abel", 33},
}
ret1 := LambdaArray(us).All(func(u user) bool { return u.age > 30 })
fmt.Println(ret1) // false
ret2 := LambdaArray(us).All(func(u user) bool { return u.age > 10 })
fmt.Println(ret2) // true
```

#### Count

Returns a number indicating how many elements in the specified Array satisfy the condition

```go
Count(express interface{}) int
```

```go
us := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
ret1 := LambdaArray(us).Count(func(u user) bool { return u.age > 30 })
fmt.Println(ret1) // 2
ret2 := LambdaArray(us).Count(func(u user) bool { return u.age > 20 })
fmt.Println(ret2) // 4
```



#### First

Returns the first element of an Array that satisfies the condition

```go
First(express interface{}) (interface{}, error)
```

```go
us := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
arr := LambdaArray(us)
if u, err := arr.First(func(u user) bool { return u.name == "Charles" }); err == nil {
    fmt.Println(u, " found")
} else {
    fmt.Println("not found")
}
// {Charles 40}  found
if u, err := arr.First(func(u user) bool { return u.name == "jack" }); err == nil {
    fmt.Println(u, " found")
} else {
    fmt.Println("not found")
}
// not found

```

#### Last

Returns the last element of an Array that satisfies the condition

```go
Last(express interface{}) (interface{}, error)
```

```go
us := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
arr := LambdaArray(us)
if u, err := arr.Last(func(u user) bool { return u.name == "Anthony" }); err == nil {
    fmt.Println(u, " found")
} else {
    fmt.Println("not found")
}
// {Anthony 26}  found
if u, err := arr.Last(func(u user) bool { return u.age > 35 }); err == nil {
    fmt.Println(u, " found")
} else {
    fmt.Println("not found")
}
// {Charles 40}  found
```



#### Index

Returns the zero based index of the first occurrence in an Array

```go
Index(i int) (interface{}, error)
```

```go
if element, err := LambdaArray([]int{1, 2, 3, 4, 5}).Index(3); err == nil {
    fmt.Println(element)
} else {
    fmt.Println(err)
}
// 4
if element, err := LambdaArray([]int{1, 2, 3, 4, 5}).Index(10); err == nil {
    fmt.Println(element)
} else {
    fmt.Println(err)
}
// 10 out of range
```



#### Take

take `count` elements start by `skip`

```go
Take(skip, count int) Array
```

```go
ret1 := LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Take(4, 10).Pointer().([]int)
fmt.Println(ret1) // [5 6 7 8 9 10]
ret2 := LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Take(10, 10).Pointer().([]int)
fmt.Println(ret2) // []
```



#### Sum

sum of the values returned by the expression

```go
Sum(express interface{}) interface{}
```

```go
us := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
arr := LambdaArray(us)
fmt.Println("total user age is ", arr.Sum(func(u user) int { return u.age }))
// total user age is 144
```



#### Average

average of the values returned by the expression

```go
Average(express interface{}) float64
```

```go
us := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
arr := LambdaArray(us)
fmt.Println("all user average age is", arr.Average(func(u user) int { return u.age }))
// all user average age is 28.8
```



#### Contains

Determines whether the array contains the specified element

```go
Contains(express interface{}) bool
```

```go
us := []user{
    {"Abraham", 20},
    {"Edith", 25},
    {"Charles", 40},
    {"Anthony", 26},
    {"Abel", 33},
}
arr2 := LambdaArray(us)
fmt.Println(arr2.Contains(func(u user) bool { return u.age > 25 })) //true

fmt.Println(LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Contains(9)) // true
fmt.Println(LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Contains(0)) // false
```

#### Pointer

array or slice pointer

```go
Pointer() interface{}
```



## Tutorial

Usage

## Questions

Please let me know if you have any questions.



