package netRetry

//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestFewTry(t *testing.T) {
//	ctx:=BackGroundContext()
//	nr:=NewNetRetry(ctx,10,5,20)
//	status:=Success
//	var nextJump uint16
//	nextJump=1
//	i:=1
//	for status==Success {
//		status=nr.Try(ctx,&nextJump)
//		fmt.Printf("第%d回合，等待时间为%d \n",i,nextJump)
//		i+=1
//	}
//}
