package client

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/adshao/go-binance/v2"
	tempLog "go.temporal.io/sdk/log"

	"go.temporal.io/sdk/client"
)

type MyClient struct {
	client client.Client
}

func New() (*MyClient, error) {
	// Create a new Temporal client
	// with a custom logger
	myLogger := tempLog.NewStructuredLogger(slog.Default())
	c, err := client.Dial(client.Options{
		Logger:   myLogger,
		HostPort: "localhost:7233",
	})
	if err != nil {
		return nil, err
	}
	// defer c.Close()
	clnt := &MyClient{
		client: c,
	}

	return clnt, nil
}

func (c *MyClient) GetBinanceData() {
	workflowOptions := client.StartWorkflowOptions{
		ID:        "binance-workflow-id",
		TaskQueue: "binance_data",
	}
	we, err := c.client.ExecuteWorkflow(
		context.Background(),
		workflowOptions,
		"BinanceWorkflow",
	)
	if err != nil {
		log.Fatalln("Error executing Workflow:", err)
	}

	log.Println("Workflow started: \n ", "WorkflowID:", we.GetID(), "\nRunID:", we.GetRunID())

	errHandler := func(err error) {
		log.Println("WebSocket error:", err)
	}
	wsKlineHandler := getWsKlineHandler(c.client)
	doneC, stopC, err := binance.WsKlineServe("BTCUSDT", "1m", wsKlineHandler, errHandler)
	if err != nil {
		log.Fatalln("WebSocket error:", err)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-doneC:
		log.Println("WebSocket done")
		err = c.client.TerminateWorkflow(context.TODO(), "binance-workflow-id", we.GetRunID(), "Workflow terminated")

		if err != nil {
			log.Println("Error terminating workflow:", err)
		} else {
			c.client.Close()
			log.Println("workflow stopped")
		}
		os.Exit(0)
		return
	case <-sig:
		log.Println("❌ Ctrl+C received, stopping binance websocket...")
		//close websocket
		close(stopC)

		//cancle workflow
		// err := c.client.CancelWorkflow(context.TODO(), "binance-workflow-id", we.GetRunID())
		// if err != nil {
		// 	log.Println("Error cancelling workflow:", err)
		// }

		//terminate workflow
		err = c.client.TerminateWorkflow(context.TODO(), "binance-workflow-id", we.GetRunID(), "Workflow terminated")

		if err != nil {
			log.Println("Error terminating workflow:", err)
		} else {
			c.client.Close()
			log.Println("workflow stopped")
		}
		os.Exit(0)
		return
	}
}

func getWsKlineHandler(c client.Client) func(event *binance.WsKlineEvent) {
	return func(event *binance.WsKlineEvent) {
		if event.Kline.IsFinal {
			err := c.SignalWorkflow(context.TODO(), "binance-workflow-id", "", "binance-signal", *event)
			if err != nil {
				log.Println("Error signaling workflow:", err)
			}
		}
	}
}
