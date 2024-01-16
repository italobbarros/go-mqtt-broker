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
                                "$ref": "#/definitions/api.Container"
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
                            "$ref": "#/definitions/api.ContainerPost"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.ContainerPost"
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
                            "$ref": "#/definitions/api.Container"
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
                            "$ref": "#/definitions/api.ContainerPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ContainerPost"
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
                                "$ref": "#/definitions/api.Publication"
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
                            "$ref": "#/definitions/api.PublicationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.PublicationRequest"
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
                            "$ref": "#/definitions/api.Publication"
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
                            "$ref": "#/definitions/api.PublicationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.PublicationRequest"
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
                                "$ref": "#/definitions/api.Session"
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
                            "$ref": "#/definitions/api.SessionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.SessionRequest"
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
                            "$ref": "#/definitions/api.Session"
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
                            "$ref": "#/definitions/api.SessionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SessionRequest"
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
                                "$ref": "#/definitions/api.Subscription"
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
                            "$ref": "#/definitions/api.Subscription"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.Subscription"
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
                            "$ref": "#/definitions/api.Subscription"
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
                            "$ref": "#/definitions/api.SubscriptionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SubscriptionRequest"
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
                                "$ref": "#/definitions/api.Topic"
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
                            "$ref": "#/definitions/api.TopicRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.Topic"
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
                                "$ref": "#/definitions/api.Topic"
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
                            "$ref": "#/definitions/api.Topic"
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
                            "$ref": "#/definitions/api.TopicRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.TopicRequest"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Container": {
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
        "api.ContainerPost": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "api.Publication": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/api.Container"
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
                    "$ref": "#/definitions/api.Topic"
                },
                "updated": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                }
            }
        },
        "api.PublicationRequest": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/api.Container"
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
                    "$ref": "#/definitions/api.Topic"
                }
            }
        },
        "api.Session": {
            "type": "object",
            "properties": {
                "clean": {
                    "type": "boolean"
                },
                "clientId": {
                    "type": "string"
                },
                "container": {
                    "$ref": "#/definitions/api.Container"
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
        "api.SessionRequest": {
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
        "api.Subscription": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/api.Container"
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
                    "$ref": "#/definitions/api.Session"
                },
                "topic": {
                    "$ref": "#/definitions/api.Topic"
                },
                "updated": {
                    "type": "string",
                    "example": "2024-01-16T12:00:00Z"
                }
            }
        },
        "api.SubscriptionRequest": {
            "type": "object",
            "properties": {
                "container": {
                    "$ref": "#/definitions/api.Container"
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
                    "$ref": "#/definitions/api.Topic"
                }
            }
        },
        "api.Topic": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/api.TopicConfig"
                },
                "container": {
                    "$ref": "#/definitions/api.Container"
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
        "api.TopicConfig": {
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
        "api.TopicRequest": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/api.TopicConfig"
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
