{{$args := parseArgs 1 "Improper Formatting! Please use -hiatus <user> <reason>" (carg "user" "user") (carg "string" "reason")}}
{{$userID := ($args.Get 0)}}
{{if targetHasRoleID $userID 1091441098330746918}}
{{giveRoleID $userID 1130238675503022140}} 
 
{{$ids := cslice 1121590212011773962 1091441098330746919 1091441098330746918 1176361199437361243 1193970513547100200 1221371016375435317 1193970513547100200 1220511784604008458 1220502078003085433}}
{{range $ids}}
    {{- takeRoleID $userID . -}}
{{- end }}
Mod put in hiatus.
 
<@&1121590212011773962>
{{else}}
User is not a moderator!
{{end}}