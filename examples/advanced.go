package main

import (
	"context"
	"fmt"
	"log"
	"time"

	browseruse "github.com/yourusername/browser-use-go-sdk"
)

func advancedExample() {
	client, err := browseruse.NewClient(&browseruse.ClientOptions{
		APIKey: "bu_...",
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Create task with custom options
	fmt.Println("=== Example 1: Task with custom options ===")
	maxSteps := 20
	startURL := "https://news.ycombinator.com"
	llm := browseruse.LLMClaudeSonnet45

	task1, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
		Task:     "Find the top 5 posts and their scores",
		LLM:      &llm,
		StartURL: &startURL,
		MaxSteps: &maxSteps,
		Metadata: map[string]string{
			"user_id": "12345",
			"source":  "api",
		},
	})
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}
	fmt.Printf("Task created: %s\n\n", task1.ID)

	// Example 2: Monitor task progress
	fmt.Println("=== Example 2: Monitor task progress ===")
	task2, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
		Task: "Navigate to GitHub trending page and list the top 3 repositories",
	})
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	// Custom polling with progress updates
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	timeout := time.After(2 * time.Minute)

	for {
		select {
		case <-timeout:
			log.Fatal("Task timed out")
		case <-ticker.C:
			taskView, err := task2.Get(ctx)
			if err != nil {
				log.Fatalf("Failed to get task: %v", err)
			}

			fmt.Printf("Status: %s | Steps: %d\n", taskView.Status, len(taskView.Steps))

			if taskView.Status == browseruse.TaskStatusFinished || taskView.Status == browseruse.TaskStatusStopped {
				if taskView.Output != nil {
					fmt.Printf("\nFinal Output:\n%s\n\n", *taskView.Output)
				}
				goto next_example
			}
		}
	}

next_example:
	// Example 3: List all tasks
	fmt.Println("=== Example 3: List all tasks ===")
	pageSize := 10
	filterBy := browseruse.TaskStatusFinished

	listResponse, err := client.Tasks.ListTasks(ctx, &browseruse.ListTasksOptions{
		PageSize: &pageSize,
		FilterBy: &filterBy,
	})
	if err != nil {
		log.Fatalf("Failed to list tasks: %v", err)
	}

	fmt.Printf("Found %d finished tasks (showing %d):\n", listResponse.TotalItems, len(listResponse.Items))
	for i, task := range listResponse.Items {
		fmt.Printf("%d. %s - %s\n", i+1, task.ID, task.Task)
	}
	fmt.Println()

	// Example 4: Stop a running task
	fmt.Println("=== Example 4: Stop a task ===")
	task3, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
		Task: "Browse multiple pages on Wikipedia",
	})
	if err != nil {
		log.Fatalf("Failed to create task: %v", err)
	}

	// Wait a bit, then stop it
	time.Sleep(5 * time.Second)
	stoppedTask, err := task3.Stop(ctx)
	if err != nil {
		log.Fatalf("Failed to stop task: %v", err)
	}
	fmt.Printf("Task stopped. Status: %s\n\n", stoppedTask.Status)

	// Example 5: Get task logs
	fmt.Println("=== Example 5: Get task logs ===")
	logs, err := task1.GetLogs(ctx)
	if err != nil {
		log.Fatalf("Failed to get logs: %v", err)
	}
	fmt.Printf("Logs download URL: %s\n", logs.DownloadURL)
}
