print <<"END";
package sifxml

import (
        "encoding/xml"
        "github.com/clbanning/mxj"
        "reflect"
        "regexp"
        "testing"
)

var re = regexp.MustCompile(`(?s:\\s*<[^>/]+></[^>]+>\\s*)`)

func stripEmptyTags(s []byte) []byte {
        s = re.ReplaceAll(s, []byte(""))
        s = re.ReplaceAll(s, []byte(""))
        s = re.ReplaceAll(s, []byte(""))
        s = re.ReplaceAll(s, []byte(""))
        s = re.ReplaceAll(s, []byte(""))
        // log.Println(string(s))
        return s
}

func errcheck(t *testing.T, err error) {
        if err != nil {
                t.Fatalf("Error %s", err)
        }
}
END

while(<>) {
  if ($var, $object) = /var (\S+) = `\s*<(\S+)/ {
print << "END";
func Test_$var(t *testing.T) {
        a := $object{}
        err := xml.Unmarshal([]byte($var), &a)
        errcheck(t, err)
        output, err := xml.Marshal(a)
        errcheck(t, err)
        want, err := mxj.NewMapXml([]byte($var))
        errcheck(t, err)
        mv, err := mxj.NewMapXml(stripEmptyTags(output))
        errcheck(t, err)
        if !reflect.DeepEqual(mv, want) {
                t.Fatalf("$var not DeepEqual:\\nmv: \%+v\\nwant: \%+v\\n", mv, want)
        }
}
END
  }

}
