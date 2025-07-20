package temporal

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/client"
)

func StartAggregatorWorkflow(input []Asset) {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "agreggator-workflow-" + time.Now().Format("20060102150405"),
		TaskQueue: "AGGREGATOR_TASK_QUEUE",
	}

	_, err = c.ExecuteWorkflow(context.Background(), workflowOptions, "AggregatorWorkflow", input)

	if err != nil {
		log.Fatalf("Failed to start workflow: %v", err)
		return
	}

	log.Printf("Started workflow with ID: %s", workflowOptions.ID)
}
