{{- $args := parseArgs 1 "Improper Formatting! Please use -unrb <user>" (carg "user" "user")}}
{{$userID := ($args.Get 0)}}
{{$reason := (dbGet $userID.ID "rbreason").Value}}
{{$mod := "none"}}
 
{{try}}
{{$mod = (getMember ((dbGet $userID.ID "rbMod").Value)).User}}
{{catch}}
{{$mod = "none stored"}}
{{end}}
 
{{if $reason}}
{{print "<@" $userID.ID ">" " Roleban reason: \n" $reason "\n\nMod Responsible: " $mod }}
{{else}}
{{print "Either user has no stored reason or is not rolebanned!"}}
{{end}}