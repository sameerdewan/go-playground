package exercise10

import (
	"encoding/json"
	"fmt"
)

type Permissions struct {
	Read  bool `json:"Read"`
	Write bool `json:"Write"`
}

type Role struct {
	Id          int         `json:"Id"`
	Permissions Permissions `json:"Permissions"`
}

func createAdmin() string {
	permissions := Permissions{Read: true, Write: true}
	admin := Role{Id: 0, Permissions: permissions}
	fmt.Println(admin)
	marshalledJson, err := json.Marshal(admin) // returns byte array
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(marshalledJson) // byte array
	stringifiedJson := string(marshalledJson)
	fmt.Println(stringifiedJson) // byte array to string (stringified json object parallel to javascript)
	return stringifiedJson
}

func Run() {
	stringifiedJsonAdmin := createAdmin()
	returnedToByteArray := []byte(stringifiedJsonAdmin)
	fmt.Println(returnedToByteArray) // same as byte array printed first
	unmarshalledJson := Role{}
	err := json.Unmarshal(returnedToByteArray, &unmarshalledJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(unmarshalledJson) // same as initial logged admin
}
