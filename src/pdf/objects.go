package pdf

import "fmt"

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
