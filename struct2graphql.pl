while(<>) {
  next if /^\s*package/;
  print $_;
}

