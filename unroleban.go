{{- $args := parseArgs 1 "Improper Formatting! Please use -unrb <user>" (carg "user" "user") (carg "string" "reason") }}
{{$userID := ($args.Get 0)}}
{{$logging_channel_id := 1091441099601621060}}
{{$Avatar := $userID.AvatarURL "256"}}
{{$userM := getMember $userID}}
{{$Roles := (dbGet $userID.ID "Roles").Value}}
{{dbDel $userID.ID "Roles"}}
{{if not $Roles}}
  {{/* handle non-existing entry */}}
Either the user is not Rolebanned or the user has no stored roles! please re-add manually ;-;
  {{return}}
{{end}}

{{setRoles $userID.ID $Roles}}
User Roleban removed!
{{$r := dbIncr .User.ID "unrbcount" 1}}
{{sendMessage $logging_channel_id (cembed "description" (print "**User Un-Rolebanned:** "  $userID "\n**By:**[ " .User "](https://discord.com/users/" .User.ID ")") "color" 0x00FF00 "thumbnail" (sdict "url" $Avatar)
)}}
{{dbDel $userID.ID "rbreason"}}
{{dbDel $userID.ID "rbMod"}}
