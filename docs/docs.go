// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/message": {
            "post": {
                "description": "Create new message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Create new message",
                "parameters": [
                    {
                        "description": "some text",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.MessageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/messages": {
            "get": {
                "description": "get list of message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "List of messages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Message"
                            }
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "auth body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/user/active-users": {
            "get": {
                "description": "get list of users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "List of active users",
                "responses": {}
            }
        },
        "/user/list": {
            "get": {
                "description": "get list of users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "List of users",
                "responses": {}
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
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
                        "description": "auth body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.AuthRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "handlers.LoginResponse": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "handlers.MessageRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                }
            }
        },
        "handlers.RegisterResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "required": [
                "body"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Lets go chat",
	Description:      "This is a go sandbox.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
