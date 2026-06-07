{{$args := parseArgs 1 "" (carg "user" "UserID to 1984") (carg "string" "junk string")}}
{{$user := ($args.Get 0)}}
{{takeRoleID $user 1098341447217983511}} {{takeRoleID $user 1101857823060664331}} 
User un-1984'd!
{{$r := dbIncr .User.ID "un1984count" 1}}