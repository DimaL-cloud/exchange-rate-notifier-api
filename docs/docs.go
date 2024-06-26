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
        "/api/subscribe": {
            "post": {
                "description": "subscribe to notifications",
                "tags": [
                    "subscription"
                ],
                "summary": "Subscribe to notifications",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "400": {
                        "description": "invalid email format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to create subscription",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/unsubscribe": {
            "post": {
                "description": "unsubscribe from notifications",
                "tags": [
                    "subscription"
                ],
                "summary": "Unsubscribe from notifications",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "400": {
                        "description": "invalid email format",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to delete subscription",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Exchange rate notifier API",
	Description:      "API server for notifying exchange rate",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
