// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol for loading resources with URL schemes that WebKit doesn't handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemehandler?language=objc
type PURLSchemeHandler interface {
	// optional
	WebViewStartURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskObject)
	HasWebViewStartURLSchemeTask() bool

	// optional
	WebViewStopURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskObject)
	HasWebViewStopURLSchemeTask() bool
}

// ensure impl type implements protocol interface
var _ PURLSchemeHandler = (*URLSchemeHandlerObject)(nil)

// A concrete type for the [PURLSchemeHandler] protocol.
type URLSchemeHandlerObject struct {
	objc.Object
}

func (u_ URLSchemeHandlerObject) HasWebViewStartURLSchemeTask() bool {
	return u_.RespondsToSelector(objc.Sel("webView:startURLSchemeTask:"))
}

// Asks your handler to begin loading the data for the specified resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemehandler/2890834-webview?language=objc
func (u_ URLSchemeHandlerObject) WebViewStartURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskObject) {
	po1 := objc.WrapAsProtocol("WKURLSchemeTask", urlSchemeTask)
	objc.Call[objc.Void](u_, objc.Sel("webView:startURLSchemeTask:"), webView, po1)
}

func (u_ URLSchemeHandlerObject) HasWebViewStopURLSchemeTask() bool {
	return u_.RespondsToSelector(objc.Sel("webView:stopURLSchemeTask:"))
}

// Asks your handler to stop loading the data for the specified resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemehandler/2890835-webview?language=objc
func (u_ URLSchemeHandlerObject) WebViewStopURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskObject) {
	po1 := objc.WrapAsProtocol("WKURLSchemeTask", urlSchemeTask)
	objc.Call[objc.Void](u_, objc.Sel("webView:stopURLSchemeTask:"), webView, po1)
}
