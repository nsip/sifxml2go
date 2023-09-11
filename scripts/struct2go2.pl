# perl struct2go.pl : run with individual objects 
# perl -s struct2go.pl -o : sets $o : run with data model file

while(<>) {
    if (m/\btype (\S+).*\{/ && !$seen && !$o) {
    $obj = $1;
    $lc_obj = lc $obj;
    print <<"END";

    type ${obj}s struct {
      ${lc_obj}s `json:"${obj}s"`
    }

    type ${lc_obj}s struct {
      ${obj} []${lc_obj} `json:"${obj}"`
    }

    type ${obj} struct {
      ${lc_obj} `xml:"${obj}" json:"${obj}"`
     }

     type ${lc_obj} struct {
END
    $seen++;
    next;
  } 
  if (m/\btype (\S+).*/ && !$seen && !$o) { # object is alias of type
    $obj = $1;
    $lc_obj = lc $obj;
    print <<"END";

    type ${obj}s struct {
      ${lc_obj}s `json:"${obj}s"`
    }

    type ${lc_obj}s struct {
      ${obj} []${lc_obj} `json:"${obj}"`
    }

    type ${lc_obj} ${obj}

END
    print;
    $seen++;
    next;
  }

  if (m/\s*\}\s*$/) {
    $in_struct = 0;
  }
  # struct inheritance
  if (m/^\s*(\S+)\s*$/) {
    $_ = lc $_;
  }
  print;
  if (m/\btype (\S+) struct \{/) {
    $in_struct = 1;
    $t = $1;
    $lct = lc $t;
    $annotation = $o ? "" : qq{`xml:"$t" json:"$t"`};
    print << "END";
  $lct $annotation
}

type $lct struct {

END
  }
}
