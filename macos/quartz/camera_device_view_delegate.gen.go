// Code generated by DarwinKit. DO NOT EDIT.

package quartz

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The IKCameraDeviceViewDelegate protocol is adopted by the delegate of the IKCameraDeviceView class. It allows downloading of camera content, handling selection changes, and handling errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate?language=objc
type PCameraDeviceViewDelegate interface {
	// optional
	CameraDeviceViewDidEncounterError(cameraDeviceView CameraDeviceView, error foundation.Error)
	HasCameraDeviceViewDidEncounterError() bool

	// optional
	CameraDeviceViewDidDownloadFileLocationFileDataError(cameraDeviceView CameraDeviceView, file objc.Object, url foundation.URL, data []byte, error foundation.Error)
	HasCameraDeviceViewDidDownloadFileLocationFileDataError() bool

	// optional
	CameraDeviceViewSelectionDidChange(cameraDeviceView CameraDeviceView)
	HasCameraDeviceViewSelectionDidChange() bool
}

// A delegate implementation builder for the [PCameraDeviceViewDelegate] protocol.
type CameraDeviceViewDelegate struct {
	_CameraDeviceViewDidEncounterError                    func(cameraDeviceView CameraDeviceView, error foundation.Error)
	_CameraDeviceViewDidDownloadFileLocationFileDataError func(cameraDeviceView CameraDeviceView, file objc.Object, url foundation.URL, data []byte, error foundation.Error)
	_CameraDeviceViewSelectionDidChange                   func(cameraDeviceView CameraDeviceView)
}

func (di *CameraDeviceViewDelegate) HasCameraDeviceViewDidEncounterError() bool {
	return di._CameraDeviceViewDidEncounterError != nil
}

// Invoked when the camera encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505239-cameradeviceview?language=objc
func (di *CameraDeviceViewDelegate) SetCameraDeviceViewDidEncounterError(f func(cameraDeviceView CameraDeviceView, error foundation.Error)) {
	di._CameraDeviceViewDidEncounterError = f
}

// Invoked when the camera encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505239-cameradeviceview?language=objc
func (di *CameraDeviceViewDelegate) CameraDeviceViewDidEncounterError(cameraDeviceView CameraDeviceView, error foundation.Error) {
	di._CameraDeviceViewDidEncounterError(cameraDeviceView, error)
}
func (di *CameraDeviceViewDelegate) HasCameraDeviceViewDidDownloadFileLocationFileDataError() bool {
	return di._CameraDeviceViewDidDownloadFileLocationFileDataError != nil
}

// Invoked for each file that is downloaded from the camera device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1503524-cameradeviceview?language=objc
func (di *CameraDeviceViewDelegate) SetCameraDeviceViewDidDownloadFileLocationFileDataError(f func(cameraDeviceView CameraDeviceView, file objc.Object, url foundation.URL, data []byte, error foundation.Error)) {
	di._CameraDeviceViewDidDownloadFileLocationFileDataError = f
}

// Invoked for each file that is downloaded from the camera device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1503524-cameradeviceview?language=objc
func (di *CameraDeviceViewDelegate) CameraDeviceViewDidDownloadFileLocationFileDataError(cameraDeviceView CameraDeviceView, file objc.Object, url foundation.URL, data []byte, error foundation.Error) {
	di._CameraDeviceViewDidDownloadFileLocationFileDataError(cameraDeviceView, file, url, data, error)
}
func (di *CameraDeviceViewDelegate) HasCameraDeviceViewSelectionDidChange() bool {
	return di._CameraDeviceViewSelectionDidChange != nil
}

// Invoked when the selection changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505308-cameradeviceviewselectiondidchan?language=objc
func (di *CameraDeviceViewDelegate) SetCameraDeviceViewSelectionDidChange(f func(cameraDeviceView CameraDeviceView)) {
	di._CameraDeviceViewSelectionDidChange = f
}

// Invoked when the selection changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505308-cameradeviceviewselectiondidchan?language=objc
func (di *CameraDeviceViewDelegate) CameraDeviceViewSelectionDidChange(cameraDeviceView CameraDeviceView) {
	di._CameraDeviceViewSelectionDidChange(cameraDeviceView)
}

// ensure impl type implements protocol interface
var _ PCameraDeviceViewDelegate = (*CameraDeviceViewDelegateObject)(nil)

// A concrete type for the [PCameraDeviceViewDelegate] protocol.
type CameraDeviceViewDelegateObject struct {
	objc.Object
}

func (c_ CameraDeviceViewDelegateObject) HasCameraDeviceViewDidEncounterError() bool {
	return c_.RespondsToSelector(objc.Sel("cameraDeviceView:didEncounterError:"))
}

// Invoked when the camera encounters an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505239-cameradeviceview?language=objc
func (c_ CameraDeviceViewDelegateObject) CameraDeviceViewDidEncounterError(cameraDeviceView CameraDeviceView, error foundation.Error) {
	objc.Call[objc.Void](c_, objc.Sel("cameraDeviceView:didEncounterError:"), cameraDeviceView, error)
}

func (c_ CameraDeviceViewDelegateObject) HasCameraDeviceViewDidDownloadFileLocationFileDataError() bool {
	return c_.RespondsToSelector(objc.Sel("cameraDeviceView:didDownloadFile:location:fileData:error:"))
}

// Invoked for each file that is downloaded from the camera device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1503524-cameradeviceview?language=objc
func (c_ CameraDeviceViewDelegateObject) CameraDeviceViewDidDownloadFileLocationFileDataError(cameraDeviceView CameraDeviceView, file objc.Object, url foundation.URL, data []byte, error foundation.Error) {
	objc.Call[objc.Void](c_, objc.Sel("cameraDeviceView:didDownloadFile:location:fileData:error:"), cameraDeviceView, file, url, data, error)
}

func (c_ CameraDeviceViewDelegateObject) HasCameraDeviceViewSelectionDidChange() bool {
	return c_.RespondsToSelector(objc.Sel("cameraDeviceViewSelectionDidChange:"))
}

// Invoked when the selection changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikcameradeviceviewdelegate/1505308-cameradeviceviewselectiondidchan?language=objc
func (c_ CameraDeviceViewDelegateObject) CameraDeviceViewSelectionDidChange(cameraDeviceView CameraDeviceView) {
	objc.Call[objc.Void](c_, objc.Sel("cameraDeviceViewSelectionDidChange:"), cameraDeviceView)
}
