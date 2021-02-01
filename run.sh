#!/bin/bash

# 1. Download specgen

#rm -fr specgen_input
#git clone https://github.com/nsip/specgen_input.git

# 2. Empty target directories

rm -rf sifxml
rm -rf sifxml_tmp
rm -rf sifgraphql
mkdir -p sifxml
mkdir -p sifxml_tmp
mkdir -p sifgraphql
echo "" > sifgraphql/sif-schema.graphql
echo "package sifxml\n" > sifxml_tmp/examples.go

# 3. Extract all necessary information from specgen into flat files

containsElement () {
  local e match="$1"
  shift
  for e; do [[ "$e" == "$match" ]] && return 0; done
  return 1
}

xsltproc scripts/included_objects.xslt specgen_input/06_DataModel/Custom/DataModel-Custom-AU.xml | perl -ne 'next unless $_ =~ /\S/; next if $_ =~ /<\?/; s/^\s+//; s/\s+$//; print "./specgen_input/06_DataModel/Custom/" . $_ . "\n"' > objs.txt
IFS=$'\n' read -d '' -r -a objectarray < objs.txt


for filename in ./specgen_input/06_DataModel/Custom/Common/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  if #[[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/StudentScoreSet.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/PersonPrivacyObligation.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/ReportAuthorityInfo.xml" ]] ; then
    continue
  fi
  xsltproc scripts/sifobject.xslt "$filename" | perl scripts/xslt_postprocess.pl | perl -s scripts/struct2go.pl -o > sifxml_tmp/$(basename "$filename" .xml).go
  xsltproc scripts/sifobject.xslt "$filename" | perl scripts/xslt_postprocess.pl | perl scripts/struct2graphql.pl >> sifgraphql/sif-schema.graphql
done
for filename in ./specgen_input/06_DataModel/Custom/AU/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  xsltproc scripts/sifobject.xslt "$filename" | perl scripts/xslt_postprocess.pl | perl -s scripts/struct2go.pl -o > sifxml_tmp/$(basename "$filename" .xml).go
  xsltproc scripts/sifobject.xslt "$filename" | perl scripts/xslt_postprocess.pl | perl scripts/struct2graphql.pl >> sifgraphql/sif-schema.graphql
done

echo '<root xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xi="http://www.w3.org/2001/XInclude" xmlns="http://sifassociation.org/SpecGen" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xhtml="http://www.w3.org/1999/xhtml" >' > data.xml
cat specgen_input/80_BackMatter/Generic-CommonTypes.xml >> data.xml
cat specgen_input/80_BackMatter/Custom/DataModel-CommonTypes-Custom.xml >> data.xml
echo '</root>' >> data.xml
xsltproc scripts/sifobject.xslt data.xml | perl scripts/xslt_postprocess.pl | perl scripts/struct2go.pl > sifxml_tmp/DataModel.go
echo "type String string\ntype Int int\ntype Float float64\ntype Bool bool" >> sifxml_tmp/DataModel.go

xsltproc scripts/sifobject.xslt data.xml | perl scripts/xslt_postprocess.pl | perl scripts/struct2graphql.pl >> sifgraphql/sif-schema.graphql

echo '<root xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xi="http://www.w3.org/2001/XInclude" xmlns="http://sifassociation.org/SpecGen" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xhtml="http://www.w3.org/1999/xhtml" >' > codesets.xml
cat specgen_input/80_BackMatter/Custom/DataModel-CodeSets-Custom.xml >> codesets.xml
cat specgen_input/80_BackMatter/Custom/DataModel-ExternalCodeSets-Custom.xml >> codesets.xml
echo '</root>' >> codesets.xml

xsltproc scripts/sifobject.xslt codesets.xml | perl scripts/xslt_postprocess.pl | perl scripts/struct2go.pl > sifxml_tmp/Codesets.go

for filename in ./sifxml_tmp/*.go; do
  cat "$filename" |  perl struct2go2.pl > sifxml/$(basename "$filename")
done

cat sifxml_tmp/[A-Z]*.go | perl scripts/goHelpers.pl > sifxml/Helpers.go

# 4. Examples

for filename in ./specgen_input/06_DataModel/Custom/Common/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  if #[[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/StudentScoreSet.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/PersonPrivacyObligation.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/ReportAuthorityInfo.xml" ]] ; then
    continue
  fi
  perl scripts/sifexamples.pl "$filename" >> sifxml/examples.go
done
for filename in ./specgen_input/06_DataModel/Custom/AU/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  #if [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/StudentScoreSet.xml" ]]; then
    #continue
  #fi
  perl scripts/sifexamples.pl "$filename" >> sifxml/examples.go
done

cat sifxml/examples.go | perl scripts/siftest.pl > sifxml/sifxml_test.go
cp scripts/doc.go sifxml/doc.go

