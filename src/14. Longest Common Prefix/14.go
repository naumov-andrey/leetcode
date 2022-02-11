func longestCommonPrefix(strs []string) string {
    p := 0
    
    for ; p < len(strs[0]); p++ {
        for j := 1; j < len(strs); j++ {
            if p >= len(strs[j]) || strs[j][p] != strs[0][p] {
                return strs[0][:p]
            }
        }
    }
    
    return strs[0][:p]
}
