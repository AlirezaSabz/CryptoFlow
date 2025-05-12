package workers

import (
	"binanceTemporal/activities"
	"binanceTemporal/workflow"

	tempWorkflow "go.temporal.io/sdk/workflow"

	"log/slog"

	"go.temporal.io/sdk/client"
	tempLog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func Start() {
	myLogger := tempLog.NewStructuredLogger(slog.Default())
	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
		Logger:    myLogger,
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()
	w := worker.New(c, "binance_data", worker.Options{})
	// w.RegisterWorkflow(workflow.BinanceWorkflow)
	w.RegisterWorkflowWithOptions(workflow.BinanceWorkflow, tempWorkflow.RegisterOptions{Name: "BinanceWorkflow"})
	w.RegisterActivity(activities.SaveToDB)
	// tempWorkflow.CancelFunc
	err = w.Run(worker.InterruptCh())
	if err != nil {
		panic(err)
	}
}
