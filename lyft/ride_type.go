/* 
 * Lyft API
 *
 * Drive your app to success with Lyft's API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-support@lyft.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package lyft

type RideType struct {

	RideType RideTypeEnum `json:"ride_type,omitempty"`

	// A human readable description of the ride type
	DisplayName string `json:"display_name,omitempty"`

	// The maximum number of seats available for rides requested with this ride type
	Seats int32 `json:"seats,omitempty"`

	// The URL of an image representing this ride type
	ImageUrl string `json:"image_url,omitempty"`

	PricingDetails PricingDetails `json:"pricing_details,omitempty"`

	ScheduledPricingDetails PricingDetails `json:"scheduled_pricing_details,omitempty"`
}
