<?xml version="1.0" encoding="UTF-8"?>
<!-- v.1.3.5 -->

<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:output method="xml" version="1.0" encoding="UTF-8" indent="yes" doctype-public="yes"/>

    <xsl:template match="@* | node()">
        <xsl:copy>
            <xsl:apply-templates select="@* | node()"/>
        </xsl:copy>
    </xsl:template>

    <xsl:template match="var[@key='id']">
        <xsl:element name="span">
            <xsl:value-of select="/generateGuarantee/statement/id"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="var[@key='sity']">
        <xsl:element name="span">
            <xsl:value-of select="/generateGuarantee/additional/sity"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="var[@key='dateGuarantee']">
        <xsl:element name="span">
            <xsl:value-of select="/generateGuarantee/statement/createDate"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="var[@key='type']">
        <xsl:element name="span">
            <xsl:value-of select="/generateGuarantee/statement/type"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="var[@key='beneficiaryName']">
        <xsl:element name="span">
            <xsl:value-of select="/generateGuarantee/statement/beneficiary/organization/name"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="var[@key='principalName']">
        <xsl:element name="span">
            <xsl:value-of select="/generateGuarantee/statement/principal/organization/name"/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="var[@key='garantName']">
        <xsl:element name="span">
            <xsl:value-of select="/generateGuarantee/statement/garant/organization/name"/>
        </xsl:element>
    </xsl:template>


    <xsl:template match="var"/>
    <xsl:template match="/generateGuarantee/statement" />
    <xsl:template match="/generateGuarantee/additional" />

    <!--
    <xsl:template match="/order/agent">
        <xsl:element name="filial">
            <xsl:value-of select="."/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="/order/agent_date">
        <xsl:element name="filial_date">
            <xsl:value-of select="."/>
        </xsl:element>
    </xsl:template>

    <xsl:template match="/order/version_protocol"/>
    <xsl:template match="/order/find_pars/pars/par/@regexp"/>
    <xsl:template match="/order/rqid"/>

    <xsl:template match="/order">
        <xsl:copy>
            <xsl:apply-templates/>

            <xsl:element name="frontIntegrationMode">
                <xsl:text>0</xsl:text>
            </xsl:element>

            <xsl:element name="pay_sys_type">
                <xsl:text>AGENT</xsl:text>
            </xsl:element>

            <xsl:element name="operator"/>

        </xsl:copy>
    </xsl:template>

    <xsl:template match="/order/services/serv">
        <xsl:copy>
            <xsl:attribute name="isRiskRecipient">0</xsl:attribute>

            <xsl:apply-templates select="@*|node()"/>

            <xsl:element name="print_doc">
                <xsl:text>1</xsl:text>
            </xsl:element>

            <xsl:element name="serv_type">
                <xsl:text>1</xsl:text>
            </xsl:element>

            <xsl:element name="no_credit_card">
                <xsl:text>0</xsl:text>
            </xsl:element>

        </xsl:copy>
    </xsl:template>


    <xsl:template match="/order/services/serv/pars/par">
        <xsl:copy>
            <xsl:attribute name="print_mask"/>
            <xsl:attribute name="input_mask"/>

            <xsl:apply-templates select="@*|node()"/>

        </xsl:copy>
    </xsl:template-->

</xsl:stylesheet>