# Create Task

POST https://api.browser-use.com/api/v2/tasks
Content-Type: application/json

You can either:
1. Start a new task (auto creates a new simple session)
2. Start a new task in an existing session (you can create a custom session before starting the task and reuse it for follow-up tasks)

Reference: https://docs.cloud.browser-use.com/api-reference/v-2-api-current/tasks/create-task-tasks-post

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Create Task
  version: endpoint_tasks.create_task_tasks_post
paths:
  /tasks:
    post:
      operationId: create-task-tasks-post
      summary: Create Task
      description: >-
        You can either:

        1. Start a new task (auto creates a new simple session)

        2. Start a new task in an existing session (you can create a custom
        session before starting the task and reuse it for follow-up tasks)
      tags:
        - - subpackage_tasks
      parameters:
        - name: X-Browser-Use-API-Key
          in: header
          required: true
          schema:
            type: string
      responses:
        '202':
          description: Successful Response
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TaskCreatedResponse'
        '400':
          description: Session is stopped or has running task
          content: {}
        '404':
          description: Session not found
          content: {}
        '422':
          description: Request validation failed
          content: {}
        '429':
          description: Too many concurrent active sessions
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateTaskRequest'
components:
  schemas:
    SupportedLLMs:
      type: string
      enum:
        - value: browser-use-llm
        - value: gpt-4.1
        - value: gpt-4.1-mini
        - value: o4-mini
        - value: o3
        - value: gemini-2.5-flash
        - value: gemini-2.5-pro
        - value: gemini-3-pro-preview
        - value: gemini-flash-latest
        - value: gemini-flash-lite-latest
        - value: claude-sonnet-4-20250514
        - value: claude-sonnet-4-5-20250929
        - value: claude-opus-4-5-20251101
        - value: gpt-4o
        - value: gpt-4o-mini
        - value: llama-4-maverick-17b-128e-instruct
        - value: claude-3-7-sonnet-20250219
    CreateTaskRequestVision:
      oneOf:
        - type: boolean
        - type: string
          enum:
            - type: stringLiteral
              value: auto
    CreateTaskRequest:
      type: object
      properties:
        task:
          type: string
        llm:
          $ref: '#/components/schemas/SupportedLLMs'
        startUrl:
          type:
            - string
            - 'null'
        maxSteps:
          type: integer
        structuredOutput:
          type:
            - string
            - 'null'
        sessionId:
          type:
            - string
            - 'null'
          format: uuid
        metadata:
          type:
            - object
            - 'null'
          additionalProperties:
            type: string
        secrets:
          type:
            - object
            - 'null'
          additionalProperties:
            type: string
        allowedDomains:
          type:
            - array
            - 'null'
          items:
            type: string
        opVaultId:
          type:
            - string
            - 'null'
        highlightElements:
          type: boolean
        flashMode:
          type: boolean
        thinking:
          type: boolean
        vision:
          $ref: '#/components/schemas/CreateTaskRequestVision'
        systemPromptExtension:
          type: string
        judge:
          type: boolean
        judgeGroundTruth:
          type:
            - string
            - 'null'
        judgeLlm:
          oneOf:
            - $ref: '#/components/schemas/SupportedLLMs'
            - type: 'null'
      required:
        - task
    TaskCreatedResponse:
      type: object
      properties:
        id:
          type: string
          format: uuid
        sessionId:
          type: string
          format: uuid
      required:
        - id
        - sessionId

```

## SDK Code Examples

```python
import requests

url = "https://api.browser-use.com/api/v2/tasks"

payload = { "task": "string" }
headers = {
    "X-Browser-Use-API-Key": "<apiKey>",
    "Content-Type": "application/json"
}

response = requests.post(url, json=payload, headers=headers)

