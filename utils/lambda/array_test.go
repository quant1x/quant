package lambda

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type user struct {
	name string
	age  int
}

type account struct {
	name string
	age  int
}

const count = 10000

func makeIntArray() []int {
	want := make([]int, count)
	for i := 0; i < count; i++ {
		want[i] = i + 1
	}
	return want
}

func makeUserArray() []user {
	want := make([]user, count)
	for i := 0; i < count; i++ {
		want[i] = user{"un:" + strconv.Itoa(i+1), i + 1}
	}
	return want
}

func report(t *testing.T, start time.Time) {
	end := time.Now()
	ms := float32(end.Nanosecond()-start.Nanosecond()) / float32(1e6)
	t.Log(fmt.Sprintf("run time %.2f ms", ms))
}

func isTrue(tv interface{}, v bool) {
	t := tv.(*testing.T)
	if !v {
		t.Fail()
		panic(v)
	}
}

func isFalse(tv interface{}, v bool) {
	t := tv.(*testing.T)
	if v {
		t.Fail()
		panic(v)
	}
}

func Test__array_Join(t *testing.T) {
	defer report(t, time.Now())
	result := LambdaArray(makeIntArray()).Join(JoinOptions{
		express: func(e int) string { return strconv.Itoa(e) },
	})
	t.Log("string length", len(result))

	arr := []int{1, 2, 3, 4, 5}
	str1 := LambdaArray(arr).Join(JoinOptions{
		express: func(e int) string { return strconv.Itoa(e) },
	})
	fmt.Println(str1)

	str2 := LambdaArray(arr).Join(JoinOptions{
		express: func(e int) string { return strconv.Itoa(e) },
		Symbol:  "|",
	})
	fmt.Println(str2)
}

func Test__array_Filter(t *testing.T) {
	defer report(t, time.Now())
	want := makeIntArray()
	ret := LambdaArray(want).Filter(
		func(ele int) bool { return ele%3 == 0 }).Pointer().([]int)
	isTrue(t, len(ret) == count/3)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	larr := LambdaArray(arr)
	ret1 := larr.Filter(func(ele int) bool { return ele > 5 }).Pointer().([]int)
	fmt.Println(ret1)

	ret2 := larr.Filter(func(ele int) bool { return ele%2 == 0 }).Pointer().([]int)
	fmt.Println(ret2)

	users := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	ret3 := LambdaArray(users).Filter(func(u user) bool { return u.age < 30 }).Pointer().([]user)
	fmt.Println(ret3)
}

func Test__array_Sort_Quick(t *testing.T) {
	defer report(t, time.Now())
	want := make([]int, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		want[i] = rand.Intn(count * 10)
	}
	t.Log(want[:10], "...", want[count-10:], " count=", len(want))
	ret := LambdaArray(want).Sort(func(e1, e2 int) bool {
		return e1 > e2
	}).Pointer().([]int)
	t.Log(ret[:10], "...", ret[count-10:], " count=", len(ret))

	arr := []int{1, 3, 8, 6, 12, 5, 9}
	// order by asc
	ret1 := LambdaArray(arr).Sort(func(a, b int) bool { return a < b }).Pointer().([]int)
	// order by desc
	ret2 := LambdaArray(arr).Sort(func(a, b int) bool { return a > b }).Pointer().([]int)

	fmt.Println(ret1)
	fmt.Println(ret2)

	users := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	ret3 := LambdaArray(users).Sort(func(a, b user) bool { return a.age < b.age }).Pointer().([]user)
	fmt.Println(ret3)
}

func Test__array_Sort_QuickMT(t *testing.T) {
	defer report(t, time.Now())
	want := make([]int, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		want[i] = rand.Intn(count * 10)
	}
	t.Log(want[:10], "...", want[count-10:], " count=", len(want))
	ret := LambdaArray(want).SortMT(func(e1, e2 int) bool {
		return e1 > e2
	}).Pointer().([]int)
	t.Log(ret[:10], "...", ret[count-10:], " count=", len(ret))
}

func Test__array_Map(t *testing.T) {
	defer report(t, time.Now())

	result := LambdaArray(makeIntArray()).Map(func(e int) int {
		return e + 1
	}).Pointer().([]int)

	isTrue(t, len(result) == count)

	arr := LambdaArray([]int{1, 2, 3, 4, 5})
	users := arr.Map(func(i int) user {
		return user{name: "un:" + strconv.Itoa(i), age: i}
	}).Pointer().([]user)
	fmt.Println(users)
}

