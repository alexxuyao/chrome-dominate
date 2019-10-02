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
                    "id": "Request",
                    "description": "HTTP request data.",
                    "type": "object",
                    "properties": [
                        {
                            "name": "url",
                            "description": "Request URL (without fragment).",
                            "type": "string"
                        },
                        {
                            "name": "urlFragment",
                            "description": "Fragment of the requested URL starting with hash, if present.",
                            "optional": true,
                            "type": "string"
                        },
                        {
                            "name": "method",
                            "description": "HTTP request method.",
                            "type": "string"
                        },
                        {
                            "name": "headers",
                            "description": "HTTP request headers.",
                            "$ref": "Headers"
                        },
                        {
                            "name": "postData",
                            "description": "HTTP POST request data.",
                            "optional": true,
                            "type": "string"
                        },
                        {
                            "name": "hasPostData",
                            "description": "True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.",
                            "optional": true,
                            "type": "boolean"
                        },
                        {
                            "name": "mixedContentType",
                            "description": "The mixed content type of the request.",
                            "optional": true,
                            "$ref": "Security.MixedContentType"
                        },
                        {
                            "name": "initialPriority",
                            "description": "Priority of the resource request at the time request is sent.",
                            "$ref": "ResourcePriority"
                        },
                        {
                            "name": "referrerPolicy",
                            "description": "The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/",
                            "type": "string",
                            "enum": [
                                "unsafe-url",
                                "no-referrer-when-downgrade",
                                "no-referrer",
                                "origin",
                                "origin-when-cross-origin",
                                "same-origin",
                                "strict-origin",
                                "strict-origin-when-cross-origin"
                            ]
                        },
                        {
                            "name": "isLinkPreload",
                            "description": "Whether is loaded via link preload.",
                            "optional": true,
                            "type": "boolean"
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
