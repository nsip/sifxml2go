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

$lines[-1] .= "\n";

@lines = indent_attrs(xsi_types(inherited_types(codesets(@lines))));
while(grep(/#/, @lines)) {
  @lines = resolve_embedded(@lines);
}
(resolve_attrs(@lines));

sub codesets(@) {
  my @lines = @_;
  my @lines1 = ();
  while ($_ = shift @lines) {
    if (/ENUM (\S+) \[/) {
      $type = $1;
      $type .= "Type" ;#unless $type =~ m/Type$/;
      push @lines1, "type $type string\n";
      $_ = shift @lines;
      s/^\s+|\s+$//g;
      @elems = split /, /;
      push @lines1, "var ${type}_values = []string{";
      push @lines1, join(", ", map("\"$_\"", @elems));
      push @lines1, "}\n";
      push @lines1, "var ${type}_map = map[string]struct{}{";
      push @lines1, join(", ", map("\"$_\"\: struct{}{}", @elems));
      push @lines1, "}";
      shift @lines;
    } else {
      push @lines1, $_;
    }
  }
  return @lines1;
}


sub inherited_types(@) {
  my @lines = @_;
  my @lines1 = ();
  foreach (@lines) {
    if(/%Inherits/) {
      $_ = $_;
    }
    if (/%Inherits: (int|float64|bool|RefIdType) ;/) {
      $type = $1;
      s/%Inherits: \S+ ;/%Value: $type `xml:",chardata"`/;
      $hold = $_;
    } elsif (/%Inherits: (string|xs:\S+|AUCodeSet\S+|URIOrBinaryType) ;/) {
      s/%Inherits: \S+ ;/%Value: string `xml:",chardata"`/;
      $hold = $_;
    } elsif (/%Inherits/) {
      s/%Inherits: (\S+) ;/$1/;
      push @lines1, $_;
    } else {
      s/^/%/ if ($hold and not /\}/);
      push @lines1, $_ unless /\}/;
    }
    if (/\}/) {
      push @lines1, $hold if $hold;
      push @lines1, $_;
      $hold = '';
    }
  }
  return @lines1;
}

sub xsi_types(@) {
  my @lines = @_;
  my @lines1 = ();
  foreach (@lines) {
    if (/xsi:/) {
      s/xsi:type/Type/g;
      s/EMPTY/string/g;
    }
    push @lines1, $_;
  }
  return @lines1;
}

sub add_json(@) {
  my @lines = @_;
  my @lines1 = ();
  my ($prev);
  my @obj = ();
  foreach (@lines) {
    if (m/\btype /) {
      push @lines1 , @obj;
      @obj = ();
    }
    s/`xml:"([^,]+),attr"`/`xml:"\1,attr" json:"\1"`/;
    s/`xml:",chardata"`/`xml:",chardata" json:"value"`/;
    s/`xml:"([^,]+)"`/`xml:"\1,omitempty" json:"\1,omitempty"`/;
    push @obj, $_;
  }
  push @lines1 , @obj;
  return @lines1;
}


