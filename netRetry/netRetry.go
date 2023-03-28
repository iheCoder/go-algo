package netRetry

//
//import (
//	"math"
//	"math/rand"
//)
//
//const RETRY_FOREVER = math.MaxUint32
//
//type RetryStatus int
//
//const (
//	Success RetryStatus = iota
//	Exhausted
//)
//
//type RetryContext struct {
//	// 连续重试之间的最大退避延迟（以毫秒为单位）
//	maxBackOffDelay uint16
//	// 重试次数
//	attemptDone      uint32
//	// 下次重试的最大退避值（以毫秒为单位）
//	nextJitterMax    uint16
//	// 最大重试次数
//	maxRetryAttempts uint32
//}
//
//type NetRetry struct {
//	maxAttempts uint32
//	maxBackOff uint16
//	backOffBase uint16
//}
//
//func (nr NetRetry) Try(ctx *RetryContext,pNextBackOff *uint16) RetryStatus {
//	status:=Success
//	if ctx.maxRetryAttempts==RETRY_FOREVER || ctx.attemptDone<ctx.maxRetryAttempts {
//		*pNextBackOff=uint16(rand.Int31n(int32(ctx.nextJitterMax +1)))
//		ctx.attemptDone+=1
//
//		if ctx.nextJitterMax < ctx.maxBackOffDelay/2 {
//			ctx.nextJitterMax +=ctx.nextJitterMax
//		}else {
//			ctx.nextJitterMax =ctx.maxBackOffDelay
//		}
//	}else {
//		status=Exhausted
//	}
//	return status
//}
//
//func NewNetRetry(ctx *RetryContext, maxRetryAttempts uint32, maxBackOffDelay uint16, nextJitterMax uint16) NetRetry {
//	ctx.nextJitterMax=nextJitterMax
//	ctx.maxBackOffDelay=maxBackOffDelay
//	ctx.maxRetryAttempts=maxRetryAttempts
//	ctx.attemptDone=0
//
//	return NetRetry{}
//}
//
//func BackGroundContext() *RetryContext {
//	return &RetryContext{}
//}
