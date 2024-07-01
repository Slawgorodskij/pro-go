package pipeline

import "net/http"

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (mws *ComponentContext) Error(err error) {
	mws.error = err
}

func (mws *ComponentContext) GetError() error {
	return mws.error
}

type MiddlewareComponent interface {
	Init()
	ProcessRequest(ctx *ComponentContext, next func(*ComponentContext))
}
type ServicesMiddlwareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}
