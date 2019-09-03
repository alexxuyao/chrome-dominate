package chromedominate

type ChromeTargetType struct {
	Description          string `json:"description"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	Id                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Url                  string `json:"url"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

type CmdRootType struct {
	Id     int64       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type ResultRootType struct {
	Id     int64       `json:"id"`
	Result interface{} `json:"result,omitempty"`
}

// Page.navigate
type CmdPageNavigate struct {
	Url            string  `json:"url"`
	Referrer       *string `json:"referrer,omitempty"`
	TransitionType *string `json:"transitionType,omitempty"`
	FrameId        *string `json:"frameId,omitempty"`
}

type ResultPageNavigate struct {
	FrameId   string  `json:"frameId"`
	LoaderId  *string `json:"loaderId,omitempty"`
	ErrorText *string `json:"errorText,omitempty"`
}

type BackendNode struct {
	NodeType      int64  `json:"nodeType,omitempty"`
	NodeName      string `json:"nodeName,omitempty"`
	BackendNodeId int64  `json:"backendNodeId,omitempty"`
}

type ResultDOMNode struct {
	NodeId           int64            `json:"nodeId,omitempty"`
	ParentId         *int64           `json:"parentId,omitempty"`
	BackendNodeId    int64            `json:"backendNodeId,omitempty"`
	NodeType         int64            `json:"nodeType,omitempty"`
	NodeName         string           `json:"nodeName,omitempty"`
	LocalName        string           `json:"localName,omitempty"`
	NodeValue        string           `json:"nodeValue,omitempty"`
	ChildNodeCount   *int64           `json:"childNodeCount,omitempty"`
	Children         *[]ResultDOMNode `json:"children,omitempty"`
	Attributes       *[]string        `json:"attributes,omitempty"`
	DocumentURL      *string          `json:"documentURL,omitempty"`
	BaseURL          *string          `json:"baseURL,omitempty"`
	PublicId         *string          `json:"publicId,omitempty"`
	SystemId         *string          `json:"systemId,omitempty"`
	InternalSubset   *string          `json:"internalSubset,omitempty"`
	XmlVersion       *string          `json:"xmlVersion,omitempty"`
	Name             *string          `json:"name,omitempty"`
	Value            *string          `json:"value,omitempty"`
	PseudoType       *string          `json:"pseudoType,omitempty"`
	ShadowRootType   *string          `json:"shadowRootType,omitempty"`
	FrameId          *string          `json:"frameId,omitempty"`
	ContentDocument  *ResultDOMNode   `json:"contentDocument,omitempty"`
	ShadowRoots      *[]ResultDOMNode `json:"shadowRoots,omitempty"`
	TemplateContent  *ResultDOMNode   `json:"templateContent,omitempty"`
	PseudoElements   *[]ResultDOMNode `json:"pseudoElements,omitempty"`
	ImportedDocument *ResultDOMNode   `json:"importedDocument,omitempty"`
	DistributedNodes *[]BackendNode   `json:"distributedNodes,omitempty"`
	IsSVG            *bool            `json:"isSVG,omitempty"`
}

type Cookie struct {
	Name     string  `json:"name"`
	Value    string  `json:"value"`
	Domain   string  `json:"domain"`
	Path     string  `json:"path"`
	Expires  float64 `json:"expires"`
	Size     int64   `json:"size"`
	HttpOnly bool    `json:"httpOnly"`
	Secure   bool    `json:"secure"`
	Session  bool    `json:"session"`
	SameSite *string `json:"sameSite,omitempty"` //["Strict", "Lax"]
}

type CookieParam struct {
	Name     string   `json:"name"`
	Value    string   `json:"value"`
	Url      *string  `json:"url,omitempty"`
	Domain   *string  `json:"domain,omitempty"`
	Path     *string  `json:"path,omitempty"`
	Secure   *bool    `json:"secure,omitempty"`
	HttpOnly *bool    `json:"httpOnly,omitempty"`
	SameSite *string  `json:"sameSite,omitempty"` //["Strict", "Lax"]
	Expires  *float64 `json:"expires,omitempty"`
}

type GetCookieResult struct {
	Cookies []Cookie `json:"cookies"`
}

type SetCookieResult struct {
	Success bool `json:"success"`
}
