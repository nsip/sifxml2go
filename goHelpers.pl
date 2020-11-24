%alias_orig = ();
%alias = ();
%types = ();
%list = ();
%inherit = ();
%codeset = ();
@object = ();

while(<>) {
  print STDERR $_;
  chomp;
  if (m/^type (\S+?)s \[\]\1$/) {
    ($name) = (m/^type (\S+?)s \[\]\1/);
    push @object, $name;
  }
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

  "github.com/qdm12/reprint"
  )

type Prop struct {
  Key string
  Value interface{}
}

func CodesetContains(codeset map[string]struct{}, value interface{}) bool {

 	vstr, ok := value.(string)
 	if !ok {
 		return ok
 	}
 	_, ok = codeset[vstr]
 	return ok
 }

END

foreach $n (@object) {
  print <<"END";
func ${n}Slice() []*$n {
  return make([]*$n, 0)
  }
END
}

foreach $n (keys %types) {
  next if %alias_orig{$n};
  print <<"END";
  func ${n}Pointer(value interface{}) (*$n, bool) {
switch t := value.(type) {
        case *$n:
                return value.(*$n), true
        case $n:
                v, _ := value.($n)
                return &v, true
        default:
                fmt.Printf("Warning: cannot resolve %T (%v) to $n\\n", t, value)
        }
        return nil, false
  }

  func (t *$n) Clone() (*$n) {
return reprint.This(t).(*$n)
  }
END
}


# Matt's Append is my AddNew
foreach $n (keys %list) {
  $emptytype= emptytype($list{$n}{TYPE});
      $cv = codeset_validate($$list{$n}{TYPE});
  print <<"END";

  func (t *$n) Append(value $list{$n}{TYPE}) *$n {
    $cv
        if t == nil {
                t, _ = ${n}Pointer(${n}\{\})
        }
        if t.$list{$n}{KEY} == nil {
                t.$list{$n}{KEY} = make([]$list{$n}{TYPE}, 0)
        }
        t.$list{$n}{KEY} = append(t.$list{$n}{KEY}, value)
        return t
}

func (t *$n) AddNew() *$n {
        if t == nil {
                t, _ = ${n}Pointer(${n}\{\})
        }
        if t.$list{$n}{KEY} == nil {
                t.$list{$n}{KEY} = make([]$list{$n}{TYPE}, 0)
        }
        t.$list{$n}{KEY} = append(t.$list{$n}{KEY}, $emptytype)
        return t
}

func (t *$n) Last() *$list{$n}{TYPE} {
  if t.$list{$n}{KEY} == nil {
    t = t.AddNew()
    }
        return &(t.$list{$n}{KEY}\[len(t.$list{$n}{KEY})-1])
}

END

  if ($list{$n}{TYPE} eq "string" or $alias{$list{$n}{TYPE}} eq "string") {
      print <<"END";
      func (t *$n) AppendString(value string) *$n {
        return t.Append(($list{$n}{TYPE})(value))
        }
END
  }
}

foreach $n (keys %alias_orig) {
  if ($alias{$n} eq 'string') {
  print <<"END";
func (t *$n) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }

func ${n}Pointer(value interface{}) (*$n, bool) {
switch t := value.(type) {
 	case *$n:
 		return value.(*$n), true
        case $n:
 		v, _ := value.($n)
 		return &v, true
 	case *string:
 		vstr, _ := value.(*string)
 		v := $n(*vstr)
 		return &v, true
 	case string:
 		vstr, _ := value.(string)
 		v := $n(vstr)
 		return &v, true
 	default:
 		fmt.Printf("Warning: cannot resolve %T (%v) to $n\\n", t, value)
 	}
 	return nil, false
  }


END
}
 if ($alias{$n} eq 'int') {
  print <<"END";
func (t *$n) Int() int {
  return int((reflect.ValueOf(*t).Interface()).($n))
  }

func ${n}Pointer(value interface{}) (*$n, bool) {
switch t := value.(type) {
        case *$n:
                return value.(*$n), true
        case $n:
                v, _ := value.($n)
                return &v, true
        case *int:
 		vstr, _ := value.(*int)
 		v := $n(*vstr)
 		return &v, true
        case int:
                vstr, _ := value.(int)
                v := $n(vstr)
                return &v, true
        default:
                fmt.Printf("Warning: cannot resolve %T (%v) to $n\\n", t, value)
        }
        return nil, false
  }

END
}
 if ($alias{$n} eq 'bool') {
  print <<"END";
func (t *$n) Bool() bool {
  return bool((reflect.ValueOf(*t).Interface()).($n))
  }

  func ${n}Pointer(value interface{}) (*$n, bool) {
switch t := value.(type) {
        case *$n:
                return value.(*$n), true
        case $n:
                v, _ := value.($n)
                return &v, true
        case *bool:
                vstr, _ := value.(*bool)
                v := $n(*vstr)
                return &v, true
        case bool:
                vstr, _ := value.(bool)
                v := $n(vstr)
                return &v, true
        default:
                fmt.Printf("Warning: cannot resolve %T (%v) to $n\\n", t, value)
        }
        return nil, false
  }

END
}
 if ($alias{$n} eq 'float64') {
  print <<"END";
func (t *$n) Float() float64 {
  return float64((reflect.ValueOf(*t).Interface()).($n))
  }

   func ${n}Pointer(value interface{}) (*$n, bool) {
switch t := value.(type) {
        case *$n:
                return value.(*$n), true
        case $n:
                v, _ := value.($n)
                return &v, true
        case *float64:
                vstr, _ := value.(*float64)
                v := $n(*vstr)
                return &v, true
        case float64:
                vstr, _ := value.(float64)
                v := $n(vstr)
                return &v, true
        case *float32:
                vstr, _ := value.(*float32)
                v := $n(float64(*vstr))
                return &v, true
        case float32:
                vstr, _ := value.(float32)
                v := $n(float64(vstr))
                return &v, true
        default:
                fmt.Printf("Warning: cannot resolve %T (%v) to $n\\n", t, value)
        }
        return nil, false
  }
END
}
}

