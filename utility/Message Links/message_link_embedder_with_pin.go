{{/*

    Made by NiteLiite#0001
    Based on BlackWolf's original message link CC and uses similar techniques --> https://github.com/BlackWolfWoof/yagpdb-cc/blob/master/Misc/message_link.yag
    
    Copyright @ MIT 2022
    
    Trigger Type: Regex
    Trigger: (\A-pin\s*)?(?:[^<]|\A)https:\/\/(?:\w+\.)?discord(?:app)?\.com\/channels\/(\d+)\/(\d+)\/(\d+)(?:[^>\d]|\z)|(\A-pin(\s+|\z))
    
    Usage: Send a message link to have it embed the content, or pin a message by prepending the link with -pin, i.e. -pin <message:link>
    
*/}}
    

{{/* CONFIGURABLE VALUES */}} 
{{$mod_role_ids := cslice 123 456}} {{/* Roles that have perms to bypass ignored channels and pin messages */}}
{{$ignoredChannels := cslice 123 456 789}} {{/* Channel IDs to not display the content of messages from. */}} 
{{$pinChannel := 123}} {{/* Channel ID to send pinned messages to. */}}
{{/* END OF CONFIGURABLE VALUES */}}

{{$index := reFindAllSubmatches `(?i)https://(?:\w+\.)?discord(?:app)?\.com/channels\/(\d+)\/(\d+)\/(\d+)` .Message.Content}}
{{$pinning := false}}
{{$bypass := false}}
{{$where := .Channel.ID}}
{{range .Member.Roles}}
  {{if in $mod_role_ids .}}
    {{$bypass = true}}
  {{end}}
{{end}}
{{if and $bypass (reFind `\A-(?i)(pin)` .Message.Content)}}
  {{$pinning = true}}
  {{$where = $pinChannel}}
{{end}}
{{if and $pinning (not $index)}}
  {{sendMessage nil (cembed "description" "***❌ Err: Message link required.***" "color" 16711680 "image" (sdict "url" "https://media.discordapp.net/attachments/881835799547609109/986180455625818112/ezgif.com-gif-maker.gif") "footer" (sdict "text" "-pin <message:link>") )}}
{{else if $index}}

{{$link := index $index 0 0}}
{{$guild_id := index $index 0 1|toInt}}
{{$channel_id := index $index 0 2|toInt}}
{{$message_id := index $index 0 3|toInt}}
{{$timestamp := ((snowflakeToTime (toInt $message_id)).Format "January 02, 2006")}}
{{$footer := printf "Invocation by %s. Message from %s." .User $timestamp}}

