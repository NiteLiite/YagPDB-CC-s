{{/*
    Trigger Type: Regex
    Trigger: \A(-|<@!?204255221017214977>\s*)(r|rand(om)?)?((set)?nick(name)?)(help)?(\s+|\z)
    
    This is a fun nickname command with some extra features! Type -nickhelp to see how to use it!
*/}}

{{$admin := roleid}} {{/* STAFF ROLE ABOVE YAG */}}
{{$botChannel := channelid}} {{/* BOT CHANNEL ID */}}

{{$random := reFind `(?i)(r|rand(om)?)` .Cmd}}
{{$help := (inFold .Cmd "help")}}

{{if not $help}}
  {{deleteResponse 15}}
  {{deleteTrigger 0}}
{{end}}

{{if $random}}
  {{$r := title (joinStr " " adjective noun)}}
  {{if not (hasRoleID $admin)}}
    {{if not (ge (len $r) 33)}}
      {{editNickname $r}}
      {{print "**" .User.Username "**, your nickname has been changed to: `" $r "`"}}
    {{else}}
      {{editNickname (joinStr "" (slice $r 0 29) "...")}}
      {{print "**" .User.Username "**, your nickname has been set to: `" (joinStr "" (slice $r 0 29) "...") "`"}}
    {{end}}
  {{else}}
    {{print "Sorry, I can't change your name, but it would have been: `" $r "`"}}
  {{end}}
{{else}}
  {{with .StrippedMsg}}
    {{if not (hasRoleID $admin)}}
      {{if not (ge (len .) 33)}}
        {{editNickname .}}
        {{print "**" $.User.Username "**, your nickname has been set to: `" . "`"}}
      {{else}}
        {{editNickname (joinStr "" (slice . 0 29) "...")}}
        {{print "**" $.User.Username "**, your nickname has been set to: `" (joinStr "" (slice . 0 29) "...") "`"}}
      {{end}}
    {{else}}
      {{print "Your perms are too strong!"}}
    {{end}}
  {{else if $help}}
    {{$e := (cembed
        "title" "Nickname Command Usage:"
        "fields" (cslice
                   (sdict "name" "Normal:" "value" (print "```\n" .ServerPrefix (joinStr (print "\n" .ServerPrefix) "nick" "setnick" "nickname" "setnickname") "```") "inline" false)
                   (sdict "name" "Random:" "value" (print "```\n" (joinStr "\n" "r" "rand" "random") "```") "inline" false)
		             )
        "description" (print "To **reset** your nickname, simply type one of the \"**Normal:**\" CMDs.\n\nTo use a **random** nickname, prepend any of the \"**Random:**\" CMDs to any of the \"**Normal:**\" CMDs, *(e.g.* **" .ServerPrefix "__rand__nickname** *)*.\n\n** **")
        "footer" (sdict "text" (joinStr " " "Invocation by:" .User))
    )}}
    {{if not (or (eq .Channel.ID $botChannel) (hasRoleID $admin))}}
      {{sendMessage $botChannel (complexMessage "content" .User.Mention "embed" $e)}}
    {{else}}
      {{sendMessage nil $e}}
    {{end}}
  {{else}}
    {{if .Member.Nick}}
      {{editNickname ""}}
      {{print "**" .User.Username "**, your nickname has been reset. For more info on how to use this command, type `-nickhelp`"}}
    {{else}}
      {{print "**" .User.Username "**, you don't have a nickname! For more info on how to use this command, type `-nickhelp`"}}
    {{end}}
  {{end}}
{{end}}