sub indent_attrs(@) {
  my @lines = @_;
  my @lines1 = ();
  my ($prev);
  foreach (@lines) {
    s/ OPT / / if m/,attr/;
    ($object) = /\btype (\w+)/;
    ($indent, $term, $array, $type, $xml) = /^\s*(?:##(\d))?(\S+)\s+(\[\])?(\S+)\s(\S*)/;
    ($indent, $term, $array, $type, $xml) = () if $term eq 'type';
    ($indent, $term, $array, $type, $xml) = () if /%/;
    $seen_first_element = 0 if($object);
    $seen_first_element = 1 if ($xml and $xml !~ /,attr/);
    if ($xml =~ /,attr/ && $prev_xml !~ /,attr/ && $seen_first_element && $prev_type eq 'EMPTY') {
      $pushed_value = 1;
      $pushed_type = 'string';
    }
    if ($xml =~ /,attr/ && $prev_xml !~ /,attr/ && $seen_first_element && $prev_type ne 'EMPTY') {
      $pushed_type = $prev_type;
    }
    if ($xml =~ /,attr/ && $seen_first_element) {
      s/##\d//;
      s/^(\s*)/sprintf("      ##%d", $prev_indent+1)/e;
    }
    if ($xml =~ /,attr/ && $prev_xml !~ /,attr/ && $seen_first_element && !$pushed_value) {
      $pending_value = sprintf(qq[      ##%dValue: %s `xml:",chardata"`;\n], $prev_indent+1, $pushed_type);
      $pushed_value = 1;
    }
    if ($xml !~ /,attr/ && $prev_xml =~ /,attr/ || $object ){
      push @lines1, $pending_value if $pending_value;
      $pending_value = '';
      $pushed_value = 0;
    }
    push @lines1, $_;
    $prev = $_;
    $prev_indent = $indent;
    $prev_term = $term;
    $prev_array = $array;
    $prev_type = $type;
    $prev_xml = $xml;
    $prev_object = $object if $object;
  }
  foreach (@lines1) { s/%//; }
  return @lines1;
}

sub resolve_embedded(@) {
  my @lines = @_;
  my @lines1 = ();
  my @lines2 = ();
  my ($prev);
  foreach (@lines) {
    ($object) = /\btype (\w+)/;
    ($indent, $term, $array, $type, $xml) = /^\s*(##\d)?(\S+)\s+(\[\])?(\S+)\s(\S*)/;
    if ($prev_indent || $indent) {
      if (substr($prev_indent,2) < substr($indent,2)) {
        if ($indent eq '##1') {
          $new_type_name = $prev_object . '%' . $prev_element . ( $prev_array ? "" : ""); # "s"
          $new_type_name =~ s/://g;
          push @lines1, sprintf ("      %s %s%s %s\n", $prev_term, $prev_array, $new_type_name, $prev_xml);
          push @lines2, sprintf ("type %s struct {\n", $new_type_name);
        } else {
          $prev =~ s/##1//;
          $prev =~ s/##(\d)/sprintf("##%d", $1-1)/e;
          $element = $xml =~ /,attr/ ? $prev_element : $term;
          $new_type_name = $prev_object;
          $new_type_name =~ s/#.*$//;
          $new_type_name .= '%' . $element . ( $array ? "List" : "");
          $new_type_name =~ s/://g;
          $prev =~ s/\[\]EMPTY/\[\]$new_type_name/;
          push @lines2, $prev;
        }
      } elsif (substr($prev_indent,2) == substr($indent,2)) {
        $prev =~ s/##1//;
        $prev =~ s/##(\d)/sprintf("##%d", $1-1)/e;
        $element = $xml =~ /,attr/ ? $prev_element : $term;
        $new_type_name = $prev_object;
        $new_type_name =~ s/#.*$//;
        $new_type_name .= '%' . $element . ( $array ? "List" : "");
        $new_type_name =~ s/://g;
        $prev =~ s/\[\]EMPTY/\[\]$new_type_name/;
        push @lines2, $prev;
      } elsif (substr($prev_indent,2) > substr($indent,2)) {
        $prev =~ s/##1//;
        $prev =~ s/##(\d)/sprintf("##%d", $1-1)/e;
        push @lines2, $prev;
=pod=
      $jumps = substr($prev_indent,2) - substr($indent,2);
      for ($i=0;$i<$jumps;$i++) {
        push @lines2, sprintf ("}\n");
      }
=cut
if ($prev_indent eq '##1' or !$indent) {
  push @lines2, sprintf ("}\n");
}
      }
    } else {
      $indent = "" if /^\stype /;
      push @lines1, $prev;
    }
    $prev = $_;
    $prev_indent = $indent;
    $prev_term = $term;
    $prev_element = $term unless $xml =~ /,attr/;
    $prev_array = $array;
    $prev_type = $type;
    $prev_xml = $xml;
    $prev_object = $object if $object;
  }
  push @lines1, $prev;
  return (@lines1, @lines2);
}


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

sub resolve_attrs(@) {
  my @lines1 = @_;
  my ($prev);
  foreach (@lines1) {
    s/%/_/g;
    print;
  }
}
