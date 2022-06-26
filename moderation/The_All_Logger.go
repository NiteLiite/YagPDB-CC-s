{{/*

	This custom command logs EVERY attachment and link sent (each only for a limited period of time due to discord limitations). 
	 
	Trigger Type: Regex
	Trigger: \A
	
	Copyright MIT 2022 NiteLiite
	Repo: https://github.com/NiteLiite/YagPDB-CC-s
	
*/}}

{{/* CHANNEL TO SEND LOGS TO */}}
{{$channel := 981278657328857150}}
{{/* END */}}
 
{{$msg := reReplace `<@!?\d{17,}>` .Message.Content "`<mention>`"}}
{{$mentions := ""}}
{{$content := ""}}
{{if .Message.Mentions}}
{{range .Message.Mentions}}
  {{- $mentions = print $mentions .String " - " .ID "\n" -}}
{{end}}
  {{$mentions = print "\n**Mentioned Users:**\n" $mentions}}
{{end}}
{{if .Message.Content}}
   {{$content = (print "\n**Content:**\n" $msg "\n")}}
{{end}}
 
{{$desc := ""}}
{{if .Message.Attachments}}
  {{range .Message.Attachments}}
    {{- $desc = print $desc .URL "\n" -}}
  {{end}}
  {{sendMessage $channel (print "** **\n>>> **__ATTACHMENT__** sent by `" .User "` | *(ID " .User.ID ")*:\n**Channel:** <#" .Channel.ID ">" $content $mentions "\n" $desc)}}
{{else if and (not .Message.Attachments) (reFind .LinkRegex .Message.Content) (not (reFind `(?:[^<]|\A)https://(?:\w+\.)?discord(?:app)?\.com/channels\/(\d+)\/(\d+)\/(\d+)(?:[^>\d]|\z)` .Message.Content)) }}
  {{sendMessage $channel (print "** **\n>>> **__LINK__** sent by `" .User "` | *(ID " .User.ID ")*:\n**Channel:** <#" .Channel.ID ">" $content $mentions)}}
{{end}}
