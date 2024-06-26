// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a sunbeams generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator?language=objc
type PSunbeamsGenerator interface {
	// optional
	SetStriationContrast(value float32)
	HasSetStriationContrast() bool

	// optional
	StriationContrast() float32
	HasStriationContrast() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetSunRadius(value float32)
	HasSetSunRadius() bool

	// optional
	SunRadius() float32
	HasSunRadius() bool

	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() Color
	HasColor() bool

	// optional
	SetTime(value float32)
	HasSetTime() bool

	// optional
	Time() float32
	HasTime() bool

	// optional
	SetMaxStriationRadius(value float32)
	HasSetMaxStriationRadius() bool

	// optional
	MaxStriationRadius() float32
	HasMaxStriationRadius() bool

	// optional
	SetStriationStrength(value float32)
	HasSetStriationStrength() bool

	// optional
	StriationStrength() float32
	HasStriationStrength() bool
}

// ensure impl type implements protocol interface
var _ PSunbeamsGenerator = (*SunbeamsGeneratorObject)(nil)

// A concrete type for the [PSunbeamsGenerator] protocol.
type SunbeamsGeneratorObject struct {
	objc.Object
}

func (s_ SunbeamsGeneratorObject) HasSetStriationContrast() bool {
	return s_.RespondsToSelector(objc.Sel("setStriationContrast:"))
}

// The contrast of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228770-striationcontrast?language=objc
func (s_ SunbeamsGeneratorObject) SetStriationContrast(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setStriationContrast:"), value)
}

func (s_ SunbeamsGeneratorObject) HasStriationContrast() bool {
	return s_.RespondsToSelector(objc.Sel("striationContrast"))
}

// The contrast of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228770-striationcontrast?language=objc
func (s_ SunbeamsGeneratorObject) StriationContrast() float32 {
	rv := objc.Call[float32](s_, objc.Sel("striationContrast"))
	return rv
}

func (s_ SunbeamsGeneratorObject) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the sunbeam pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228767-center?language=objc
func (s_ SunbeamsGeneratorObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ SunbeamsGeneratorObject) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the sunbeam pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228767-center?language=objc
func (s_ SunbeamsGeneratorObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}

func (s_ SunbeamsGeneratorObject) HasSetSunRadius() bool {
	return s_.RespondsToSelector(objc.Sel("setSunRadius:"))
}

// The radius of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228772-sunradius?language=objc
func (s_ SunbeamsGeneratorObject) SetSunRadius(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setSunRadius:"), value)
}

func (s_ SunbeamsGeneratorObject) HasSunRadius() bool {
	return s_.RespondsToSelector(objc.Sel("sunRadius"))
}

// The radius of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228772-sunradius?language=objc
func (s_ SunbeamsGeneratorObject) SunRadius() float32 {
	rv := objc.Call[float32](s_, objc.Sel("sunRadius"))
	return rv
}

func (s_ SunbeamsGeneratorObject) HasSetColor() bool {
	return s_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228768-color?language=objc
func (s_ SunbeamsGeneratorObject) SetColor(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setColor:"), value)
}

func (s_ SunbeamsGeneratorObject) HasColor() bool {
	return s_.RespondsToSelector(objc.Sel("color"))
}

// The color of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228768-color?language=objc
func (s_ SunbeamsGeneratorObject) Color() Color {
	rv := objc.Call[Color](s_, objc.Sel("color"))
	return rv
}

func (s_ SunbeamsGeneratorObject) HasSetTime() bool {
	return s_.RespondsToSelector(objc.Sel("setTime:"))
}

// The duration of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228773-time?language=objc
func (s_ SunbeamsGeneratorObject) SetTime(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setTime:"), value)
}

func (s_ SunbeamsGeneratorObject) HasTime() bool {
	return s_.RespondsToSelector(objc.Sel("time"))
}

// The duration of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228773-time?language=objc
func (s_ SunbeamsGeneratorObject) Time() float32 {
	rv := objc.Call[float32](s_, objc.Sel("time"))
	return rv
}

func (s_ SunbeamsGeneratorObject) HasSetMaxStriationRadius() bool {
	return s_.RespondsToSelector(objc.Sel("setMaxStriationRadius:"))
}

// The radius of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228769-maxstriationradius?language=objc
func (s_ SunbeamsGeneratorObject) SetMaxStriationRadius(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxStriationRadius:"), value)
}

func (s_ SunbeamsGeneratorObject) HasMaxStriationRadius() bool {
	return s_.RespondsToSelector(objc.Sel("maxStriationRadius"))
}

// The radius of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228769-maxstriationradius?language=objc
func (s_ SunbeamsGeneratorObject) MaxStriationRadius() float32 {
	rv := objc.Call[float32](s_, objc.Sel("maxStriationRadius"))
	return rv
}

func (s_ SunbeamsGeneratorObject) HasSetStriationStrength() bool {
	return s_.RespondsToSelector(objc.Sel("setStriationStrength:"))
}

// The intensity of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228771-striationstrength?language=objc
func (s_ SunbeamsGeneratorObject) SetStriationStrength(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setStriationStrength:"), value)
}

func (s_ SunbeamsGeneratorObject) HasStriationStrength() bool {
	return s_.RespondsToSelector(objc.Sel("striationStrength"))
}

// The intensity of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228771-striationstrength?language=objc
func (s_ SunbeamsGeneratorObject) StriationStrength() float32 {
	rv := objc.Call[float32](s_, objc.Sel("striationStrength"))
	return rv
}
