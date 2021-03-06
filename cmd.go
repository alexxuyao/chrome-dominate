package chromedominate

type NodeId int64
type TimeSinceEpoch float64
type RequestId string
type LoaderId string
type MonotonicTime float64
type ResourceType string // "Document","Stylesheet", "Image", "Media", "Font", "Script", "TextTrack", "XHR", "Fetch", "EventSource", "WebSocket", "Manifest", "SignedExchange", "Ping", "CSPViolationReport", "Other"
type FrameId string
type TargetID string
type BrowserContextID string
type BackendNodeId int64
type RemoteObjectId string
type ScriptIdentifier string
type InterceptionStage string // Request, HeadersReceived
type RequestStage string      // Request, Response
type ErrorReason string       // "Failed"",Aborted"",TimedOut"",AccessDenied"",ConnectionClosed"",ConnectionReset"",ConnectionRefused"",ConnectionAborted"",ConnectionFailed"",NameNotResolved"",InternetDisconnected"",AddressUnreachable"",BlockedByClient"",BlockedByResponse"
type MixedContentType string  // "blockable","optionally-blockable","none"
type ResourcePriority string  // "VeryLow","Low","Medium","High","VeryHigh"

type ChromeTargetType struct {
	Description          string   `json:"description"`
	DevtoolsFrontendUrl  string   `json:"devtoolsFrontendUrl"`
	Id                   TargetID `json:"id"`
	Title                string   `json:"title"`
	Type                 string   `json:"type"`
	Url                  string   `json:"url"`
	WebSocketDebuggerUrl string   `json:"webSocketDebuggerUrl"`
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

type ResultError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ResultRootType struct {
	Id     int64       `json:"id"`
	Result interface{} `json:"result,omitempty"`
	Error  ResultError `json:"error,omitempty"`
}

// Page.navigate
type NavigateParam struct {
	Url            string   `json:"url"`
	Referrer       *string  `json:"referrer,omitempty"`
	TransitionType *string  `json:"transitionType,omitempty"`
	FrameId        *FrameId `json:"frameId,omitempty"`
}

type ResultPageNavigate struct {
	FrameId   FrameId `json:"frameId"`
	LoaderId  *string `json:"loaderId,omitempty"`
	ErrorText *string `json:"errorText,omitempty"`
}

type BackendNode struct {
	NodeType      int64  `json:"nodeType,omitempty"`
	NodeName      string `json:"nodeName,omitempty"`
	BackendNodeId int64  `json:"backendNodeId,omitempty"`
}

type GetRootResult struct {
	Root ResultDOMNode `json:"root"`
}

type ResultDOMNode struct {
	NodeId           NodeId           `json:"nodeId,omitempty"`
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

type CompileScriptParam struct {
	Expression         string `json:"expression"`
	SourceURL          string `json:"sourceURL"`
	PersistScript      bool   `json:"persistScript"`
	ExecutionContextId *int64 `json:"executionContextId,omitempty"`
}

type CompileScriptResult struct {
	ScriptId         *string           `json:"scriptId,omitempty"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type ExceptionDetails struct {
	ExceptionId        int64         `json:"exceptionId"`
	Text               string        `json:"text"`
	LineNumber         int64         `json:"lineNumber"`
	ColumnNumber       int64         `json:"columnNumber"`
	ScriptId           *string       `json:"scriptId,omitempty"`
	Url                *string       `json:"url,omitempty"`
	StackTrace         *StackTrace   `json:"stackTrace,omitempty"`
	Exception          *RemoteObject `json:"exception,omitempty"`
	ExecutionContextId *int64        `json:"executionContextId,omitempty"`
}

type StackTrace struct {
	Description *string       `json:"description,omitempty"`
	CallFrames  []CallFrame   `json:"callFrames"`
	Parent      *StackTrace   `json:"parent,omitempty"`
	ParentId    *StackTraceId `json:"parentId,omitempty"`
}

type StackTraceId struct {
	Id         string  `json:"id"`
	DebuggerId *string `json:"debuggerId,omitempty"`
}

type CallFrame struct {
	CallFrameId      string        `json:"callFrameId"`
	FunctionName     string        `json:"functionName"`
	FunctionLocation *Location     `json:"functionLocation,omitempty"`
	Location         Location      `json:"location"`
	Url              string        `json:"url"`
	ScopeChain       []Scope       `json:"scopeChain"`
	This             RemoteObject  `json:"this"`
	ReturnValue      *RemoteObject `json:"returnValue,omitempty"`
}

type Location struct {
	ScriptId     string `json:"scriptId"`
	LineNumber   int64  `json:"lineNumber"`
	ColumnNumber *int64 `json:"columnNumber,omitempty"`
}

type RemoteObject struct {
	Type                string         `json:"type"`
	Subtype             *string        `json:"subtype,omitempty"`
	ClassName           *string        `json:"className,omitempty"`
	Value               *interface{}   `json:"value,omitempty"`
	UnserializableValue *string        `json:"unserializableValue,omitempty"`
	Description         *string        `json:"description,omitempty"`
	ObjectId            *string        `json:"objectId,omitempty"`
	Preview             *ObjectPreview `json:"preview,omitempty"`
	CustomPreview       *CustomPreview `json:"customPreview,omitempty"`
}

type ObjectPreview struct {
	Type        string            `json:"type"`
	Subtype     *string           `json:"subtype,omitempty"`
	Description *string           `json:"description,omitempty"`
	Overflow    bool              `json:"overflow"`
	Properties  []PropertyPreview `json:"properties"`
	Entries     []EntryPreview    `json:"entries,omitempty"`
}

type PropertyPreview struct {
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Value        *string        `json:"value,omitempty"`
	ValuePreview *ObjectPreview `json:"valuePreview,omitempty"`
	Subtype      *string        `json:"subtype,omitempty"`
}

type EntryPreview struct {
	Key   *ObjectPreview `json:"key,omitempty"`
	Value ObjectPreview  `json:"value"`
}

type CustomPreview struct {
	Header       string  `json:"header"`
	BodyGetterId *string `json:"bodyGetterId,omitempty"`
}

type Scope struct {
	Type          string       `json:"type"`
	Object        RemoteObject `json:"object"`
	Name          *string      `json:"name,omitempty"`
	StartLocation *Location    `json:"startLocation,omitempty"`
	EndLocation   *Location    `json:"endLocation,omitempty"`
}

type RunScriptParam struct {
	ScriptId              string  `json:"scriptId"`
	ExecutionContextId    *int64  `json:"executionContextId,omitempty"`
	ObjectGroup           *string `json:"objectGroup,omitempty"`
	Silent                *bool   `json:"silent,omitempty"`
	IncludeCommandLineAPI *bool   `json:"includeCommandLineAPI,omitempty"`
	ReturnByValue         *bool   `json:"returnByValue,omitempty"`
	GeneratePreview       *bool   `json:"generatePreview,omitempty"`
	AwaitPromise          *bool   `json:"awaitPromise,omitempty"`
}

type RunScriptResult struct {
	Result           RemoteObject      `json:"result"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type EvaluateParam struct {
	Expression            string  `json:"expression"`
	ObjectGroup           *string `json:"objectGroup,omitempty"`
	IncludeCommandLineAPI *bool   `json:"includeCommandLineAPI,omitempty"`
	Silent                *bool   `json:"silent,omitempty"`
	ContextId             *int64  `json:"contextId,omitempty"`
	ReturnByValue         *bool   `json:"returnByValue,omitempty"`
	GeneratePreview       *bool   `json:"generatePreview,omitempty"`
	UserGesture           *bool   `json:"userGesture,omitempty"`
	AwaitPromise          *bool   `json:"awaitPromise,omitempty"`
	ThrowOnSideEffect     *bool   `json:"throwOnSideEffect,omitempty"`
	Timeout               *int64  `json:"timeout,omitempty"`
}

type EvaluateResult struct {
	Result           RemoteObject      `json:"result"`
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`
}

type QuerySelectorParam struct {
	NodeId   NodeId `json:"nodeId"`
	Selector string `json:"selector"`
}

type QuerySelectorResult struct {
	NodeId NodeId `json:"nodeId"`
}

type GetContentQuadsParam struct {
	NodeId        *NodeId `json:"nodeId,omitempty"`
	BackendNodeId *int64  `json:"backendNodeId,omitempty"`
	ObjectId      *string `json:"objectId,omitempty"`
}

type GetBoxModelParam struct {
	NodeId        *NodeId `json:"nodeId,omitempty"`
	BackendNodeId *int64  `json:"backendNodeId,omitempty"`
	ObjectId      *string `json:"objectId,omitempty"`
}

type GetOuterHTMLParam struct {
	NodeId        *NodeId         `json:"nodeId,omitempty"`
	BackendNodeId *BackendNodeId  `json:"backendNodeId,omitempty"`
	ObjectId      *RemoteObjectId `json:"objectId,omitempty"`
}

type GetOuterHTMLResult struct {
	OuterHTML string `json:"outerHTML"`
}

type Quad []float64

type GetContentQuadsResult struct {
	Quads []Quad `json:"quads"`
}

type DispatchMouseEventParam struct {
	Type        string   `json:"type"`
	X           float64  `json:"x"`
	Y           float64  `json:"y"`
	Modifiers   *int64   `json:"modifiers,omitempty"`
	Timestamp   *float64 `json:"timestamp,omitempty"`
	Button      *string  `json:"button,omitempty"`
	Buttons     *int64   `json:"buttons,omitempty"`
	ClickCount  *int64   `json:"clickCount,omitempty"`
	DeltaX      *float64 `json:"deltaX,omitempty"`
	DeltaY      *float64 `json:"deltaY,omitempty"`
	PointerType *string  `json:"pointerType,omitempty"`
}

type GetBoxModelResult struct {
	Model BoxModel `json:"model"`
}

type BoxModel struct {
	Content      Quad              `json:"content"`
	Padding      Quad              `json:"padding"`
	Border       Quad              `json:"border"`
	Margin       Quad              `json:"margin"`
	Width        int64             `json:"width"`
	Height       int64             `json:"height"`
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}

type ShapeOutsideInfo struct {
	Bounds      Quad          `json:"bounds"`
	Shape       []interface{} `json:"shape"`
	MarginShape []interface{} `json:"marginShape"`
}

type SetAttributeValueParam struct {
	NodeId NodeId `json:"nodeId"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type DispatchKeyEventParam struct {
	Type                  string          `json:"type"`
	Modifiers             *int64          `json:"modifiers,omitempty"` // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp             *TimeSinceEpoch `json:"timestamp,omitempty"`
	Text                  *string         `json:"text,omitempty"`
	UnmodifiedText        *string         `json:"unmodifiedText,omitempty"`
	KeyIdentifier         *string         `json:"keyIdentifier,omitempty"`
	Code                  *string         `json:"code,omitempty"`
	Key                   *string         `json:"key,omitempty"`
	WindowsVirtualKeyCode *int64          `json:"windowsVirtualKeyCode,omitempty"`
	NativeVirtualKeyCode  *int64          `json:"nativeVirtualKeyCode,omitempty"`
	AutoRepeat            *bool           `json:"autoRepeat,omitempty"`
	IsKeypad              *bool           `json:"isKeypad,omitempty"`
	IsSystemKey           *bool           `json:"isSystemKey,omitempty"`
	Location              *int64          `json:"location,omitempty"`
}

type ResponseReceivedParam struct {
	RequestId RequestId     `json:"requestId"`
	LoaderId  LoaderId      `json:"loaderId"`
	Timestamp MonotonicTime `json:"timestamp"`
	Type      ResourceType  `json:"type"`
	Response  Response      `json:"response"`
	FrameId   *FrameId      `json:"frameId,omitempty"`
}

type GetResponseBodyParam struct {
	RequestId RequestId `json:"requestId"`
}

type GetResponseBodyResult struct {
	Body          string `json:"body"`
	Base64Encoded bool   `json:"base64Encoded"`
}

type GetTargetInfoResult struct {
	TargetInfo TargetInfo `json:"targetInfo"`
}

type TargetInfo struct {
	TargetId         TargetID          `json:"targetId"`
	Type             string            `json:"type"`
	Title            string            `json:"title"`
	Url              string            `json:"url"`
	Attached         bool              `json:"attached"`
	OpenerId         *TargetID         `json:"openerId,omitempty"`
	BrowserContextId *BrowserContextID `json:"browserContextId,omitempty"`
}

type Frame struct {
	Id             FrameId  `json:"id"`
	ParentId       *string  `json:"parentId,omitempty"`
	LoaderId       LoaderId `json:"loaderId"`
	Name           *string  `json:"name,omitempty"`
	Url            string   `json:"url"`
	UrlFragment    *string  `json:"urlFragment,omitempty"`
	SecurityOrigin string   `json:"securityOrigin"`
	MimeType       string   `json:"mimeType"`
	UnreachableUrl *string  `json:"unreachableUrl,omitempty"`
}

type FrameNavigatedParam struct {
	Frame Frame `json:"frame"`
}

type AddScriptToEvaluateOnNewDocumentParam struct {
	Source    string  `json:"source"`
	WorldName *string `json:"worldName,omitempty"`
}

type AddScriptToEvaluateOnNewDocumentResult struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

type CanEmulateResult struct {
	Result bool `json:"result"`
}

type SetUserAgentOverrideParam struct {
	UserAgent      string  `json:"userAgent"`
	AcceptLanguage *string `json:"acceptLanguage,omitempty"`
	Platform       *string `json:"platform,omitempty"`
}

type RequestPattern struct {
	UrlPattern        *string            `json:"urlPattern,omitempty"`
	ResourceType      *ResourceType      `json:"resourceType,omitempty"`
	InterceptionStage *InterceptionStage `json:"interceptionStage,omitempty"` // 新版本
	RequestStage      *RequestStage      `json:"requestStage,omitempty"`      // 老版本
}

type FetchEnableParam struct {
	Patterns           []RequestPattern `json:"patterns,omitempty"`
	HandleAuthRequests bool             `json:"handleAuthRequests,omitempty"`
}

type FailRequestParam struct {
	RequestId   RequestId   `json:"requestId"`
	ErrorReason ErrorReason `json:"errorReason"`
}

type FulfillRequestParam struct {
	RequestId             RequestId     `json:"requestId"`
	ResponseCode          int64         `json:"responseCode"`
	ResponseHeaders       []HeaderEntry `json:"responseHeaders"`
	BinaryResponseHeaders *string       `json:"binaryResponseHeaders,omitempty"`
	Body                  *string       `json:"body,omitempty"`
	ResponsePhrase        *string       `json:"responsePhrase,omitempty"`
}

type HeaderEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ContinueRequestParam struct {
	RequestId RequestId     `json:"requestId"`
	Url       *string       `json:"url,omitempty"`
	Method    *string       `json:"method,omitempty"`
	PostData  *string       `json:"postData,omitempty"`
	Headers   []HeaderEntry `json:"headers,omitempty"`
}

type RequestPausedParam struct {
	RequestId           RequestId     `json:"requestId"`
	Request             Request       `json:"request"`
	FrameId             FrameId       `json:"frameId"`
	ResourceType        ResourceType  `json:"resourceType"`
	ResponseErrorReason *ErrorReason  `json:"responseErrorReason,omitempty"`
	ResponseStatusCode  *int64        `json:"responseStatusCode,omitempty"`
	ResponseHeaders     []HeaderEntry `json:"responseHeaders,omitempty"`
	NetworkId           *RequestId    `json:"networkId,omitempty"`
}

type Request struct {
	Url              string            `json:"url"`
	UrlFragment      *string           `json:"urlFragment,omitempty"`
	Method           string            `json:"method"`
	Headers          Headers           `json:"headers"`
	PostData         *string           `json:"postData,omitempty"`
	HasPostData      *bool             `json:"hasPostData,omitempty"`
	MixedContentType *MixedContentType `json:"mixedContentType,omitempty"`
	InitialPriority  ResourcePriority  `json:"initialPriority"`
	ReferrerPolicy   string            `json:"referrerPolicy"`
	IsLinkPreload    *bool             `json:"isLinkPreload,omitempty"`
}