foreach $n (keys %types) {
  next if $list{$n};
  print <<"END";

func (n *$n) Unset(key string) *$n {
        switch key {
END
  foreach $s (keys %{$types{$n}}) {
    unsetswitch($n, $s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (keys %{$types{$inherit{$n}}}) {
      unsetswitch($n, $s, $types{$inherit{$n}}{$s});
    }
  }
  print <<"END";
        default:
          log.Fatalf("%s is not a valid element name in %s\\n", key, "$n")
        }
        return n
}

func (n *$n) SetProperty(key string, value interface{}) *$n {
        if n == nil {
                n, _ = ${n}Pointer(${n}\{\})
        }
        switch key {
END
  foreach $s (keys %{$types{$n}}) {
    setswitch($n, $s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (keys %{$types{$inherit{$n}}}) {
      setswitch($n, $s, $types{$inherit{$n}}{$s});
    }
  }
  print <<"END";
        default:
          log.Fatalf("%s is not a valid element name in %s\\n", key, "$n")
        }
        return n
}

END

  foreach $s (keys %{$types{$n}}) {
    readfunction($n, $s, $types{$n}{$s});
    isnil($n, $s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (keys %{$types{$inherit{$n}}}) {
      readfunction($n, $s, $types{$inherit{$n}}{$s});
      isnil($n, $s, $types{$inherit{$n}}{$s});
    }
  }
}

sub unsetswitch($$$) {
  my ($n, $s, $t) = @_;
  $lcn = lc $n;
  print << "END";
  case "$s":
   n.$lcn.$s = nil
END
}

sub isnil($$$) {
  my ($n, $s, $t) = @_;
  $lcn = lc $n;
  print <<"END";
  func (s *$n) ${s}_IsNil() bool {
    return s.$s == nil || s.$lcn.$s == nil
    }
END
}

sub readfunction($$$){
  my ($n, $s, $t) = @_;
  $emptytype = emptytype($t);
  $lcn = lc $n;
  print <<"END";
  func (s *$n) ${s}() *$t {
    if s.$lcn.$s == nil {
END
  if($alias{$t}) {
    $cr = typecreate($alias{$t});
  print <<"END";
    if v, ok:= ${cr}Pointer($emptytype); ok {
      s.$lcn.$s = ((*$t)(v))
      }
END
  } else {
     $cr = typecreate($t);
  print <<"END";
     if v, ok:= ${cr}Pointer($emptytype); ok {
      s.$lcn.$s = v
      }
END
  }
  print <<"END";
      }
      return s.$lcn.$s
    }
END
}

sub setswitch($$$) {
  my ($n, $s, $t) = @_;
  $lcn = lc $n;
  print <<"END";
    case "$s":
END
  if ($alias{$t}) {
    $cr = typecreate($alias{$t});
    $cv = codeset_validate($t);
    print <<"END";
    $cv
    if v, ok:= ${cr}Pointer(value); ok {
      n.$lcn.$s = ((*$t)(v))
      }
END
  } else {
    $cr = typecreate($t);
    print <<"END";
    if v, ok:= ${cr}Pointer(value); ok {
      n.$lcn.$s = v
      }
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

sub codeset_validate($) {
  my ($t) = @_;
  $t = $alias_orig{$t} if $alias_orig{$t} && $codeset{$alias_orig{$t}};
  return "" unless $codeset{$t};
return << "END";
    if !CodesetContains(${t}_map, value) {
      log.Fatalf("%s is not present in %s\\n", value, "${t}_values")
      }
END
}       

sub emptytype($) {
  my ($t) = @_;
  $emptytype= "$t\{\}";
  if ($t eq "string" or $alias{$t} eq "string") {
    $emptytype= "\"\"";
  }
  if ($t eq "float64" or $alias{$t} eq "float64") {
    $emptytype= "0";
  }
  if ($t eq "int" or $alias{$t} eq "int") {
    $emptytype= "0";
  }
  if ($t eq "bool" or $alias{$t} eq "bool") {
    $emptytype= "false";
  }
  return $emptytype;
}
