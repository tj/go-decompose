package decompose

import (
	"testing"
)

var result = `{
  "definitions": {
    "Error": {
      "properties": {
        "code": {
          "format": "int32",
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      },
      "required": [
        "code",
        "message"
      ],
      "type": "object"
    },
    "Pet": {
      "properties": {
        "id": {
          "format": "int64",
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "tag": {
          "type": "string"
        }
      },
      "required": [
        "id",
        "name"
      ],
      "type": "object"
    }
  },
  "name": "api",
  "paths": {
    "/pets": {
      "get": {
        "responses": {
          "200": {
            "description": "pet response",
            "headers": {
              "x-expires": {
                "type": "string"
              }
            },
            "schema": {
              "items": {
                "$ref": "#/definitions/Pet"
              },
              "type": "array"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "summary": "finds pets in the system"
      }
    }
  }
}`

var pathsResult = `{
  "/pets": {
    "get": {
      "summary": "finds pets in the system",
      "responses": {
        "200": {
          "description": "pet response",
          "schema": {
            "type": "array",
            "items": {
              "$ref": "#/definitions/Pet"
            }
          },
          "headers": {
            "x-expires": {
              "type": "string"
            }
          }
        },
        "default": {
          "description": "unexpected error",
          "schema": {
            "$ref": "#/definitions/Error"
          }
        }
      }
    }
  }
}`

func TestDecompose_directory(t *testing.T) {
	b, err := Decompose("_fixtures/gateway")

	if err != nil {
		t.Fatal(err)
	}

	if string(b) != result {
		t.Fatalf("got: `%s`", b)
	}
}

func TestDecompose_singleFile(t *testing.T) {
	b, err := Decompose("_fixtures/gateway/paths.json")

	if err != nil {
		t.Fatal(err)
	}

	if string(b) != pathsResult {
		t.Fatalf("got: `%s`", b)
	}
}
