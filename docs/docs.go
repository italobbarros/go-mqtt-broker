// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/containers": {
            "get": {
                "description": "Get all containers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Container"
                ],
                "summary": "Get all containers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Container"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new container",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Container"
                ],
                "summary": "Create a new container",
                "parameters": [
                    {
                        "description": "Container object that needs to be added",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ContainerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ContainerRequest"
                        }
                    }
                }
            }
        },
        "/containers/{id}": {
            "get": {
                "description": "Get a container by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Container"
                ],
                "summary": "Get a container by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Container ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Container"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a container by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Container"
                ],
                "summary": "Update a container by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Container ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated container object",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ContainerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ContainerRequest"
                        }
                    }
                }
            }
        },
        "/publications": {
            "get": {
                "description": "Get all publications",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publications"
                ],
                "summary": "Get all publications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Publication"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new publication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publications"
                ],
                "summary": "Create a new publication",
                "parameters": [
                    {
                        "description": "Publication object that needs to be added",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PublicationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.PublicationRequest"
                        }
                    }
                }
            }
        },
        "/publications/{id}": {
            "get": {
                "description": "Get a publication by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publications"
                ],
                "summary": "Get a publication by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Publication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Publication"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a publication by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Publications"
                ],
                "summary": "Update a publication by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Publication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated publication object",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PublicationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PublicationRequest"
                        }
                    }
                }
            }
        },
        "/sessions": {
            "get": {
                "description": "Get all sessions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Get all sessions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Session"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Create a new session",
                "parameters": [
                    {
                        "description": "Session object that needs to be added",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SessionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SessionRequest"
                        }
                    }
                }
            }
        },
        "/sessions/{id}": {
            "get": {
                "description": "Get a session by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Get a session by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Session"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a session by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Update a session by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated session object",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SessionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SessionRequest"
                        }
                    }
                }
            }
        },
        "/subscriptions": {
            "get": {
                "description": "Get all subscriptions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Get all subscriptions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Subscription"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Create a new subscription",
                "parameters": [
                    {
                        "description": "Subscription object that needs to be added",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Subscription"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Subscription"
                        }
                    }
                }
            }
        },
        "/subscriptions/{id}": {
            "get": {
                "description": "Get a subscription by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Get a subscription by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Subscription ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Subscription"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a subscription by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Update a subscription by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Subscription ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated subscription object",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SubscriptionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SubscriptionRequest"
                        }
                    }
                }
            }
        },
        "/topics": {
            "get": {
                "description": "Get all topics",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "Get all topics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Topic"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new topic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "Create a new topic",
                "parameters": [
                    {
                        "description": "Topic object that needs to be added",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TopicRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Topic"
                        }
                    }
                }
            }
        },
        "/topics/ByIdContainer/{IdContainer}": {
            "get": {
                "description": "Get all topics",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "Get all topics",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Topic by IdContainer",
                        "name": "IdContainer",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Topic"
                            }
                        }
                    }
                }
            }
        },
        "/topics/{id}": {
            "get": {
                "description": "Get a topic by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "Get a topic by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Topic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Topic"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a topic by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "Update a topic by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Topic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated topic object",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TopicRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TopicRequest"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Container": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ContainerRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Publication": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/models.Container"
                },
                "created": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                },
                "deleted": {
                    "type": "string",
                    "example": "2024-01-16T12:45:00Z"
                },
                "id": {
                    "type": "integer"
                },
                "idContainer": {
                    "type": "integer"
                },
                "idTopic": {
                    "type": "integer"
                },
                "payload": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "integer"
                },
                "topic": {
                    "$ref": "#/definitions/models.Topic"
                },
                "updated": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                }
            }
        },
        "models.PublicationRequest": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/models.Container"
                },
                "idContainer": {
                    "type": "integer"
                },
                "idTopic": {
                    "type": "integer"
                },
                "payload": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "integer"
                },
                "topic": {
                    "$ref": "#/definitions/models.Topic"
                }
            }
        },
        "models.Session": {
            "type": "object",
            "properties": {
                "clean": {
                    "type": "boolean"
                },
                "clientId": {
                    "type": "string"
                },
                "container": {
                    "$ref": "#/definitions/models.Container"
                },
                "created": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                },
                "deleted": {
                    "type": "string",
                    "example": "2024-01-16T12:45:00Z"
                },
                "id": {
                    "type": "integer"
                },
                "idContainer": {
                    "type": "integer"
                },
                "keepAlive": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updated": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.SessionRequest": {
            "type": "object",
            "properties": {
                "clean": {
                    "type": "boolean"
                },
                "clientId": {
                    "type": "string"
                },
                "idContainer": {
                    "type": "integer"
                },
                "keepAlive": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Subscription": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/models.Container"
                },
                "created": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                },
                "deleted": {
                    "type": "string",
                    "example": "2024-01-16T12:45:00Z"
                },
                "id": {
                    "type": "integer"
                },
                "idContainer": {
                    "type": "integer"
                },
                "idSession": {
                    "type": "integer"
                },
                "idTopic": {
                    "type": "integer"
                },
                "session": {
                    "$ref": "#/definitions/models.Session"
                },
                "topic": {
                    "$ref": "#/definitions/models.Topic"
                },
                "updated": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                }
            }
        },
        "models.SubscriptionRequest": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/models.Container"
                },
                "idContainer": {
                    "type": "integer"
                },
                "idTopic": {
                    "type": "integer"
                },
                "sessionId": {
                    "type": "string"
                },
                "topic": {
                    "$ref": "#/definitions/models.Topic"
                }
            }
        },
        "models.Topic": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/models.TopicConfig"
                },
                "container": {
                    "$ref": "#/definitions/models.Container"
                },
                "created": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                },
                "deleted": {
                    "type": "string",
                    "example": "2024-01-16T12:45:00Z"
                },
                "id": {
                    "type": "integer"
                },
                "idContainer": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                }
            }
        },
        "models.TopicConfig": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "string"
                },
                "qos": {
                    "type": "integer"
                },
                "retained": {
                    "description": "Indica se a mensagem é retida ou não",
                    "type": "boolean"
                },
                "securityRule": {
                    "description": "Regra de segurança aplicada ao tópico",
                    "type": "string"
                }
            }
        },
        "models.TopicRequest": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/models.TopicConfig"
                },
                "idContainer": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
