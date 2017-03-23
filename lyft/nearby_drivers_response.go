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

type NearbyDriversResponse struct {

	NearbyDrivers []NearbyDriversByRideType `json:"nearby_drivers,omitempty"`
}
