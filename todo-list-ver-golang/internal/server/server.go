package server

import (
    "context"
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "go.uber.org/fx"
    "net/http"
)

func GinHttpServerOpt() fx.Option {
    return fx.Options(
        fx.Provide(NewGinEngine),
        fx.Provide(NewHTTPServer),
        fx.Invoke(OnStartHttpServerHook),
    )
}

func OnStopHttpServerOpt() fx.Option {
    return fx.Invoke(OnStopHttpServerHook)
}

func NewGinEngine() *gin.Engine {
    r := gin.Default()
    return r
}

func NewHTTPServer(engine *gin.Engine) *http.Server {
    return &http.Server{
        Addr:    fmt.Sprintf(":%d", 8080),
        Handler: engine,
    }
}

func OnStartHttpServerHook(lc fx.Lifecycle, engine *gin.Engine) {
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            fmt.Println("Starting HTTP server at port:", 80)
            go engine.Run(":8082")
            return nil
        },
        OnStop: func(ctx context.Context) error {
            fmt.Println("Stopping HTTP Server at port:", 80)
            return errors.New("Stopping HTTP Server at port")
        },
    })
}

func OnStopHttpServerHook(lc fx.Lifecycle, httpServer *http.Server) {
    lc.Append(fx.Hook{
        OnStop: func(ctx context.Context) error {
            return httpServer.Shutdown(ctx)
        },
    })
}
