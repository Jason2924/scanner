package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Jason2924/scanner/backend/models"
	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeLmt time.Duration) gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		newCtxt, cacl := context.WithTimeout(ctxt.Request.Context(), timeLmt)
		defer cacl()
		ctxt.Request = ctxt.Request.WithContext(newCtxt)
		done := make(chan struct{}, 1)
		go func() {
			ctxt.Next()
			done <- struct{}{}
		}()
		select {
		case <-newCtxt.Done():
			if newCtxt.Err() == context.Canceled {
				errCtnt := "Client has canceled the request"
				log.Println(errCtnt)
				// need not sent the response and status code to client
				return
			}
			errCtnt := "The request has timedout"
			resp := models.NewResponse(errCtnt, nil)
			log.Println(errCtnt)
			ctxt.AbortWithStatusJSON(http.StatusRequestTimeout, resp)
		case <-done:
			// success
		}
	}
}
