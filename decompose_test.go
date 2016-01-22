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

func TestDecompose(t *testing.T) {
	b, err := Decompose("_fixtures/gateway")

	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if string(b) != result {
		t.Fatalf("got: `%s`", b)
	}
}
