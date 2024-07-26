package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/client"
)

// Initialize Temporal client
var temporalClient client.Client

func init() {
	var err error
	temporalClient, err = client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Failed to create Temporal client: %v", err)
	}
}

// Trigger Email Workflow
// func TriggerEmailWorkflow(email, subject, body string) error {
// 	params := map[string]interface{}{
// 		"To":      email,
// 		"Subject": subject,
// 		"Body":    body,
// 	}
// 	workflowID := fmt.Sprintf("workflow-%d", time.Now().UnixNano())
// 	_, err := temporalClient.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
// 		ID:        workflowID,
// 		TaskQueue: "email-task-queue",
// 	}, "SendEmailWorkflow", params)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

type EmailParams struct {
	To      string
	Subject string
	Body    string
}

func TriggerEmailWorkflow(email, subject, body string) error {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return err
	}
	defer c.Close()
	workflowID := fmt.Sprintf("workflow-%d", time.Now().UnixNano())
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "email-task-queue",
	}

	params := EmailParams{
		To:      email,
		Subject: subject,
		Body:    body,
	}

	_, err = temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, "SendEmailWorkflow", params)
	if err != nil {
		return err
	}

	return nil
}
