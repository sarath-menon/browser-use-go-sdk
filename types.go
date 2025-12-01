package browseruse

import (
	"time"
)

// LLMModel represents supported LLM models
type LLMModel string

const (
	LLMBrowserUse           LLMModel = "browser-use-llm"
	LLMGPT41                LLMModel = "gpt-4.1"
	LLMGPT41Mini            LLMModel = "gpt-4.1-mini"
	LLMO4Mini               LLMModel = "o4-mini"
	LLMO3                   LLMModel = "o3"
	LLMGemini25Flash        LLMModel = "gemini-2.5-flash"
	LLMGemini25Pro          LLMModel = "gemini-2.5-pro"
	LLMGemini3ProPreview    LLMModel = "gemini-3-pro-preview"
	LLMGeminiFlashLatest    LLMModel = "gemini-flash-latest"
	LLMGeminiFlashLiteLatest LLMModel = "gemini-flash-lite-latest"
	LLMClaudeSonnet4        LLMModel = "claude-sonnet-4-20250514"
	LLMClaudeSonnet45       LLMModel = "claude-sonnet-4-5-20250929"
	LLMClaudeOpus45         LLMModel = "claude-opus-4-5-20251101"
	LLMGPT4o                LLMModel = "gpt-4o"
	LLMGPT4oMini            LLMModel = "gpt-4o-mini"
	LLMLlama4Maverick       LLMModel = "llama-4-maverick-17b-128e-instruct"
	LLMClaude37Sonnet       LLMModel = "claude-3-7-sonnet-20250219"
)

// TaskStatus represents the current state of a task
type TaskStatus string

const (
	TaskStatusCreated  TaskStatus = "created"
	TaskStatusStarted  TaskStatus = "started"
	TaskStatusFinished TaskStatus = "finished"
	TaskStatusStopped  TaskStatus = "stopped"
)

// TaskUpdateAction represents actions that can be performed on a task
type TaskUpdateAction string

const (
	TaskActionStop               TaskUpdateAction = "stop"
	TaskActionStopTaskAndSession TaskUpdateAction = "stop_task_and_session"
)

// Vision can be either a boolean or the string "auto"
type Vision struct {
	Auto  bool
	Value *bool
}

// CreateTaskRequest represents a task creation request
type CreateTaskRequest struct {
	Task                  string            `json:"task"`
	LLM                   *LLMModel         `json:"llm,omitempty"`
	StartURL              *string           `json:"startUrl,omitempty"`
	MaxSteps              *int              `json:"maxSteps,omitempty"`
	StructuredOutput      *string           `json:"structuredOutput,omitempty"`
	SessionID             *string           `json:"sessionId,omitempty"`
	Metadata              map[string]string `json:"metadata,omitempty"`
	Secrets               map[string]string `json:"secrets,omitempty"`
	AllowedDomains        []string          `json:"allowedDomains,omitempty"`
	OpVaultID             *string           `json:"opVaultId,omitempty"`
	HighlightElements     *bool             `json:"highlightElements,omitempty"`
	FlashMode             *bool             `json:"flashMode,omitempty"`
	Thinking              *bool             `json:"thinking,omitempty"`
	Vision                *Vision           `json:"vision,omitempty"`
	SystemPromptExtension *string           `json:"systemPromptExtension,omitempty"`
	Judge                 *bool             `json:"judge,omitempty"`
	JudgeGroundTruth      *string           `json:"judgeGroundTruth,omitempty"`
	JudgeLLM              *LLMModel         `json:"judgeLlm,omitempty"`
}

// TaskCreatedResponse is returned when a task is created
type TaskCreatedResponse struct {
	ID        string `json:"id"`
	SessionID string `json:"sessionId"`
}

// TaskStepView represents a single step in task execution
type TaskStepView struct {
	Number                 int      `json:"number"`
	Memory                 string   `json:"memory"`
	EvaluationPreviousGoal string   `json:"evaluationPreviousGoal"`
	NextGoal               string   `json:"nextGoal"`
	URL                    string   `json:"url"`
	ScreenshotURL          *string  `json:"screenshotUrl"`
	Actions                []string `json:"actions"`
}

// FileView represents a file output from a task
type FileView struct {
	ID       string `json:"id"`
	FileName string `json:"fileName"`
}

// TaskView represents detailed task information
type TaskView struct {
	ID                string         `json:"id"`
	SessionID         string         `json:"sessionId"`
	LLM               string         `json:"llm"`
	Task              string         `json:"task"`
	Status            TaskStatus     `json:"status"`
	CreatedAt         time.Time      `json:"createdAt"`
	StartedAt         *time.Time     `json:"startedAt"`
	FinishedAt        *time.Time     `json:"finishedAt"`
	Metadata          map[string]any `json:"metadata"`
	Steps             []TaskStepView `json:"steps"`
	Output            *string        `json:"output"`
	OutputFiles       []FileView     `json:"outputFiles"`
	BrowserUseVersion *string        `json:"browserUseVersion"`
	IsSuccess         *bool          `json:"isSuccess"`
	Judgement         *string        `json:"judgement"`
	JudgeVerdict      *bool          `json:"judgeVerdict"`
}

// UpdateTaskRequest represents a task update request
type UpdateTaskRequest struct {
	Action TaskUpdateAction `json:"action"`
}

// TaskLogFileResponse contains the download URL for task logs
type TaskLogFileResponse struct {
	DownloadURL string `json:"downloadUrl"`
}

// TaskItemView represents a task in a list response
type TaskItemView struct {
	ID                string         `json:"id"`
	SessionID         string         `json:"sessionId"`
	LLM               string         `json:"llm"`
	Task              string         `json:"task"`
	Status            TaskStatus     `json:"status"`
	CreatedAt         time.Time      `json:"createdAt"`
	StartedAt         *time.Time     `json:"startedAt"`
	FinishedAt        *time.Time     `json:"finishedAt"`
	Metadata          map[string]any `json:"metadata"`
	Output            *string        `json:"output"`
	BrowserUseVersion *string        `json:"browserUseVersion"`
	IsSuccess         *bool          `json:"isSuccess"`
	Judgement         *string        `json:"judgement"`
	JudgeVerdict      *bool          `json:"judgeVerdict"`
}

// TaskListResponse represents a paginated list of tasks
type TaskListResponse struct {
	Items      []TaskItemView `json:"items"`
	TotalItems int            `json:"totalItems"`
	PageNumber int            `json:"pageNumber"`
	PageSize   int            `json:"pageSize"`
}

// ListTasksOptions represents options for listing tasks
type ListTasksOptions struct {
	PageSize   *int        `json:"pageSize,omitempty"`
	PageNumber *int        `json:"pageNumber,omitempty"`
	SessionID  *string     `json:"sessionId,omitempty"`
	FilterBy   *TaskStatus `json:"filterBy,omitempty"`
	After      *time.Time  `json:"after,omitempty"`
	Before     *time.Time  `json:"before,omitempty"`
}
