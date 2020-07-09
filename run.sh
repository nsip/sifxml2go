#!/bin/bash

# 1. Download specgen

#rm -fr specgen_input
#git clone https://github.com/nsip/specgen_input.git

# 2. Empty target directories

rm -rf sifxml
rm -rf sifgraphql
mkdir -p sifxml
mkdir -p sifgraphql
echo "" > sifgraphql/sif-schema.graphql
echo "package sifxml\n" > sifxml/examples.go

# 3. Extract all necessary information from specgen into flat files

containsElement () {
  local e match="$1"
  shift
  for e; do [[ "$e" == "$match" ]] && return 0; done
  return 1
}

xsltproc included_objects.xslt specgen_input/06_DataModel/Custom/DataModel-Custom-AU.xml | perl -ne 'next unless $_ =~ /\S/; next if $_ =~ /<\?/; s/^\s+//; s/\s+$//; print "./specgen_input/06_DataModel/Custom/" . $_ . "\n"' > objs.txt
IFS=$'\n' read -d '' -r -a objectarray < objs.txt


for filename in ./specgen_input/06_DataModel/Custom/Common/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  if [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/StudentScoreSet.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/PersonPrivacyObligation.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/ReportAuthorityInfo.xml" ]] ; then
    continue
  fi
  xsltproc sifobject.xslt "$filename" | perl xslt_postprocess.pl | perl -s struct2go.pl -o > sifxml/$(basename "$filename" .xml).go
  xsltproc sifobject.xslt "$filename" | perl xslt_postprocess.pl | perl struct2graphql.pl >> sifgraphql/sif-schema.graphql
done
for filename in ./specgen_input/06_DataModel/Custom/AU/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  xsltproc sifobject.xslt "$filename" | perl xslt_postprocess.pl | perl -s struct2go.pl -o > sifxml/$(basename "$filename" .xml).go
  xsltproc sifobject.xslt "$filename" | perl xslt_postprocess.pl | perl struct2graphql.pl >> sifgraphql/sif-schema.graphql
done

echo '<root xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xi="http://www.w3.org/2001/XInclude" xmlns="http://sifassociation.org/SpecGen" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xhtml="http://www.w3.org/1999/xhtml" >' > data.xml
cat specgen_input/80_BackMatter/Generic-CommonTypes.xml >> data.xml
cat specgen_input/80_BackMatter/Custom/DataModel-CommonTypes-Custom.xml >> data.xml
echo '</root>' >> data.xml
xsltproc sifobject.xslt data.xml | perl xslt_postprocess.pl | perl struct2go.pl > sifxml/DataModel.go
xsltproc sifobject.xslt data.xml | perl xslt_postprocess.pl | perl struct2graphql.pl >> sifgraphql/sif-schema.graphql

echo '<root xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xi="http://www.w3.org/2001/XInclude" xmlns="http://sifassociation.org/SpecGen" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xhtml="http://www.w3.org/1999/xhtml" >' > codesets.xml
cat specgen_input/80_BackMatter/Custom/DataModel-CodeSets-Custom.xml >> codesets.xml
cat specgen_input/80_BackMatter/Custom/DataModel-ExternalCodeSets-Custom.xml >> codesets.xml
echo '</root>' >> codesets.xml

xsltproc sifobject.xslt codesets.xml | perl xslt_postprocess.pl | perl struct2go.pl > sifxml/Codesets.go

# 4. Examples

for filename in ./specgen_input/06_DataModel/Custom/Common/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  if [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/StudentScoreSet.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/PersonPrivacyObligation.xml" ]] ||
       [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/ReportAuthorityInfo.xml" ]] ; then
    continue
  fi
  perl sifexamples.pl "$filename" >> sifxml/examples.go
done
for filename in ./specgen_input/06_DataModel/Custom/AU/*.xml; do
  if containsElement "$filename" "${objectarray[@]}" ; then
    :
  else
    echo "Excluded:" $filename;
    continue;
  fi
  if [[ "$filename" == "./specgen_input/06_DataModel/Custom/Common/StudentScoreSet.xml" ]]; then
    continue
  fi
  perl sifexamples.pl "$filename" >> sifxml/examples.go
done

cat sifxml/examples.go | perl siftest.pl > sifxml/sifxml_test.go
