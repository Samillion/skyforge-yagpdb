{{ $region := "EU" }}
{{ $colorPurple := 5793266 }}
{{ $out := "Void" }}
{{ $post := false }}
{{ $tz := "Europe/Berlin" }}
{{ $showTz := "CET" }}
{{ if ( eq ( lower $region ) "na" ) }}
	{{ $tz = "US/Pacific" }}
	{{ $showTz = "PST" }}
{{ else }}
	{{ $tz = "Europe/Berlin" }}
	{{ $showTz = "CET" }}
{{ end }}
{{ $location := (newDate 0 0 0 0 0 0 $tz).Location }}
{{ $time := currentTime.In $location }}
{{ $hour := ( toInt ( $time.Format "15" ) ) }}
{{ $day := ( $time.Format "Monday" ) }}
{{ $at := 0 }}
{{/* Trick bot to think it's tomorrow noon if current hour is more than 21. Kiss my ass, Laama! */}}
{{ if and ( ge $hour 21 ) ( le $hour 23 ) }}
	{{ $day = ( $time.AddDate 0 0 1 ).Weekday.String }}
	{{ $hour = 14 }}
	{{ $at = 14 }}
	{{ $post = true }}
{{ end }}
{{/* Boring and lazy logic to determine map type */}}
{{ if ( eq $day "Saturday" "Sunday" ) }}
		{{ if and ( le $hour 14 ) ( lt $hour 17 ) }}
			{{ $at = 14 }}
		{{ else if and ( le $hour 17 ) ( lt $hour 20 ) }}
			{{ $at = 17 }}
		{{ else if and ( le $hour 20 ) ( gt $hour 17 ) }}
			{{ $at = 20 }}
		{{ end }}
        {{ $out = "All" }}
        {{ $post = true }}
{{ else if and ( le $hour 14 ) ( lt $hour 17 ) }}
    {{ if ( eq $day "Monday" ) }}
        {{ $out = "10v10" }}
    {{ else if ( eq $day "Tuesday" ) }}
        {{ $out = "3v3" }}
    {{ else if ( eq $day "Wednesday" ) }}
        {{ $out = "10v10" }}
    {{ else if ( eq $day "Thursday" ) }}
        {{ $out = "5v5" }}
    {{ else if ( eq $day "Friday" ) }}
        {{ $out = "3v3" }}
    {{ end }}
	{{ $at = 14 }}
    {{ $post = true }}
{{ else if and ( le $hour 17 ) ( lt $hour 20 ) }}
    {{ if ( eq $day "Monday" ) }}
        {{ $out = "5v5" }}
    {{ else if ( eq $day "Tuesday" ) }}
        {{ $out = "5v5" }}
    {{ else if ( eq $day "Wednesday" ) }}
        {{ $out = "3v3" }}
    {{ else if ( eq $day "Thursday" ) }}
        {{ $out = "10v10" }}
    {{ else if ( eq $day "Friday" ) }}
        {{ $out = "5v5" }}
    {{ end }}
	{{ $at = 17 }}
    {{ $post = true }}
{{ else if and ( le $hour 20 ) ( gt $hour 17 ) }}
    {{ if ( eq $day "Monday" ) }}
        {{ $out = "3v3" }}
    {{ else if ( eq $day "Tuesday" ) }}
        {{ $out = "10v10" }}
    {{ else if ( eq $day "Wednesday" ) }}
        {{ $out = "5v5" }}
    {{ else if ( eq $day "Thursday" ) }}
        {{ $out = "3v3" }}
    {{ else if ( eq $day "Friday" ) }}
        {{ $out = "10v10" }}
    {{ end }}
	{{ $at = 20 }}
    {{ $post = true }}
{{ else }}
    {{ $out = "Not set" }}
{{ end }}
{{ if $post }}
	{{ $descP := ( joinStr "" "Next active map type: **" $out "** at **" $at ":00 " $showTz "**." ) }}
	{{ $embedPvPmain := cembed 
		"title" "PvP"
		"description" $descP
		"image" (sdict "url" "https://i.imgur.com/QDqUvJA.jpeg")
		"color" $colorPurple
		"footer" (sdict "text" "Sub-option: -pvp next")
	}}
	{{ if ( eq (len .Args) 2 ) }}
		{{ $pvpOpt := ( cslice "next" ) }}
		{{ $pvpArg := index .CmdArgs 0 }}
		{{ $pvpArg := ( lower $pvpArg ) }}
		{{ if in $pvpOpt $pvpArg }}
			{{ if ( eq $pvpArg "next" ) }}
				{{ $embedPvP := cembed 
					"title" "PvP"
					"description" $descP
					"color" $colorPurple
				}}
				{{ sendMessage nil $embedPvP }}
			{{ else }}
				{{ sendMessage nil $embedPvPmain }}
			{{ end }}
		{{ end }}
	{{ else }}
		{{ sendMessage nil $embedPvPmain }}
	{{ end }}
{{ else }}
	{{ $embedPvP := cembed 
		"title" "PvP"
		"image" (sdict "url" "https://i.imgur.com/QDqUvJA.jpeg")
		"color" $colorPurple
	}}
	{{ sendMessage nil $embedPvP }}
{{ end }}
