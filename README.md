# sifxml2go
Convert SIF XML to Go structs

Uses the SIF XML definitions in the https://github.com/nsip/specgen_input format, and produces Golang structs, with XML annotations—meaning they can be roundtripped back to XML. Note that the Golang structs ignore enums, and treat all primitive types as strings: they take a purely textual approach to the data. 

The conversion is done with XSLT, with a stateful script doing cleanup in two aspects:

* The structs process the Russian Doll structures, that are still in place in the Australian SIF Data model. 
* The structs have difficulty with XML elements that have both attributes and values; this is a difficulty shared with JSON. The structs use the solution of injecting a "Value" element to represent the value of the node itself, as distinct from the attributes; e.g.

````
OtherLEA        struct {
                SIF_RefObject string `xml:"SIF_RefObject,attr"`
                Value         string `xml:",chardata"`
        } `xml:"OtherLEA"
````

corresponding to

````
<OtherLEA SIF_RefObject="OtherLEA">xxxx-xxxx-xxxx-xxxx</OtherLEA>
````

TODO: Create custom Marshallers and UnMarshallers, to make the Value element transparent (move attributes for such structs to a separately suffixed container).

TODO: Generate GraphQL schemas based on simplification of the Golang structs.
