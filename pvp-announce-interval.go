{{ $region := "EU" }}
{{ $out := "Void" }}
{{ $post := false }}
{{ $tz := "Europe/Berlin" }}
{{ if ( eq ( lower $region ) "na" ) }}
	{{ $region = "NA" }}
	{{ $tz = "US/Pacific" }}
{{ else }}
	{{ $region = "EU" }}
	{{ $tz = "Europe/Berlin" }}
{{ end }}
{{ $location := (newDate 0 0 0 0 0 0 $tz).Location }}
{{ $time := currentTime.In $location }}
{{ $hour := toInt ( $time.Format "15" ) }}
{{ $day := $time.Format "Monday" }}
{{ if eq $day "Saturday" "Sunday" }}
	{{ if or ( eq $hour 14 ) ( eq $hour 17 ) ( eq $hour 20 ) }}
		{{ $out = "All" }}
		{{ $post = true }}
	{{ end }}
{{ else if eq $hour 14 }}
	{{ if eq $day "Monday" }}
		{{ $out = "10v10" }}
	{{ else if eq $day "Tuesday" }}
		{{ $out = "3v3" }}
	{{ else if eq $day "Wednesday" }}
		{{ $out = "10v10" }}
	{{ else if eq $day "Thursday" }}
		{{ $out = "5v5" }}
	{{ else if eq $day "Friday" }}
		{{ $out = "3v3" }}
	{{ end }}
	{{ $post = true }}
{{ else if eq $hour 17 }}
	{{ if eq $day "Monday" }}
		{{ $out = "5v5" }}
	{{ else if eq $day "Tuesday" }}
		{{ $out = "5v5" }}
	{{ else if eq $day "Wednesday" }}
		{{ $out = "3v3" }}
	{{ else if eq $day "Thursday" }}
		{{ $out = "10v10" }}
	{{ else if eq $day "Friday" }}
		{{ $out = "5v5" }}
	{{ end }}
	{{ $post = true }}
{{ else if eq $hour 20 }}
	{{ if eq $day "Monday" }}
		{{ $out = "3v3" }}
	{{ else if eq $day "Tuesday" }}
		{{ $out = "10v10" }}
	{{ else if eq $day "Wednesday" }}
		{{ $out = "5v5" }}
	{{ else if eq $day "Thursday" }}
		{{ $out = "3v3" }}
	{{ else if eq $day "Friday" }}
		{{ $out = "10v10" }}
	{{ end }}
	{{ $post = true }}
{{ else }}
	{{ $out = "Not set" }}
{{ end }}
{{ if $post }}
**Skyforge {{ $region }}:** Immortals, it is PvP Happy Hour now. Active map: **{{ $out }}**.
{{ end }}
