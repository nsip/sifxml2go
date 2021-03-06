# sifxml2go
Convert SIF XML specification to golang code for generating and populating SIF objects

Code consists of two packages, sifxml and populate

* `sifxml`, code representing SIF objects for the full SIF-AU specification, allowing it to be ingested and output in XML and JSON
* `populate`, using `sifxml` code to populate SIF objects with plausible values for all objects required in the NSIP HITS test harness

## sifxml

[![GoDoc](https://godoc.org/github.com/nsip/sifxml2go/sifxml?status.svg)](https://godoc.org/github.com/nsip/sifxml2go/sifxml)

The code in `sifxml` is autogenerated, and because of the lack of golang generics, it is extensive. The scripts in `scripts/`, executed by `sh run.sh`, download the SIF XML specification from https://github.com/nsip/specgen_input repository, and use it to generate the code in sifxml.

* The structs are annotated to generate conformant SIF XML and SIF PESC JSON. (For more on SIF JSON, see https://github.com/nsip/sifxml2pescjson.)
* The structs process the (very few) Russian Doll structures that are still in place in the Australian SIF Data model. 
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

The `run.sh` script also generates unit testing for the XML struct definitions, by round-tripping the generated XML against XML read in from all XML examples in the SIF-AU specification. To run the unit tests, `cd sifxml; go test`.

There is also currently code to generate GraphQL schemas (output to `/sifgraphql`). However those conversions do not currently deal with type extensions (for example, NameOfRecordType is an extension of BaseNameType, but is not populated in the GraphQL schema. So right now, no names work), and they have not been updated in line with the major refactoring of code undertaken late 2020. So it is not production-ready.

## populate

[![GoDoc](https://godoc.org/github.com/nsip/sifxml2go/populate?status.svg)](https://godoc.org/github.com/nsip/sifxml2go/populate)

The code for populating objects is currently invoked as `MakeUsecaseObjects()` in usecases.go
