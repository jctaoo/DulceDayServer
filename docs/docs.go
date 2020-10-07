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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/static/{key}": {
            "get": {
                "summary": "获取静态资源",
                "parameters": [
                    {
                        "type": "string",
                        "description": "资源路径",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/user/login/email": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "使用邮箱登陆",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.loginWithEmailParameter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/login/sensitive/email": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "使用邮箱进行敏感登陆验证, 需要事先登录",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.sEmailLoginParameter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/login/username": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "使用用户名登陆",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.loginWithUsernameParameter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取登录用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_profile.getProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/profile/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "userProfile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_profile.updateProfileParameter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_profile.updateProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/profile/update/avatar": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新头像 (文件传输 go-swagger 无法胜任，请使用 postman 等工具)",
                "parameters": [
                    {
                        "type": "file",
                        "description": "头像图片",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_profile.updateAvatarResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/profile/{username}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_profile.getProfileResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.registerParameter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.registerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/register/sensitive/email": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "敏感注册，用于生成验证码等, 需要事先登录",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.sEmailRegisterParameter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.BaseResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UserProfile": {
            "type": "object",
            "properties": {
                "avatar_file_key": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.loginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "user.loginWithEmailParameter": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "device_name": {
                    "description": "登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”",
                    "type": "string",
                    "example": "bob的iPhone"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "haha@test.com"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "qwerty123"
                }
            }
        },
        "user.loginWithUsernameParameter": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "device_name": {
                    "description": "登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”",
                    "type": "string",
                    "example": "bob的iPhone"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "qwerty123"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "bob"
                }
            }
        },
        "user.registerParameter": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "haha@test.com"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "qwerty123"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "bob"
                }
            }
        },
        "user.registerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.sEmailLoginParameter": {
            "type": "object",
            "required": [
                "email",
                "verificationCode"
            ],
            "properties": {
                "device_name": {
                    "description": "登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”",
                    "type": "string",
                    "example": "bob的iPhone"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "haha@test.com"
                },
                "verificationCode": {
                    "description": "验证码",
                    "type": "string",
                    "example": "623597"
                }
            }
        },
        "user.sEmailRegisterParameter": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "device_name": {
                    "description": "登录的设备的名字，浏览器为 “浏览器(ip所在的城市)”",
                    "type": "string",
                    "example": "bob的iPhone"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "haha@test.com"
                }
            }
        },
        "user_profile.getProfileResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "profile": {
                    "type": "object",
                    "$ref": "#/definitions/models.UserProfile"
                }
            }
        },
        "user_profile.updateAvatarResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user_profile.updateProfileParameter": {
            "type": "object",
            "properties": {
                "nickname": {
                    "description": "昵称",
                    "type": "string",
                    "example": "jc😄taoo"
                }
            }
        },
        "user_profile.updateProfileResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "nickname": {
                    "description": "更改后的昵称",
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
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "DulceDayServer",
	Description: "一个轻社区app, 内置直播, 帮助用户找到身边的兴趣圈",
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
