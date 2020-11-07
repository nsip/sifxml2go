print "package sifxml\n\n";

while(<>) {
  s/Inherits://;

  # perl -s -struct2go.pl -o
  if (m/\btype (\S+)/ && !$seen && $o) {
    print "type $1s []$1\n\n";
    $seen++;
  }


  @parts = split /(`)/;
  $parts[0] =~ s/://g;
  $_ = join('', @parts);
  s/;//g;
  s/`xml:"([^,]+),attr"`/`xml:"\1,attr" json:"\1"`/;
  s/`xml:",chardata"`/`xml:",chardata" json:"value"`/;
  $omitempty = "";
  $omitempty = ",omitempty" if m/ OPT/ && !/(float64|int|bool) OPT/;
  s/`xml:"([^,]+)"`/`xml:"\1$omitempty" json:"\1"`/;
  # We use pointer for all values, so we can introduce NULL for omitempty
  # distinct from 0 for numeric: https://stackoverflow.com/a/38487668
  s/ OPT//;
  s/^(\s*\S+\s+)(\S+\s+)`/\1*\2`/; # name type `...`
  s/^(\s*\S+\s+)(\S+\s*)$/\1*\2/; # name alias
  print;
}
