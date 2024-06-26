// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A protocol you implement to track download progress and handle redirects, authentication challenges, and failures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate?language=objc
type PDownloadDelegate interface {
	// optional
	DownloadDidReceiveAuthenticationChallengeCompletionHandler(download Download, challenge foundation.URLAuthenticationChallenge, completionHandler func(arg0 foundation.URLSessionAuthChallengeDisposition, arg1 foundation.URLCredential))
	HasDownloadDidReceiveAuthenticationChallengeCompletionHandler() bool

	// optional
	DownloadDidFailWithErrorResumeData(download Download, error foundation.Error, resumeData []byte)
	HasDownloadDidFailWithErrorResumeData() bool

	// optional
	DownloadDidFinish(download Download)
	HasDownloadDidFinish() bool

	// optional
	DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download Download, response foundation.URLResponse, suggestedFilename string, completionHandler func(destination foundation.URL))
	HasDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler() bool

	// optional
	DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download Download, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(arg0 DownloadRedirectPolicy))
	HasDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler() bool
}

// A delegate implementation builder for the [PDownloadDelegate] protocol.
type DownloadDelegate struct {
	_DownloadDidReceiveAuthenticationChallengeCompletionHandler               func(download Download, challenge foundation.URLAuthenticationChallenge, completionHandler func(arg0 foundation.URLSessionAuthChallengeDisposition, arg1 foundation.URLCredential))
	_DownloadDidFailWithErrorResumeData                                       func(download Download, error foundation.Error, resumeData []byte)
	_DownloadDidFinish                                                        func(download Download)
	_DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler func(download Download, response foundation.URLResponse, suggestedFilename string, completionHandler func(destination foundation.URL))
	_DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler              func(download Download, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(arg0 DownloadRedirectPolicy))
}

func (di *DownloadDelegate) HasDownloadDidReceiveAuthenticationChallengeCompletionHandler() bool {
	return di._DownloadDidReceiveAuthenticationChallengeCompletionHandler != nil
}

// Asks the delegate to respond to an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727346-download?language=objc
func (di *DownloadDelegate) SetDownloadDidReceiveAuthenticationChallengeCompletionHandler(f func(download Download, challenge foundation.URLAuthenticationChallenge, completionHandler func(arg0 foundation.URLSessionAuthChallengeDisposition, arg1 foundation.URLCredential))) {
	di._DownloadDidReceiveAuthenticationChallengeCompletionHandler = f
}

// Asks the delegate to respond to an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727346-download?language=objc
func (di *DownloadDelegate) DownloadDidReceiveAuthenticationChallengeCompletionHandler(download Download, challenge foundation.URLAuthenticationChallenge, completionHandler func(arg0 foundation.URLSessionAuthChallengeDisposition, arg1 foundation.URLCredential)) {
	di._DownloadDidReceiveAuthenticationChallengeCompletionHandler(download, challenge, completionHandler)
}
func (di *DownloadDelegate) HasDownloadDidFailWithErrorResumeData() bool {
	return di._DownloadDidFailWithErrorResumeData != nil
}

// Tells the delegate that the download failed, with error information and data you can use to restart the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727345-download?language=objc
func (di *DownloadDelegate) SetDownloadDidFailWithErrorResumeData(f func(download Download, error foundation.Error, resumeData []byte)) {
	di._DownloadDidFailWithErrorResumeData = f
}

// Tells the delegate that the download failed, with error information and data you can use to restart the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727345-download?language=objc
func (di *DownloadDelegate) DownloadDidFailWithErrorResumeData(download Download, error foundation.Error, resumeData []byte) {
	di._DownloadDidFailWithErrorResumeData(download, error, resumeData)
}
func (di *DownloadDelegate) HasDownloadDidFinish() bool {
	return di._DownloadDidFinish != nil
}

// Tells the delegate that the download finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727348-downloaddidfinish?language=objc
func (di *DownloadDelegate) SetDownloadDidFinish(f func(download Download)) {
	di._DownloadDidFinish = f
}

// Tells the delegate that the download finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727348-downloaddidfinish?language=objc
func (di *DownloadDelegate) DownloadDidFinish(download Download) {
	di._DownloadDidFinish(download)
}
func (di *DownloadDelegate) HasDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler() bool {
	return di._DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler != nil
}

