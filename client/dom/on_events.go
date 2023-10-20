//go:build js && wasm

package dom

type OnEvent string

const (
	// Window Events
	//
	OnAfterPrint   OnEvent = "onafterprint"
	OnBeforePrint  OnEvent = "onbeforeprint"
	OnBeforeUnload OnEvent = "onbeforeunload"
	OnHashChange   OnEvent = "onhashchange"
	Onload         OnEvent = "onload"
	OnMessage      OnEvent = "onmessage"
	OnOffline      OnEvent = "onoffline"
	OnOnline       OnEvent = "ononline"
	OnPageHide     OnEvent = "onpagehide"
	OnPageShow     OnEvent = "onpageshow"
	OnPopState     OnEvent = "onpopstate"
	OnResize       OnEvent = "onresize"
	OnStorage      OnEvent = "onstorage"
	OnUnload       OnEvent = "onunload"
	//
	// Form Events
	//
	OnBlur        OnEvent = "onblur"
	OnChange      OnEvent = "onchange"
	OnContextMenu OnEvent = "oncontextmenu"
	OnFocus       OnEvent = "onfocus"
	OnInput       OnEvent = "oninput"
	OnInvalid     OnEvent = "oninvalid"
	OnReset       OnEvent = "onreset"
	OnSearch      OnEvent = "onsearch"
	OnSelect      OnEvent = "onselect"
	OnSubmit      OnEvent = "onsubmit"
	//
	// Keyboard Events
	//
	OnKeyDown  OnEvent = "onkeydown"
	OnKeyPress OnEvent = "onkeypress"
	OnKeyUp    OnEvent = "onkeyup"
	//
	// Mouse Events
	//
	OnClick       OnEvent = "onclick"
	OnDoubleClick OnEvent = "ondblclick"
	OnMouseDown   OnEvent = "onmousedown"
	OnMouseMove   OnEvent = "onmousemove"
	OnMouseOut    OnEvent = "onmouseout"
	OnMouseOver   OnEvent = "onmouseover"
	OnMouseUp     OnEvent = "onmouseup"
	OnMouseWheel  OnEvent = "onmousewheel"
	OnWheel       OnEvent = "onwheel"
	//
	// Drag Events
	//
	OnDrag      OnEvent = "ondrag"
	OnDragEnd   OnEvent = "ondragend"
	OnDragEnter OnEvent = "ondragenter"
	OnDragLeave OnEvent = "ondragleave"
	OnDragOver  OnEvent = "ondragover"
	OnDragStart OnEvent = "ondragstart"
	OnDrop      OnEvent = "ondrop"
	OnScroll    OnEvent = "onscroll"
	//
	// Clipboard Events
	//
	OnCopy         OnEvent = "oncopy"
	OnCut          OnEvent = "oncut"
	OnPasteOnEvent         = "onpaste"
	//
	// Media Events
	//
	OnAbort          OnEvent = "onabort"
	OnCanPlay        OnEvent = "oncanplay"
	OnCanPlayThrough OnEvent = "oncanplaythrough"
	OnCueChange      OnEvent = "oncuechange"
	OnDurationChange OnEvent = "ondurationchange"
	OnEmptied        OnEvent = "onemptied"
	OnEnded          OnEvent = "onended"
	OnLoadedData     OnEvent = "onloadeddata"
	OnLoadedMetaData OnEvent = "onloadedmetadata"
	OnLoadStart      OnEvent = "onloadstart"
	OnPause          OnEvent = "onpause"
	OnPlay           OnEvent = "onplay"
	OnPlaying        OnEvent = "onplaying"
	OnProgress       OnEvent = "onprogress"
	OnRateChange     OnEvent = "onratechange"
	OnSeeked         OnEvent = "onseeked"
	OnSeeking        OnEvent = "onseeking"
	OnStalled        OnEvent = "onstalled"
	OnSuspended      OnEvent = "onsuspend"
	OnTimeUpdate     OnEvent = "ontimeupdate"
	OnVolumeChange   OnEvent = "onvolumechange"
	OnWaiting        OnEvent = "onwaiting"
	//
	// Misc
	//
	OnToggle OnEvent = "ontoggle"
	OnError  OnEvent = "onerror"
)
