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

type SandboxDriverAvailability struct {

	// The latitude component of a location
	Lat float64 `json:"lat,omitempty"`

	// The longitude component of a location
	Lng float64 `json:"lng,omitempty"`

	// The availability of driver in a region
	DriverAvailability bool `json:"driver_availability"`
}
