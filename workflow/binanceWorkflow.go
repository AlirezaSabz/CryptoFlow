package workflow

import (
	"binanceTemporal/activities"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"go.temporal.io/sdk/workflow"
)

func BinanceWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Binance Workflow started successfully!")
	signalChan := workflow.GetSignalChannel(ctx, "binance-signal")

	ao := workflow.ActivityOptions{
		TaskQueue:           "binance_data",
		StartToCloseTimeout: 5 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	for {
		var event binance.WsKlineEvent
		signalChan.Receive(ctx, &event)
		logger.Info(fmt.Sprintf("%+v  \n \n ", event))
		workflow.ExecuteActivity(ctx, activities.SaveToDB, event)
	}
}
