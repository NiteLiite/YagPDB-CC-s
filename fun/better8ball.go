{{/*
    This command sucks.
    Trigger: 8ball
    Trigger Type: Command
    Usage: you know already
    
    Repo: https://github.com/NiteLiite/YagPDB-CC-s/
    License: not for this one, it doesnt deserve one
*/}}

{{$answers := dict
0 "● It is certain."
1 "● It is decidedly so."
2 "● Without a doubt."
3 "● Yes definitely."
4 "● You may rely on it."
5 "● As I see it, yes."
6 "● Most likely."
7 "● Outlook good."
8 "● Yes."
9 "● Signs point to yes."
10 "● Reply hazy, try again."
11 "● Ask again later."
12 "● Better not tell you now."
13 "● Cannot predict now."
14 "● Concentrate and ask again."
15 "● Don't count on it."
16 "● My reply is no."
17 "● My sources say no."
18 "● Outlook not so good."
19 "● Very doubtful."
}}
 
{{$rand := randInt 0 20}}
{{$reply := print ($answers.Get $rand)}}
{{$col := ""}}
 
{{if (lt $rand 5)}}
  {{$col = 4568196}}
{{else if (lt $rand 10)}}
  {{$col = 4568196}}
{{else if (lt $rand 15)}}
  {{$col = 16762945}}
{{else if (lt $rand 20)}}
  {{$col = 16711680}}
{{end}}
 
{{sendMessage nil (cembed 
	"title" "***Magic 8-Ball***" 	
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/881835799547609109/985802661792206899/0cfd4882c0646d504900c90166d80cf8.png")
	"fields" (cslice
			    (sdict "name" "Question:" "value" (print "● " .StrippedMsg) "inline" false)
			    (sdict "name" "Response:" "value" $reply "inline" false) 
             )
	"footer" (sdict "text" (print "Asked by " .User) "icon_url" (.Member.AvatarURL "256"))
	"color" $col
)}}
