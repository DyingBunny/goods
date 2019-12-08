// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-24 16:15:49.5942104 +0800 CST m=+0.086997401

package docs

import (
	"bytes"
	"encoding/json"

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
        "/auth": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "token获取b",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username,password",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"fail\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username,password",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":code,\t\"msg\" : e.GetMsg(code),\t\"Logintime\":time.Now()}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\":code,\t\"msg\" : e.GetMsg(code),}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "信息查看",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\" : code,\t\"msg\" : e.GetMsg(code),\t\"user_id\":auth.User_id,\t\"username\":auth.Username,\t\"password\":auth.Password,\t\"gender\":auth.Gender,\t\"email\":auth.Email,\t\"login\":auth.Login,\t\"role\":auth.Role,}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\" : code,\t\"msg\" : e.GetMsg(code),}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}