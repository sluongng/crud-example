// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-06-02 18:39:34.257379628 +0700 +07 m=+0.070981765

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Return general service information",
        "title": "CRUD example API",
        "contact": {
            "name": "Son Luong",
            "url": "https://sluongng.gitlab.io/",
            "email": "sluongng@gmail.com"
        },
        "license": {},
        "version": "0.1"
    },
    "paths": {
        "/": {
            "get": {
                "description": "Return general service information",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Root handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "get list of current users from Database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get list of users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create a new user to DB",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get a user by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get 1 user detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            },
            "put": {
                "description": "update a user by id",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update a user's detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user by id",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
