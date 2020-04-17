// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-16 20:36:01.257830005 -0400 EDT m=+0.112551187

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
        "contact": {},
        "license": {
            "name": "MIT License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health/hello": {
            "get": {
                "description": "Non-authenticated endpoint that returns 200 with hello message. Used to validate that the service is responsive.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Hello sanity endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.Ping"
                        }
                    }
                }
            }
        },
        "/portfolio": {
            "get": {
                "description": "Non-authenticated endpoint that returns array of all stored portfolios.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "portfolio"
                ],
                "summary": "Get portfolios endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AllPortfoliosViewModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Insert portfolio. Returns the portfolio ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "portfolio"
                ],
                "summary": "Creates portfolio a unique ID",
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "portfolio",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PortfolioCreateModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PortfolioID"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a portfolio with the specified ID. Returns 200 if the resource did not already exist.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "portfolio"
                ],
                "summary": "Deletes a portfolio at the specified ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Portfolio ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/portfolio/{id}": {
            "get": {
                "description": "Non-authenticated endpoint that returns a portfolio with matching key.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "portfolio"
                ],
                "summary": "Get portfolios endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Portfolio ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PortfolioViewModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "health.Ping": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "hello"
                }
            }
        },
        "model.AllPortfoliosViewModel": {
            "type": "object",
            "properties": {
                "portfolios": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PortfolioViewModel"
                    }
                }
            }
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.MetadataViewModel": {
            "type": "object",
            "properties": {
                "createTime": {
                    "type": "string",
                    "example": "02/01/2020 11:12:00"
                },
                "id": {
                    "type": "string",
                    "example": "123884"
                },
                "lastUpdated": {
                    "type": "string",
                    "example": "02/01/2020 11:12:00"
                }
            }
        },
        "model.PortfolioCreateModel": {
            "type": "object",
            "properties": {
                "stocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StockViewModel"
                    }
                }
            }
        },
        "model.PortfolioID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "model.PortfolioViewModel": {
            "type": "object",
            "properties": {
                "metadata": {
                    "type": "object",
                    "$ref": "#/definitions/model.MetadataViewModel"
                },
                "stocks": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/model.StockViewModel"
                    }
                }
            }
        },
        "model.StockViewModel": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string",
                    "example": "CAD"
                },
                "currentPrice": {
                    "type": "number",
                    "example": 105
                },
                "name": {
                    "type": "string",
                    "example": "Canadian Pacific Railway Limited"
                },
                "purchaseDate": {
                    "type": "string",
                    "example": "02/03/2020"
                },
                "purchasePrice": {
                    "type": "number",
                    "example": 10000
                },
                "quantity": {
                    "type": "integer",
                    "example": 100
                },
                "ticker": {
                    "type": "string",
                    "example": "CP.TO"
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
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Portfolio Service API",
	Description: "Stores and fetches user and model data",
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
