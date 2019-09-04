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

type ChromeMainTargetType struct {
	Browser              string `json:"Browser"`
	ProtocolVersion      string `json:"Protocol-Version"`
	UserAgent            string `json:"User-Agent"`
	V8Version            string `json:"V8-Version"`
	WebKitVersion        string `json:"WebKit-Version"`
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

type NetworkEnableParam struct {
	MaxTotalBufferSize    *int64 `json:"maxTotalBufferSize,omitempty"`
	MaxResourceBufferSize *int64 `json:"maxResourceBufferSize,omitempty"`
	MaxPostDataSize       *int64 `json:"maxPostDataSize,omitempty"`
}

type NetworkResponseReceived struct {
	RequestId string   `json:"requestId"`
	LoaderId  string   `json:"loaderId"`
	Timestamp float64  `json:"timestamp"`
	Type      string   `json:"type"` // Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Ping, CSPViolationReport, Other
	Response  Response `json:"response"`
	FrameId   *string  `json:"frameId,omitempty"`
}

type Headers map[string]string

type ResourceTiming struct {
	RequestTime       float64 `json:"requestTime"`
	ProxyStart        float64 `json:"proxyStart"`
	ProxyEnd          float64 `json:"proxyEnd"`
	DnsStart          float64 `json:"dnsStart"`
	DnsEnd            float64 `json:"dnsEnd"`
	ConnectStart      float64 `json:"connectStart"`
	ConnectEnd        float64 `json:"connectEnd"`
	SslStart          float64 `json:"sslStart"`
	SslEnd            float64 `json:"sslEnd"`
	WorkerStart       float64 `json:"workerStart"`
	WorkerReady       float64 `json:"workerReady"`
	SendStart         float64 `json:"sendStart"`
	SendEnd           float64 `json:"sendEnd"`
	PushStart         float64 `json:"pushStart"`
	PushEnd           float64 `json:"pushEnd"`
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"`
}

type SignedCertificateTimestamp struct {
	Status             string  `json:"status"`
	Origin             string  `json:"origin"`
	LogDescription     string  `json:"logDescription"`
	LogId              string  `json:"logId"`
	Timestamp          float64 `json:"timestamp"`
	HashAlgorithm      string  `json:"hashAlgorithm"`
	SignatureAlgorithm string  `json:"signatureAlgorithm"`
	SignatureData      string  `json:"signatureData"`
}

type SecurityDetails struct {
	Protocol                       string                       `json:"protocol"`
	KeyExchange                    string                       `json:"keyExchange"`
	KeyExchangeGroup               *string                      `json:"keyExchangeGroup,omitempty"`
	Cipher                         string                       `json:"cipher"`
	Mac                            *string                      `json:"mac,omitempty"`
	CertificateId                  int64                        `json:"certificateId"`
	SubjectName                    string                       `json:"subjectName"`
	SanList                        []string                     `json:"sanList"`
	Issuer                         string                       `json:"issuer"`
	ValidFrom                      float64                      `json:"validFrom"`
	ValidTo                        float64                      `json:"validTo"`
	SignedCertificateTimestampList []SignedCertificateTimestamp `json:"signedCertificateTimestampList"`
}

type Response struct {
	Url                string           `json:"url"`
	Status             int64            `json:"status"`
	StatusText         string           `json:"statusText"`
	Headers            Headers          `json:"headers"`
	HeadersText        *string          `json:"headersText,omitempty"`
	MimeType           string           `json:"mimeType"`
	RequestHeaders     *Headers         `json:"requestHeaders,omitempty"`
	RequestHeadersText *string          `json:"requestHeadersText,omitempty"`
	ConnectionReused   bool             `json:"connectionReused"`
	ConnectionId       float64          `json:"connectionId"`
	RemoteIPAddress    *string          `json:"remoteIPAddress,omitempty"`
	RemotePort         *int64           `json:"remotePort,omitempty"`
	FromDiskCache      *bool            `json:"fromDiskCache,omitempty"`
	FromServiceWorker  *bool            `json:"fromServiceWorker,omitempty"`
	EncodedDataLength  *float64         `json:"encodedDataLength,omitempty"`
	Timing             *ResourceTiming  `json:"timing,omitempty"`
	Protocol           *string          `json:"protocol,omitempty"`
	SecurityState      string           `json:"securityState"` //["unknown", "neutral", "insecure", "secure", "info"]
	SecurityDetails    *SecurityDetails `json:"securityDetails,omitempty"`
}
