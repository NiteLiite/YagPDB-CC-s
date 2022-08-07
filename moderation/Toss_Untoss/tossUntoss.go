{{/*
	Trigger: \A-(((add)?(rb|toss|roleban))|(urb|untoss|unroleban))
	Trigger Type: Regex
	
	NOTE: THIS COMMAND ASSUMES YOU HAVE - AS YOUR PREFIX.
	If not, simply swap the - in the trigger with your prefix!
*/}}

{{$staff := 507368219632205828}} {{/* ID OF STAFF ROLE (NOTE: You still have to lock this command to a staff role) */}}
{{$tossed := 954490752417796136}} {{/* ID OF ROLE THAT LOCKS THE USER OUT OF ALL CHANNELS */}}


{{if (ge (len .CmdArgs) 1)}}
	{{if (getMember (index .CmdArgs 0))}}
		{{$member := (getMember (index .CmdArgs 0))}}
		{{if (not (targetHasRoleID $member.User.ID $staff))}}
 
{{/* TOSS */}}
{{if reFind `(?i)(((add)?)(rb|toss|roleban))` .Cmd}}
	{{if not (targetHasRoleID $member.User.ID $tossed)}}
		{{dbSet $member.User.ID "autoban" (str $member.User.ID)}}
		{{giveRoleID ($member.User) $tossed}}

		{{if not (reFind `(?i)(rb)` .Cmd)}}
			{{$embed := sdict
				"author" (sdict "name" (print "✅  " $member.User.Username "#" $member.User.Discriminator " has been tossed.") )
				"footer" (sdict "text" "Note: Questioning or commenting on rolebans is discouraged.")
				"color" 4568196
			}} {{if ge (len .CmdArgs) 2}} {{$embed.Set "description" (print "📄 **Reason: **" (joinStr " " (slice .CmdArgs 1)) )}}{{end}}
			{{sendMessage nil (cembed $embed)}}
		{{else}}
			{{deleteTrigger 0}}
		{{end}}

		{{if not (reFind `(?i)(add)` .Cmd)}}
			{{try}}
				{{$shh := createTicket $member.User.ID (print "🔕┃" $member.User.ID "┃") }}
			{{catch}}
				{{sendMessage nil (cembed "title" "Error!" "description" (print "**❌ " .Error "**") "color" 16727808)}}
			{{end}}
		{{else}}
			{{$hush := execAdmin "ticket adduser" $member.User.ID}}
			{{sleep 1}}
			{{print $member.User.Mention " you have been added to this roleban case."}}
		{{end}}
	{{else}}
		{{print "User is already rolebanned."}}
	{{end}}
{{end}}
{{/* END OF TOSS */}}
 
{{/* UNTOSS */}}
{{if reFind `(?i)(ub|untoss|unroleban)` .Cmd}}
	{{if (targetHasRoleID $member.User.ID $tossed)}}
		{{dbDel $member.User.ID "autoban"}}
		{{$Uembed := sdict
			"author" (sdict "name" (print "✅  " $member.User.Username "#" $member.User.Discriminator " is no longer tossed.") )
			"color" 4568196
			"timestamp" currentTime
		}}{{if ge (len .CmdArgs) 2}} {{$Uembed.Set "description" (print "📄 **Reason: **" (joinStr " " (slice .CmdArgs 1)) )}}{{end}}
			{{sendMessage nil (cembed $Uembed)}}
		{{$hush := execAdmin "ticket removeuser" $member.User.ID}}
		{{takeRoleID ($member.User) $tossed}}
	{{else}}
		{{print "User is not rolebanned."}}
	{{end}} 
{{end}}
{{/* END OF UNTOSS */}}

{{/* INVALID SYNTAX RESPONSE HANDLING */}}
		{{else}}
			{{$Eembed := sdict
				"author" (sdict "name" (print "✅  " $member.User.Username "#" $member.User.Discriminator " has been thrown against a fucking wall.") )
				"footer" (sdict "text" "Note: Questioning or commenting on rolebans is discouraged.")
				"color" 4568196
			}}{{if ge (len .CmdArgs) 2}} {{$Eembed.Set "description" (print "📄 **Reason: **" (joinStr " " (slice .CmdArgs 1)) )}}{{end}}
			{{sendMessage nil (cembed $Eembed)}}
		{{end}}
	{{else}}
		{{sendMessage nil (cembed "description" "***❌ Err: Invalid mention/ID.***" "color" 16711680)}}
	{{end}}
{{else}}
	{{sendMessage nil (cembed 
		"color" 8690687
		"fields" (cslice
 			(sdict "name" "Roleban CMDs:" "value" (print "```\n-rb\n-toss\n-roleban```") "inline" true)
			(sdict "name" "Unroleban CMDs:" "value" (print "```\n-urb\n-untoss\n-unroleban```") "inline" true) 
			(sdict "name" "Other:" "value" (print "```\n-addtoss : adds a user to the current case.\n-purge <reason:optional> : closes the case.```") "inline" true))
		"footer" (sdict "text" (print "Format: -CMD <member>"))
	)}}
{{end}}
{{/* END OF INVALID SYNTAX RESPONSE HANDLING */}}
