// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-07-12 14:40:35.5594665 +0800 CST m=+0.080948801

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/job": {
            "get": {
                "description": "获取所有任务",
                "tags": [
                    "v1/job"
                ],
                "summary": "获取所有任务",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.Job"
                        }
                    }
                }
            },
            "post": {
                "description": "新建任务",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/job"
                ],
                "summary": "新建任务",
                "parameters": [
                    {
                        "type": "string",
                        "description": "任务名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.Job"
                        }
                    }
                }
            }
        },
        "/v1/job/{id}": {
            "get": {
                "description": "获取单个任务",
                "tags": [
                    "v1/job"
                ],
                "summary": "获取单个任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.Job"
                        }
                    }
                }
            },
            "put": {
                "description": "更新任务",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1/job"
                ],
                "summary": "更新任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务结果",
                        "name": "result",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除任务",
                "tags": [
                    "v1/job"
                ],
                "summary": "删除任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务id",
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
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Job": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键",
                    "type": "integer"
                },
                "name": {
                    "description": "任务名称",
                    "type": "string"
                },
                "result": {
                    "description": "任务结果",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "任务管理",
            "name": "v1/job"
        }
    ]
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "任务管理 API",
	Description: "任务管理系统",
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
