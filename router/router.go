package router

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	_registerHandlers []func(engine *gin.Engine) = nil
	_registerHandlerMutex sync.Mutex
)

func addRegisterHandler(handler func(engine *gin.Engine)){
	_registerHandlerMutex.Lock()
	_registerHandlerMutex.Unlock()

	if _registerHandlers == nil{
		_registerHandlers = make([]func(engine *gin.Engine), 0)
	}
	_registerHandlers = append(_registerHandlers, handler)
}

func Register(engin *gin.Engine){
	_registerHandlerMutex.Lock()
	_registerHandlerMutex.Unlock()

	for i := range _registerHandlers{
		_registerHandlers[i](engin)
	}
}
