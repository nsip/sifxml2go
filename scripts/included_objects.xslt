<xsl:stylesheet version="1.0"
  xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
  xmlns:xi="http://www.w3.org/2001/XInclude">

  <xsl:template match="//xi:include">
     <xsl:value-of select="./@href"/>
  </xsl:template>

</xsl:stylesheet>
