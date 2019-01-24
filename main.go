package main;

import (
	"fmt"
	"time"
	"context"
)

func testContext(timeout time.Duration) (*context.Context, func()){
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	return &ctx, cancel
}

func main(){
	xd := make(chan *context.Context)
	xdd := make(chan func())
	go func(){
		var timeout time.Duration = 5
		ctx, cancel := testContext(timeout)
		xd <- ctx
		xdd <- cancel
	}()

	mainContext := <- xd

	// Comment one code block and uncomment the other

	// Code block 1:
	
	mainCancel := <- xdd
	for{
		select{
		case <-(*mainContext).Done():
			fmt.Println("Done!")
			return
		default:
			fmt.Println(".")
			time.Sleep(1*time.Second)
			mainCancel()
		}
	}

	// Code block 2:

	// for{
	// 	select{
	// 	case <-(*mainContext).Done():
	// 		fmt.Println("Done!")
	// 		return
	// 	default:
	// 		fmt.Println(".")
	// 		time.Sleep(1*time.Second)
	// 	}
	// }
}