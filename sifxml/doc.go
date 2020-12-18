// sifxml contains types to represent the SIF-AU data model, and accessor methods to manipulate them as objects.
//
// Purpose
//
// Objects in the SIF data model are highly interdependent, and synthesising SIF objects is a complex activity
// involving iterating through the various dependencies involved. The SIF objects generated have to comply with
// the SIF schema, including respecting codesets where applicable. This package ensures that SIF objects are
// populated consistently for data generation, and can be output as well-formed SIF XML and SIF JSON.
// (The JSON output follows the PESC convention, rather than the Goessner convention, and is idiomatic
// JSON.)
//
// The package has been created for NSIP to populate its HITS test harness with the objects users need
// to consume, in order to confirm their compliance with the various test cases it supports. However it can be used
// by anyone wishing to generate SIF objects.
//
// Provenance
//
// The contents of the sifxml package are automatically generated from the specgen files specifying the data model
// for the current release of the SIF-AU data model, and are updated as the specification itself is updated.
// Updating sifxml to reflect the current SIF-AU data model is the responsibility of the NSIP Program.
// The current specification of SIF-AU may be found at http://specification.sifassociation.org/Implementation/AU/
//
// Because it relies on specgen input, the code generating this package can also be run in other locales using
// the same input format.
//
// Approach
//
// Like most XML and JSON schemas, SIF contains a large number of highly nested structures, and most of its elements
// are optional. In order for XML and JSON to be generated reliably, without spurious empty tags, Go relies on the
// use of pointers to structs and types.
//
// In order to ensure the consistent use of types, without forcing developers to remember to use pointers,
// the package is highly object oriented, and elements are only accessed via accessor methods.
//
// There is a huge number of types (because of the lack of Go generics),
// and their contents can only be inspected in the source code
// (in order to enforce the use of accessors, struct members are contained in private mirror structs).
// In order to keep documentation manageable, we summarise the structure of the code here.
//
// Types
//
// sifxml primitive values are of four types: String, Int, Float, and Bool. These are local aliases of the go primitive types
// string, int, float64, and bool. However, string routinely appears
// aliased to other types, in order to reflect aliasing in the underlying XML structure, or to enforce codeset restrictions.
// string aliases representing codesets are termed codesets here. The accessors take care of any necessary type conversions.
//
// All elements contain either primitive values (primitives), nested structures (containers), or lists of
// either primitives or containers. All objects are containers. Containers and lists both use structs.
// In order to enforce the use of accessors, struct members are contained in private mirror structs.
//
// Because of the need for pointers to guarantee omitempty behaviour, primitives are realised as
// pointers to primitive types, containers as pointers to structs (containing hidden structs with
// the actual struct members), and lists as pointers to structs (containing a hidden struct with a slice of the list elements).
// The accessors make sure that developers do not need to use any pointers, outside of object declarations.
//
// For example, the following is the definition of the AgggregateCharacteristicInfo object:
//
//  type AggregateCharacteristicInfo struct {
//          aggregatecharacteristicinfo `xml:"AggregateCharacteristicInfo" json:"AggregateCharacteristicInfo"`
//  }
//
//  type aggregatecharacteristicinfo struct {
//          RefId                *RefIdType                `xml:"RefId,attr" json:"RefId"`
//          Description          *String                   `xml:"Description,omitempty" json:"Description,omitempty"`
//          Definition           *String                   `xml:"Definition" json:"Definition"`
//          ElementName          *String                   `xml:"ElementName,omitempty" json:"ElementName,omitempty"`
//          LocalCodeList        *LocalCodeListType        `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
//          SIF_Metadata         *SIF_MetadataType         `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
//          SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
//  }
//
// RefIdType and String are both aliases of string. RefId, Description, Definition, and ElementName are all primitives.
//
// LocalCodeListType is a list:
//
//  type LocalCodeListType struct {
//          localcodelisttype `xml:"LocalCodeListType" json:"LocalCodeListType"`
//  }
//
//  type localcodelisttype struct {
//          LocalCode []LocalCodeType `xml:"LocalCode" json:"LocalCode"`
//  }
//
// LocalCode, as well as SIF_Metadata and SIF_ExtendedElements, are containers:
//
//  type LocalCodeType struct {
//  	localcodetype `xml:"LocalCodeType" json:"LocalCodeType"`
//  }
//
//  type localcodetype struct {
//  	LocalisedCode *String `xml:"LocalisedCode" json:"LocalisedCode"`
//  	Description   *String `xml:"Description,omitempty" json:"Description,omitempty"`
//  	Element       *String `xml:"Element,omitempty" json:"Element,omitempty"`
//  	ListIndex     *Int    `xml:"ListIndex,omitempty" json:"ListIndex,omitempty"`
//  }
//
// In the following definitions, X stands in for an applicable type name.
//
// Objects
//
// sifxml.XSlice: creates a slice of pointers to the object type.
//
//  func ActivitySlice() []*Activity
//
// String
//
// X.String: Returns string value
//
//   func (t *CountryType) String() string
//
// Codeset
//
// Each codeset type X has two associated variables: X_values, a slice of all allowed values of the codeset,
// and X_map, a map of each allowed codeset value to struct{}{} (used to determine efficiently whether a value is allowed).
// For example:
//
//  type AUCodeSetsWellbeingCharacteristicClassificationType string
//
//  var AUCodeSetsWellbeingCharacteristicClassificationType_values = []string{
//      "M", "D", "S", "O", "DP",
//  }
//
//  var AUCodeSetsWellbeingCharacteristicClassificationType_map = map[string]struct{}{
//      "M":  struct{}{},
//      "D":  struct{}{},
//      "S":  struct{}{},
//      "O":  struct{}{},
//      "DP": struct{}{},
//  }
//
// sifxml.CodesetContains: Uses the codeset map to determine whether a string is allowed for a codeset.
//
//   func CodesetContains(codeset map[string]struct{}, value interface{}) bool
//
// Int
//
// X.Int: Returns int value
//
//   func (t *Int) Int()
//
// Float
//
// X.Float: Returns float64 value
//
//   func (t *Float) Float() float64
//
// Bool
//
// X.Bool: Returns bool value
//
//   func (t *Bool) Bool() bool
//
// List
//
// X.Append: Appends value to the list. Creates list if it is empty. Aborts if the list is a list of codeset values,
// and the value does not match the codeset.
//
//   func (t *SIF_MetadataType_TimeElements) Append(value TimeElementType) *SIF_MetadataType_TimeElements
//
// X.AddNew: Appends an empty value to the list. This value can then be populated through accessors on Last().
//
//  func (t *SIF_MetadataType_TimeElements) AddNew() *SIF_MetadataType_TimeElements
//
// X.Last: Retrieve the last value of the list. Calls AddNew() if the list is empty.
//
//  func (t *SIF_MetadataType_TimeElements) Last()
//
// X.Index: Retrieves the nth value in the list. Raises error if index is out of bounds.
//
//   func (t *SIF_MetadataType_TimeElements) Index(n int) (*TimeElementType, error)
//
// X.Len: Length of the list.
//
//   func (t *SIF_MetadataType_TimeElements) Len() int
//
// X.AppendString: Append a single string to the list. Only defined for lists of strings or of types aliased to string.
//
//   func (t *LearningStandardsType) AppendString(value string) *LearningStandardsType
//
// Container
//
// In the following, Y is the name of an element of the container.
//
// X.Y(): Return the element value (as a pointer to the container, list, or primitive representing it).
//
// NOTE: This function creates an empty element value if it is called on a nil element, so that it can be used to populate that element. You should use X.Y_IsNil() to confirm the element is not nil, before calling any functions which should leave it alone if it is nil --- for example, X.Y().Clone()
//
//   func (s *NAPWritingRubricType) ScoreList() *ScoreListType
//
// X.Y_IsNil(): Returns whether the element value is nil in that container.
//
//   func (s *NAPWritingRubricType) ScoreList_IsNil() bool
//
// X.Unset: Set the value of a property to nil
//
//   func (n *NAPWritingRubricType) Unset(key string) *NAPWritingRubricType
//
// X.SetProperty: Set a property to a value. Aborts if property name is undefined for the type. Aborts if the list is a list of codeset values,
// and the value does not match the codeset.
//
//   func (n *NAPWritingRubricType) SetProperty(key string, value interface{}) *NAPWritingRubricType
//
// X.SetProperties: Set a sequence of properties
//
//   func (n *NAPWritingRubricType) SetProperties(props ...Prop) *NAPWritingRubricType
//
// All Types
//
// X.Clone(): Performs a deep clone on the type, and is used to duplicate an element into another container (particularly if the element is itself nested)
//
//   func (t *MediumOfInstructionType) Clone() (*MediumOfInstructionType)
//
// sifxml.XPointer: Generates a pointer to the given value (unless it already is a pointer), and returns an error in case
// the value mismatches X. In the case of aliased types, accepts primitive values and converts them to the required alias.
// Also deals with both Float32 and Float64 values.
//
//   func MediumOfInstructionTypePointer(value interface{}) (*MediumOfInstructionType, bool)
//
//
//
package sifxml
