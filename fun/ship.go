{{/*
  Trigger Type: Command
  Trigger: ship
  Source: https://github.com/NiteLiite/YagPDB-CC-s/
  Copyright MIT
*/}}

{{/* CONFIGURABLE VALUES */}}
{{$filtering := true}} {{/* Set this to false if you don't want possible slurs to be filtered out */}}
{{/* END OF CONFIGURABLE VALUES */}}

{{if eq (len .CmdArgs) 2}}
  {{$rand := randInt 101}}

  {{$max1 := div (len (index .CmdArgs 0)) 2}}
  {{$max2 := div (len (index .CmdArgs 1)) 2}}
  {{$output := joinStr "" (slice (index .CmdArgs 0) 0 $max1) "" (slice (index .CmdArgs 1) $max2)}}
  {{$output = reReplace `\s` $output ""}}

  {{$score := ""}}
  {{$emote := ""}}
  {{$word := adjective}}
    {{if and (ge $rand 80) (not (lt $rand 80)) (not (eq $rand 100))}}
      {{$score = "Amazing! It's a match made in Heaven!"}}
      {{$emote = "💖"}}
    {{else if (and (ge $rand 60) (not (lt $rand 60)) (not (ge $rand 80)) )}}
      {{$score = "How sweet! What a good couple!"}}
      {{$emote = "💕"}}
    {{else if (and (ge $rand 40) (not (lt $rand 40)) (not (ge $rand 60)) )}}
      {{$score = "It could use some work, but it's still good!"}}
      {{$emote = "❤️"}}
    {{else if (and (ge $rand 20) (not (lt $rand 20)) (not (ge $rand 40)) )}}
      {{$score = "It's... ok, sort of..."}}
      {{$emote = "💔"}}
    {{else if (and (lt $rand 20) (not (eq $rand 0))) }}
      {{$score = "Aw hell nah"}}
      {{$emote = "💀"}}
    {{else if (eq $rand 100)}}
      {{$score = `N-NOOOO!!! YOUR DATING POWER...!!!`}}
      {{$emote = "💛"}}
    {{else if (eq $rand 0)}}
      {{$word = "hollow"}}
      {{$score = `Connection terminated.`}}
      {{$emote = "<:gasterIsSus:970850477820481536>"}}
    {{end}}

  {{$embed := cembed
      "author" (sdict "name" (print "Invoked by " .User) "icon_url" (print (.Member.AvatarURL "256")))
      "description" (print "A word that could be used to describe this couple is... \"" $word ".\"")
      "title" "Ship Machine"
      "fields" (cslice
      (sdict "name" "The Shipped:" "value" (print "```" (title (index .CmdArgs 0)) " and " (title (index .CmdArgs 1)) "```") "inline" true)
      (sdict "name" "Ship Name:" "value" (print $emote " " $output " " $emote) "inline" true)
      	)
      "footer" (sdict "text" (print "Score: " $rand "%  |  " $score))
      "color" 16746162
  }}

  {{if and $filtering (not (reFind `(?i)(ni+g|fa+g|ta+rd|g{2,}er|ki+ke|gg+a|retar|^ra+pe)` $output))}}
    {{sendMessage nil $embed}}
  {{else}}
    {{print "I found something bad in that combination..."}}
  {{end}}
{{else}}
{{print "Not enough/too many arguments passed!\n**Usage:** `" .ServerPrefix "ship <arg1> <arg2>`\nFor multi-word args, use \"quotes like this.\""}}
{{end}}
