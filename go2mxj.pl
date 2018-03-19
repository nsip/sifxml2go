$out = "";

while(<>) {
  s/`xml:"([^,]+),attr"`/`xml:"\1,attr" json:"-\1"`/;
  s/`xml:",chardata"`/`xml:",chardata" json:"Value"`/;
  s/`xml:"([^,]+)"`/`xml:"\1,omitempty" json:"\1"`/;
  if(m/^\s*type /) {
    ($type) = m/^\s*type\s+(\S+)/;
    $output_unmarshaller = 0;
  }
  if(m/\[\]/) {
    $output_unmarshaller = 1;
  }
  $out .= $_;
  # assumed defined already: globals
  #     var re = regexp.MustCompile(`^(\\s*\\{"[^"]+"\\s*:)`)
  #     var re2 = regexp.MustCompile(`\\}\\s*\$`)
  # aliasing: see http://choly.ca/post/go-json-marshalling/

  if(m/^\s*\}/ && $output_unmarshaller) {
    $out .= <<"END_MESSAGE";
    func (this *$type) UnmarshalJSON(data []byte) error {
        type Alias $type
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "\$1["), "]}")))
}
END_MESSAGE

  }
}

if ($out =~ m/UnmarshalJSON/) {
  $out =~ s/package sifxml\n/package sifxml\nimport (\n"bytes"\n"encoding\/json"\n"log"\n)\n/;
}
print $out;
