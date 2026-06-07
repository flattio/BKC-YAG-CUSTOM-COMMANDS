{{$args := parseArgs 0 "Improper Formatting! Please use -rb <user> <reason>" (carg "user" "user") (carg "string" "reason")}}
{{$userID := ($args.Get 0).ID}}
{{if (targetHasRoleID $userID 1199490103127588934)}}
User already verified!
{{return}}
{{end}}
{{giveRoleID $userID 1199490103127588934}}
{{takeRoleID $userID 1101857932813008948}}
{{$r := dbIncr .User.ID "verifycount" 1}}
User Verified!