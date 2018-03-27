while(<>) {
  if(m#</xhtml:Example>#) {
    $example = 0;
    $out =~ s/ xml:lang="en"//;
    ($object) = $out =~ m/<(\S+)/;
    print "var test_$name = `$out`\n";
    $name = "";
  }
  $out .= $_ if $example;
  if(/<xhtml:Example/){
    ($name) = /name="([^"]+)"/;
    $name = sprintf("example%04d", int(rand(1000))) unless $name;
    $name =~ s/[^A-Za-z0-9]//g;
    $example = 1;
    $out = "";
  }
}
