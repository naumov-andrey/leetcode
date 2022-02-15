func climbStairs(n int) int {
    cache := [2]int{1, 2}
    
    for i := 2; i < n; i++ {
        cache[i % 2] = cache[0] + cache[1]
    }
    
    return cache[(n - 1) % 2]
}
