package chi

import (
    "log"

    "go.uber.org/zap"
)
type greeting string

func (g greeting) Greet() {
	logger, err := zap.NewProduction()
    if err != nil {
        log.Fatal(err)
    }

    sugar := logger.Sugar()

    sugar.Info("你好宇宙 with ZAG")
	
}

var Greeter greeting
