package koa

import (
	"fmt"
	"io"
	"net/http"
)

// Next assaf
type Next func() interface{}

// Middleware asdasf
type Middleware func(*http.Request, Next) interface{}

// Application asf
type Application struct {
	// http.HandleFunc
	middlewares []Middleware
}

// Use asd
func (app *Application) Use(middleware Middleware) {
	app.middlewares = append(app.middlewares, middleware)
}

// Callback asdasf
func (app *Application) Callback(w http.ResponseWriter, r *http.Request) {
	// app.middlewares
	current, count := 0, len(app.middlewares)
	var next Next

	// 这里实现好像有一些问题
	// 应该使用compose，将数组函数变为嵌套函数。
	next = func() interface{} {
		if current < count {
			middleware := app.middlewares[current]
			current = current + 1
			return middleware(r, next)
		}
		return nil
	}
	switch t := next().(type) {
	case io.Reader:
		{
			p := make([]byte, 10)
			for {
				n, err := t.Read(p)
				if err != nil {
					break
				}
				w.Write(p[:n])
			}
		}
	default:
		fmt.Fprint(w, t)
	}

	// app.middlewares[current](w, r)
}

// Listen asdsaf
func (app *Application) Listen(addr string) error {
	http.HandleFunc("/", app.Callback)
	return http.ListenAndServe(addr, nil)
}