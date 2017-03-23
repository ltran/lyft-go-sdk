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

type RideLocation struct {

	// The latitude component of a location
	Lat float64 `json:"lat,omitempty"`

	// The longitude component of a location
	Lng float64 `json:"lng,omitempty"`

	// A human readable address at/near the given location
	Address string `json:"address,omitempty"`

	// Estimated seconds for a driver to pickup or reach destination based on ride status
	EtaSeconds int32 `json:"eta_seconds,omitempty"`
}
