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

type PassengerDetail struct {

	// The passenger's first name
	FirstName string `json:"first_name,omitempty"`

	// The passenger's profile image
	ImageUrl string `json:"image_url,omitempty"`

	// The passenger's rating
	Rating string `json:"rating,omitempty"`

	// The passenger's last name
	LastName string `json:"last_name,omitempty"`

	// The passenger's lyft user id
	UserId string `json:"user_id,omitempty"`
}