print(response.json())
```

```javascript
const url = 'https://api.browser-use.com/api/v2/tasks';
const options = {
  method: 'POST',
  headers: {'X-Browser-Use-API-Key': '<apiKey>', 'Content-Type': 'application/json'},
  body: '{"task":"string"}'
};

try {
  const response = await fetch(url, options);
  const data = await response.json();
  console.log(data);
} catch (error) {
  console.error(error);
}
```

```go
package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {

	url := "https://api.browser-use.com/api/v2/tasks"

	payload := strings.NewReader("{\n  \"task\": \"string\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("X-Browser-Use-API-Key", "<apiKey>")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
```

```ruby
require 'uri'
require 'net/http'

url = URI("https://api.browser-use.com/api/v2/tasks")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Post.new(url)
request["X-Browser-Use-API-Key"] = '<apiKey>'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"task\": \"string\"\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.post("https://api.browser-use.com/api/v2/tasks")
  .header("X-Browser-Use-API-Key", "<apiKey>")
  .header("Content-Type", "application/json")
  .body("{\n  \"task\": \"string\"\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('POST', 'https://api.browser-use.com/api/v2/tasks', [
  'body' => '{
  "task": "string"
}',
  'headers' => [
    'Content-Type' => 'application/json',
    'X-Browser-Use-API-Key' => '<apiKey>',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.browser-use.com/api/v2/tasks");
var request = new RestRequest(Method.POST);
request.AddHeader("X-Browser-Use-API-Key", "<apiKey>");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"task\": \"string\"\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "X-Browser-Use-API-Key": "<apiKey>",
  "Content-Type": "application/json"
]
let parameters = ["task": "string"] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.browser-use.com/api/v2/tasks")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "POST"
request.allHTTPHeaderFields = headers
request.httpBody = postData as Data

let session = URLSession.shared
let dataTask = session.dataTask(with: request as URLRequest, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    print(error as Any)
  } else {
    let httpResponse = response as? HTTPURLResponse
    print(httpResponse)
  }
})

dataTask.resume()
```


# Get Task

GET https://api.browser-use.com/api/v2/tasks/{task_id}

Get detailed task information including status, progress, steps, and file outputs.

Reference: https://docs.cloud.browser-use.com/api-reference/v-2-api-current/tasks/get-task-tasks-task-id-get

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Task
  version: endpoint_tasks.get_task_tasks__task_id__get
paths:
  /tasks/{task_id}:
    get:
      operationId: get-task-tasks-task-id-get
      summary: Get Task
      description: >-
        Get detailed task information including status, progress, steps, and
        file outputs.
      tags:
        - - subpackage_tasks
      parameters:
        - name: task_id
          in: path
          required: true
          schema:
            type: string
            format: uuid
        - name: X-Browser-Use-API-Key
          in: header
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Successful Response
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TaskView'
        '404':
          description: Task not found
          content: {}
        '422':
          description: Validation Error
          content: {}
components:
  schemas:
    TaskStatus:
      type: string
      enum:
        - value: created
        - value: started
        - value: finished
        - value: stopped
    TaskStepView:
      type: object
      properties:
        number:
          type: integer
        memory:
          type: string
        evaluationPreviousGoal:
          type: string
        nextGoal:
          type: string
        url:
          type: string
        screenshotUrl:
          type:
            - string
            - 'null'
        actions:
          type: array
          items:
            type: string
      required:
        - number
        - memory
        - evaluationPreviousGoal
        - nextGoal
        - url
        - actions
    FileView:
      type: object
      properties:
        id:
          type: string
          format: uuid
        fileName:
          type: string
      required:
        - id
        - fileName
    TaskView:
      type: object
      properties:
        id:
          type: string
          format: uuid
        sessionId:
          type: string
          format: uuid
        llm:
          type: string
        task:
          type: string
        status:
          $ref: '#/components/schemas/TaskStatus'
        createdAt:
          type: string
          format: date-time
        startedAt:
          type:
            - string
            - 'null'
          format: date-time
        finishedAt:
          type:
            - string
            - 'null'
          format: date-time
        metadata:
          type: object
          additionalProperties:
            description: Any type
        steps:
          type: array
          items:
            $ref: '#/components/schemas/TaskStepView'
        output:
          type:
            - string
            - 'null'
        outputFiles:
          type: array
          items:
            $ref: '#/components/schemas/FileView'
        browserUseVersion:
          type:
            - string
            - 'null'
        isSuccess:
          type:
            - boolean
            - 'null'
        judgement:
          type:
            - string
            - 'null'
        judgeVerdict:
          type:
            - boolean
            - 'null'
      required:
        - id
        - sessionId
        - llm
        - task
        - status
        - createdAt
        - steps
        - outputFiles

```

