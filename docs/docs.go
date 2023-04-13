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
        "/employee": {
            "post": {
                "description": "Create a new employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Create a new employee",
                "parameters": [
                    {
                        "description": "Create a new employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EmployeeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/employee/{id}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get details of employee corresponding to the input id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Get details for a given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Update an employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update an employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EmployeeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Delete an employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.EmployeeReq": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 20
                },
                "division": {
                    "type": "string",
                    "example": "Tech"
                },
                "email": {
                    "type": "string",
                    "example": "marwan@example.com"
                },
                "full_name": {
                    "type": "string",
                    "example": "Marwan"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
