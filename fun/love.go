{{/*

	Trigger Type: Regex
	Trigger: \A-(hug|((head)?pat|pet)|(bonk|smack)|(boop|poke))(\s+|\z)
	
	Usage:
	-hug <user/text>
	-headpat/pat <user/text>
	-bonk <user/text>
	-boop/poke <user/text>

	Made by @NiteLiite#0001, I hate myself thoroughly.
  Copyright MIT

*/}}

{{$boops := cslice
"o9X9XXVCm-MAAAAd/bird-cute"
"QQDEdB7Y1xoAAAAC/poke-usagi"
"8RAKw8XAMqAAAAAC/assassination-classroom-assassination-class"
"SioUZX9lIJoAAAAd/nami-one-piece"
"4VxzO9Gg5VYAAAAC/puella-magi-madoka-magica-poke"
"NjIdfk7i3bsAAAAS/poke-poke-poke"
"6rS1x-dVUwEAAAAS/ishtar-ishtar-fgo"
"3dOqO4vVlr8AAAAS/poke-anime"
"y4R6rexNEJIAAAAC/boop-anime"
"4OHxyGd4qp0AAAAC/boop-nose"
"1YMrMsCtxLQAAAAC/anime-poke"
"fxIMcE41WpgAAAAd/anime-boop"
"APqauOtznp4AAAAC/boop-poke"
"RmQElPHERIoAAAAC/boop-anime"
"_vVL5fuzj4cAAAAC/nagi-no"
"G5u3bfszOPMAAAAC/anime-picking-nose"
"B1ohHuPJIpgAAAAd/anime-cuteness"
"kir_655L4uoAAAAC/cat-boop"
"i5Pqj03zv9kAAAAd/boop-nose-cat"
"t6ABAaRJEA0AAAAC/oreimo-ore-no-im%C5%8Dto-ga-konna-ni-kawaii-wake-ga-nai"
"ySdxnfxoTrUAAAAC/ascendence-of-a-bookworm-bookworm-anime"
"JoA5p9DuIkwAAAAC/poke"
"-qArCsdI9ekAAAAC/poke-anime"
"5QSqLIF7BHQAAAAC/anime-anime-girl"
"V0hkhPDG1KkAAAAC/platelets-cell-at-work"
}}

{{$hugs := cslice 
"sZPpQ_kn2UsAAAAd/puuung-love"
"_Ip7XSmd8M8AAAAC/clannad-after-story-anime"
"aG0pA87t0dMAAAAC/anime-chino"
"_ACTAjo1sZ4AAAAC/anime-friends"
"ZzraYDaXSjAAAAAC/hug"
"ZWrKJEYcmREAAAAC/tonari-no-kaibutsu-kun-my-little-monster"
"q5GWONzp2h0AAAAC/abra%C3%A7o-hug"
"mBmkxNFl5acAAAAd/anime-happy"
"aW94765-Z_4AAAAC/chi-chobits"
"LpYeBetRjCQAAAAC/anime-pat"
"4l0cz9CqXuYAAAAd/girls-hugging-yuri-hug"
"AslEAmkLHbIAAAAC/hug-head-pat"
"OIKUxVk2Sm8AAAAC/hug"
"CfwNyKCP-L4AAAAC/anime-cute"
"pHc46mKF83wAAAAd/zutara-hug"
"UIZHJoSeIjMAAAAC/sushichaeng-adventure-time"
"LDIPMQBFcD8AAAAC/tonari-no-kaibutsu-kun-my-little-monster"
"KHUhRSyp03EAAAAC/miss-you"
"ZPC39kDaCcYAAAAC/nolas-pfp"
"LRDMe1QpqFYAAAAC/cat-cute"
"fE_7w7DHRf8AAAAC/bears-hug"
"ZzorehuOxt8AAAAC/hug-cats"
"8Jk1ueYnyYUAAAAC/hug" 
"YQwIdPJlxocAAAAd/spy-x-family-anya-forger"
"gKlGEBBkliwAAAAC/anime-yuru-yuri" 
"L3FiTfr6d-wAAAAC/spy-x-family-loid-forger"
"WpbZhwwj6zAAAAAC/happy-hug" 
"gqM9rl1GKu8AAAAC/kitsune-upload-hug" 
"qF7mO4nnL0sAAAAC/abra%C3%A7o-hug" 
"-3I0yCd6L6AAAAAC/anime-hug-anime" 
"O3qIam1dAQQAAAAC/hug-cuddle" 
"0T3_4tv71-kAAAAC/anime-happy" 
"LTv9iKl2MlYAAAAC/one-piece-yamato" 
"FJRptuv4oX8AAAAC/hug-k-on"
"xmPGxvkD-Q8AAAAC/tight-hug" 
"i9MDVwZtSOcAAAAd/anime-hug" 
"X7esiQqlZCwAAAAC/luffy"
"xGNW9f0crgwAAAAC/one-piece-monkey-d-luffy"
"C9o3SCSGj60AAAAC/nami-luffy"
"bSf01Gv-5HIAAAAC/hug"
"h3CevNZffv0AAAAC/one-piece-carrot"
"6TW8llDClMkAAAAC/monkey-d-luffy-and"
"-ydwjrgf6ZAAAAAC/foxy-one-piece."
"beokMsbOxBwAAAAd/yamato-one-piece"
"r0SS5rof6aAAAAAC/littlenightmares-six"
"TC5S_OBFwOwAAAAC/hug-undertale"
"6peoLF2DowEAAAAC/sending-hugs-sending-virtual-hug"
"0aqQjKi1ETIAAAAC/deltarune-ralsei"
"Upo7pbGvd_IAAAAC/omori-hug"
"l2zfxZANNbkAAAAC/akko-hug"
}}

