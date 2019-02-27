func romanToInt(s string) int {
    riMap := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    last := 0
    m := 0
    for i := 0; i < len(s); i++{
        if i < len(s) -1{
            if riMap[s[i+1]] <= riMap[s[i]] {
                m += riMap[s[i]]
                last = riMap[s[i]]
            }else{
                m +=(riMap[s[i+1]] - riMap[s[i]])
                last = riMap[s[i]]
                i++
            }
        }else if riMap[s[i]] <= last{
             m += riMap[s[i]]
        }
    }
    if m == 0{
        return riMap[s[0]]
    }
    return m
}
