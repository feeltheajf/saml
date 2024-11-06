package saml

import (
	"html/template"
)

var defaultResponseTemplate = template.Must(template.New("").Parse(`<html>
<form action="{{ .URL }}" method="post" id="form">
<input type="hidden" name="SAMLResponse" value="{{ .SAMLResponse }}" />
<input type="hidden" name="RelayState" value="{{ .RelayState }}" />
<input type="submit" value="Continue" />
</form>
<script>document.getElementById('form').submit();</script>
</html>`))

var defaultRequestTemplate = template.Must(template.New("").Parse(`<html>
<form action="{{ .URL }}" method="post" id="form">
<input type="hidden" name="SAMLRequest" value="{{ .SAMLRequest }}" />
<input type="hidden" name="RelayState" value="{{ .RelayState }}" />
<input type="submit" value="Continue" />
</form>
<script>document.getElementById('form').submit();</script>
</html>`))
