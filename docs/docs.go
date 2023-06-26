// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/glyph/{matchID}": {
            "post": {
                "description": "Get glyphs using match id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "glyph"
                ],
                "summary": "Get glyphs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Match ID",
                        "name": "matchID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Glyphs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Glyph"
                            }
                        }
                    },
                    "400": {
                        "description": "Glyphs parse error",
                        "schema": {
                            "$ref": "#/definitions/dtos.MessageResponseType"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.MessageResponseType": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Glyph": {
            "type": "object",
            "properties": {
                "heroID": {
                    "description": "ID of hero (https://liquipedia.net/dota2/MediaWiki:Dota2webapi-heroes.json)",
                    "type": "integer"
                },
                "matchID": {
                    "type": "integer"
                },
                "minute": {
                    "type": "integer"
                },
                "second": {
                    "type": "integer"
                },
                "team": {
                    "description": "Radiant team is 2 and dire team is 3",
                    "type": "integer"
                },
                "userSteamID": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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
