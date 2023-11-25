// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/api/events": {
            "get": {
                "description": "This endpoint returns all events matching the search terms. Note that only events from today on will be returned, ie no past events.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get all events.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title search string",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "location search string",
                        "name": "location",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "city search string",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "country search string",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "radius around given city in kilometers",
                        "name": "radius",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "date search string",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Event"
                            }
                        }
                    },
                    "404": {
                        "description": "No events found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Add new events to the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Add new events.",
                "parameters": [
                    {
                        "description": "Event Info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Event"
                            }
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "failed to parse body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to insert events",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete events.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Delete events.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sourceUrl string",
                        "name": "sourceUrl",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "datetime string",
                        "name": "datetime",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A success message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to delete events",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/events/today/slack": {
            "post": {
                "description": "This endpoint returns today's events in a format that slack needs for its slash command. Currently, Zurich is hardcoded as city (will be changed).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get today's events.",
                "responses": {
                    "200": {
                        "description": "A json with the results",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/events/{field}": {
            "get": {
                "description": "This endpoint returns all distinct values for the given field. Note that past events are not considered for this query.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get distinct field values.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "field name, can only be location or city",
                        "name": "field",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to retrieve values",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/notifications/activate": {
            "get": {
                "description": "This endpoint activates a notification that has been added previously if the inactive notification hasn't expired yet (expires after 24h).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notifications"
                ],
                "summary": "Activate notification.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "400": {
                        "description": "failed to activate notification",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed to activate notification",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/notifications/add": {
            "get": {
                "description": "Add new notification to the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notifications"
                ],
                "summary": "Add new notification.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title search string",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "location search string",
                        "name": "location",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "city search string",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "country search string",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "radius around given city in kilometers",
                        "name": "radius",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "date search string",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Failed to parse body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to insert notification",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/notifications/delete": {
            "get": {
                "description": "This endpoint deletes a notification that has been added previously based on the email address and the token.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notifications"
                ],
                "summary": "Delete notification.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "500": {
                        "description": "Failed to delete notification",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/notifications/deleteInactive": {
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "This endpoint deletes all inactive notification that are older than 24h.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notifications"
                ],
                "summary": "Delete inactive notifications.",
                "responses": {
                    "500": {
                        "description": "Failed to delete notifications",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/notifications/send": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "This endpoint sends an email for every active notification whose query returns a result.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notifications"
                ],
                "summary": "Send notifications.",
                "responses": {
                    "500": {
                        "description": "failed to send notifications",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Event": {
            "type": "object",
            "required": [
                "city",
                "date",
                "location",
                "sourceUrl",
                "title",
                "type",
                "url"
            ],
            "properties": {
                "city": {
                    "type": "string",
                    "example": "SuperCity"
                },
                "comment": {
                    "type": "string",
                    "example": "Super exciting comment."
                },
                "country": {
                    "type": "string",
                    "example": "SuperCountry"
                },
                "date": {
                    "type": "string",
                    "example": "2021-10-31T19:00:00.000Z"
                },
                "geolocation": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    },
                    "example": [
                        7.4514512,
                        46.9482713
                    ]
                },
                "imageUrl": {
                    "type": "string",
                    "example": "http://link.to/concert/image.jpg"
                },
                "location": {
                    "type": "string",
                    "example": "SuperLocation"
                },
                "offset": {
                    "type": "integer"
                },
                "sourceUrl": {
                    "type": "string",
                    "example": "http://link.to/source"
                },
                "title": {
                    "type": "string",
                    "example": "ExcitingTitle"
                },
                "type": {
                    "type": "string",
                    "example": "concert"
                },
                "url": {
                    "type": "string",
                    "example": "http://link.to/concert/page"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