## SDK Code Examples

```python
import requests

url = "https://api.browser-use.com/api/v2/tasks/task_id"

headers = {"X-Browser-Use-API-Key": "<apiKey>"}

response = requests.get(url, headers=headers)

print(response.json())
```

```javascript
const url = 'https://api.browser-use.com/api/v2/tasks/task_id';
const options = {method: 'GET', headers: {'X-Browser-Use-API-Key': '<apiKey>'}};

try {
  const response = await fetch(url, options);
  const data = await response.json();
  console.log(data);
} catch (error) {
  console.error(error);
}
```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.browser-use.com/api/v2/tasks/task_id"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Browser-Use-API-Key", "<apiKey>")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
```

```ruby
require 'uri'
require 'net/http'

url = URI("https://api.browser-use.com/api/v2/tasks/task_id")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["X-Browser-Use-API-Key"] = '<apiKey>'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.browser-use.com/api/v2/tasks/task_id")
  .header("X-Browser-Use-API-Key", "<apiKey>")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.browser-use.com/api/v2/tasks/task_id', [
  'headers' => [
    'X-Browser-Use-API-Key' => '<apiKey>',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.browser-use.com/api/v2/tasks/task_id");
var request = new RestRequest(Method.GET);
request.AddHeader("X-Browser-Use-API-Key", "<apiKey>");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["X-Browser-Use-API-Key": "<apiKey>"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.browser-use.com/api/v2/tasks/task_id")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "GET"
request.allHTTPHeaderFields = headers

let session = URLSession.shared
let dataTask = session.dataTask(with: request as URLRequest, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    print(error as Any)
  } else {
    let httpResponse = response as? HTTPURLResponse
    print(httpResponse)
  }
})