{{$bonks := cslice
"iDdGxlZZfGoAAAAC/powerful-head-slap"
"1T5bgBYtMgUAAAAC/head-hit-anime"
"EiFGi9dZXSAAAAAC/toradora-taiga"
"CrmEU2LKix8AAAAC/anime-bonk"
"JS0ZB_b3kjoAAAAC/sugiura-ayano-funami-yui"
"qZIzOW1VSnEAAAAC/shinji-shinji-broom"
"7gieVV7wfEQAAAAC/bonk-anime"
"5YrUft9OXfUAAAAS/bonk-doge"
"Xr8J9quvUHgAAAAd/bonk"
"tfgcD7qcy1cAAAAC/bonk"
"NS2upjfHs6wAAAAC/bonk"
"N3ZdP-hTVHYAAAAC/anime-anime-bonk"
"D_-8tx--KDAAAAAC/chuunibyou-anime"
"U6vSI52F4jwAAAAC/anime-hit"
"U6vSI52F4jwAAAAC/anime-hit"
"b5tAzjbsG68AAAAC/minmi-bonk"
"31WOy2yRK3QAAAAC/chuunibyou-hit"
"E6njrpISBV4AAAAC/bonk-hit"
"tf2de2mbRgwAAAAC/touka-kirishima-tokyo-ghoul"
"CsXEC2e1F6MAAAAd/klee-klee-bonk"
"mKX_7m0GsVAAAAAC/anime-blends"
"g_9NDHUmUdgAAAAC/anime"
"M0Vi6oBi7RcAAAAS/ranma-akane-tendo"
"c2vv1C1apy0AAAAC/head-hit"
"TCxz2fAU75IAAAAC/love-lab-head-smack-anime-smack"
"AlTCWaCOju0AAAAS/one-piece"
"pT_4p9x2KDkAAAAS/tama-otama"
"XiYuU9h44-AAAAAC/anime-slap-mad"
"bW9sL6u6V7AAAAAS/fly-away-slap"
"yJmrNruFNtEAAAAC/slap"
"2-r7BEc-cb8AAAAC/slap-smack"
"OuYAPinRFYgAAAAC/anime-slap"
"1lemb3ZmGf8AAAAC/anime-slap"
"_5FoWyBcuzwAAAAC/ranma-slapping"
"ra17G61QRQQAAAAC/tapa-slap"
"CvBTA0GyrogAAAAC/anime-slap"
"hNa8BhraaXsAAAAC/anime-nagatoro"
"QYxiYpG1lp8AAAAC/shenmue-shenmue-mai-and-goro"
"WcYvM-SqPkoAAAAC/baka-slap"
"ObgxhbfdVCAAAAAd/luffy-anime"
"YNqc67MOWVMAAAAC/one-piece-luffy"
"VOK9zSk09KMAAAAC/one-piece-luffy"
"9hBwXT9okiAAAAAC/one-piece-anime"
"84iZ9IP_Kw4AAAAC/vivi-usopp"
"h_qFkmXJnYQAAAAC/cat-attack"
"ZlisNbF5uh8AAAAC/toradora-attack"
"53_ElD98OSQAAAAd/nichijou-mio"
"DB5qwXcuerAAAAAd/anime-attack-ohto-ai"
"blYwpMNaaCUAAAAd/nichijou-uppercut"
"wOCOTBGZJyEAAAAC/chikku-neesan-girl-hit-wall"
}}

