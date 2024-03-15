{{ $region := "EU" }}
{{ $image := "https://cdn.discordapp.com/attachments/913303283110649916/920981799834898472/pvp-hh-table-en.jpeg" }}

{{ $tz := "Europe/Berlin" }}
{{ $zone := "CET" }}

{{ if eq (lower $region) "na" }}
	{{ $tz = "US/Pacific" }}
	{{ $zone = "PST" }}
{{ end }}

{{ $time := currentTime.In (loadLocation $tz) }}
{{ $hour := $time.Format "15" | toInt }}
{{ $day := $time.Format "Monday" }}

{{ $post := false }}
{{ $map := "" }}
{{ $at := 0 }}

{{ if and (ge $hour 21) (le $hour 23) }}
	{{ $day = ($time.AddDate 0 0 1).Weekday.String }}
	{{ $hour = 14 }}
	{{ $at = 14 }}
	{{ $post = true }}
{{ end }}

{{ if eq $day "Saturday" "Sunday" }}
		{{ if and (le $hour 14) (lt $hour 17) }}
			{{ $at = 14 }}
		{{ else if and (le $hour 17) (lt $hour 20) }}
			{{ $at = 17 }}
		{{ else if and (le $hour 20) (gt $hour 17) }}
			{{ $at = 20 }}
		{{ end }}
        {{ $map = "All" }}
        {{ $post = true }}
{{ else if and (le $hour 14) (lt $hour 17) }}
	{{ $set := sdict 
		"Monday" "10v10"
		"Tuesday" "3v3"
		"Wednesday" "10v10"
		"Thursday" "5v5"
		"Friday" "3v3"
	}}
	
	{{ if $set.HasKey $day }}
		{{ $map = $set.Get $day }}
	{{ end }}

	{{ $at = 14 }}
    {{ $post = true }}
{{ else if and (le $hour 17) (lt $hour 20) }}
	{{ $set := sdict 
		"Monday" "5v5"
		"Tuesday" "5v5"
		"Wednesday" "3v3"
		"Thursday" "10v10"
		"Friday" "5v5"
	}}
	
	{{ if $set.HasKey $day }}
		{{ $map = $set.Get $day }}
	{{ end }}

	{{ $at = 17 }}
    {{ $post = true }}
{{ else if and (le $hour 20) (gt $hour 17) }}
	{{ $set := sdict 
		"Monday" "3v3"
		"Tuesday" "10v10"
		"Wednesday" "5v5"
		"Thursday" "3v3"
		"Friday" "10v10"
	}}
	
	{{ if $set.HasKey $day }}
		{{ $map = $set.Get $day }}
	{{ end }}

	{{ $at = 20 }}
    {{ $post = true }}
{{ else }}
    {{ $map = "Not set" }}
{{ end }}

{{ $onTime := newDate $time.Year 0 0 $at 0 0 $tz }}
{{ $readable := $onTime.Format "3:04 PM" }}

{{ $fields := cslice 
	(sdict "name" "Next Map" "value" $map "inline" true)
	(sdict "name" "Time" "value" (print $readable " " $zone) "inline" true)
	(sdict "name" "Your Time" "value" (print "<t:" $onTime.Unix ":t>") "inline" true)
}}

{{ $main := sdict 
	"description" "Information about the next map and timings of PvP."
	"fields" $fields
	"image" (sdict "url" $image)
	"footer" (sdict "text" "Sub-option: -pvp next")
}}
{{ $short := sdict "description" "Information about the next PvP map." "fields" $fields }}
{{ $basic := sdict "image" (sdict "url" $image) }}

{{ $embed := $main }}

{{ if $post }}
	{{ if .CmdArgs }}
		{{ $arg := index .CmdArgs 0 | lower }}
		{{ if eq $arg "next" }}
			{{ $embed = $short }}
		{{ end }}
	{{ end }}
{{ else }}
	{{ $embed = $basic }}
{{ end }}

{{ $embed.Set "title" "PVP" }}
{{ $embed.Set "color" 5793266 }}
{{ sendMessage nil (cembed $embed) }}
