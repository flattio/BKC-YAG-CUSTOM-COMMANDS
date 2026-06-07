{{$args := parseArgs 1 "" (carg "user" "UserID to anti image perm") (carg "string" "junk string")}}
{{$userID := ($args.Get 0)}}
{{$logging_channel_id := 1091441099601621060}}
{{$Avatar := $userID.AvatarURL "256"}}
{{takeRoleID $userID 1133141559546740860}}
User can now send images again!
 
{{$r := dbIncr .User.ID "unaipcount" 1}}
 
{{sendMessage $logging_channel_id (cembed "description" (print "**User Un-Imagebanned:** "  $userID "\n**By:**[ " .User "](https://discord.com/users/" .User.ID ")") "color" 0x00FF00 "thumbnail" (sdict "url" $Avatar)
)}}
 
{{dbDel $userID.ID "aipreason"}}
{{dbDel $userID.ID "aipmod"}}