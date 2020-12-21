while(<>) {
  s/`[^`]+`//;
  s/ struct / /;
  s/\bstring\b/String/g;
  s/\bInherits: (\S+)\b/\1 \1/g;
  s/\[\](\w+)/\[\1\]/g;
  s/;//;
  print $_;
}