func Test__array_Append(t *testing.T) {
	defer report(t, time.Now())
	want := LambdaArray(makeIntArray())
	want.Append(count + 1)
	isTrue(t, count+1 == want.Count(nil))

	arr := LambdaArray([]int{1, 2, 3})
	arr.Append(4)
	fmt.Println(arr.Pointer().([]int))
	arr.Append(5, 6)
	fmt.Println(arr.Pointer().([]int))
}

func (p account) CompareTo(a interface{}) int {
	return p.age - a.(account).age
}

func Test__array_Max(t *testing.T) {
	defer report(t, time.Now())
	want := []int{1, 5, 6, 3, 8, 9, 3, 12, 56, 186, 4, 9, 14}

	var iArr = LambdaArray(want)

	ret := iArr.Max(nil).(int)
	t.Log(ret)
	ret = iArr.Max(func(ele int) int { return ele }).(int)
	t.Log(ret)

	wantUsers := iArr.Map(func(ele int) account {
		s := fmt.Sprintf("%d", ele)
		return account{"zzz" + s, ele}
	})

	ret2 := wantUsers.Max(func(u account) int { return u.age })
	t.Log(ret2)

	ret3 := wantUsers.Max(nil)
	t.Log(ret3)

	users := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	eldest := LambdaArray(users).Max(func(u user) int { return u.age }).(user)
	fmt.Println(eldest.name + " is the eldest")
}

func Test__array_Sort_Min(t *testing.T) {

	defer report(t, time.Now())

	want := []int{1, 5, 6, 3, 8, 9, 3, 12, 56, 186, 4, 9, 14}

	var iArr = LambdaArray(want)

	ret := iArr.Min(nil).(int)
	t.Log(ret)
	ret = iArr.Min(func(ele int) int { return ele }).(int)
	t.Log(ret)

	wantUsers := iArr.Map(func(ele int) account {
		s := fmt.Sprintf("%d", ele)
		return account{"zzz" + s, ele}
	})

	ret2 := wantUsers.Min(func(u account) int { return u.age })
	t.Log(ret2)

	ret3 := wantUsers.Min(nil)
	t.Log(ret3)

	users := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	eldest := LambdaArray(users).Min(func(u user) int { return u.age }).(user)
	fmt.Println(eldest.name + " is the Charles")
}

func Test__array_Any(t *testing.T) {
	defer report(t, time.Now())
	ints := LambdaArray(makeIntArray())
	users := LambdaArray(makeUserArray())
	ret := []bool{
		ints.Any(nil),
		ints.Any(func(ele int) bool { return ele > 99999999 }),
		users.Any(func(u user) bool { return u.name == "un:1997" }),
	}
	isTrue(t, ret[0])
	isFalse(t, ret[1])
	isTrue(t, ret[2])

	us := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	ret1 := LambdaArray(us).Any(func(u user) bool { return u.age > 30 })
	fmt.Println(ret1)
	ret2 := LambdaArray(us).Any(func(u user) bool { return u.age < 0 })
	fmt.Println(ret2)
}

func Test__array_All(t *testing.T) {
	defer report(t, time.Now())
	ints := LambdaArray(makeIntArray())
	users := LambdaArray(makeUserArray())
	ret := []bool{
		ints.All(nil),
		ints.All(func(ele int) bool { return ele > 0 }),
		users.All(func(u user) bool { return u.name == "un:1997" }),
	}
	isTrue(t, ret[0])
	isTrue(t, ret[1])
	isFalse(t, ret[2])

	us := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	ret1 := LambdaArray(us).All(func(u user) bool { return u.age > 30 })
	fmt.Println(ret1)
	ret2 := LambdaArray(us).All(func(u user) bool { return u.age > 10 })
	fmt.Println(ret2)
}

func Test__array_Count(t *testing.T) {
	defer report(t, time.Now())
	ints := LambdaArray(makeIntArray())
	ret := []int{
		ints.Count(nil),
		ints.Count(func(ele int) bool { return ele%2 == 0 }),
	}
	isTrue(t, ret[0] == count)
	isTrue(t, ret[1]*2 == count)

	us := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	ret1 := LambdaArray(us).Count(func(u user) bool { return u.age > 30 })
	fmt.Println(ret1)
	ret2 := LambdaArray(us).Count(func(u user) bool { return u.age > 20 })
	fmt.Println(ret2)
}

