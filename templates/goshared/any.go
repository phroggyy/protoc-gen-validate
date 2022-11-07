package goshared

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	if a := {{ accessor . }}; a != nil {
		{{ if $r.In }}
			if _, ok := {{ lookup $f "InLookup" }}[a.GetTypeUrl()]; !ok {
				err := {{ err . "any.in" "type URL must be in list " $r.In }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else if $r.NotIn }}
			if _, ok := {{ lookup $f "NotInLookup" }}[a.GetTypeUrl()]; ok {
				err := {{ err . "any.not_in" "type URL must not be in list " $r.NotIn }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	}
`
