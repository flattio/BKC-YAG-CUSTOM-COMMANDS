{{- $args := parseArgs 1 "Improper Formatting! Please use -unaip <user>" (carg "user" "user")}}
{{$userID := ($args.Get 0)}}
{{$reason := (dbGet $userID.ID "aipreason").Value}}
{{$mod := "none"}}
 
{{try}}
{{$mod = (getMember ((dbGet $userID.ID "aipmod").Value)).User}}
{{catch}}
{{$mod = "none stored"}}
{{end}}
 
{{if $reason}}
{{print "<@" $userID.ID ">" " Anti-image perm reason: \n" $reason "\n\nMod responsible: " $mod }}
{{else}}
{{print "Either user has no stored reason or is not anti image perm'd!"}}
{{end}}