package format

import "strings"

func GetDomainNames(list string) (retval string) {

    if len(list) < 3 {
        return list
    }

    tmp := strings.Fields(list)
    sb := strings.Builder{}

    for _, s := range tmp {
        for strings.Contains(s, "@") {
            s = s[strings.Index(s, "@")+1:]
        }
        sb.WriteString(" ")
        sb.WriteString(s)
    }

    return strings.TrimSpace(sb.String())
}
