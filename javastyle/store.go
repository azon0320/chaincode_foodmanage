package javastyle

import "github.com/hyperledger/fabric/core/chaincode/shim"

/*
 * TODO Avoid Compile in Windows because of the gcc
 * TODO Be sure the code is tagged when compile in Windows
 */
/*
 * No error will be thrown on Java Style Value Returns
 * Generally used on a Released ChainCode
 * Otherwise, use Logger or throw errors directly
 * TODO Add Debug Logger
 */
func GetStateBytesByKey(key string, stub shim.ChaincodeStubInterface) []byte {
	dat, err := stub.GetState(key)
	if err != nil {
		return nil
	}
	return dat
}

func SaveStateBytesByKey(key string, dat []byte, stub shim.ChaincodeStubInterface) bool {
	return stub.PutState(key, dat) == nil
}
