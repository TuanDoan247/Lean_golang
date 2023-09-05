package main

import (
    "go.uber.org/fx"
    "todo-list-ver-golang/internal/bootstrap"
)

func main() {
    fx.New(bootstrap.All()).Run()
}
