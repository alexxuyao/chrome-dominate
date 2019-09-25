package chromedominate

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type TheProperties struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Optional    *bool   `json:"optional,omitempty"`
	Ref         *string `json:"$ref,omitempty"`
}

type TheModel struct {
	Id          string          `json:"id"`
	Type        string          `json:"type"`
	Description string          `json:"description"`
	Properties  []TheProperties `json:"properties"`
}

func upperFirst(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func GenType(m TheModel) {

	tmap := map[string]string{
		"string":  "string",
		"number":  "float64",
		"integer": "int64",
		"boolean": "bool",
	}

	tName := upperFirst(m.Id)
	txt := "type " + tName + " struct{\n"

	for _, v := range m.Properties {

		txt += upperFirst(v.Name)
		txt += " "

		if v.Optional != nil {
			txt += "*"
		}

		if v.Ref != nil {
			txt += upperFirst(*v.Ref)
		} else {
			txt += tmap[v.Type]
		}

		// json
		txt += " "
		txt += "`json:\""
		txt += v.Name
		if v.Optional != nil {
			txt += ",omitempty"
		}

		txt += "\"`\n"

	}

	txt += "}"

	fmt.Println(txt)
}

func TestTypeGen(t *testing.T) {
	txt := `{
                    "id": "ShapeOutsideInfo",
                    "description": "CSS Shape Outside details.",
                    "type": "object",
                    "properties": [
                        {
                            "name": "bounds",
                            "description": "Shape bounds",
                            "$ref": "Quad"
                        },
                        {
                            "name": "shape",
                            "description": "Shape coordinate details",
                            "type": "array",
                            "items": {
                                "type": "any"
                            }
                        },
                        {
                            "name": "marginShape",
                            "description": "Margin shape bounds",
                            "type": "array",
                            "items": {
                                "type": "any"
                            }
                        }
                    ]
                }`

	m := TheModel{}
	err := json.Unmarshal([]byte(txt), &m)
	if err != nil {
		t.Error(err)
		return
	}

	GenType(m)

}
