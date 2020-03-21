// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-20 21:37:28.403186952 -0400 EDT m=+0.085229714

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
                "Portfolios": {
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
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastUpdated": {
                    "type": "string"
                }
            }
        },
        "model.PortfolioViewModel": {
            "type": "object",
            "properties": {
                "Metadata": {
                    "type": "object",
                    "$ref": "#/definitions/model.MetadataViewModel"
                },
                "stocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StockViewModel"
                    }
                }
            }
        },
        "model.StockViewModel": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "currentPrice": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "purchaseDate": {
                    "type": "string"
                },
                "purchasePrice": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "ticker": {
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
