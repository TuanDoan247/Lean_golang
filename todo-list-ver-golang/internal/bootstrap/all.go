package bootstrap

import (
    "go.uber.org/fx"
    "todo-list-ver-golang/adapter/repository/connection"
    "todo-list-ver-golang/adapter/repository/mysql"
    service "todo-list-ver-golang/adapter/service"
    "todo-list-ver-golang/core/usecase"
    "todo-list-ver-golang/internal/controller"
    "todo-list-ver-golang/internal/router"
    "todo-list-ver-golang/internal/server"
)

func All() fx.Option {
    return fx.Options(
        //Provide database connection
        fx.Provide(connection.NewConnection),

        // Provide port's implements
        fx.Provide(mysql.NewTodoRepositoryAdapter),
        fx.Provide(mysql.NewDatabaseTransaction),

        // Provide use cases
        fx.Provide(usecase.NewDatabaseTransactionUseCase),
        fx.Provide(usecase.NewCreateTodoUseCase),

        // Provide services
        fx.Provide(service.NewTodoService),

        //Provider controllers, these controllers will be used
        // when register router was invoke
        fx.Provide(controller.NewTodoController),

        // Provide gin engine, register core handlers,
        // actuator endpoints and application routers
        server.GinHttpServerOpt(),
        fx.Invoke(router.RegisterGinRouters),

        // Graceful shutdown.
        // OnStop hooks will run in reverse order.
        server.OnStopHttpServerOpt(),
    )
}
