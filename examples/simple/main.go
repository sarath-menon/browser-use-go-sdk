package main

import (
	"context"
	"fmt"
	"log"

	browseruse "github.com/yourusername/browser-use-go-sdk"
)

func main() {
	// Create a new Browser Use client
	// Reads API key from BROWSER_USE_API_KEY environment variable
	client, err := browseruse.NewClient(nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}

	ctx := context.Background()

	// Create a new task
	task, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
		Task: "Search for the top 10 Hacker News posts and return the title and url.",
	})
	if err != nil {
		log.Fatalf("Failed to create task: %v\n", err)
	}

	fmt.Printf("Task created with ID: %s\n", task.ID)
	fmt.Println("Waiting for task to complete...")

	// Wait for the task to complete (similar to TypeScript's task.complete())
	result, err := task.Complete(ctx, nil)
	if err != nil {
		log.Fatalf("Task failed: %v\n", err)
	}

	// Print the output
	fmt.Println("\n" + "================================================================================")
	fmt.Println("TASK COMPLETED")
	fmt.Println("================================================================================")

	if result.Output != nil {
		fmt.Printf("\nOutput:\n%s\n", *result.Output)
	} else {
		fmt.Println("\nNo output returned")
	}

	fmt.Printf("\nTask Details:\n")
	fmt.Printf("  Status: %s\n", result.Status)
	fmt.Printf("  Steps taken: %d\n", len(result.Steps))
	fmt.Printf("  Session ID: %s\n", result.SessionID)
	fmt.Printf("  LLM used: %s\n", result.LLM)

	if result.IsSuccess != nil {
		fmt.Printf("  Success: %v\n", *result.IsSuccess)
	}
}
