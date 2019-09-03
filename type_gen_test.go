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
                "id": "CookieParam",
                "type": "object",
                "description": "Cookie parameter object",
                "properties": [
                    { "name": "name", "type": "string", "description": "Cookie name." },
                    { "name": "value", "type": "string", "description": "Cookie value." },
                    { "name": "url", "type": "string", "optional": true, "description": "The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie." },
                    { "name": "domain", "type": "string", "optional": true, "description": "Cookie domain." },
                    { "name": "path", "type": "string", "optional": true, "description": "Cookie path." },
                    { "name": "secure", "type": "boolean", "optional": true, "description": "True if cookie is secure." },
                    { "name": "httpOnly", "type": "boolean", "optional": true, "description": "True if cookie is http-only." },
                    { "name": "sameSite", "$ref": "CookieSameSite", "optional": true, "description": "Cookie SameSite type." },
                    { "name": "expires", "$ref": "TimeSinceEpoch", "optional": true, "description": "Cookie expiration date, session cookie if not set" }
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
