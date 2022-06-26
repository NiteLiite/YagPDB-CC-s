{{/*
  Trigger Type: Regex
  Trigger \A(-|<@!?204255221017214977>\s*)(server|s|g|guild)?(avatar|av|pfp)(\s+|\z)

  Usage: Type "(prefix)av help" to see how to use this! Also, if your prefix is not - replace the first - in the Trigger with your prefix!
*/}}

{{$server := reFind `(?i)(server|s|g|guild)` .Cmd}}
{{$help := false}}
{{$break := false}}

{{$target := .User}}
{{if ge (len .CmdArgs) 1}}
  {{with (getMember (index .CmdArgs 0))}}
    {{$target = .User}}
    {{if $server}}
      {{$target = .}}
    {{end}}
  {{else if (inFold .CmdArgs "help" )}}
    {{$help = true}}
  {{else if reFind `((?i).*#\d{4})|(\d{17,})` (joinStr " " .CmdArgs) }}
    {{print "Hmm... did you enter that correctly? Try `" .ServerPrefix "pfp help` for help!"}}
    {{$break = true}}
  {{end}}
{{else}}
  {{if $server}}
    {{$target = .Member}}
  {{end}}
{{end}}

{{if not $break}}
  {{if not $help}}
    {{$col := 3092790}}
    {{$dynamic := ""}}
    {{$pos := 0}}
      {{if $server}}
        {{$dynamic = $target}}
      {{else}}
        {{$dynamic = (getMember $target)}}
      {{end}}
    {{$r := $dynamic.Roles}}
      {{range .Guild.Roles}}
        {{- if and (in $r .ID) (.Color) (lt $pos .Position) -}}
          {{- $pos = .Position -}}
          {{- $col = .Color -}}
        {{- end -}}
      {{end}}
    {{$e := sdict
        "title" (print (or $dynamic.Nick $dynamic.User.Username) "'s Avatar:")
        "image" (sdict "url" (print ($target.AvatarURL "4096") ))
        "color" $col
        "footer" (sdict "text" (print "User ID: " $dynamic.User.ID))
    }}
    {{if ne $dynamic.User.ID .User.ID}}
      {{$e.Set "description" (joinStr " " "Invoked by:" .User)}}
    {{end}}
{{sendMessage nil (cembed $e)}}
  {{else}}
    {{sendMessage nil (cembed
        "fields" (cslice
        (sdict "name" "User Avatar CMDs" "value" (joinStr (print "\n" .ServerPrefix) (print .ServerPrefix "av") "avatar" "pfp"))
        (sdict "name" "Server Avatar CMDs" "value" (joinStr "\n" "s" "g" "guild" "server")) 
            )
        "description" (print "To get a user's **server** pfp, prepend any of the \"**Server Avatar CMDs**\" to any of the \"**User Avatar CMDs**\" *(e.g.* **" .ServerPrefix "__" (index (shuffle (cslice "s" "g" "server" "guild")) 0) "__" (index (shuffle (cslice "av" "pfp" "avatar")) 0) "** *)*")
        "title" "Avatar Command Usage:"
        "footer" (sdict "text" (print "Usage: " .ServerPrefix "CMD <member:optional>")) 
    )}}
  {{end}}
{{end}}
