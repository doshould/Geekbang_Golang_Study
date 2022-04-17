package hystrixDemo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Wrapper(
	size int,
	reqThreshold int,
	failThreshold float64,
	duration time.Duration,
) gin.HandlerFunc {
	r := NewRollingWindow(size, reqThreshold, failThreshold, duration)
	r.Start()
	r.Monitor()
	r.ShowStatus()
	return func(ctx *gin.Context) {
		if r.Broken() {
			ctx.String(http.StatusInternalServerError, "reject <---- hystrixDemo")
			ctx.Abort()
			return
		}
		ctx.Next()
		if ctx.Writer.Status() != http.StatusOK {
			r.RecordReqRequest(false)
		} else {
			r.RecordReqRequest(true)
		}
	}
}
