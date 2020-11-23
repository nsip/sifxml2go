print "package sifxml\n\n";

# perl -s struct2go.pl -o : sets $o : run with individual objects
# perl struct2go.pl : run with data model file

while(<>) {
  s/Inherits://;

  # perl -s struct2go.pl -o
  if (m/\btype (\S+)/ && !$seen && $o) {
    print "type $1s []$1\n\n";
    $seen++;
  }


  @parts = split /(`)/;
  $parts[0] =~ s/://g unless $parts[0] =~ /map\[/;
  $_ = join('', @parts);
  s/;//g;
  s/`xml:"([^,]+),attr"`/`xml:"\1,attr" json:"\1"`/;
  s/`xml:",chardata"`/`xml:",chardata" json:"value"`/;
  $omitempty = "";
  $omitempty = ",omitempty" if m/ OPT/;
  s/`xml:"([^,]+)"`/`xml:"\1$omitempty" json:"\1"`/;
  # We use pointer for all values, so we can introduce NULL for omitempty
  s/ OPT//;
  s/^(\s*\S+\s+)([^\[\s ]\S*\s+)`/\1*\2`/; # name type `...`
  s/^(\s*\S+\s+)(\S+\s*)$/\1*\2/; # name alias
  # We will make String a local type, so that we can apply methods to it shared by the alias types like codesets
  s/ \*string / *String /;
  # Ditto the other primitives
  s/ \*int / *Int /;
  s/ \*float64 / *Float /;
  s/ \*bool / *Bool /;
  print;
}
