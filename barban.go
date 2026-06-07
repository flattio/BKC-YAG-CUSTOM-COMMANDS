{$args := parseArgs 1 "Improper Formatting! Please use -bb <user> <reason>" (carg "user" "user") (carg "string" "reason")}}
 
{{$userID := ($args.Get 0)}}
{{$userM := getMember $userID}}
{{$Avatar := $userID.AvatarURL "256"}}
{{$modID := .user.ID}}
{{$reason := $args.Get 1}} 
{{/* Configurable values */}}
{{$logging_channel_id := 1091441099601621060}}
{{$Roles := $userM.Roles}}
 
{{if or (targetHasRoleID $userID.ID 1091441098330746918) (targetHasRoleID $userID.ID 1121590212011773962) (targetHasRoleID $userID.ID 1091441098330746919) (targetHasRoleID $userID.ID 1193970513547100200)}}
Target is immune!
{{else}}
 
{{if not ($args.IsSet 1)}} 
Please provide a reason!
{{return}}
{{end}}
 
{{if targetHasRoleID $userID.ID 1279595746382844077 }}
User already barbanned!
{{return}}
{{end}}
 
{{takeRoleID $userID 1199490103127588934}} {{giveRoleID $userID 1279595746382844077}} {{giveRoleID $userID 1101857932813008948}} 
 
{{sendMessage $logging_channel_id (cembed "description" (print "**User Barbanned:** "  $userID " (" $userID.ID ")" "\n**By:**[ " .User "](https://discord.com/users/" .User.ID ")\n \n**Reason:**\n" $reason) "color" 0xFF0000 "thumbnail" (sdict "url" $Avatar)
)}}
User barbanned!
 
{{$r := dbIncr .User.ID "bbcount" 1}}
{{dbSetExpire $userID.ID "bbreason" $reason 2678400}}
{{dbSetExpire $userID.ID "bbmod" (str .User.ID) 2678400}}
{{end}}