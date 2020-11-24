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
print << "END";
  $lct `xml:"$t" json:"$t"`
}

type $lct struct {
END
    }
}
