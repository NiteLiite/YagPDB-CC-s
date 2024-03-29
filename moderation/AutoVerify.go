{{/*
	Trigger: None
	Trigger Type: None
 
	Description: YAGPDB will detect suspicious accounts based on 4 key algorithms.
 
	Note: Please configure the variables after pasting the code into your control panel Custom Commands. Then, paste {{execCC CCIDHERE anychannelID 0 .ExecData}} into your join feed message (or just paste the full code into your join message).
 
	Repository: https://github.com/NiteLiite/YagPDB-CC-s
 
*/}}
 
{{/* Configurable Values */}}
 
{{$days := 7}} {{/* This will change how many days of account age are required for 3 total points out of 6. (i.e. 1 point per 7, 7 x 5, and 7 x 17 days) */}}
{{$mod_joinlog := 962838110557839420}} {{/* Where regular join messages will go (channel ID) */}}
{{$alerts := 961127009436696626}} {{/* Where suspicious account alerts will go (channel ID) */}}
{{$memberrole := 946624674345943120}} {{/* Put the ID of your member role here */}}
 
{{define "sus-action"}} {{/* <-- Don't edit this one */}}
{{$susaction := "" }} {{/* Replace the "" with the code you want to execute upon members that do not pass verification (e.g: {{execAdmin "kick" .User "Failed verification requirements, rejoin when <issue(s)> have been resolved"}} ) */}}
{{end}} {{/* <-- Don't edit this one either */}}
 
{{$level := 3}} {{/* For this one, there are specific numbers that can (and should) be entered.
			0 = No action on distrusted users.
			1 = Action on extremely distrusted users.
			2 = Action on mostly distrusted users.
			3 = Action on somewhat distrusted users.
			4 = Action on neutral users. (Recommended)
			5 = Action on trusted users. (Anything from this point and beyond is NOT recommended.)
			6 = Action on all users.
		*/}}
 
 
{{/* DO NOT edit any of the code below unless you know what you are doing. */}}
 
{{$age := currentTime.Sub currentUserCreated}}
{{$status := ""}}

    {{try}}
        {{$status = (index (exec "whois" .User.ID).Fields 6).Value}}
    {{catch}}
        {{$status = "Could not retrieve status."}}
    {{end}}
 
{{$createdAt := div .User.ID 4194304 | add 1420070400000 | mult 1000000 | toDuration | (newDate 1970 1 1 0 0 0).Add }}
{{$when := $createdAt.Format "Monday, January 2, 2006 at 3:04 AM"}}
 
  {{$check := sdict
    "Custom avatar" .User.Avatar
    "No numbers"    (not (reFind `[0-9]` .User.Username))
    "No non-ASCII"   (not (reFind `[^[:ascii:]]` .User.Username))
    "Old account"   (not (lt (toInt $age.Minutes) (mult 24 $days 60)))
  }}
 
  {{$trust := 0}}
  {{$desc := ""}}
 
{{range $type, $val := $check}}
    {{- $curr := "❌"}}
    {{- if $val}}{{$curr = "✅"}}{{$trust = add $trust 1}}{{end}}
    {{- $desc = print $desc "\n" $type ": " $curr -}}
{{end}}
 
{{if (not (lt (toInt $age.Minutes) (mult 24 (mult $days 5) 60)))}}
  {{$trust = add $trust 1}}
{{end}}
{{if (not (lt (toInt $age.Minutes) (mult 24 (mult $days 17) 60)))}}
  {{$trust = add $trust 1}}
{{end}}

{{$trustiness := dict
    0 "Completely Distrusted"
    1 "Very Distrusted"
    2 "Distrusted"
    3 "Neutral"
    4 "Trusted"
    5 "Very Trusted"
    6 "Completely Trusted"}}
{{$trustiness = ($trustiness.Get (toInt $trust))}}
{{$color := 4325195 }}
{{$thing := "✅"}}
{{if (eq $trust 3) }}
    {{$color = 16765184}}
    {{$thing = "⚠️"}}
{{end}}
{{if or (eq $trust 0) (eq $trust 1) (eq $trust 2) }}
    {{$thing = "⛔"}}
    {{$color = 16715776}}
{{end}}
 
{{$id := sendMessageRetID $mod_joinlog (complexMessage "content" .User.ID "embed" (cembed 
    "author" (sdict "name" (print .User.Username "#" .User.Discriminator) "icon_url" (print (.User.AvatarURL "256")))
    "description" (print $thing " "  .User.Mention " *(ID " .User.ID ")*" " has joined." " Account created on " $when ". *(" (humanizeDurationMinutes $age) " old)*") 
    "color" $color
    "footer" (sdict "text" (print "User Trust: " (print $trust) "/6  |  " (print $trustiness)))
))}}
 
{{if ge $trust $level}}
    {{try}}
        {{giveRoleID .User.ID $memberrole}}
	{{catch}}
        {{editMessage $mod_joinlog $id (complexMessage "content" .User.ID "embed" (cembed 
          "author" (sdict "name" (print .User.Username "#" .User.Discriminator) "icon_url" (print (.User.AvatarURL "256")))
          "description" (print $thing " "  .User.Mention " *(ID " .User.ID ")*" " has joined." " Account created on " $when ". *(" (humanizeDurationMinutes $age) " old)*\n\n> USER LEFT SERVER BEFORE MEMBER ROLE COULD BE ADDED.") 
          "color" $color
          "footer" (sdict "text" (print "User Trust: " (print $trust) "/6  |  " (print $trustiness)))
        ))}}
	{{end}}
{{else}}
    {{$hush := sendTemplate nil "sus-action"}}
    {{sendMessage $alerts (complexMessage "content" (print .User.ID) "embed" (cembed
        "title" (print .User.String " is suspicious!")
        "description" (print "```" $desc "```\n**Other info:**\n> Custom status: " $status "\n> Account age: " (humanizeDurationMinutes $age))
        "thumbnail" (sdict "url" (.User.AvatarURL "256"))
        "footer" (sdict "text" (print "User Trust: " $thing " " $trust "/6 | " $trustiness))
        "color" $color                                                                        
    ))}}
{{end}}