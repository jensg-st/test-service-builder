// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Run testme in Direktiv",
    "title": "testme",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "unknown"
      ],
      "container": "gcr.io/direktiv/apps/testme",
      "issues": "https://github.com/direktiv-apps/testme/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "Run testme in Direktiv as a function",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/testme"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "command": {
                        "description": "Command to run",
                        "type": "string"
                      },
                      "continue": {
                        "description": "Stops excecution if command fails, otherwise proceeds with next command",
                        "type": "boolean"
                      },
                      "print": {
                        "description": "If set to false the command will not print the full command with arguments to logs.",
                        "type": "boolean",
                        "default": true
                      },
                      "silent": {
                        "description": "If set to false the command will not print output to logs.",
                        "type": "boolean",
                        "default": false
                      }
                    }
                  }
                },
                "files": {
                  "description": "File to create before running commands.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/direktivFile"
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "properties": {
                "testme": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "success",
                      "result"
                    ],
                    "properties": {
                      "result": {
                        "additionalProperties": false
                      },
                      "success": {
                        "type": "boolean"
                      }
                    }
                  }
                }
              }
            },
            "examples": {
              "testme": [
                {
                  "result": null,
                  "success": true
                },
                {
                  "result": null,
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"testme\": {{ index . 0 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: testme\n  type: action\n  action:\n    function: testme\n    input: \n      commands:\n      - command: Example of running testme",
            "title": "Basic"
          },
          {
            "content": "- id: testme\n  type: action\n  action:\n    function: testme\n    input: \n      files:\n      - name: hello.txt\n        data: Hello World\n        mode: '0755'\n      commands:\n      - command: Example of running testme",
            "title": "Advanced"
          }
        ],
        "x-direktiv-function": "functions:\n- id: testme\n  image: gcr.io/direktiv/apps/testme:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "This is a secret value",
            "name": "testmeSecret"
          },
          {
            "description": "gerke",
            "name": "jens"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Run testme in Direktiv",
    "title": "testme",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "unknown"
      ],
      "container": "gcr.io/direktiv/apps/testme",
      "issues": "https://github.com/direktiv-apps/testme/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "Run testme in Direktiv as a function",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/testme"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "testme": [
                {
                  "result": null,
                  "success": true
                },
                {
                  "result": null,
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"testme\": {{ index . 0 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: testme\n  type: action\n  action:\n    function: testme\n    input: \n      commands:\n      - command: Example of running testme",
            "title": "Basic"
          },
          {
            "content": "- id: testme\n  type: action\n  action:\n    function: testme\n    input: \n      files:\n      - name: hello.txt\n        data: Hello World\n        mode: '0755'\n      commands:\n      - command: Example of running testme",
            "title": "Advanced"
          }
        ],
        "x-direktiv-function": "functions:\n- id: testme\n  image: gcr.io/direktiv/apps/testme:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "This is a secret value",
            "name": "testmeSecret"
          },
          {
            "description": "gerke",
            "name": "jens"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "testme": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/postOKBodyTestmeItems"
          }
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodyTestmeItems": {
      "type": "object",
      "required": [
        "success",
        "result"
      ],
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "commands": {
          "description": "Array of commands.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/postParamsBodyCommandsItems"
          }
        },
        "files": {
          "description": "File to create before running commands.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/direktivFile"
          }
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyCommandsItems": {
      "type": "object",
      "properties": {
        "command": {
          "description": "Command to run",
          "type": "string"
        },
        "continue": {
          "description": "Stops excecution if command fails, otherwise proceeds with next command",
          "type": "boolean"
        },
        "print": {
          "description": "If set to false the command will not print the full command with arguments to logs.",
          "type": "boolean",
          "default": true
        },
        "silent": {
          "description": "If set to false the command will not print output to logs.",
          "type": "boolean",
          "default": false
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
