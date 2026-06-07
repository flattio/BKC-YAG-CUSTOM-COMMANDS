{{$args := parseArgs 1 "Improper Formatting! Please use -vcb <user> <reason>" (carg "user" "user") (carg "string" "reason")}}
{{$userID := ($args.Get 0)}}
{{$userM := getMember $userID}}
{{$Avatar := $userID.AvatarURL "256"}}
{{$modID := .user.ID}}
{{$reason := $args.Get 1}} 
{{/* Configurable values */}}
{{$logging_channel_id := 1091441099601621060}}
{{$Roles := $userM.Roles}}
 
{{if or (targetHasRoleID $userID.ID 1091441098330746918) (targetHasRoleID $userID.ID 1121590212011773962) (targetHasRoleID $userID.ID 1091441098330746919)}}
Target is Staff!
{{else}}
 
{{if not ($args.IsSet 1)}} {{- /*comment out these lines to remove reason requirement*/ -}} 
Please provide a reason!
{{return}}
{{end}}
 
{{if targetHasRoleID $userID.ID 1174194288071020574 }}
User already Voice Chat Banned!
{{return}}
{{end}}
 
 {{giveRoleID $userID 1174194288071020574}}
 
{{sendMessage $logging_channel_id (cembed "description" (print "**User Voice chat banned:** "  $userID " (" $userID.ID ")" "\n**By:**[ " .User "](https://discord.com/users/" .User.ID ")\n \n**Reason:**\n" $reason) "color" 0xFF0000 "thumbnail" (sdict "url" $Avatar)
)}}
User Voice Chat banned!
{{$r := dbIncr .User.ID "VCBcount" 1}}
{{dbSetExpire $userID.ID "vcbreason" $reason 2678400}}
{{dbSetExpire $userID.ID "VcbMod" (str .User.ID) 2678400}}
{{end}}