dataTask.resume()
```

# Update Task

PATCH https://api.browser-use.com/api/v2/tasks/{task_id}
Content-Type: application/json

Control task execution with stop, pause, resume, or stop task and session actions.

Reference: https://docs.cloud.browser-use.com/api-reference/v-2-api-current/tasks/update-task-tasks-task-id-patch

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Update Task
  version: endpoint_tasks.update_task_tasks__task_id__patch
paths:
  /tasks/{task_id}:
    patch:
      operationId: update-task-tasks-task-id-patch
      summary: Update Task
      description: >-
        Control task execution with stop, pause, resume, or stop task and
        session actions.
      tags:
        - - subpackage_tasks
      parameters:
        - name: task_id
          in: path
          required: true
          schema:
            type: string
            format: uuid
        - name: X-Browser-Use-API-Key
          in: header
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Successful Response
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TaskView'
        '404':
          description: Task not found
          content: {}
        '422':
          description: Request validation failed
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/UpdateTaskRequest'
components:
  schemas:
    TaskUpdateAction:
      type: string
      enum:
        - value: stop
        - value: stop_task_and_session
    UpdateTaskRequest:
      type: object
      properties:
        action:
          $ref: '#/components/schemas/TaskUpdateAction'
      required:
        - action
    TaskStatus:
      type: string
      enum:
        - value: created
        - value: started
        - value: finished
        - value: stopped
    TaskStepView:
      type: object
      properties:
        number:
          type: integer
        memory:
          type: string
        evaluationPreviousGoal:
          type: string
        nextGoal:
          type: string
        url:
          type: string
        screenshotUrl:
          type:
            - string
            - 'null'
        actions:
          type: array
          items:
            type: string
      required:
        - number
        - memory
        - evaluationPreviousGoal
        - nextGoal
        - url
        - actions
    FileView:
      type: object
      properties:
        id:
          type: string
          format: uuid
        fileName:
          type: string
      required:
        - id
        - fileName
    TaskView:
      type: object
      properties:
        id:
          type: string
          format: uuid
        sessionId:
          type: string
          format: uuid
        llm:
          type: string
        task:
          type: string
        status:
          $ref: '#/components/schemas/TaskStatus'
        createdAt:
          type: string
          format: date-time
        startedAt:
          type:
            - string
            - 'null'
          format: date-time
        finishedAt:
          type:
            - string
            - 'null'
          format: date-time
        metadata:
          type: object
          additionalProperties:
            description: Any type
        steps:
          type: array
          items:
            $ref: '#/components/schemas/TaskStepView'
        output:
          type:
            - string
            - 'null'
        outputFiles:
          type: array
          items:
            $ref: '#/components/schemas/FileView'
        browserUseVersion:
          type:
            - string
            - 'null'
        isSuccess:
          type:
            - boolean
            - 'null'
        judgement:
          type:
            - string
            - 'null'
        judgeVerdict:
          type:
            - boolean
            - 'null'
      required:
        - id
        - sessionId
        - llm
        - task
        - status
        - createdAt
        - steps
        - outputFiles

```

## SDK Code Examples

```python
import requests

url = "https://api.browser-use.com/api/v2/tasks/task_id"

payload = { "action": "stop" }
headers = {
    "X-Browser-Use-API-Key": "<apiKey>",
    "Content-Type": "application/json"
}

response = requests.patch(url, json=payload, headers=headers)

print(response.json())
```

```javascript
const url = 'https://api.browser-use.com/api/v2/tasks/task_id';
const options = {
  method: 'PATCH',
  headers: {'X-Browser-Use-API-Key': '<apiKey>', 'Content-Type': 'application/json'},
  body: '{"action":"stop"}'
};

try {
  const response = await fetch(url, options);
  const data = await response.json();
  console.log(data);
} catch (error) {
  console.error(error);
}
```

```go
package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {

	url := "https://api.browser-use.com/api/v2/tasks/task_id"

	payload := strings.NewReader("{\n  \"action\": \"stop\"\n}")

	req, _ := http.NewRequest("PATCH", url, payload)

	req.Header.Add("X-Browser-Use-API-Key", "<apiKey>")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
```

```ruby
require 'uri'
require 'net/http'

url = URI("https://api.browser-use.com/api/v2/tasks/task_id")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Patch.new(url)
request["X-Browser-Use-API-Key"] = '<apiKey>'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"action\": \"stop\"\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.patch("https://api.browser-use.com/api/v2/tasks/task_id")
  .header("X-Browser-Use-API-Key", "<apiKey>")
  .header("Content-Type", "application/json")
  .body("{\n  \"action\": \"stop\"\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('PATCH', 'https://api.browser-use.com/api/v2/tasks/task_id', [
  'body' => '{
  "action": "stop"
}',
  'headers' => [
    'Content-Type' => 'application/json',
    'X-Browser-Use-API-Key' => '<apiKey>',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.browser-use.com/api/v2/tasks/task_id");
var request = new RestRequest(Method.PATCH);
request.AddHeader("X-Browser-Use-API-Key", "<apiKey>");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"action\": \"stop\"\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "X-Browser-Use-API-Key": "<apiKey>",
  "Content-Type": "application/json"
]
let parameters = ["action": "stop"] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.browser-use.com/api/v2/tasks/task_id")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "PATCH"
request.allHTTPHeaderFields = headers
request.httpBody = postData as Data

let session = URLSession.shared
let dataTask = session.dataTask(with: request as URLRequest, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    print(error as Any)
  } else {
    let httpResponse = response as? HTTPURLResponse
    print(httpResponse)
  }
})

dataTask.resume()
```

