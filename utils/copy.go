package utils

import "encoding/json"

// CopyJSON copies the JSON data from src to dest.
//
// @param dest the destination to copy the JSON data to
// @param src the source to copy the JSON data from
// @return error if the JSON data cannot be copied
func CopyJSON(dest any, src []byte) error {
	return json.Unmarshal(src, dest)
}
