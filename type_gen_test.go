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
                    "id": "DispatchMouseEventParam",
                    "description": "Scope description.",
                    "type": "object",
                    "properties": [
                          {
                            "name": "type",
                            "description": "Type of the mouse event.",
                            "type": "string",
                            "enum": [
                                "mousePressed",
                                "mouseReleased",
                                "mouseMoved",
                                "mouseWheel"
                            ]
                        },
                        {
                            "name": "x",
                            "description": "X coordinate of the event relative to the main frame's viewport in CSS pixels.",
                            "type": "number"
                        },
                        {
                            "name": "y",
                            "description": "Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to\nthe top of the viewport and Y increases as it proceeds towards the bottom of the viewport.",
                            "type": "number"
                        },
                        {
                            "name": "modifiers",
                            "description": "Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8\n(default: 0).",
                            "optional": true,
                            "type": "integer"
                        },
                        {
                            "name": "timestamp",
                            "description": "Time at which the event occurred.",
                            "optional": true,
                            "$ref": "TimeSinceEpoch"
                        },
                        {
                            "name": "button",
                            "description": "Mouse button (default: \"none\").",
                            "optional": true,
                            "type": "string",
                            "enum": [
                                "none",
                                "left",
                                "middle",
                                "right",
                                "back",
                                "forward"
                            ]
                        },
                        {
                            "name": "buttons",
                            "description": "A number indicating which buttons are pressed on the mouse when a mouse event is triggered.\nLeft=1, Right=2, Middle=4, Back=8, Forward=16, None=0.",
                            "optional": true,
                            "type": "integer"
                        },
                        {
                            "name": "clickCount",
                            "description": "Number of times the mouse button was clicked (default: 0).",
                            "optional": true,
                            "type": "integer"
                        },
                        {
                            "name": "deltaX",
                            "description": "X delta in CSS pixels for mouse wheel event (default: 0).",
                            "optional": true,
                            "type": "number"
                        },
                        {
                            "name": "deltaY",
                            "description": "Y delta in CSS pixels for mouse wheel event (default: 0).",
                            "optional": true,
                            "type": "number"
                        },
                        {
                            "name": "pointerType",
                            "description": "Pointer type (default: \"mouse\").",
                            "optional": true,
                            "type": "string",
                            "enum": [
                                "mouse",
                                "pen"
                            ]
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
