{{$args := parseArgs 1 "" (carg "user" "UserID to Ticket-ban") (carg "string" "junk string")}}
 
{{$user := ($args.Get 0)}}
 
{{giveRoleID $user 1291841420377718885}}
 
User Ticket-banned!
 
{{$r := dbIncr .User.ID "TBcount" 1}}