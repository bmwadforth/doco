package to_be_named

import (
	"bytes"
	"fmt"
)

type ObjectType string
const (
	Catalog ObjectType = "/Catalog"
	Pages   ObjectType = "/Pages"
	Page    ObjectType = "/Page"
)

type ObjectReference struct {
	ObjectNumber uint32
	VersionNumber uint32
}

func (ref *ObjectReference) Format() string {
	return fmt.Sprintf("%d %d R", ref.ObjectNumber, ref.VersionNumber)
}

type Object struct {
	Type ObjectType
}


func WriteObject(objects map[string]interface{}) []byte {
	objectBytes := bytes.NewBuffer(make([]byte, 16))
	objectBytes.Write([]byte("<<"))
	for key, value := range objects {
		objectBytes.Write([]byte(fmt.Sprintf("%s %s", key, value)))
	}
	objectBytes.Write([]byte(">>"))

	return objectBytes.Bytes()
}