{{$pets := cslice
"cBE_Y6PUiwUAAAAC/loki-videoyun"
"uNT-3nchwysAAAAd/headpats-anime"
"Fxku5ndWrN8AAAAC/head-pat-anime"
"s7lGkoIAieYAAAAC/so-cute-cat"
"N41zKEDABuUAAAAC/anime-head-pat-anime-pat"
"GC9rg-v-wvMAAAAC/anime-pat"
"5ySiT5zF21oAAAAC/anime-head-pat"
"g49m0bqB0eMAAAAd/cat-pankocat"
"Ls2uiad4RRUAAAAC/anime-anime-headpat"
"u6SihPmmUbwAAAAd/cat-kitten"
"1bBIALbG0ikAAAAC/anime-anime-head-rub"
"Y7B6npa9mXcAAAAC/rikka-head-pat-pat-on-head"
"zBPha3hhm7QAAAAC/anime-girl"
"V5pe5qbYJT8AAAAC/head-pat"
"JWx5wmF6bqAAAAAC/charlotte-pat"
"dLdNYQrLsp4AAAAC/umaru-frown"
"R6wPLGZT9YAAAAAC/natsume-pat-on-the-head"
"nwbxEGQINOsAAAAd/pet-dog"
"tikIIf6NJvEAAAAC/cat-sleep"
"n6km1_0i97kAAAAC/anime-cat"
"bfpRP4Feg3oAAAAC/pet-anime"
"K1FJ_-7NDPQAAAAC/toad-pet"
"rN7iVS6kqGcAAAAC/kris-deltarune"
"CmB3aYpaJnAAAAAS/dog-petting"
"UTZOodcQt3IAAAAS/doggo-petting"
"HNZMnQj1DDUAAAAd/tate-no-yuusha-no-nariagari-the-rising-of-the-shield-hero"
"4Sy0Q_i8LgcAAAAd/anime-girl-pet"
"bBa-TswgT2oAAAAC/anime-cuddle"
"I1EbzmnZ0rgAAAAC/anime-pet"
"icOfHB5CcTQAAAAS/otama-one-piece"
"9HHCiDIdCkIAAAAS/cat-pet"
"tVwc20r-GwQAAAAd/rascal-does-not-dream-of-bunny-girl-senpai-seishun-buta-yar%C5%8D"
"5dezHjhjBucAAAAC/nagi-no-asukara-manaka-mukaido"
"-hkJYNs7tUkAAAAC/anime-pat"
"x7CVs1j6OdIAAAAC/kanan-lovelive"
"pB5LKEouppgAAAAC/pat-pat-on-head"
"aXEcXKRxXcEAAAAC/nezuko"
"jEUscYib2GkAAAAC/ogiue-chika"
"ZQ38A6xCjdUAAAAC/holo-headpat-spice"
"8cY2q5UaGccAAAAd/azumanga-daioh-azumanga"
"8w4TYd2tsKcAAAAC/anime-pat"
"ZeYnQ6uBq8MAAAAC/anime-pat"
"oqmgNNIJQqAAAAAC/anime-pat"
"zzb9MGIkmtkAAAAC/anime-pat"
"Jj-vHGZOgT4AAAAC/anime-anime-girl"
"i88aIexbldkAAAAd/hlop"
"HbKkTMht1VkAAAAC/anime-cute"
"l868YC4J2_MAAAAC/chobit-pat"
"jEr3aB39yfUAAAAC/pat-rem"
"TtcyuOii6gYAAAAC/luffy-one-piece"
"AQa-sGBW6JYAAAAC/luffy-gordo-luffy-gordinho"
}}

{{$col := 3092790}}
{{if (getMember .User.ID)}}
   {{$pos := 0}}
    {{$r := (getMember .User.ID).Roles}}
{{range .Guild.Roles}}
    {{if and (in $r .ID) (.Color) (lt $pos .Position)}}
  {{$pos = .Position}}{{$col = .Color}}
    {{end}}
 {{end}}
{{end}}

{{ $pet := index (shuffle $pets) 0 }}
{{ $hug := index (shuffle $hugs) 0 }}
{{ $bonk := index (shuffle $bonks) 0 }}
{{ $boop := index (shuffle $boops) 0 }}

{{if ge (len .CmdArgs) 1}}

{{$aww := ""}}
{{$self := false}}
{{with (getMember (index .CmdArgs 0))}}
	{{$aww = (or .Nick .User.Username)}}
	{{$self = eq .User.ID $.User.ID}}
{{else}}
	{{$aww = .StrippedMsg}}
{{end}}

{{ $image := "" }}
{{if reFind `(?i)(hug)` .Cmd}}
   {{$image = $hug}}
{{else if reFind `(?i)((head)?pat|pet)` .Cmd}}
   {{$image = $pet}}
{{else if reFind `(?i)(bonk|smack)` .Cmd}}
   {{$image = $bonk}}
{{else if reFind `(?i)(boop|poke)` .Cmd}}
   {{$image = $boop}}
{{end}}

{{if not $self}}
   {{$aww = (print "** " (lower (reReplace `-` .Cmd "")) "s **" $aww "**")}}
{{sendMessage nil (cembed
	"author" (sdict "name" "" "icon_url" (.Member.AvatarURL "256"))
	"description" (print "**" (or .Member.Nick .User.Username) $aww)
	"image" (sdict "url" (print "https://c.tenor.com/" $image ".gif"))
	"color" $col
)}}
{{else}}
D-do you need a {{lower (reReplace `-` .Cmd "")}} {{or .Member.Nick .User.Username}}...?
{{end}}

{{else}}
Who do you wanna {{lower (reReplace `-` .Cmd "")}}?
{{end}}
