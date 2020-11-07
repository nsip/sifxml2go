print <<"END";
package sifxml

import (
        "encoding/xml"
        "github.com/clbanning/mxj"
        "github.com/creasty/defaults"
        "reflect"
        "regexp"
        "strings"
        "testing"
)

var emptytag1 = regexp.MustCompile(`(?s:\\s*<[^>/]+></[^>]+>\\s*)`)
var emptytag2 = regexp.MustCompile(`(?s:\\s*<[^>/]+/>\\s*)`)
var emptytag3 = regexp.MustCompile(`(?s:\\s+[^>='" ]+=(''|""))`)
var pointzero1 = regexp.MustCompile(`(?s:>(\\d*)\\.0+<)`)
var pointzero2 = regexp.MustCompile(`(?s:>(\\d*\\.\\d+?)0+<)`)
var noleadpoint = regexp.MustCompile(`(?s:>(\\.\\d+)<)`)

func stripEmptyTags(s []byte) []byte {
        s = pointzero1.ReplaceAll(s, []byte(">\$1<"))
        s = pointzero2.ReplaceAll(s, []byte(">\$1<"))
        s = noleadpoint.ReplaceAll(s, []byte(">0\$1<"))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag1.ReplaceAll(s, []byte(""))
        s = emptytag2.ReplaceAll(s, []byte(""))
        arr := strings.SplitAfter(string(s), ">")
        for i, _ := range arr {
          if strings.HasPrefix(arr[i], "<") {
            // only works because we don't have mixed tags and text in SIF
            arr[i] = emptytag3.ReplaceAllString(arr[i], "")
          }
        }
        s = []byte(strings.Join(arr, ""))
        return s
}

func errcheck(t *testing.T, err error) {
        if err != nil {
                t.Fatalf("Error %s", err)
        }
}
END

while(<>) {
  if (m/^var (\S+) = `\s*<(\S+)/) {
    ($var, $object) = m/^var (\S+) = `\s*<(\S+)/;
    print << "END";
func Test_$var(t *testing.T) {
        a := $object\{\}
        err := xml.Unmarshal([]byte($var), &a)
        errcheck(t, err)
        err = defaults.Set(&a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml(stripEmptyTags([]byte($var)))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("$var not DeepEqual:\\nmv: \%+v\\nwant: \%+v\\n%s\\n%s\\n", mv, want, string(stripEmptyTags([]byte($var))), string(stripEmptyTags(output)))
        }
}
END
  }

}
