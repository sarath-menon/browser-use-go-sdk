# Browser Use Go SDK

Official Go SDK for the [Browser Use](https://browser-use.com) API - Control web browsers with AI.

## Installation

```bash
go get github.com/sarath-menon/browser-use-go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    browseruse "github.com/sarath-menon/browser-use-go-sdk"
)

func main() {
    // Create client (API key from env: BROWSER_USE_API_KEY)
    client, err := browseruse.NewClient(&browseruse.ClientOptions{
        APIKey: "bu_...", // Optional if BROWSER_USE_API_KEY is set
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Create a task
    task, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
        Task: "Search for the top 10 Hacker News posts and return the title and url.",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Wait for completion
    result, err := task.Complete(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(*result.Output)
}
```

## Features

- ✅ Simple and idiomatic Go API
- ✅ Context support for timeouts and cancellation
- ✅ Automatic task completion polling
- ✅ Type-safe request/response models
- ✅ Comprehensive error handling
- ✅ Support for all Browser Use API endpoints

## API Reference

### Client Initialization

```go
// Using environment variable (BROWSER_USE_API_KEY)
client, err := browseruse.NewClient(nil)

// With explicit API key
client, err := browseruse.NewClient(&browseruse.ClientOptions{
    APIKey: "bu_your_api_key",
})

// With custom HTTP client and timeout
client, err := browseruse.NewClient(&browseruse.ClientOptions{
    APIKey:     "bu_your_api_key",
    HTTPClient: customHTTPClient,
    Timeout:    60 * time.Second,
})
```

### Creating Tasks

```go
// Simple task
task, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
    Task: "Find the weather in San Francisco",
})

// Task with options
llm := browseruse.LLMClaudeSonnet45
maxSteps := 20
startURL := "https://google.com"

task, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
    Task:     "Search for AI news",
    LLM:      &llm,
    StartURL: &startURL,
    MaxSteps: &maxSteps,
    Metadata: map[string]string{
        "user_id": "12345",
    },
})
```

### Task Operations

```go
// Wait for task to complete (with auto-polling)
result, err := task.Complete(ctx, nil)

// Custom completion options
result, err := task.Complete(ctx, &browseruse.CompleteOptions{
    PollInterval: 3 * time.Second,
    Timeout:      10 * time.Minute,
})

// Get current task status
taskView, err := task.Get(ctx)

// Stop a task
taskView, err := task.Stop(ctx)

// Stop task and session
taskView, err := task.StopWithSession(ctx)

// Get task logs
logs, err := task.GetLogs(ctx)
fmt.Println(logs.DownloadURL)
```

### Listing Tasks

```go
// List all tasks
tasks, err := client.Tasks.ListTasks(ctx, nil)

// With pagination and filters
pageSize := 20
pageNumber := 1
filterBy := browseruse.TaskStatusFinished

tasks, err := client.Tasks.ListTasks(ctx, &browseruse.ListTasksOptions{
    PageSize:   &pageSize,
    PageNumber: &pageNumber,
    FilterBy:   &filterBy,
})

for _, task := range tasks.Items {
    fmt.Printf("%s: %s\n", task.ID, task.Task)
}
```

### Direct Task Operations

```go
// Create and get task info separately
task, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
    Task: "Navigate to example.com",
})

// Get task by ID
taskView, err := client.Tasks.GetTask(ctx, task.ID)

// Stop task by ID
taskView, err := client.Tasks.StopTask(ctx, task.ID)

// Get logs by ID
logs, err := client.Tasks.GetTaskLogs(ctx, task.ID)
```

## Supported LLM Models

```go
browseruse.LLMBrowserUse           // "browser-use-llm"
browseruse.LLMGPT41                // "gpt-4.1"
browseruse.LLMGPT41Mini            // "gpt-4.1-mini"
browseruse.LLMO4Mini               // "o4-mini"
browseruse.LLMO3                   // "o3"
browseruse.LLMGemini25Flash        // "gemini-2.5-flash"
browseruse.LLMGemini25Pro          // "gemini-2.5-pro"
browseruse.LLMClaudeSonnet45       // "claude-sonnet-4-5-20250929"
browseruse.LLMClaudeOpus45         // "claude-opus-4-5-20251101"
browseruse.LLMGPT4o                // "gpt-4o"
browseruse.LLMGPT4oMini            // "gpt-4o-mini"
// ... and more
```

## Error Handling

```go
task, err := client.Tasks.CreateTask(ctx, &browseruse.CreateTaskRequest{
    Task: "Some task",
})
if err != nil {
    if browseruse.IsNotFound(err) {
        // Handle 404 error
    } else if browseruse.IsRateLimitError(err) {
        // Handle rate limit (429)
    } else if browseruse.IsValidationError(err) {
        // Handle validation error (422)
    } else {
        // Handle other errors
    }
}
```

## Examples

See the [examples](./examples) directory for more usage examples:

- [basic.go](./examples/main.go) - Simple task creation and completion
- [advanced.go](./examples/advanced.go) - Advanced features, monitoring, and options

## API Documentation

For complete API documentation, visit: https://docs.cloud.browser-use.com/api-reference

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
