// populate contains functions to populate SIF objects from the sifxml package with plausible values.
//
// Purpose
//
// The populate package provides functions to populate SIF XML objects (as specified in the sifxml package)
// with plausible values (objects.go). It provides helper functions to make that population more straightforward (base.go).
// Objects are generated in the first instance to populate the NSIP HITS test harness http://hits.nsip.edu.au/dashboard/
// with input data, for clients to consume as part of use case validation. The code to generate bundles of interrelated
// object (usecases.go) parametrises object generation according to the HITS use cases it fulfils.
//
// The assumptions and default values behind the generation of each object are documented in for each method involved.
package populate
