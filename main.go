package main

import (
	"fmt"
	"syscall/js"
)

var (
	count = 0
	done = make(chan string)
)




//func main()  {
//	sayHelloCallback := js.NewCallback(sayHello)
//	defer sayHelloCallback.Release()
//	sayHello := js.Global().Get("sayHello")
//	sayHello.Invoke(sayHelloCallback)
//	<- done
//}
//
//func sayHello(args []js.Value)  {
//	alert := js.Global().Get("alert")
//	alert.Invoke(args[0].String())
//	done <- "done"
//}




//func main()  {
//	alert := js.Global().Get("alert")
//	alert.Invoke("hello world")
//}

//func main()  {
//	fmt.Println("hello world")
//}


//var(
//	count = 0
//	done = make(chan string)
//)

//  func main()  {
//	callback := js.NewCallback(sayHello)
//	defer callback.Release()
//	sayMessage := js.Global().Get("sayMessage")
//	sayMessage.Invoke(callback)
//	<- done
//}
//
//
//func sayHello(args []js.Value) {
//	value := args[0].String()
//	fmt.Println(value)
//	done <- "done"
//}

func main()  {
	sayMessageCallback := js.NewCallback(sayMessageFunc)
	defer sayMessageCallback.Release()
	sayMessage := js.Global().Get("sayMessage")
	sayMessage.Invoke(sayMessageCallback)

	unbindCallback := js.NewEventCallback(0, unbindFunc)
	defer unbindCallback.Release()
	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", unbindCallback)

	<- done
}

func sayMessageFunc(args []js.Value) {
	fmt.Println(args[0].String(), count)
	count++
}

func unbindFunc(event js.Value) {
	done <- "done"
}

