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
                    "id": "Frame",
                    "description": "Information about the Frame on the page.",
                    "type": "object",
                    "properties": [
                        {
                            "name": "id",
                            "description": "Frame unique identifier.",
                            "$ref": "FrameId"
                        },
                        {
                            "name": "parentId",
                            "description": "Parent frame identifier.",
                            "optional": true,
                            "type": "string"
                        },
                        {
                            "name": "loaderId",
                            "description": "Identifier of the loader associated with this frame.",
                            "$ref": "Network.LoaderId"
                        },
                        {
                            "name": "name",
                            "description": "Frame's name as specified in the tag.",
                            "optional": true,
                            "type": "string"
                        },
                        {
                            "name": "url",
                            "description": "Frame document's URL without fragment.",
                            "type": "string"
                        },
                        {
                            "name": "urlFragment",
                            "description": "Frame document's URL fragment including the '#'.",
                            "experimental": true,
                            "optional": true,
                            "type": "string"
                        },
                        {
                            "name": "securityOrigin",
                            "description": "Frame document's security origin.",
                            "type": "string"
                        },
                        {
                            "name": "mimeType",
                            "description": "Frame document's mimeType as determined by the browser.",
                            "type": "string"
                        },
                        {
                            "name": "unreachableUrl",
                            "description": "If the frame failed to load, this contains the URL that could not be loaded. Note that unlike url above, this URL may contain a fragment.",
                            "experimental": true,
                            "optional": true,
                            "type": "string"
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
