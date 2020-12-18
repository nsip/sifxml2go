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
// In order to enforce the use of accessors, struct members are contained in private mirror structs
//
package sifxml
