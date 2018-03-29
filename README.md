# sifxml2go
Convert SIF XML to Go structs

Uses the SIF XML definitions in the https://github.com/nsip/specgen_input format, and produces Golang structs, with XML annotations—meaning they can be roundtripped back to XML. Note that the Golang structs ignore enums, and treat all primitive types as strings: they take a purely textual approach to the data. 

The conversion is done with XSLT (`sh run.sh`, output to `/sifxml`). The script downloads the specgen_input repository as input to the conversion, and is currently set up to consume SIF AU objects and types.

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

The `run.sh` script also extracts all examples from specgen, and generates unit testing for the XML struct definitions. To run the unit tests, `cd sifxml; go test`.

The `run.sh` script also converts these XML definitions into GraphQL schemas (output to `/sifgraphql`). However those conversions do not currently deal with type extensions (for example, NameOfRecordType is an extension of BaseNameType, but is not populated in the GraphQL schema. So right now, no names work.)

The project further converts those XML struct definitions into XML structs that can be consumed by the [https://github.com/clbanning/mxj](MXJ) library, which maps between string maps and structs (`sh mxj.sh`, output to `/sifxmlmxj`). These definitions are used by the https://github.com/nsip/nias3 project, which uses the map structures to map SIF/XML documents to triples and back.

The project is not currently set up to generate JSON output from XML structs. (The sifxmlmxj structures are tagged for JSON, but their JSON is specific to the MXJ map structures: a variant of Goessner notation, with `-` prefixing XML attributes. We have had to use `Value` instead of `#text` for values of elements with attributes, because Golang does not permit navigating object paths with hash prefixes.)
