package main

import (
	"github.com/pkg/errors"
	"time"
)
type errorString string;
func main()  {
/*	println("进入main")
	printFuture := func() (interface{},error){println("进入hello ");time.Sleep(50 * time.Millisecond);println("hello");return nil,nil}
	printAFutureA := func() (interface{},error){println("进入world");time.Sleep(10 * time.Millisecond);println("world ");return nil,nil}

	println("两个异步goroutine声明完毕")
	requestFuture := NewFuture(printFuture)
	requestFutureA := NewFuture(printAFutureA)
	println("两个异步goroutine等待完毕")
	 result1,_:=requestFuture.get();
	 result2,_:=requestFutureA.get()
	 println(result1,result2)*/
//rpc.Accept()
//	if ErrNameTyep == New("EOF"){
//		println("named type error")
//	}
//	fmt.Printf("adqqqd:%T \n",ErrNameTyep )
//
//	fmt.Printf("adsfad:%T \n",ErrStructType )
//	if ErrStructType == errors.New("EOF"){
//		println("struct type o")
//	}
//最多改变一个元素情况下 对于数列 实现非递减数列
	//[3,4,2,3] 无法实现。最简单的实现判断是否只有一次单调下降 ==>会导致bug
	//[4,2,3] = > 4->1 [1,2,3] ==> true
	//[4,2,3] => [1,2,3]
	//aa := []int{4,2,3};
	//println(checkPossibility(aa))
//早餐组合
	staple := []int{10,20,5};
	drinks := []int{5,5,2}
	x := 15;
	println(breakfastNumber(staple,drinks,x))

}
/*
  这道题给了我们一个数组，说我们最多有1次修改某个数字的机会，
  问能不能将数组变为非递减数组。题目中给的例子太少，不能覆盖所有情况，我们再来看下面三个例子：
	4，2，3
	-1，4，2，3
	2，3，3，2，4
我们通过分析上面三个例子可以发现，当我们发现后面的数字小于前面的数字产生冲突后，
[1]有时候需要修改前面较大的数字(比如前两个例子需要修改4)，
[2]有时候却要修改后面较小的那个数字(比如前第三个例子需要修改2)，
那么有什么内在规律吗？是有的，判断修改那个数字其实跟再前面一个数的大小有关系，
首先如果再前面的数不存在，比如例子1，4前面没有数字了，我们直接修改前面的数字为当前的数字2即可。
而当再前面的数字存在，并且小于当前数时，比如例子2，-1小于2，我们还是需要修改前面的数字4为当前数字2；
如果再前面的数大于当前数，比如例子3，3大于2，我们需要修改当前数2为前面的数3。

*/
func checkPossibility(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1{
		return true;
	}
	//4,2,3
	var count int = 0;
	for i:=1;i<len(nums) && count<2; i++ {
		if nums[i-1] <= nums[i]{
			continue
		}
		count++;
		if i>=2 && nums[i-2] > nums[i] {
			println("count1",count)
			nums[i] = nums[i-1];
		}else {
			println("count2",count)
			nums[i-1] = nums[i];
		}
	}
	return count <= 1;
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	count := 0;
	for _,v := range staple{
		for _,b := range drinks{
			if v+b <=x{
				//println ("b",b,v)
				count++;
			}
		}
	}
	return count;
}
func (e errorString) Error() string {
	return string(e)
}
func New(test string) error{
	return errorString(test)
}
var ErrNameTyep = New("EOF");
var ErrStructType = errors.New("EOF")

type Future struct {
	result      interface{}
	err         error
	signal      chan struct{}
	isCompleted bool
}
type FutureFunc func()(interface{},error)
func (f *Future) getA(d time.Duration)(result interface{},err error,timeout bool){
	select {
	case <- time.After(d):
		return nil,nil,true
	case <-f.signal:
		return f.result,f.err,false;
	}
}

func (f *Future) get()(result interface{},err error){
	<- f.signal
	return f.result, err
}

func NewFuture(fun FutureFunc) *Future  {
	 f := new(Future);
	f.signal = make(chan struct{},1);
	go func() {
		defer close(f.signal);
		result ,err := fun();
		f.result = result;
		f.err = err;
		f.isCompleted = true;
	}()
	return f;
}

type aa struct {//class
	id int;
	name string;
}




/*
public boolean checkPossibility(int[] nums) {
if (nums == null || nums.length <= 1) {
return true;
}
int cnt = 0;
for (int i = 1; i < nums.length && cnt < 2; i++) {
if (nums[i-1] <= nums[i]) {
continue;
}
cnt++;
if (i-2>=0 && nums[i-2] > nums[i]) {
nums[i] = nums[i-1];
}else {
nums[i-1] = nums[i];
}
}
return cnt <= 1;
}
*/