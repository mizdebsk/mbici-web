package main

import (
	"html/template"
)

var Template = template.Must(template.New("tpl").Parse(`
<!doctype html>
<html lang="en">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">
    <meta name="generator" content="">
    <title>MBICI</title>
    <link href="https://mbi-artifacts.s3.eu-central-1.amazonaws.com/static/fontawesome/css/all.min.css" type="text/css" rel="stylesheet" />
    <link href="https://mbi-artifacts.s3.eu-central-1.amazonaws.com/static/bootstrap/css/bootstrap.min.css" type="text/css" rel="stylesheet" />
    <link href="https://mbi-artifacts.s3.eu-central-1.amazonaws.com/static/custom.css" rel="stylesheet" />
  </head>
  <body>
    <nav class="navbar navbar-expand-md navbar-dark bg-primary fixed-top">
      <div class="container-fluid">
        <svg class='logo' style='font-size:65px' height='47' width='150'>
          <svg y='0' height='4' width='150' viewbox='0 0 150 4'><text x='-6' y='47'>MBI</text></svg>
          <svg y='6' height='4' width='150' viewbox='0 6 150 4'><text x='-6' y='47'>MBI</text></svg>
          <svg y='12' height='4' width='150' viewbox='0 12 150 4'><text x='-6' y='47'>MBI</text></svg>
          <svg y='18' height='4' width='150' viewbox='0 18 150 4'><text x='-6' y='47'>MBI</text></svg>
          <svg y='24' height='4' width='150' viewbox='0 24 150 4'><text x='-6' y='47'>MBI</text></svg>
          <svg y='30' height='4' width='150' viewbox='0 30 150 4'><text x='-6' y='47'>MBI</text></svg>
          <svg y='36' height='4' width='150' viewbox='0 36 150 4'><text x='-6' y='47'>MBI</text></svg>
          <svg y='42' height='4' width='150' viewbox='0 42 150 4'><text x='-6' y='47'>MBI</text></svg>
        </svg>
	<a class="navbar-brand" href="/">MBICI</a>
	<div class="collapse navbar-collapse" id="navbarsExampleDefault"></div>
      </div>
    </nav>
    <main class='container'>

<svg
   width="1270"
   height="220"
   viewBox="0 0 336.02082 58.208335"
   version="1.1"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:svg="http://www.w3.org/2000/svg">
  <style>
    path {
	fill:none;stroke:#000000;stroke-width:0.264583px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1;
    }
    .state-Succeeded path {
	fill: lightgreen
    }
    .state-Failed path {
	fill: red
    }
    .left {
        font-size:3.52778px;
	stroke-width:0.264583;
	text-align:start;
	text-anchor:start;
    }
    .right {
        font-size:3.52778px;
	stroke-width:0.264583;
	text-align:end;
	text-anchor:end;
    }
    .caption {
	font-size:7.76111px;
	line-height:1.25;
	font-family:sans-serif;
	text-align:center;
	text-anchor:middle;
	stroke-width:0.264583;
	font-weight:bold;
	font-size:7.76111px;
	stroke-width:0.264583;
    }
    .state {
	font-size:4.93889px;
	line-height:1.25;
	font-family:sans-serif;
	text-align:center;
	text-anchor:middle;
	stroke-width:0.264583;
    }
  </style>

  <g>
    <g class="state-{{ .SCM.State }}">
      <path d="M 2.6458333,55.562499 H 55.562499 v -31.75 H 2.6458333 Z"/>
      <path d="M 2.6458333,23.812499 V 2.6458333 H 55.562499 l 5.291667,10.5833327 -5.291667,10.583333 z"/>
    </g>
    <g class="state-{{ .SRPM.State }}">
      <path d="M 58.208333,55.562499 H 111.125 v -31.75 H 58.208333 Z"/>
      <path d="M 58.208333,2.6458333 63.499999,13.229166 58.208333,23.812499 H 111.125 L 116.41666,13.229166 111.125,2.6458333 Z"/>
    </g>
    <g class="state-{{ .JPB.State }}">
      <path d="m 113.77083,55.562499 h 52.91667 v -31.75 h -52.91667 z"/>
      <path d="m 113.77083,2.6458333 5.29167,10.5833327 -5.29167,10.583333 H 166.6875 L 171.97916,13.229166 166.6875,2.6458333 Z"/>
    </g>
    <g class="state-{{ .Bootstrap.State }}">
      <path d="M 169.33333,55.562499 H 222.25 v -31.75 h -52.91667 z"/>
      <path d="M 169.33333,2.6458333 174.625,13.229166 169.33333,23.812499 H 222.25 L 227.54166,13.229166 222.25,2.6458333 Z"/>
    </g>
    <g class="state-{{ .Rebuild.State }}">
      <path d="m 224.89583,55.562499 h 52.91667 v -31.75 h -52.91667 z"/>
      <path d="m 224.89583,2.6458333 5.29167,10.5833327 -5.29167,10.583333 H 277.8125 L 283.10416,13.229166 277.8125,2.6458333 Z"/>
    </g>
    <g class="state-{{ .Validate.State }}">
      <path d="M 280.45833,55.562499 H 333.375 v -31.75 h -52.91667 z"/>
      <path d="M 280.45833,23.812499 285.75,13.229166 280.45833,2.6458333 H 333.375 V 23.812499 Z"/>
    </g>

    <g class="caption">
      <text x="29.104166" y="15.875">SCM</text>
      <text x="87.3125" y="15.875">SRPM</text>
      <text x="142.875" y="15.875">JPB</text>
      <text x="198.4375" y="15.875">Bootstrap</text>
      <text x="254" y="15.875">Rebuild</text>
      <text x="309.5625" y="15.875">Validate</text>
    </g>

    <g class="state">
      <text x="29.104166" y="31.749998">{{ .SCM.State }}</text>
      <text x="84.666664" y="31.749998">{{ .SRPM.State }}</text>
      <text x="140.22916" y="31.749998">{{ .JPB.State }}</text>
      <text x="195.79166" y="31.749998">{{ .Bootstrap.State }}</text>
      <text x="251.35416" y="31.749998">{{ .Rebuild.State }}</text>
      <text x="306.91666" y="31.749998">{{ .Validate.State }}</text>
    </g>

    <g class="left">
      <text x="14.811508" y="39.687496">Total tasks:</text>
      <text x="14.811508" y="44.097221">Pending:</text>
      <text x="14.811508" y="48.506947">Failed:</text>
      <text x="14.811508" y="52.916672">Complete:</text>
      <text x="70.374008" y="39.687496">Total tasks:</text>
      <text x="70.374008" y="44.097221">Pending:</text>
      <text x="70.374008" y="48.506947">Failed:</text>
      <text x="70.374008" y="52.916672">Complete:</text>
      <text x="125.93651" y="39.687496">Total tasks:</text>
      <text x="125.93651" y="44.097221">Pending:</text>
      <text x="125.93651" y="48.506947">Failed:</text>
      <text x="125.93651" y="52.916672">Complete:</text>
      <text x="181.49901" y="39.687496">Total tasks:</text>
      <text x="181.49901" y="44.097221">Pending:</text>
      <text x="181.49901" y="48.506947">Failed:</text>
      <text x="181.49901" y="52.916672">Complete:</text>
      <text x="237.06151" y="39.687496">Total tasks:</text>
      <text x="237.06151" y="44.097221">Pending:</text>
      <text x="237.06151" y="48.506947">Failed:</text>
      <text x="237.06151" y="52.916672">Complete:</text>
      <text x="292.62399" y="39.687496">Total tasks:</text>
      <text x="292.62399" y="44.097221">Pending:</text>
      <text x="292.62399" y="48.506947">Failed:</text>
      <text x="292.62399" y="52.916672">Complete:</text>
    </g>
    <g class="right">
      <text x="42.160404" y="39.687496">{{ .SCM.NTasks }}</text>
      <text x="42.160404" y="44.097221">{{ .SCM.NPending }}</text>
      <text x="42.160404" y="48.506947">{{ .SCM.NFailed }}</text>
      <text x="42.160404" y="52.916672">{{ .SCM.NComplete }}</text>
      <text x="97.7229" y="39.687496">{{ .SRPM.NTasks }}</text>
      <text x="97.7229" y="44.097221">{{ .SRPM.NPending }}</text>
      <text x="97.7229" y="48.506947">{{ .SRPM.NFailed }}</text>
      <text x="97.7229" y="52.916672">{{ .SRPM.NComplete }}</text>
      <text x="153.2854" y="39.687496">{{ .JPB.NTasks }}</text>
      <text x="153.2854" y="44.097221">{{ .JPB.NPending }}</text>
      <text x="153.2854" y="48.506947">{{ .JPB.NFailed }}</text>
      <text x="153.2854" y="52.916672">{{ .JPB.NComplete }}</text>
      <text x="208.8479" y="39.687496">{{ .Bootstrap.NTasks }}</text>
      <text x="208.8479" y="44.097221">{{ .Bootstrap.NPending }}</text>
      <text x="208.8479" y="48.506947">{{ .Bootstrap.NFailed }}</text>
      <text x="208.8479" y="52.916672">{{ .Bootstrap.NComplete }}</text>
      <text x="264.4104" y="39.687496">{{ .Rebuild.NTasks }}</text>
      <text x="264.4104" y="44.097221">{{ .Rebuild.NPending }}</text>
      <text x="264.4104" y="48.506947">{{ .Rebuild.NFailed }}</text>
      <text x="264.4104" y="52.916672">{{ .Rebuild.NComplete }}</text>
      <text x="319.9729" y="39.687496">{{ .Validate.NTasks }}</text>
      <text x="319.9729" y="44.097221">{{ .Validate.NPending }}</text>
      <text x="319.9729" y="48.506947">{{ .Validate.NFailed }}</text>
      <text x="319.9729" y="52.916672">{{ .Validate.NComplete }}</text>
    </g>
  </g>
</svg>

<h2>Workflow details</h2>

{{ with .Workflow }}

<ul>
{{ range .Tasks }}
<li>
Task: <strong>{{.Id}}</strong><br/>
Handler: {{.Handler}}<br/>

{{ with .Dependencies }}
Dependencies:<br/>
<ul>
{{ range . }}
<li>{{.}}</li>
{{ end }}
</ul>
{{ else }}
No dependencies<br/>
{{ end }}

{{ with .Parameters }}
Parameters:<br/>
<ul>
{{ range . }}
<li>{{.Name}}: {{.Value}}</li>
{{ end }}
</ul>
{{ else }}
No parameters<br/>
{{ end }}

{{ with .Result }}

State: finished<br/>
Time started: {{.TimeStarted}}<br/>
Time finished: {{.TimeFinished}}<br/>
Outcome: {{.Outcome}}<br/>
Outcome reason: {{.OutcomeReason}}<br/>

{{ with .Artifacts }}
Artifacts:<br/>
<ul>
{{ range . }}
<li>{{.Type}}: {{.Name}}</li>
{{ end }}
</ul>
{{ else }}
No artifacts<br/>
{{ end }}

{{ else }}
State: not finished<br/>
{{ end }}

</li>
{{ end }}
</ul>

{{ end }}

    </main>
    <script src="https://mbi-artifacts.s3.eu-central-1.amazonaws.com/static/bootstrap/js/bootstrap.bundle.min.js"></script>
  </body>
</html>
`))
