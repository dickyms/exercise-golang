package main

import(
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey`s value is %s\n", ctx.Value("myKey"))
	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)
	fmt.Printf("doSomething: myKey`s value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doSomething: myKey`s value is %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	resultsCh := make(chan *WorkResult)
	for {
		select {
		case <- ctx.done():
			return
		case result := <- resultsCh
		}
	}
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}

/*For example, if a web page request comes to your Go web server, the user may end up hitting the “Stop” button or closing their browser before the page finishes loading. If the page they requested requires running a few database queries, the server may run the queries even though the data won’t be used. However, if your functions are using a context.Context, your functions will know the context is done because Go’s web server will cancel it, and that they can skip running any other database queries they haven’t already run. This will free up web server and database server processing time so it can be used on a different request instead.
https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
*/