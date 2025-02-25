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
        "/v1/deals": {
            "get": {
                "description": "Fetches all deals from Pipedrive",
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
                        "type": "integer",
                        "example": 456,
                        "name": "filter_id",
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
                        "example": 1,
                        "name": "owned_by_you",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "update_time DESC",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 2,
                        "name": "stage_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "open",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 123,
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetAllDealsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new deal in Pipedrive",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deals"
                ],
                "summary": "Add a new deal",
                "parameters": [
                    {
                        "description": "Deal data to be created. Required fields: title",
                        "name": "deal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateDealRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.CreateDealResponse"
                        }
                    }
                }
            }
        },
        "/v1/deals/{id}": {
            "put": {
                "description": "Updates an existing deal in Pipedrive",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deals"
                ],
                "summary": "Update an existing deal",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Deal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Deal data to update. Optional fields can be omitted.",
                        "name": "deal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateDealRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.UpdateDealResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CreateDealRequest": {
            "type": "object"
        },
        "api.CreateDealResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "api.Deal": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean",
                    "example": true
                },
                "activities_count": {
                    "type": "integer",
                    "example": 1
                },
                "acv_currency": {
                    "type": "string",
                    "example": "EUR"
                },
                "add_time": {
                    "type": "string",
                    "example": "2019-05-29 04:21:51"
                },
                "arr_currency": {
                    "type": "string",
                    "example": "EUR"
                },
                "channel": {
                    "type": "integer"
                },
                "channel_id": {
                    "type": "string"
                },
                "close_time": {
                    "type": "string"
                },
                "currency": {
                    "type": "string",
                    "example": "EUR"
                },
                "deleted": {
                    "type": "boolean",
                    "example": false
                },
                "done_activities_count": {
                    "type": "integer",
                    "example": 0
                },
                "email_messages_count": {
                    "type": "integer",
                    "example": 4
                },
                "expected_close_date": {
                    "type": "string",
                    "example": "2019-06-29"
                },
                "files_count": {
                    "type": "integer",
                    "example": 0
                },
                "first_won_time": {
                    "type": "string"
                },
                "followers_count": {
                    "type": "integer",
                    "example": 0
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "label": {
                    "type": "string",
                    "example": "11"
                },
                "last_activity_date": {
                    "type": "string"
                },
                "last_activity_id": {
                    "type": "integer"
                },
                "last_incoming_mail_time": {
                    "type": "string"
                },
                "last_outgoing_mail_time": {
                    "type": "string"
                },
                "lost_reason": {
                    "type": "string"
                },
                "lost_time": {
                    "type": "string"
                },
                "mrr": {
                    "type": "number"
                },
                "mrr_currency": {
                    "type": "string",
                    "example": "EUR"
                },
                "next_activity_date": {
                    "type": "string",
                    "example": "2019-11-29"
                },
                "next_activity_id": {
                    "type": "integer"
                },
                "next_activity_time": {
                    "type": "string",
                    "example": "11:30:00"
                },
                "notes_count": {
                    "type": "integer",
                    "example": 2
                },
                "origin": {
                    "type": "string",
                    "example": "ManuallyCreated"
                },
                "participants_count": {
                    "type": "integer",
                    "example": 1
                },
                "pipeline_id": {
                    "type": "integer",
                    "example": 1
                },
                "probability": {
                    "type": "integer"
                },
                "products_count": {
                    "type": "integer",
                    "example": 2
                },
                "stage_change_time": {
                    "type": "string",
                    "example": "2019-11-28 15:41:22"
                },
                "stage_id": {
                    "type": "integer",
                    "example": 2
                },
                "stage_order_nr": {
                    "type": "integer"
                },
                "status": {
                    "type": "string",
                    "example": "open"
                },
                "title": {
                    "type": "string",
                    "example": "Deal One"
                },
                "undone_activities_count": {
                    "type": "integer",
                    "example": 1
                },
                "update_time": {
                    "type": "string",
                    "example": "2019-11-28 16:19:50"
                },
                "value": {
                    "type": "number",
                    "example": 130
                },
                "visible_to": {
                    "type": "string",
                    "example": "1"
                },
                "won_time": {
                    "type": "string"
                }
            }
        },
        "api.GetAllDealsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Deal"
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "api.UpdateDealRequest": {
            "type": "object"
        },
        "api.UpdateDealResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean",
                    "example": true
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
