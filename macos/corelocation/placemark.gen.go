// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/contacts"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Placemark] class.
var PlacemarkClass = _PlacemarkClass{objc.GetClass("CLPlacemark")}

type _PlacemarkClass struct {
	objc.Class
}

// An interface definition for the [Placemark] class.
type IPlacemark interface {
	objc.IObject
	InlandWater() string
	AreasOfInterest() []string
	Ocean() string
	Name() string
	SubLocality() string
	Location() Location
	Locality() string
	PostalCode() string
	ISOcountryCode() string
	TimeZone() foundation.TimeZone
	AdministrativeArea() string
	PostalAddress() contacts.PostalAddress
	SubThoroughfare() string
	Region() Region
	SubAdministrativeArea() string
	Country() string
	Thoroughfare() string
}

// A user-friendly description of a geographic coordinate, often containing the name of the place, its address, and other relevant information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark?language=objc
type Placemark struct {
	objc.Object
}

func PlacemarkFrom(ptr unsafe.Pointer) Placemark {
	return Placemark{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlacemarkClass) PlacemarkWithLocationNamePostalAddress(location ILocation, name string, postalAddress contacts.IPostalAddress) Placemark {
	rv := objc.Call[Placemark](pc, objc.Sel("placemarkWithLocation:name:postalAddress:"), objc.Ptr(location), name, objc.Ptr(postalAddress))
	return rv
}

// Creates and initializes a placemark object using the specified location and address information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/2132103-placemarkwithlocation?language=objc
func Placemark_PlacemarkWithLocationNamePostalAddress(location ILocation, name string, postalAddress contacts.IPostalAddress) Placemark {
	return PlacemarkClass.PlacemarkWithLocationNamePostalAddress(location, name, postalAddress)
}

func (p_ Placemark) InitWithPlacemark(placemark IPlacemark) Placemark {
	rv := objc.Call[Placemark](p_, objc.Sel("initWithPlacemark:"), objc.Ptr(placemark))
	return rv
}

// Initializes and returns a placemark object from another placemark object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423818-initwithplacemark?language=objc
func NewPlacemarkWithPlacemark(placemark IPlacemark) Placemark {
	instance := PlacemarkClass.Alloc().InitWithPlacemark(placemark)
	instance.Autorelease()
	return instance
}

func (pc _PlacemarkClass) Alloc() Placemark {
	rv := objc.Call[Placemark](pc, objc.Sel("alloc"))
	return rv
}

func Placemark_Alloc() Placemark {
	return PlacemarkClass.Alloc()
}

func (pc _PlacemarkClass) New() Placemark {
	rv := objc.Call[Placemark](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlacemark() Placemark {
	return PlacemarkClass.New()
}

func (p_ Placemark) Init() Placemark {
	rv := objc.Call[Placemark](p_, objc.Sel("init"))
	return rv
}

// The name of the inland water body associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423738-inlandwater?language=objc
func (p_ Placemark) InlandWater() string {
	rv := objc.Call[string](p_, objc.Sel("inlandWater"))
	return rv
}

// The relevant areas of interest associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423673-areasofinterest?language=objc
func (p_ Placemark) AreasOfInterest() []string {
	rv := objc.Call[[]string](p_, objc.Sel("areasOfInterest"))
	return rv
}

// The name of the ocean associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423619-ocean?language=objc
func (p_ Placemark) Ocean() string {
	rv := objc.Call[string](p_, objc.Sel("ocean"))
	return rv
}

// The name of the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423634-name?language=objc
func (p_ Placemark) Name() string {
	rv := objc.Call[string](p_, objc.Sel("name"))
	return rv
}

// Additional city-level information for the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423794-sublocality?language=objc
func (p_ Placemark) SubLocality() string {
	rv := objc.Call[string](p_, objc.Sel("subLocality"))
	return rv
}

// The location object containing latitude and longitude information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423603-location?language=objc
func (p_ Placemark) Location() Location {
	rv := objc.Call[Location](p_, objc.Sel("location"))
	return rv
}

// The city associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423507-locality?language=objc
func (p_ Placemark) Locality() string {
	rv := objc.Call[string](p_, objc.Sel("locality"))
	return rv
}

// The postal code associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423851-postalcode?language=objc
func (p_ Placemark) PostalCode() string {
	rv := objc.Call[string](p_, objc.Sel("postalCode"))
	return rv
}

// The abbreviated country or region name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423796-isocountrycode?language=objc
func (p_ Placemark) ISOcountryCode() string {
	rv := objc.Call[string](p_, objc.Sel("ISOcountryCode"))
	return rv
}

// The time zone associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423707-timezone?language=objc
func (p_ Placemark) TimeZone() foundation.TimeZone {
	rv := objc.Call[foundation.TimeZone](p_, objc.Sel("timeZone"))
	return rv
}

// The state or province associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423628-administrativearea?language=objc
func (p_ Placemark) AdministrativeArea() string {
	rv := objc.Call[string](p_, objc.Sel("administrativeArea"))
	return rv
}

// The postal address associated with the location, formatted for use with the Contacts framework. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/2890749-postaladdress?language=objc
func (p_ Placemark) PostalAddress() contacts.PostalAddress {
	rv := objc.Call[contacts.PostalAddress](p_, objc.Sel("postalAddress"))
	return rv
}

// Additional street-level information for the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423782-subthoroughfare?language=objc
func (p_ Placemark) SubThoroughfare() string {
	rv := objc.Call[string](p_, objc.Sel("subThoroughfare"))
	return rv
}

// The geographic region associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423808-region?language=objc
func (p_ Placemark) Region() Region {
	rv := objc.Call[Region](p_, objc.Sel("region"))
	return rv
}

// Additional administrative area information for the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423776-subadministrativearea?language=objc
func (p_ Placemark) SubAdministrativeArea() string {
	rv := objc.Call[string](p_, objc.Sel("subAdministrativeArea"))
	return rv
}

// The name of the country or region associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423800-country?language=objc
func (p_ Placemark) Country() string {
	rv := objc.Call[string](p_, objc.Sel("country"))
	return rv
}

// The street address associated with the placemark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clplacemark/1423814-thoroughfare?language=objc
func (p_ Placemark) Thoroughfare() string {
	rv := objc.Call[string](p_, objc.Sel("thoroughfare"))
	return rv
}