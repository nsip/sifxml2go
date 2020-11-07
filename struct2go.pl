print "package sifxml\n\nimport _ \"github.com/creasty/defaults\"\n\n";

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
  #s/^(\s*\S+\s+)(\S+\s+)`/\1*\2`/;

  # no, instead we presuppose https://github.com/creasty/defaults, 
  # and introduce defaults for numeric optional values. We will leave boolean options
  # as false for data generation, although this is problematic
  s/ OPT$//;
  s/ OPT / / if (/ OPT / && !/(float64|int|bool) OPT /);
  s/ OPT / / if (/ OPT / && /\[\](float64|int|bool) OPT /);
  s/ OPT / / if (/ OPT / && /bool OPT /);
  if(/(float64|int) OPT /) {
    s/ OPT / /;
    s/`/`default:"-2147483648" /;
  }
  print;
}
