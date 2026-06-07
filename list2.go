{{$args := parseArgs 1 "Improper Formatting! Please use -list2 <user>" (carg "user" "user")}}
 
{{$user := ($args.Get 0)}}
{{$userID := $user.ID}}
 
{{$verifycount := (dbGet $userID "verifycount").Value}}
{{$un1984count := (dbGet $userID "un1984count").Value}}
{{$bbcount := (dbGet $userID "bbcount").Value}}
{{$unbbcount := (dbGet $userID "unbbcount").Value}}
 
{{print "Stats for " "<@" $userID ">" 
"\nㅤ"
"\n Un-1984 Count: " $un1984count 
"\n Verification Count: " $verifycount
"\n Barban Count: " $bbcount 
"\n Un-barban Count: " $unbbcount }}