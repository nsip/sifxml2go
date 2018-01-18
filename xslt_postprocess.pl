=pod=
      SpanGaps string `xml:"SpanGaps"`
       ##1SpanGap []string `xml:"SpanGap"`
       ##2Type string `xml:"Type"`
       ##2Code string `xml:"Code"`
       ##2Name string `xml:"Name"`
       ##2Value string `xml:"Value"`
       ##2StartDateTime string `xml:"StartDateTime"`
       ##2EndDateTime string `xml:"EndDateTime"`
      IsCurrent string `xml:"IsCurrent"`

to

      SpanGaps struct {
        SpanGap []struct {
          Type string `xml:"Type"`
        } `xml:"SpanGap"`
      } `xml:"SpanGaps"`

    type ScheduledActivityOverrideType struct {

=cut=

@lines = ();
@lines1 = ();
while(<>) {
  push @lines, $_;
}
foreach (@lines) {
  ($indent, $term, $array, $type, $xml) = /^\s*(##\d)?(\S+)\s+(\[\])?(\S+)\s(\S*)/;
  if ($prev_indent || $indent) {
    if (substr($prev_indent,2) < substr($indent,2)) {
      # printf "%s %sstruct {\n", $prev_term, $prev_array;
      push @lines1, sprintf ("%s %sstruct {\n", $prev_term, $prev_array);
      push @tags, $prev_xml;
    } elsif (substr($prev_indent,2) == substr($indent,2)) {
      $prev =~ s/##\d//;
      push @lines1, $prev;
    } elsif (substr($prev_indent,2) > substr($indent,2)) {
      $prev =~ s/##\d//;
      $jumps = substr($prev_indent,2) - substr($indent,2);
      push @lines1, $prev;
      for ($i=0;$i<$jumps;$i++) {
        # printf "} %s\n", pop(@tags);
        push @lines1, sprintf ("} %s\n", pop(@tags));
      }
    }
  } else {
    $indent = "" if /^\stype /;
    push @lines1, $prev;
  }
  $prev = $_;
  $prev_indent = $indent;
  $prev_term = $term;
  $prev_array = $array;
  $prev_type = $type;
  $prev_xml = $xml;
}
push @lines1, $prev;

=pod=
        GetBookings     string `xml:"getBookings"`
        From            time.Time       `xml:"from,attr"`
        Location        string          `xml:"location,attr"`

to

        GetBookings     struct {
                From            time.Time       `xml:"from,attr"`
                Location        string          `xml:"location,attr"`
                Value string `xml:",chardata"`
        }       `xml:"getBookings"`

    type ScheduledActivityOverrideType struct {
      DateOfOverride string `xml:"DateOfOverride,attr"`
      }

=cut=

foreach (@lines1) {
  ($indent, $term, $array, $type, $xml) = /^\s*(##\d)?(\S+)\s+(\[\])?(\S+)\s(\S*)/;
  ($attr) = /(,attr"`)/;
  if ($attr and !$prev_attr) {
    if ($prev =~ m/^\s*type /) {
      print $prev;
      $attr = "";
    } else {
      push @tags, $prev_xml;
      printf "%s struct {\n", $prev_term;
    }
  }
  elsif (!$attr and $prev_attr) {
    print $prev;
    print qq{Value string `xml:",chardata"`\n};
    printf "} %s\n", pop(@tags);
  } else {
    print $prev;
  }
  $prev_attr = $attr;
  $prev = $_;
  $prev_indent = $indent;
  $prev_term = $term;
  $prev_array = $array;
  $prev_type = $type;
  $prev_xml = $xml;

}
print $_;
