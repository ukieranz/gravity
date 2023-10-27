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
        "/api/v1/auth/signin": {
            "post": {
                "description": "Sign in account for admin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Sign In",
                        "name": "sign_in",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.SignInResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/signup": {
            "post": {
                "description": "Register account for admin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Sign Up",
                        "name": "sign_up",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.SignUpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/lab/activity": {
            "post": {
                "description": "Save activity type in/out.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lab"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit records",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/lab/history": {
            "get": {
                "description": "Get History.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lab"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit records",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.DataResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/lab/register": {
            "post": {
                "description": "Register lab.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lab"
                ],
                "parameters": [
                    {
                        "description": "RegisterLab",
                        "name": "sign_in",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.RegistrationLabRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/lab/user": {
            "get": {
                "description": "Get User for checkin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lab"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "student id",
                        "name": "sid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "student id",
                        "name": "room",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.DataResponse"
                        }
                    }
                }
            }
        },
        "/common/capture": {
            "get": {
                "description": "test sentry.io capture.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/common/email": {
            "get": {
                "description": "test email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/common/healthcheck": {
            "get": {
                "description": "health check api server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/common/sentry": {
            "get": {
                "description": "test sentry.io.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/common/workercheck": {
            "get": {
                "description": "health check worker.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.DataResponse": {
            "type": "object",
            "required": [
                "data",
                "success"
            ],
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.RegistrationLabRequest": {
            "type": "object",
            "required": [
                "email",
                "end_day",
                "first_name",
                "last_name",
                "phone_number",
                "registration_time",
                "room_id",
                "start_day",
                "student_id",
                "supervisor"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "end_day": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "registration_time": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                },
                "start_day": {
                    "type": "string"
                },
                "student_id": {
                    "type": "string"
                },
                "supervisor": {
                    "type": "string"
                }
            }
        },
        "schema.Response": {
            "type": "object",
            "required": [
                "msg",
                "success"
            ],
            "properties": {
                "msg": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.SignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "schema.SignInResponse": {
            "type": "object",
            "required": [
                "email",
                "success",
                "token"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "schema.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "id",
                "last_name",
                "password",
                "phone_number"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "schema.SignUpResponse": {
            "type": "object",
            "required": [
                "msg",
                "success"
            ],
            "properties": {
                "msg": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Student Checkin System",
	Description:      "This is a documentation for the Student Checkin System API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