// Asks the delegate to provide a file destination where the system should write the download data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727344-download?language=objc
func (di *DownloadDelegate) SetDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(f func(download Download, response foundation.URLResponse, suggestedFilename string, completionHandler func(destination foundation.URL))) {
	di._DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler = f
}

// Asks the delegate to provide a file destination where the system should write the download data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727344-download?language=objc
func (di *DownloadDelegate) DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download Download, response foundation.URLResponse, suggestedFilename string, completionHandler func(destination foundation.URL)) {
	di._DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download, response, suggestedFilename, completionHandler)
}
func (di *DownloadDelegate) HasDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler() bool {
	return di._DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler != nil
}

// Asks the delegate to respond to the download’s redirect response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727347-download?language=objc
func (di *DownloadDelegate) SetDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(f func(download Download, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(arg0 DownloadRedirectPolicy))) {
	di._DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler = f
}

// Asks the delegate to respond to the download’s redirect response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727347-download?language=objc
func (di *DownloadDelegate) DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download Download, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(arg0 DownloadRedirectPolicy)) {
	di._DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download, response, request, decisionHandler)
}

// ensure impl type implements protocol interface
var _ PDownloadDelegate = (*DownloadDelegateObject)(nil)

// A concrete type for the [PDownloadDelegate] protocol.
type DownloadDelegateObject struct {
	objc.Object
}

func (d_ DownloadDelegateObject) HasDownloadDidReceiveAuthenticationChallengeCompletionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("download:didReceiveAuthenticationChallenge:completionHandler:"))
}

// Asks the delegate to respond to an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727346-download?language=objc
func (d_ DownloadDelegateObject) DownloadDidReceiveAuthenticationChallengeCompletionHandler(download Download, challenge foundation.URLAuthenticationChallenge, completionHandler func(arg0 foundation.URLSessionAuthChallengeDisposition, arg1 foundation.URLCredential)) {
	objc.Call[objc.Void](d_, objc.Sel("download:didReceiveAuthenticationChallenge:completionHandler:"), download, challenge, completionHandler)
}

func (d_ DownloadDelegateObject) HasDownloadDidFailWithErrorResumeData() bool {
	return d_.RespondsToSelector(objc.Sel("download:didFailWithError:resumeData:"))
}

// Tells the delegate that the download failed, with error information and data you can use to restart the download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727345-download?language=objc
func (d_ DownloadDelegateObject) DownloadDidFailWithErrorResumeData(download Download, error foundation.Error, resumeData []byte) {
	objc.Call[objc.Void](d_, objc.Sel("download:didFailWithError:resumeData:"), download, error, resumeData)
}

func (d_ DownloadDelegateObject) HasDownloadDidFinish() bool {
	return d_.RespondsToSelector(objc.Sel("downloadDidFinish:"))
}

// Tells the delegate that the download finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727348-downloaddidfinish?language=objc
func (d_ DownloadDelegateObject) DownloadDidFinish(download Download) {
	objc.Call[objc.Void](d_, objc.Sel("downloadDidFinish:"), download)
}

func (d_ DownloadDelegateObject) HasDownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("download:decideDestinationUsingResponse:suggestedFilename:completionHandler:"))
}

// Asks the delegate to provide a file destination where the system should write the download data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727344-download?language=objc
func (d_ DownloadDelegateObject) DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download Download, response foundation.URLResponse, suggestedFilename string, completionHandler func(destination foundation.URL)) {
	objc.Call[objc.Void](d_, objc.Sel("download:decideDestinationUsingResponse:suggestedFilename:completionHandler:"), download, response, suggestedFilename, completionHandler)
}

func (d_ DownloadDelegateObject) HasDownloadWillPerformHTTPRedirectionNewRequestDecisionHandler() bool {
	return d_.RespondsToSelector(objc.Sel("download:willPerformHTTPRedirection:newRequest:decisionHandler:"))
}

// Asks the delegate to respond to the download’s redirect response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate/3727347-download?language=objc
func (d_ DownloadDelegateObject) DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download Download, response foundation.HTTPURLResponse, request foundation.URLRequest, decisionHandler func(arg0 DownloadRedirectPolicy)) {
	objc.Call[objc.Void](d_, objc.Sel("download:willPerformHTTPRedirection:newRequest:decisionHandler:"), download, response, request, decisionHandler)
}
