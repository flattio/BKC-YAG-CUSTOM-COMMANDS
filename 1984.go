{{$args := parseArgs 1 "" (carg "user" "UserID to 1984") (carg "string" "junk string")}}
{{$user := ($args.Get 0)}}
{{takeRoleID $user 1101857932813008948}} {{giveRoleID $user 1098341447217983511}} {{giveRoleID $user 1101857823060664331}}
User 1984'd!
{{$r := dbIncr .User.ID "1984count" 1}}