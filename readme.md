
# testme 1.0

Run testme in Direktiv

---
- #### Categories: unknown
- #### Image: gcr.io/direktiv/apps/testme 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/testme/issues
- #### URL: https://github.com/direktiv-apps/testme
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About testme

Run testme in Direktiv as a function

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: testme
  image: gcr.io/direktiv/apps/testme:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: testme
  type: action
  action:
    function: testme
    input: 
      commands:
      - command: Example of running testme
```
   #### Advanced
```yaml
- id: testme
  type: action
  action:
    function: testme
    input: 
      files:
      - name: hello.txt
        data: Hello World
        mode: '0755'
      commands:
      - command: Example of running testme
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": null,
    "success": true
  },
  {
    "result": null,
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| testme | [][PostOKBodyTestmeItems](#post-o-k-body-testme-items)| `[]*PostOKBodyTestmeItems` |  | |  |  |


#### <span id="post-o-k-body-testme-items"></span> postOKBodyTestmeItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run |  |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