{{if not (eq (toInt64 $guild_id) .Guild.ID)}}
  {{(cembed
    "author" (sdict
    "name" "Unknown User"
    "icon_url" "https://i.imgur.com/jNz2Dwp.png")
    "description" (print "\n\n**[Message](" $link "/) in <#" $channel_id ">**\n" "<:excl:565142262401728512> Unknown Message")
    "color" 0xF04747
    "footer" (sdict
    "text" $footer)
  ) | sendMessage nil }}
{{else}}
  {{$msg := getMessage $channel_id $message_id}}
  {{if not $msg}}
    {{(cembed "color" 0xDD2E44 "description" "❌ `That message does not exist.`") | sendMessage nil}}
  {{else}}
    {{$col := 3092790}}
    {{$av := print "https://cdn.discordapp.com/embed/avatars/" (randInt 0 6) ".png?size=256"}}
    {{if (getMember $msg.Author.ID)}}
      {{$av = (getMember $msg.Author.ID).AvatarURL "256"}}
      {{$pos := 0}}{{$r := (getMember $msg.Author.ID).Roles}}
      {{- range .Guild.Roles}}
        {{- if and (in $r .ID) (.Color) (lt $pos .Position) -}}
          {{- $pos = .Position -}}
          {{- $col = .Color -}}
        {{- end -}}
      {{- end -}}
    {{end}}
    {{if or (and (eq $msg.Content "") (not $msg.Embeds) (not $msg.Attachments)) (and (not $bypass) (in $ignoredChannels $channel_id))}}
      {{(cembed
        "author" (sdict "name" $msg.Author.String "icon_url" $av)
        "description" (print "\n\n**[Message](" $link "/) in <#" $channel_id ">**\n" "`Bot: No Content Displayed`")
        "footer" (sdict "text" $footer)
        "color" $col
      ) | sendMessage $where}}
    {{else}}
      {{$e := sdict
        "color" $col
        "author" (sdict "name" (print $msg.Author.String " (ID " $msg.Author.ID ")") "icon_url" $av)
        "description" (printf "**[Message](%s) in** <#%d>" $link $channel_id)
        "footer" (sdict "text" $footer) 
        "fields" (cslice )
      }}

      {{with $msg.Content}}
        {{if lt (len .) 2001}}
          {{$e.Set "description" (joinStr "\n" ($e.Get "description") .)}}
        {{else}}
          {{$e.Set "description" (joinStr "\n" ($e.Get "description") (print (slice . 0 1993) "`[...]`"))}}
        {{end}}
      {{end}}
      {{with $msg.ReferencedMessage}}
        {{$e.Set "fields" (($e.Get "fields").Append (sdict
          "name" "Replied To"
          "value" (print "**[" .Author.String ":](https://discord.com/channels/" $.Guild.ID "/" $channel_id "/" .ID ")** " .Content) "inline" false)
        )}}
      {{end}}

      {{if $pinning}}
        {{$e.Set "title" "New pinned message!"}}
        {{$e.Set "thumbnail" (sdict "url" "https://media.discordapp.net/attachments/881835799547609109/988190263900012615/pin.png")}}
        {{sendMessage nil (cembed "color" $col "description" (print "**" (or .Member.Nick .User.Username) "** has pinned a [**message**](" $msg.Link ") to <#" $where ">**!**"))}}
      {{end}}

      {{$total := add (len $msg.Embeds) (len $msg.Attachments)}}{{$add := 0}}
      {{$imagePattern := `(?i)\.(jpg|jpeg|png|gif|webp)\z`}}

      {{range $msg.Embeds}}
        {{if or (or .Title .Description .Author .Footer .Fields) (eq .Type "rich")}}
          {{$total = sub $total 1}}
          {{if eq .Type "rich"}}
            {{sendMessage $where (cembed (. | structToSdict))}}
          {{else}}
            {{sendMessage $where .URL}}
          {{end}}
        {{end}}
      {{end}}
      {{if eq $add $total}}
        {{sendMessage $where (cembed $e)}}
      {{else}}
        {{range $msg.Embeds}}
          {{if and (ne .Type "rich") (not (or .Title .Description .Author .Footer .Fields)) (or .Image .Thumbnail (eq .Type "video"))}}
            {{$add = add $add 1}}
            {{if reFind $imagePattern .URL}}
              {{$e2 := sdict "image" (sdict "url" .URL) "color" $col "footer" (sdict "text" (print $add "/" $total " - " .Type))}}
              {{if eq $add 1}}
                {{$e.Set "image" (sdict "url" .URL)}}
                {{if ne $add $total}}
                  {{$e.Set "footer" (sdict "text" (print $add "/" $total " - " .Type))}}{{else}}{{$e.Set "footer" (sdict "text" (print $add "/" $total " - " .Type "\n" $footer))}}
                {{end}}
                {{sendMessage $where (cembed $e)}}
              {{else}}
                {{if eq $add $total}}
                  {{$e2.Set "footer" (sdict "text" (print $add "/" $total " - " .Type "\n" $footer))}}
                {{end}}
                {{sendMessage $where (cembed $e2)}}
              {{end}}
            {{else}}
              {{if eq $add 1}}
                {{$e.Del "footer"}}
                {{sendMessage $where (cembed $e)}}
              {{end}}
                {{sendMessage $where .URL}}
              {{if eq $add $total}}
                {{sendMessage $where (cembed "footer" (sdict "text" (print $add "/" $total " - " .Type "\n" $footer)) "color" $col)}}
              {{end}}
            {{end}}
          {{end}}
        {{end}}
        {{range $msg.Attachments}}
          {{$add = add $add 1}}
          {{if reFind $imagePattern .URL}}
            {{$e2 := sdict "image" (sdict "url" .URL) "color" $col "footer" (sdict "text" (print $add "/" $total " - " .Filename))}}
            {{if eq $add 1}}
              {{$e.Set "image" (sdict "url" .URL)}}
              {{if ne $add $total}}
                {{$e.Set "footer" (sdict "text" (print $add "/" $total " - " .Filename))}}{{else}}{{$e.Set "footer" (sdict "text" (print $add "/" $total " - " .Type "\n" $footer))}}
              {{end}}{{sendMessage $where (cembed $e)}}
            {{else}}
              {{if eq $add $total}}
                {{$e2.Set "footer" (sdict "text" (print $add "/" $total " - " .Filename "\n" $footer))}}
              {{end}}{{sendMessage nil (cembed $e2)}}
            {{end}}
          {{else if reFind `(?i)\.(mov|mp4|webm)\z` .URL}}
            {{$e.Del "footer"}}
            {{if eq $add 1}}
              {{sendMessage $where (cembed $e)}}
            {{end}}
              {{sendMessage $where .URL}}
              {{if eq $add $total}}
                {{sendMessage $where (cembed "footer" (sdict "text" (print $add "/" $total " - " .Filename "\n" $footer)) "color" $col)}}
              {{end}}
          {{else}}
            {{if eq $add 1}}
              {{$e.Del "footer"}}
              {{sendMessage $where (cembed $e)}}
            {{end}}
            {{$f := sdict 
              "color" $col 
              "fields" (cslice
                         (sdict "name" "File:" "value" (printf "[%s](%s)" .Filename .URL) "inline" false)
                         (sdict "name" "Size:" "value" (print "`" (div .Size 1000) "KB, " .Size "B`") "inline" false)
                       )
              "footer" (sdict "text" (print $add "/" $total))
            }}
            {{if eq $add $total}}
              {{$f.Set "footer" (sdict "text" (print $add "/" $total " - " .Filename "\n" $footer))}}
            {{end}}
            {{sendMessage $where (cembed $f)}}
          {{end}}
        {{end}}
      {{end}}

      {{if and (eq (len $link) (len .Message.Content)) (not .Message.Attachments)}}
        {{deleteTrigger 1}}
      {{end}}

    {{end}}
  {{end}}
{{end}}

{{if and (eq (len $link) (len .Message.Content)) (not .Message.Attachments)}}
  {{deleteTrigger 1}}
{{end}}

{{end}}
