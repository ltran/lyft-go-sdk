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

// Types of rides offered via Lyft
type RideTypeEnum string

const (
	RideTypeLyft     RideTypeEnum = "lyft"
	RideTypeLyftLine RideTypeEnum = "lyft_line"
	RideTypeLyftPlus RideTypeEnum = "lyft_plus"
	RideTypeLyftSUV  RideTypeEnum = "lyft_suv"
)
