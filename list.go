{{$args := parseArgs 1 "Improper Formatting! Please use -list <user>" (carg "user" "user")}}
 
{{$user := ($args.Get 0)}}
{{$userID := $user.ID}}
 
{{$rbcount := (dbGet $userID "rbcount").Value}}
{{$unrbcount := (dbGet $userID "unrbcount").Value}}
{{$verifiedartcount := (dbGet $userID "verifiedartcount").Value}}
{{$1984count := (dbGet $userID "1984count").Value}}
{{$TBcount := (dbGet $userID "TBcount").Value}}
{{$VCBcount := (dbGet $userID "VCBcount").Value}}
{{$unvcbcount := (dbGet $userID "unvcbcount").Value}}
{{$VABcount := (dbGet $userID "VABcount").Value}}
{{$cbcount := (dbGet $userID "countingbancount").Value}}
 
{{print "Stats for " "<@" $userID ">" 
"\nㅤ"
"\n Roleban Count: " $rbcount  
"\n Un-roleban Count: " $unrbcount  
"\n Verified-art Count: " $verifiedartcount  
"\n 1984 Count: " $1984count 
"\n Ticket ban Count: " $TBcount
"\n Voicechat ban Count: " $VCBcount
"\n Un-voicechat ban Count: " $unvcbcount
"\n Voice-activity Count: " $VABcount 
"\n Counting ban Count: " $cbcount}}