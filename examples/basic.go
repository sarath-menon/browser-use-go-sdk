package main

import (
	"context"
	"fmt"
	"log"

	browseruse "github.com/yourusername/browser-use-go-sdk"
)

func main() {
	// Create a new Browser Use client
	// API key can be provided here or via BROWSER_USE_API_KEY environment variable
	client, err := browseruse.NewClient(&browseruse.ClientOptions{
		APIKey: "bu_...", // Optional if BROWSER_USE_API_KEY is set
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a new task
	task, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
		Task: "Search for the top 10 Hacker News posts and return the title and url.",
	})
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	fmt.Printf("Task created with ID: %s\n", task.ID)

	// Wait for the task to complete (similar to TypeScript's task.complete())
	result, err := task.Complete(ctx, nil)
	if err != nil {
		log.Fatalf("Task failed: %v", err)
	}

	// Print the output
	if result.Output != nil {
		fmt.Printf("\nTask Output:\n%s\n", *result.Output)
	} else {
		fmt.Println("\nNo output returned")
	}

	fmt.Printf("\nTask Status: %s\n", result.Status)
	fmt.Printf("Steps taken: %d\n", len(result.Steps))
}
