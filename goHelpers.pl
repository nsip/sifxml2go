%alias_orig = ();
%alias = ();
%types = ();
%list = ();
%inherit = ();
%codeset = ();

while(<>) {
  print STDERR $_;
  chomp;
  if (m/^\s*var\s+\S+\s+=\s+\[\]/) {
    ($name) = m/^\s*var\s+(\S+)\s+=\s+\[\]/;
    $name =~ s/_values//;
    $codeset{$name}++;
  }
  if (m/^\s*type\s/) {
    ($name, $type) = m/^\s*type\s*(\S+)\s+(\S+)/;
    if ($type eq 'struct') {
      $types{$name} = {};
      $linecount = 0;
      while(<>) {
        print STDERR $_;
        chomp;
        last if /^\s*\}\s*$/;
        next unless /\S/;
        $linecount++;
        ($subname, $subtype) = /^\s*(\S+)\s*(\S*)/;
        if ($subtype eq "") {
          $inherit{$name} = $subname;
        } else {
          $types{$name}{$subname} = $subtype;
          $types{$name}{$subname} =~ s/^(\[\]|\*)//;
        }
      }
      if ($linecount == 1 && $subtype =~ /^\[\]/) {
        $subtype =~ s/^\[\]//;
        $list{$name} = {KEY => $subname, TYPE => $subtype};
      }
    } elsif ($type =~ m/^\[\]/) {
    } else {
      $alias{$name} = $type;
    }
  }
}

foreach $k (keys %alias) {
  $alias_orig{$k} = $alias{$k};
}
foreach $k (keys %alias) {
  if ($alias{$alias{$k}}) {
    $alias{$k} = $alias{$alias{$k}};
  }
} 

print <<"END";
package sifxml

import (
  "fmt"
  "log"
  "reflect"

      "github.com/ulule/deepcopier"
  )

func IntCreate(x int) *int {
  return &x
}

func FloatCreate(x float64) *float64 {
  return &x
}

func StringCreate(x string) *string {
  return &x
}

func BoolCreate(x bool) *bool {
  return &x
}

END

foreach $n (keys %types) {
  print <<"END";
func ${n}Create(x $n) *$n {
        return &x
}

func (n *$n) New() *$n {
        return ${n}Create($n\{\})
}

func (n *$n) Clone() $n {
  resource := &$n\{\}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

END
}

foreach $n (keys %list) {
  $emptytype= "$list{$n}{TYPE}\{\}";
  if ($list{$n}{TYPE} eq "string" or $alias{$list{$n}{TYPE}} eq "string") {
    $emptytype= "\"\"";
  }
  if ($list{$n}{TYPE} eq "float64" or $alias{$list{$n}{TYPE}} eq "float64") {
    $emptytype= "0";
  }
  if ($list{$n}{TYPE} eq "int" or $alias{$list{$n}{TYPE}} eq "int") {
    $emptytype= "0";
  }
  if ($list{$n}{TYPE} eq "bool" or $alias{$list{$n}{TYPE}} eq "bool") {
    $emptytype= "false";
  }
  $cv = codeset_validate($list{$n}{TYPE}, 0);
  print <<"END";
  func (t *$n) Append(value $list{$n}{TYPE}) *$n {
    $cv
        if t == nil {
                t = ${n}Create(${n}\{\})
        }
        if t.$list{$n}{KEY} == nil {
                t.$list{$n}{KEY} = make([]$list{$n}{TYPE}, 0)
        }
        t.$list{$n}{KEY} = append(t.$list{$n}{KEY}, value)
        return t
}

func (t *$n) AddNew() *$n {
        if t == nil {
                t = ${n}Create(${n}\{\})
        }
        if t.$list{$n}{KEY} == nil {
                t.$list{$n}{KEY} = make([]$list{$n}{TYPE}, 0)
        }
        t.$list{$n}{KEY} = append(t.$list{$n}{KEY}, $emptytype)
        return t
}

func (t *$n) Last() *$list{$n}{TYPE} {
        return &(t.$list{$n}{KEY}\[len(t.$list{$n}{KEY})-1])
}

END

  if ($list{$n}{TYPE} eq "string" or $alias{$list{$n}{TYPE}} eq "string") {
      print <<"END";
      func (t *$n) AppendString(value interface{}) *$n {
        return t.Append(($list{$n}{TYPE})(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
END
  }
}

foreach $n (keys %alias_orig) {
  print <<"END";
func (t *$n) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
END
}

foreach $n (keys %types) {
  next if $list{$n};
  print <<"END";

func (t *$n) CopyString(key string, value interface{}) *$n {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *$n) Unset(key string) *$n {
        switch key {
END
  foreach $s (keys %{$types{$n}}) {
    unsetswitch($s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (keys %{$types{$inherit{$n}}}) {
      unsetswitch($s, $types{$inherit{$n}}{$s});
    }
  }
  print <<"END";
        default:
          log.Fatalf("%s is not a valid element name in %s\\n", key, "$n")
        }
        return n
}

func (n *$n) Set(key string, value interface{}) *$n {
        if n == nil {
                n = ${n}Create($n\{\})
        }
        switch key {
END
  foreach $s (keys %{$types{$n}}) {
    setswitch($s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (keys %{$types{$inherit{$n}}}) {
      setswitch($s, $types{$inherit{$n}}{$s});
    }
  }
  print <<"END";
        default:
          log.Fatalf("%s is not a valid element name in %s\\n", key, "$n")
        }
        return n
}

END
}

sub unsetswitch($) {
  my ($s) = @_;
  print << "END";
  case "$s":
   n.$s = nil
END
}

sub setswitch($$) {
  my ($s, $t) = @_;
  print <<"END";
    case "$s":
END
  if ($alias{$t}) {
    $cr = typecreate($alias{$t});
    $cv = codeset_validate($t, 1);
    print <<"END";
    $cv
                      n.$s = ((*$t)(${cr}Create(value.($alias{$t}))))
END
  } else {
    $cr = typecreate($t);
    print <<"END";
                      n.$s = ${cr}Create(value.($t))
END
  }
}

sub typecreate($) {
  my ($s) = @_;
  if ($s eq "string") { return "String"; }
  elsif ($s eq "bool") { return "Bool"; }
  elsif ($s eq "float64") { return "Float"; }
  elsif ($s eq "int") { return "Int"; }
  else { return $s; }
}

sub codeset_validate($$) {
  my ($t, $interface) = @_;
  if($t eq "CountryType") {
    $t=$t;
  }
  $t = $alias_orig{$t} if $alias_orig{$t} && $codeset{$alias_orig{$t}};
  return "" unless $codeset{$t};
  my $cast = $interface ? "value.(string)" : "string(value)";
  return <<"END";
  present := false
  for _, b := range ${t}_values {
        if b == $cast {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\\n", value, "${t}_values")
      }
END
}       