# Get Task Logs

GET https://api.browser-use.com/api/v2/tasks/{task_id}/logs

Get secure download URL for task execution logs with step-by-step details.

Reference: https://docs.cloud.browser-use.com/api-reference/v-2-api-current/tasks/get-task-logs-tasks-task-id-logs-get

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Task Logs
  version: endpoint_tasks.get_task_logs_tasks__task_id__logs_get
paths:
  /tasks/{task_id}/logs:
    get:
      operationId: get-task-logs-tasks-task-id-logs-get
      summary: Get Task Logs
      description: >-
        Get secure download URL for task execution logs with step-by-step
        details.
      tags:
        - - subpackage_tasks
      parameters:
        - name: task_id
          in: path
          required: true
          schema:
            type: string
            format: uuid
        - name: X-Browser-Use-API-Key
          in: header
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Successful Response
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TaskLogFileResponse'
        '404':
          description: Task not found
          content: {}
        '422':
          description: Validation Error
          content: {}
        '500':
          description: Failed to generate download URL
          content: {}
components:
  schemas:
    TaskLogFileResponse:
      type: object
      properties:
        downloadUrl:
          type: string
      required:
        - downloadUrl

```

## SDK Code Examples

```python
import requests

url = "https://api.browser-use.com/api/v2/tasks/task_id/logs"

headers = {"X-Browser-Use-API-Key": "<apiKey>"}

response = requests.get(url, headers=headers)

print(response.json())
```

```javascript
const url = 'https://api.browser-use.com/api/v2/tasks/task_id/logs';
const options = {method: 'GET', headers: {'X-Browser-Use-API-Key': '<apiKey>'}};

try {
  const response = await fetch(url, options);
  const data = await response.json();
  console.log(data);
} catch (error) {
  console.error(error);
}
```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.browser-use.com/api/v2/tasks/task_id/logs"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Browser-Use-API-Key", "<apiKey>")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
```

```ruby
require 'uri'
require 'net/http'

url = URI("https://api.browser-use.com/api/v2/tasks/task_id/logs")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["X-Browser-Use-API-Key"] = '<apiKey>'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.browser-use.com/api/v2/tasks/task_id/logs")
  .header("X-Browser-Use-API-Key", "<apiKey>")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.browser-use.com/api/v2/tasks/task_id/logs', [
  'headers' => [
    'X-Browser-Use-API-Key' => '<apiKey>',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.browser-use.com/api/v2/tasks/task_id/logs");
var request = new RestRequest(Method.GET);
request.AddHeader("X-Browser-Use-API-Key", "<apiKey>");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["X-Browser-Use-API-Key": "<apiKey>"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.browser-use.com/api/v2/tasks/task_id/logs")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "GET"
request.allHTTPHeaderFields = headers

let session = URLSession.shared
let dataTask = session.dataTask(with: request as URLRequest, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    print(error as Any)
  } else {
    let httpResponse = response as? HTTPURLResponse
    print(httpResponse)
  }
})

dataTask.resume()
```

# List Tasks

GET https://api.browser-use.com/api/v2/tasks

Get paginated list of AI agent tasks with optional filtering by session and status.

Reference: https://docs.cloud.browser-use.com/api-reference/v-2-api-current/tasks/list-tasks-tasks-get

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: List Tasks
  version: endpoint_tasks.list_tasks_tasks_get
