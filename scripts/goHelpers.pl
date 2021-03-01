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
  "errors"
  "fmt"
  "log"
  "strconv"
  "encoding/json"
  "reflect"

  "github.com/qdm12/reprint"
  )

// Property/Value pair, used to set multiple properties of a container together
type Prop struct {
  Key string
  Value interface{}
}

// Check whether value is allowed in a codeset, based on the associated codeset values map
func CodesetContains(codeset map[string]struct{}, value interface{}) bool {

 	vstr, ok := value.(string)
 	if !ok {
 		return ok
 	}
 	_, ok = codeset[vstr]
 	return ok
}

func (a *Int) UnmarshalJSON(b []byte) error {
  log.Printf("Unmarshal %s to Int\\n", string(b))
  type temporary Int
	err := json.Unmarshal(b, (*temporary)(a))
  if err != nil {
    log.Println("try #1")
    var str string
    err = json.Unmarshal(b, &str)
    log.Println("try #2")
      if err != nil {
        return err
        }
        aI, err:=strconv.Atoi(str)
      if err != nil {
        return err
        }
    log.Printf("try #3: %d\\n", aI)
        *a = Int(aI)
    }
  return nil
}

func (a *Float) UnmarshalJSON(b []byte) error {
  type temporary Float
	err := json.Unmarshal(b, (*temporary)(a))
  if err != nil {
    var str string
    err = json.Unmarshal(b, &str)
      if err != nil {
        return err
        }
        aI, err:=strconv.ParseFloat(str, 64)
      if err != nil {
        return err
        }
        *a = Float(aI)
    }
  return nil
}

func (a *Bool) UnmarshalJSON(b []byte) error {
  type temporary Bool
	err := json.Unmarshal(b, (*temporary)(a))
  if err != nil {
    var str string
    err = json.Unmarshal(b, &str)
      if err != nil {
        return err
        }
        *a = Bool(str == "true")
    }
  return nil
}


END

foreach $n (@object) {
  print <<"END";
// Create a slice of pointers to the object type
func ${n}Slice() []*$n {
  return make([]*$n, 0)
  }
END
}

foreach $n (sort keys %types) {
print << "END";
// Performs a deep clone on the type, and is used to duplicate an element into another container (particularly if the element is itself nested)
  func (t *$n) Clone() (*$n) {
return reprint.This(t).(*$n)
}
END
  next if %alias_orig{$n};
  print <<"END";
// Generates a pointer to the given value (unless it already is a pointer), and returns an error in case
// the value mismatches X.
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
END
}


# Matt's Append is my AddNew
foreach $n (sort keys %list) {
  $emptytype= emptytype($list{$n}{TYPE});
      $cv = codeset_validate($$list{$n}{TYPE});
  print <<"END";

// Appends value to the list. Creates list if it is empty. Aborts if the list is a list of codeset values,
// and the value does not match the codeset.
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

// Appends an empty value to the list. This value can then be populated through accessors on Last().
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

// Retrieve the last value of the list. Calls AddNew() if the list is empty.
func (t *$n) Last() *$list{$n}{TYPE} {
  if t.$list{$n}{KEY} == nil {
    t = t.AddNew()
    }
        return &(t.$list{$n}{KEY}\[len(t.$list{$n}{KEY})-1])
}

// Retrieves the nth value in the list. Raises error if index is out of bounds.
func (t *$n) Index(n int) (*$list{$n}{TYPE}, error) {
  if (n >= t.Len() || n < 0) {
    return nil, errors.New("subscript out of range on list")
    }
  if t.$list{$n}{KEY} == nil {
    t = t.AddNew()
    }
        return &(t.$list{$n}{KEY}\[n]), nil
}

// Length of the list.
func (t *$n) Len() int {
  if t.$list{$n}{KEY} == nil {
    t = t.AddNew()
    }
        return len(t.$list{$n}{KEY})
}


END

  if ($list{$n}{TYPE} eq "string" or $alias{$list{$n}{TYPE}} eq "string") {
      print <<"END";
// Append a single string to the list. Only defined for lists of strings or of types aliased to string.
      func (t *$n) AppendString(value string) *$n {
        return t.Append(($list{$n}{TYPE})(value))
        }
END
  }
}

foreach $n (sort keys %alias_orig) {
  if ($alias{$n} eq 'string') {
  print <<"END";
// Return string value
func (t *$n) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }

// Generates a pointer to the given value (unless it already is a pointer), and returns an error in case
// the value mismatches $n. In the case of aliased types, accepts primitive values and converts them to the required alias.
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
// Returns int value
func (t *$n) Int() int {
  return int((reflect.ValueOf(*t).Interface()).($n))
  }

// Generates a pointer to the given value (unless it already is a pointer), and returns an error in case
// the value mismatches $n. In the case of aliased types, accepts primitive values and converts them to the required alias.
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
// 
// Returns bool value
func (t *$n) Bool() bool {
  return bool((reflect.ValueOf(*t).Interface()).($n))
  }

// Generates a pointer to the given value (unless it already is a pointer), and returns an error in case
// the value mismatches $n. In the case of aliased types, accepts primitive values and converts them to the required alias.
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
// Returns float64 value
func (t *$n) Float() float64 {
  return float64((reflect.ValueOf(*t).Interface()).($n))
  }

// Generates a pointer to the given value (unless it already is a pointer), and returns an error in case
// the value mismatches $n. In the case of aliased types, accepts primitive values and converts them to the required alias.
// Also deals with both Float32 and Float64 values.
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

foreach $n (sort keys %types) {
  next if $list{$n};
  print <<"END";

// Set the value of a property to nil
func (n *$n) Unset(key string) *$n {
        switch key {
END
  foreach $s (sort keys %{$types{$n}}) {
    unsetswitch($n, $s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (sort keys %{$types{$inherit{$n}}}) {
      unsetswitch($n, $s, $types{$inherit{$n}}{$s});
    }
  }
  print <<"END";
        default:
          log.Fatalf("%s is not a valid element name in %s\\n", key, "$n")
        }
        return n
}

// Set a sequence of properties
func (n *$n) SetProperties(props ...Prop) *$n {
 	for _, p := range props {
 		n.SetProperty(p.Key, p.Value)
 	}
 	return n
 }

// Set a property to a value. Aborts if property name is undefined for the type. Aborts if the list is a list of codeset values,
// and the value does not match the codeset.
func (n *$n) SetProperty(key string, value interface{}) *$n {
        if n == nil {
                n, _ = ${n}Pointer(${n}\{\})
        }
        switch key {
END
  foreach $s (sort keys %{$types{$n}}) {
    setswitch($n, $s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (sort keys %{$types{$inherit{$n}}}) {
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

  foreach $s (sort keys %{$types{$n}}) {
    readfunction($n, $s, $types{$n}{$s});
    isnil($n, $s, $types{$n}{$s});
  }
  if ($inherit{$n}) {
    foreach $s (sort keys %{$types{$inherit{$n}}}) {
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
// Returns whether the element value for $s is nil in the container $n.
  func (s *$n) ${s}_IsNil() bool {
    return s.$lcn.$s == nil
    }
END
}

sub readfunction($$$){
  my ($n, $s, $t) = @_;
  $emptytype = emptytype($t);
  $lcn = lc $n;
  print <<"END";
// Return the element value (as a pointer to the container, list, or primitive representing it).
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
