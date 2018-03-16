print "package sifxml\n\n";

while(<>) {
  s/Inherits://;
  @parts = split /(`)/;
  $parts[0] =~ s/://g;
  $_ = join('', @parts);
  s/;//g;
  print;
}
