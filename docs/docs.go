// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todo": {
            "post": {
                "description": "Create a new todo with the input payload",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Create a new todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Create todo",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Todo"
                        }
                    }
                }
            }
        },
        "/todos": {
            "get": {
                "description": "Get details of all todos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Get details of all todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Todo"
                            }
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Get specific todo",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Get specific todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Get todos by Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Todo"
                        }
                    }
                }
            },
            "put": {
                "description": "Update specific todo",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Update specific todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Update todos by Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Todo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete specific todo",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todos"
                ],
                "summary": "Delete specific todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete todos by Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Todo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Todo": {
            "type": "object",
            "properties": {
                "complete": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/todos",
	Schemes:     []string{},
	Title:       "Todo API",
	Description: "This is a sample API.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
