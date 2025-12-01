package browseruse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// TasksService handles task-related API operations
type TasksService struct {
	client *Client
}

// Task represents a created task with methods to interact with it
type Task struct {
	ID        string
	SessionID string
	service   *TasksService
}

// CreateTask creates a new task and returns a Task handle
func (s *TasksService) CreateTask(ctx context.Context, req *CreateTaskRequest) (*Task, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	var resp TaskCreatedResponse
	if err := s.client.doRequest(ctx, "POST", "/tasks", bytes.NewReader(body), &resp); err != nil {
		return nil, err
	}

	return &Task{
		ID:        resp.ID,
		SessionID: resp.SessionID,
		service:   s,
	}, nil
}

// GetTask retrieves detailed information about a task
func (s *TasksService) GetTask(ctx context.Context, taskID string) (*TaskView, error) {
	var task TaskView
	if err := s.client.doRequest(ctx, "GET", "/tasks/"+taskID, nil, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask updates a task with the specified action
func (s *TasksService) UpdateTask(ctx context.Context, taskID string, req *UpdateTaskRequest) (*TaskView, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	var task TaskView
	if err := s.client.doRequest(ctx, "PATCH", "/tasks/"+taskID, bytes.NewReader(body), &task); err != nil {
		return nil, err
	}
	return &task, nil
}

// StopTask stops a running task
func (s *TasksService) StopTask(ctx context.Context, taskID string) (*TaskView, error) {
	return s.UpdateTask(ctx, taskID, &UpdateTaskRequest{
		Action: TaskActionStop,
	})
}

// StopTaskAndSession stops a task and its associated session
func (s *TasksService) StopTaskAndSession(ctx context.Context, taskID string) (*TaskView, error) {
	return s.UpdateTask(ctx, taskID, &UpdateTaskRequest{
		Action: TaskActionStopTaskAndSession,
	})
}

// GetTaskLogs retrieves the download URL for task logs
func (s *TasksService) GetTaskLogs(ctx context.Context, taskID string) (*TaskLogFileResponse, error) {
	var logs TaskLogFileResponse
	if err := s.client.doRequest(ctx, "GET", "/tasks/"+taskID+"/logs", nil, &logs); err != nil {
		return nil, err
	}
	return &logs, nil
}

// ListTasks retrieves a paginated list of tasks
func (s *TasksService) ListTasks(ctx context.Context, opts *ListTasksOptions) (*TaskListResponse, error) {
	path := "/tasks"

	if opts != nil {
		params := url.Values{}
		if opts.PageSize != nil {
			params.Add("pageSize", fmt.Sprintf("%d", *opts.PageSize))
		}
		if opts.PageNumber != nil {
			params.Add("pageNumber", fmt.Sprintf("%d", *opts.PageNumber))
		}
		if opts.SessionID != nil {
			params.Add("sessionId", *opts.SessionID)
		}
		if opts.FilterBy != nil {
			params.Add("filterBy", string(*opts.FilterBy))
		}
		if opts.After != nil {
			params.Add("after", opts.After.Format(time.RFC3339))
		}
		if opts.Before != nil {
			params.Add("before", opts.Before.Format(time.RFC3339))
		}
		if len(params) > 0 {
			path += "?" + params.Encode()
		}
	}

	var list TaskListResponse
	if err := s.client.doRequest(ctx, "GET", path, nil, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get retrieves the current state of the task
func (t *Task) Get(ctx context.Context) (*TaskView, error) {
	return t.service.GetTask(ctx, t.ID)
}

// Stop stops the task execution
func (t *Task) Stop(ctx context.Context) (*TaskView, error) {
	return t.service.StopTask(ctx, t.ID)
}

// StopWithSession stops the task and its associated session
func (t *Task) StopWithSession(ctx context.Context) (*TaskView, error) {
	return t.service.StopTaskAndSession(ctx, t.ID)
}

// GetLogs retrieves the download URL for task logs
func (t *Task) GetLogs(ctx context.Context) (*TaskLogFileResponse, error) {
	return t.service.GetTaskLogs(ctx, t.ID)
}

// CompleteOptions configures the behavior of the Complete method
type CompleteOptions struct {
	// PollInterval is the time to wait between status checks
	PollInterval time.Duration
	// Timeout is the maximum time to wait for task completion
	Timeout time.Duration
}

// Complete waits for the task to finish and returns the final result
// This is similar to the TypeScript SDK's task.complete() method
func (t *Task) Complete(ctx context.Context, opts *CompleteOptions) (*TaskView, error) {
	if opts == nil {
		opts = &CompleteOptions{
			PollInterval: 2 * time.Second,
			Timeout:      5 * time.Minute,
		}
	}

	if opts.PollInterval == 0 {
		opts.PollInterval = 2 * time.Second
	}
	if opts.Timeout == 0 {
		opts.Timeout = 5 * time.Minute
	}

	// Create a context with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	ticker := time.NewTicker(opts.PollInterval)
	defer ticker.Stop()

	// Check immediately first
	task, err := t.Get(timeoutCtx)
	if err != nil {
		return nil, err
	}

	if task.Status == TaskStatusFinished || task.Status == TaskStatusStopped {
		return task, nil
	}

	// Poll for completion
	for {
		select {
		case <-timeoutCtx.Done():
			return nil, fmt.Errorf("task did not complete within timeout: %w", timeoutCtx.Err())
		case <-ticker.C:
			task, err := t.Get(timeoutCtx)
			if err != nil {
				return nil, err
			}

			if task.Status == TaskStatusFinished || task.Status == TaskStatusStopped {
				return task, nil
			}
		}
	}
}
