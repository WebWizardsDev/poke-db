// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jacob Potter",
            "email": "pttr.jcb@gmail.com"
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
        "/evolution": {
            "get": {
                "description": "List Evolutions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution"
                ],
                "summary": "List Evolutions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Evolutions"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new evolution",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution"
                ],
                "summary": "Creates Evolution",
                "parameters": [
                    {
                        "description": "The new evolution",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Evolution"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Evolution"
                        }
                    }
                }
            }
        },
        "/evolution/{id}": {
            "get": {
                "description": "Get evolution by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution"
                ],
                "summary": "Get Evolution",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Evolution ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Evolution"
                        }
                    }
                }
            },
            "put": {
                "description": "Update evolution by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution"
                ],
                "summary": "Update Evolution",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Evolution ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The updated evolution",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Evolution"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Evolution"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete evolution by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution"
                ],
                "summary": "Delete Evolution",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Evolution ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/evolution_requirements": {
            "get": {
                "description": "List Evolution Requirements",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution_requirements"
                ],
                "summary": "List Evolution Requirements",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EvolutionRequirement"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new evolution requirement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution_requirements"
                ],
                "summary": "Creates Evolution Requirement",
                "parameters": [
                    {
                        "description": "The new evolution requirement",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EvolutionRequirement"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.EvolutionRequirement"
                        }
                    }
                }
            }
        },
        "/evolution_requirements/{id}": {
            "get": {
                "description": "Get evolution requirement by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution_requirements"
                ],
                "summary": "Get Evolution Requirement",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Evolution Requirement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EvolutionRequirement"
                        }
                    }
                }
            },
            "put": {
                "description": "Update evolution requirement by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution_requirements"
                ],
                "summary": "Update Evolution Requirement",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Evolution Requirement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The updated evolution requirement",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EvolutionRequirement"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.EvolutionRequirement"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete evolution requirement by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "evolution_requirements"
                ],
                "summary": "Delete Evolution Requirement",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Evolution Requirement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/moves": {
            "get": {
                "description": "List Moves",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "moves"
                ],
                "summary": "List Moves",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Moves"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new move",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "moves"
                ],
                "summary": "Creates Move",
                "parameters": [
                    {
                        "description": "The new move",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Move"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Move"
                        }
                    }
                }
            }
        },
        "/moves/{id}": {
            "get": {
                "description": "Get move by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "moves"
                ],
                "summary": "Get Move",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Move ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Move"
                        }
                    }
                }
            },
            "put": {
                "description": "Update move by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "moves"
                ],
                "summary": "Update Move",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Move ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The updated move",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Move"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Move"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete move by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "moves"
                ],
                "summary": "Delete Move",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Move ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pokemon": {
            "get": {
                "description": "List Pokemon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "List Pokemon",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Pokemons"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new pokemon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Creates Pokemon",
                "parameters": [
                    {
                        "description": "The new pokemon",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Pokemon"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Pokemon"
                        }
                    }
                }
            }
        },
        "/pokemon/{id}": {
            "get": {
                "description": "Get pokemon by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Get Pokemon",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pokemon ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Pokemon"
                        }
                    }
                }
            },
            "put": {
                "description": "Update pokemon by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Update Pokemon",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pokemon ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The updated pokemon",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Pokemon"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Pokemon"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete pokemon by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Delete Pokemon",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pokemon ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/type": {
            "get": {
                "description": "List Types",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "type"
                ],
                "summary": "List Types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Types"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "type"
                ],
                "summary": "Creates Type",
                "parameters": [
                    {
                        "description": "The new type",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Type"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Type"
                        }
                    }
                }
            }
        },
        "/type/{id}": {
            "get": {
                "description": "Get type by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "type"
                ],
                "summary": "Get Type",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Type ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Type"
                        }
                    }
                }
            },
            "put": {
                "description": "Update type by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "type"
                ],
                "summary": "Update Type",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Type ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The updated type",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Type"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Type"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete type by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "type"
                ],
                "summary": "Delete Type",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Type ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Evolution": {
            "type": "object",
            "properties": {
                "evolution_requirement": {
                    "$ref": "#/definitions/models.EvolutionRequirement"
                },
                "evolution_requirement_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "level_requirement": {
                    "type": "integer"
                },
                "pokemon_id": {
                    "type": "integer"
                },
                "target_pokemon_id": {
                    "type": "integer"
                }
            }
        },
        "models.EvolutionRequirement": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Evolutions": {
            "type": "object",
            "properties": {
                "evolutions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Evolution"
                    }
                }
            }
        },
        "models.Move": {
            "type": "object",
            "properties": {
                "damage": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/models.Type"
                },
                "type_id": {
                    "type": "integer"
                }
            }
        },
        "models.Moves": {
            "type": "object",
            "properties": {
                "moves": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Move"
                    }
                }
            }
        },
        "models.Pokemon": {
            "type": "object",
            "properties": {
                "evolutions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Evolution"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "move_set": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Move"
                    }
                },
                "name": {
                    "type": "string"
                },
                "previous_evolution": {
                    "$ref": "#/definitions/models.Evolution"
                },
                "primary_type": {
                    "$ref": "#/definitions/models.Type"
                },
                "primary_type_id_id": {
                    "type": "integer"
                },
                "secondary_type": {
                    "$ref": "#/definitions/models.Type"
                },
                "secondary_type_id": {
                    "type": "integer"
                }
            }
        },
        "models.Pokemons": {
            "type": "object",
            "properties": {
                "pokemon": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Pokemon"
                    }
                }
            }
        },
        "models.Type": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "no_effect": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "not_very_effective": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "super_effective": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "models.Types": {
            "type": "object",
            "properties": {
                "moveTypes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Type"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "PokeDB API",
	Description:      "Pokemon Database API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
