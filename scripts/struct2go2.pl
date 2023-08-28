# perl struct2go.pl : run with individual objects
# perl -s struct2go.pl -o : sets $o : run with data model file

while(<>) {
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
    $annotation = $o ? "" : qq{`xml:"$t"`};
    print << "END";
  $lct $annotation
}

type $lct struct {
END
  }
}
