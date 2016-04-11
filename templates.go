package icecubeapp

import "html/template"

var writeSigsProtoTemplate = template.Must(template.New("wSigsP").Parse(
	`{{range .}}
    set_{{.SigName}} {<br>
    &emsp;out "{{.SerialStr}}{{.DataType}}\n";<br>
    &emsp;ExtraInput = Ignore;<br>
    }<br>
  {{end}}
  `))

var readSigsProtoTemplate = template.Must(template.New("rSigsP").Parse(
	`{{range .}}
      get_{{.SigName}} {<br>
      &emsp;out "{{.SerialStr}}";<br>
      &emsp;in "{{.SerialStr}} {{.DataType}}"<br>
      &emsp;ExtraInput = Ignore;<br>
      }<br>
    {{end}}
    `))

var writeSigsDBTemplate = template.Must(template.New("wSigsDB").Parse(
	`{{range .}}
        record(ao, {{.SigName}}:set) {<br>
        &emsp;field(DESC, "{{.SigName}} input")<br>
        &emsp;field(DTYP, "stream")<br>
        &emsp;field(INP, "@arduino.proto set_{{.SigName}}() $(PORT)")<br>
        }<br>
      {{end}}
      `))

var readSigsDBTemplate = template.Must(template.New("rSigsDB").Parse(
	`{{range .}}
        record(ai, {{.SigName}}:get) {<br>
        &emsp;field(DESC, "{{.SigName}} output")<br>
        &emsp;field(DTYP, "stream")<br>
        &emsp;field(INP, "@arduino.proto get_{{.SigName}}() $(PORT)")<br>
        }<br>
      {{end}}
      `))

var definedSigsTemplate = template.Must(template.New("sigs").Parse(
	`<br>
    {{range .}}
      <b>{{.SigName}}</b>
      {{.SerialStr}}
      {{.DataType}}<br>
    {{end}}
    `))
