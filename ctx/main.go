package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
最佳实践
	1、context.Background 只应用在最高等级，作为所有派生 context 的根。
	2、context.TO DO 应用在不确定要使用什么的地方，或者当前函数以后会更新以便使用 context。
	3、context 取消是建议性的，这些函数可能需要一些时间来清理和退出。
	4、context.Value 应该很少使用，它不应该被用来传递可选参数。这使得 API 隐式的并且可以引起错误。取而代之的是，这些值应该作为参数传递。
	5、不要将 context 存储在结构中，在函数中显式传递它们，最好是作为第一个参数。
	6、永远不要传递不存在的 context 。相反，如果您不确定使用什么，使用一个 To Do context。
	7、Context 结构没有取消方法，因为只有派生 context 的函数才应该取消 context。
*/
func main() {
	/*
		这个函数返回一个空 context。
		这只能用于高等级（在 main 或顶级请求处理中）。
		这能用于派生我们稍后谈及的其他 context 。
	*/
	backgroundCtx := context.Background()
	fmt.Println(backgroundCtx.Value(1))

	/*
		这个函数也是创建一个空 context。
		也只能用于高等级或当您不确定使用什么 context，或函数以后会更新以便接收一个 context 。
		这意味您（或维护者）计划将来要添加 context 到函数。
		有趣的是，查看代码，它与 background 完全相同。
		不同的是，静态分析工具可以使用它来验证 context 是否正确传递，这是一个重要的细节，因为静态分析工具可以帮助在早期发现潜在的错误，并且可以连接到 CI/CD 管道。
	*/
	todoCtx := context.TODO()
	fmt.Println(todoCtx.Value(1))

	/*
		此函数接收 context 并返回派生 context，其中值 val 与 key 关联，并通过 context 树与 context 一起传递。
		这意味着一旦获得带有值的 context，从中派生的任何 context 都会获得此值。
		不建议使用 context 值传递关键参数，而是函数应接收签名中的那些值，使其显式化。
	*/
	valueCtx := context.WithValue(context.Background(), 1, 2)
	fmt.Println(valueCtx.Value(1))

	valueCtx2 := context.WithValue(valueCtx, 2, 3)
	fmt.Println(valueCtx2.Value(1))
	fmt.Println(valueCtx2.Value(2))

	valueCtx3 := context.WithValue(valueCtx, 1, 3)
	fmt.Println(valueCtx3.Value(1))
	fmt.Println(valueCtx3.Value(2))

	/*
		返回派生 context 和取消函数。只有创建它的函数才能调用取消函数来取消此 context。
		如果您愿意，可以传递取消函数，但是，强烈建议不要这样做。
		这可能导致取消函数的调用者没有意识到取消 context 的下游影响。
		可能存在源自此的其他 context，这可能导致程序以意外的方式运行。简而言之，永远不要传递取消函数。
	*/
	cancelCtx, cancel := context.WithCancel(valueCtx)
	fmt.Println(cancelCtx.Value(1))
	cancel() //执行完业务操作后，可以取消此上下文

	/*
	   此函数返回其父项的派生 context，当截止日期超过或取消函数被调用时，该 context 将被取消。
	   例如，您可以创建一个将在以后的某个时间自动取消的 context，并在子函数中传递它。
	   当因为截止日期耗尽而取消该 context 时，获此 context 的所有函数都会收到通知去停止运行并返回。
	*/
	deadlineCtx, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	fmt.Println(deadlineCtx.Value(1))
	select {
	case <-deadlineCtx.Done():
		fmt.Println("done")
	case <-time.Tick(time.Millisecond * 1001):
		deadlineCancel()
		fmt.Println("deadCancel ")
	}

	/*
		此函数类似于 context.WithDeadline。
		不同之处在于它将持续时间作为参数输入而不是时间对象。
		此函数返回派生 context，如果调用取消函数或超出超时持续时间，则会取消该派生 context。
	*/
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Println(timeoutCtx.Value(1))
	select {
	case <-deadlineCtx.Done():
		fmt.Println("timeout done")
	case <-time.Tick(time.Millisecond * 1001):
		timeoutCancel()
		fmt.Println("timeoutCancel ")
	}

	demo()
}

//Slow function
func sleepRandom(fromFunction string, ch chan int) {
	//defer cleanup
	defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()
	//Perform a slow task
	//For illustration purpose,
	//Sleep here for random ms
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100
	fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")
	//write on the channel if it was passed in
	if ch != nil {
		ch <- sleeptime
	}
}

//Function that does slow processing with a context
//Note that context is the first argument
func sleepRandomContext(ctx context.Context, ch chan bool) {
	//Cleanup tasks
	//There are no contexts being created here
	//Hence, no canceling needed
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()
	//Make a channel
	sleeptimeChan := make(chan int)
	//Start slow processing in a goroutine
	//Send a channel for communication
	go sleepRandom("sleepRandomContext", sleeptimeChan)
	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//If context is cancelled, this case is selected
		//This can happen if the timeout doWorkContext expires or
		//doWorkContext calls cancelFunction or main calls cancelFunction
		//Free up resources that may no longer be needed because of aborting the work
		//Signal all the goroutines that should stop work (use channels)
		//Usually, you would send something on channel,
		//wait for goroutines to exit and then return
		//Or, use wait groups instead of channels for synchronization
		fmt.Println("sleepRandomContext: Time to return")
	case sleeptime := <-sleeptimeChan:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}

//A helper function, this can, in the real world do various things.
//In this example, it is just calling one function.
//Here, this could have just lived in main
func doWorkContext(ctx context.Context) {
	//Derive a timeout context from context with cancel
	//Timeout in 150 ms
	//All the contexts derived from this will returns in 150 ms
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	//Cancel to release resources once the function is complete
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()
	//Make channel and call context function
	//Can use wait groups as well for this particular case
	//As we do not use the return value sent on channel
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)
	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//This case is selected when the passed in context notifies to stop work
		//In this example, it will be notified when main calls cancelFunction
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("sleepRandomContext returned")
	}
}

func demo() {
	//Make a background context
	ctx := context.Background()
	//Derive a context with cancel
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	//defer canceling so that all the resources are freed up
	//For this and the derived contexts
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunction()
	}()
	//Cancel context after a random time
	//This cancels the request after a random timeout
	//If this happens, all the contexts derived from this should return
	go func() {
		sleepRandom("Main", nil)
		cancelFunction()
		fmt.Println("Main Sleep complete. canceling context")
	}()
	//Do work
	doWorkContext(ctxWithCancel)
}