paths:
  /tasks:
    get:
      operationId: list-tasks-tasks-get
      summary: List Tasks
      description: >-
        Get paginated list of AI agent tasks with optional filtering by session
        and status.
      tags:
        - - subpackage_tasks
      parameters:
        - name: pageSize
          in: query
          required: false
          schema:
            type: integer
        - name: pageNumber
          in: query
          required: false
          schema:
            type: integer
        - name: sessionId
          in: query
          required: false
          schema:
            type:
              - string
              - 'null'
            format: uuid
        - name: filterBy
          in: query
          required: false
          schema:
            oneOf:
              - $ref: '#/components/schemas/TaskStatus'
              - type: 'null'
        - name: after
          in: query
          required: false
          schema:
            type:
              - string
              - 'null'
            format: date-time
        - name: before
          in: query
          required: false
          schema:
            type:
              - string
              - 'null'
            format: date-time
        - name: X-Browser-Use-API-Key
          in: header
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Successful Response
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/TaskListResponse'
        '422':
          description: Validation Error
          content: {}
components:
  schemas:
    TaskStatus:
      type: string
      enum:
        - value: created
        - value: started
        - value: finished
        - value: stopped
    TaskItemView:
      type: object
      properties:
        id:
          type: string
          format: uuid
        sessionId:
          type: string
          format: uuid
        llm:
          type: string
        task:
          type: string
        status:
          $ref: '#/components/schemas/TaskStatus'
        createdAt:
          type: string
          format: date-time
        startedAt:
          type:
            - string
            - 'null'
          format: date-time
        finishedAt:
          type:
            - string
            - 'null'
          format: date-time
        metadata:
          type: object
          additionalProperties:
            description: Any type
        output:
          type:
            - string
            - 'null'
        browserUseVersion:
          type:
            - string
            - 'null'
        isSuccess:
          type:
            - boolean
            - 'null'
        judgement:
          type:
            - string
            - 'null'
        judgeVerdict:
          type:
            - boolean
            - 'null'
      required:
        - id
        - sessionId
        - llm
        - task
        - status
        - createdAt
    TaskListResponse:
      type: object
      properties:
        items:
          type: array
          items:
            $ref: '#/components/schemas/TaskItemView'
        totalItems:
          type: integer
        pageNumber:
          type: integer
        pageSize:
          type: integer
      required:
        - items
        - totalItems
        - pageNumber
        - pageSize

```

## SDK Code Examples

```python
import requests

url = "https://api.browser-use.com/api/v2/tasks"

headers = {"X-Browser-Use-API-Key": "<apiKey>"}

response = requests.get(url, headers=headers)

print(response.json())
```

```javascript
const url = 'https://api.browser-use.com/api/v2/tasks';
const options = {method: 'GET', headers: {'X-Browser-Use-API-Key': '<apiKey>'}};

try {
  const response = await fetch(url, options);
  const data = await response.json();
  console.log(data);
} catch (error) {
  console.error(error);
}
```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.browser-use.com/api/v2/tasks"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Browser-Use-API-Key", "<apiKey>")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
```

```ruby
require 'uri'
require 'net/http'

url = URI("https://api.browser-use.com/api/v2/tasks")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["X-Browser-Use-API-Key"] = '<apiKey>'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.browser-use.com/api/v2/tasks")
  .header("X-Browser-Use-API-Key", "<apiKey>")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.browser-use.com/api/v2/tasks', [
  'headers' => [
    'X-Browser-Use-API-Key' => '<apiKey>',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.browser-use.com/api/v2/tasks");
var request = new RestRequest(Method.GET);
request.AddHeader("X-Browser-Use-API-Key", "<apiKey>");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["X-Browser-Use-API-Key": "<apiKey>"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.browser-use.com/api/v2/tasks")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "GET"
request.allHTTPHeaderFields = headers

let session = URLSession.shared
let dataTask = session.dataTask(with: request as URLRequest, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    print(error as Any)
  } else {
    let httpResponse = response as? HTTPURLResponse
    print(httpResponse)
  }
})

dataTask.resume()
```
