// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/": {
            "get": {
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "index"
                ],
                "summary": "Index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/auth/token": {
            "post": {
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "auth user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AuthUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.JWTToken"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "User logout",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/auth/user": {
            "post": {
                "description": "Create user and storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatedUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/containers": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "List container",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "container"
                ],
                "summary": "List container",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Container"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create container",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "container"
                ],
                "summary": "Create container",
                "parameters": [
                    {
                        "description": "container info",
                        "name": "container",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatedContainer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Container"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/containers/:id": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update container",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "container"
                ],
                "summary": "Update container",
                "parameters": [
                    {
                        "description": "container info",
                        "name": "container",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatedContainer"
                        }
                    },
                    {
                        "type": "string",
                        "description": "container id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Container"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/containers/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get container",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "container"
                ],
                "summary": "Get container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Container"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Operate container",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "container"
                ],
                "summary": "Operate container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "verb: start/stop/restart",
                        "name": "verb",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete container",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "container"
                ],
                "summary": "Delete container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/containers/{id}/exec": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "exec container",
                "tags": [
                    "container"
                ],
                "summary": "Exec container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "shell, sh or bash",
                        "name": "shell",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/containers/{id}/log": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get container log",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "container"
                ],
                "summary": "Get container log",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "follow log",
                        "name": "follow",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "tail log all or number",
                        "name": "tail",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/containers/{id}/proxy": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "proxy container",
                "tags": [
                    "container"
                ],
                "summary": "Proxy container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "shell, sh or bash",
                        "name": "shell",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "List user and storage",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "List user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.User"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create user and storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatedUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get user and storage",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update user and storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatedUser"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete user and storage",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthz"
                ],
                "summary": "Healthz",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.AuthUser": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "setCookie": {
                    "type": "boolean"
                }
            }
        },
        "model.Container": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "cmd": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "startAt": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.CreatedContainer": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "model.CreatedUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.JWTToken": {
            "type": "object",
            "properties": {
                "describe": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "model.UpdatedUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "127.0.0.1:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "gin示例 API",
	Description:      "相关接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}