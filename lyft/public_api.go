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

import (
	"net/url"
	"net/http"
	"strings"
	"errors"
	"golang.org/x/net/context"
	"encoding/json"
)

var _ context.Context

type PublicApiService service


// PublicApiService Cost estimates
// Estimate the cost of taking a Lyft between two points. 
//
// @param startLat Latitude of the starting location
// @param startLng Longitude of the starting location
// @param optional (nil or map[string]interface{}) with one or more of:
//     @param "rideType" (string) ID of a ride type
//     @param "endLat" (float64) Latitude of the ending location
//     @param "endLng" (float64) Longitude of the ending location
// @return CostEstimateResponse
func (a PublicApiService) GetCost(startLat float64, startLng float64, localVarOptionals map[string]interface{}) (CostEstimateResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  CostEstimateResponse
	)

	// create path and map variables
	localVarPath := "https://api.lyft.com/v1/cost"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarTempParam, localVarOk := localVarOptionals["rideType"].(string); localVarOk {
		localVarQueryParams.Add("ride_type", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("start_lat", parameterToString(startLat, ""))
	localVarQueryParams.Add("start_lng", parameterToString(startLng, ""))
	if localVarTempParam, localVarOk := localVarOptionals["endLat"].(float64); localVarOk {
		localVarQueryParams.Add("end_lat", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endLng"].(float64); localVarOk {
		localVarQueryParams.Add("end_lng", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	 if err != nil {
		  return successPayload, nil, err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

// PublicApiService Available drivers nearby
// The drivers endpoint returns a list of nearby drivers&#39; lat and lng at a given location. 
//
// @param lat Latitude of a location
// @param lng Longitude of a location
// @return NearbyDriversResponse
func (a PublicApiService) GetDrivers(lat float64, lng float64) (NearbyDriversResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NearbyDriversResponse
	)

	// create path and map variables
	localVarPath := "https://api.lyft.com/v1/drivers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("lat", parameterToString(lat, ""))
	localVarQueryParams.Add("lng", parameterToString(lng, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	 if err != nil {
		  return successPayload, nil, err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

// PublicApiService Pickup ETAs
// The ETA endpoint lets you know how quickly a Lyft driver can come get you 
//
// @param lat Latitude of a location
// @param lng Longitude of a location
// @param optional (nil or map[string]interface{}) with one or more of:
//     @param "destinationLat" (float64) Latitude of destination location
//     @param "destinationLng" (float64) Longitude of destination location
//     @param "rideType" (string) ID of a ride type
// @return EtaEstimateResponse
func (a PublicApiService) GetETA(lat float64, lng float64, localVarOptionals map[string]interface{}) (EtaEstimateResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  EtaEstimateResponse
	)

	// create path and map variables
	localVarPath := "https://api.lyft.com/v1/eta"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("lat", parameterToString(lat, ""))
	localVarQueryParams.Add("lng", parameterToString(lng, ""))
	if localVarTempParam, localVarOk := localVarOptionals["destinationLat"].(float64); localVarOk {
		localVarQueryParams.Add("destination_lat", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destinationLng"].(float64); localVarOk {
		localVarQueryParams.Add("destination_lng", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["rideType"].(string); localVarOk {
		localVarQueryParams.Add("ride_type", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	 if err != nil {
		  return successPayload, nil, err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

// PublicApiService Types of rides
// The ride types endpoint returns information about what kinds of Lyft rides you can request at a given location. 
//
// @param lat Latitude of a location
// @param lng Longitude of a location
// @param optional (nil or map[string]interface{}) with one or more of:
//     @param "rideType" (string) ID of a ride type
// @return RideTypesResponse
func (a PublicApiService) GetRideTypes(lat float64, lng float64, localVarOptionals map[string]interface{}) (RideTypesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  RideTypesResponse
	)

	// create path and map variables
	localVarPath := "https://api.lyft.com/v1/ridetypes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("lat", parameterToString(lat, ""))
	localVarQueryParams.Add("lng", parameterToString(lng, ""))
	if localVarTempParam, localVarOk := localVarOptionals["rideType"].(string); localVarOk {
		localVarQueryParams.Add("ride_type", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	 if err != nil {
		  return successPayload, nil, err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

