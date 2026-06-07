{{$args := parseArgs 1 "" (carg "user" "UserID to barban") (carg "string" "junk string")}}
{{$userID := ($args.Get 0)}}
{{$logging_channel_id := 1091441099601621060}}
{{$Avatar := $userID.AvatarURL "256"}}
{{takeRoleID $userID 1279595746382844077}}
Users Barban has been removed!
 
{{$r := dbIncr .User.ID "unbbcount" 1}}
 
{{sendMessage $logging_channel_id (cembed "description" (print "**User Un-Barbanned:** "  $userID "\n**By:**[ " .User "](https://discord.com/users/" .User.ID ")") "color" 0x00FF00 "thumbnail" (sdict "url" $Avatar)
)}}
 
{{dbDel $userID.ID "bbreason"}}
{{dbDel $userID.ID "bbmod"}}