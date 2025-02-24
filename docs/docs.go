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
        "/deals": {
            "get": {
                "description": "Returns all deals from Pipedrive",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deals"
                ],
                "summary": "Get all deals",
                "parameters": [
                    {
                        "type": "string",
                        "example": "cursor_token",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "custom_field_1,custom_field_2",
                        "name": "custom_fields",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 123,
                        "name": "filter_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "1,2,3",
                        "name": "ids",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "next_activity_id,last_activity_id",
                        "name": "include_fields",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 100,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 101,
                        "name": "org_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 456,
                        "name": "owner_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 789,
                        "name": "person_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 202,
                        "name": "pipeline_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "update_time",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "desc",
                        "name": "sort_direction",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 303,
                        "name": "stage_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "open,winner",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2025-01-01T10:20:00Z",
                        "name": "updated_since",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2025-01-02T10:20:00Z",
                        "name": "updated_until",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetDealsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.GetDealsResponse": {
            "type": "object",
            "properties": {
                "additional_data": {
                    "description": "Additional metadata"
                },
                "data": {
                    "description": "Deals data from Pipedrive"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
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
	Title:            "DevOps Challenge API",
	Description:      "This is a simple API for the DevOps challenge.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