func Test__array_First(t *testing.T) {
	defer report(t, time.Now())
	want := []int{1, 5, 6, 3, 8, 9, 3, 12, 56, 186, 4, 9, 14}
	if c, err := LambdaArray(want).First(func(e int) bool { return e > 30 }); err == nil {
		t.Log(c)
		isTrue(t, c == 56)
	} else {
		t.Fail()
	}

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

	if u, err := arr.First(func(u user) bool { return u.name == "jack" }); err == nil {
		fmt.Println(u, " found")
	} else {
		fmt.Println("not found")
	}
}

func Test__array_Index(t *testing.T) {
	defer report(t, time.Now())
	ints := LambdaArray(makeIntArray())
	ret := []int{
		ints.Count(nil),
		ints.Count(func(ele int) bool { return ele%2 == 0 }),
	}
	isTrue(t, ret[0] == count)
	isTrue(t, ret[1]*2 == count)

	if element, err := LambdaArray([]int{1, 2, 3, 4, 5}).Index(3); err == nil {
		fmt.Println(element)
	} else {
		fmt.Println(err)
	}
	if element, err := LambdaArray([]int{1, 2, 3, 4, 5}).Index(10); err == nil {
		fmt.Println(element)
	} else {
		fmt.Println(err)
	}
}

func Test__array_Last(t *testing.T) {
	defer report(t, time.Now())
	want := []int{1, 5, 6, 3, 8, 9, 3, 12, 56, 186, 4, 9, 14}
	if c, err := LambdaArray(want).Last(func(e int) bool { return e > 30 }); err == nil {
		t.Log(c)
		isTrue(t, c == 186)
	} else {
		t.Fail()
	}

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

	if u, err := arr.Last(func(u user) bool { return u.age > 35 }); err == nil {
		fmt.Println(u, " found")
	} else {
		fmt.Println("not found")
	}
}

func Test__array_Take(t *testing.T) {
	defer report(t, time.Now())
	ints := LambdaArray(makeIntArray())
	ret := ints.Take(200, 10).Pointer().([]int)
	isTrue(t, ret[0] == 201)
	isTrue(t, ret[9] == 210)

	ret1 := LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Take(4, 10).Pointer().([]int)
	fmt.Println(ret1)
	ret2 := LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Take(10, 10).Pointer().([]int)
	fmt.Println(ret2)
}

func Test__array_Sum(t *testing.T) {
	defer report(t, time.Now())
	ret := LambdaArray(makeIntArray()).Sum(nil).(int)
	ret2 := LambdaArray(makeUserArray()).Sum(func(u user) int { return u.age })
	t.Log(ret, ret2)

	us := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	arr := LambdaArray(us)
	fmt.Println("total user age is", arr.Sum(func(u user) int { return u.age }))
}

func Test__array_Avg(t *testing.T) {
	defer report(t, time.Now())
	ints := LambdaArray(makeIntArray())
	ret := ints.Average(nil)
	t.Log(ret)

	us := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	arr := LambdaArray(us)
	fmt.Println("all user average age is", arr.Average(func(u user) int { return u.age }))
}

func Test__array_Contain(t *testing.T) {

	defer report(t, time.Now())

	want := makeIntArray()
	arr := LambdaArray(want)
	ret := []bool{arr.Contains(7777), arr.Contains(count + 1)}
	isTrue(t, ret[0])
	isFalse(t, ret[1])

	users := LambdaArray(makeUserArray())
	ret = []bool{
		users.Contains(user{"un:18", 18}),
		users.Contains(user{"zzz", 18}),
	}
	isTrue(t, ret[0])
	isFalse(t, ret[1])

	ret = []bool{
		users.Contains(func(u user) bool { return u.age > 5000 }),
		users.Contains(func(u user) bool { return u.age > count+1 }),
	}
	isTrue(t, ret[0])
	isFalse(t, ret[1])

	us := []user{
		{"Abraham", 20},
		{"Edith", 25},
		{"Charles", 40},
		{"Anthony", 26},
		{"Abel", 33},
	}
	arr2 := LambdaArray(us)
	fmt.Println(arr2.Contains(func(u user) bool { return u.age > 25 }))

	fmt.Println(LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Contains(9))
	fmt.Println(LambdaArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Contains(0))
}

func (u user) Equals(obj interface{}) bool {
	if c, ok := obj.(user); ok {
		return u.name == c.name && u.age == c.age
	}
	return false
}
