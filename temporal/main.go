package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type EmailParams struct {
	To      string
	Subject string
	Body    string
}

func SendEmailWorkflow(ctx workflow.Context, params EmailParams) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, SendEmailActivity, params).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailActivity(ctx context.Context, params EmailParams) error {
	auth := smtp.PlainAuth(
		"",
		"prityr@moneymul.com",
		"rweh lphl hexk egbq",
		"smtp.gmail.com",
	)

	to := []string{params.To}
	msg := []byte("To: " + params.To + "\r\n" +
		"Subject: " + params.Subject + "\r\n" +
		"\r\n" +
		params.Body + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "prityr@moneymul.com", to, msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	fmt.Println("Email sent successfully")
	return nil
}

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	w := worker.New(c, "email-task-queue", worker.Options{})
	w.RegisterWorkflow(SendEmailWorkflow)
	w.RegisterActivity(SendEmailActivity)

	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			panic(err)
		}
	}()

	select {}
}
