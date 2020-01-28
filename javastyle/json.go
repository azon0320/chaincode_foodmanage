package javastyle

import "encoding/json"

/*
 * No error will be thrown on Java Style Value Returns
 * Generally used on a Released ChainCode
 * Otherwise, use Logger or throw errors directly
 * TODO Add Debug Logger
 */
func JsonEncode(v interface{}) []byte {
	dat, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return dat
}
