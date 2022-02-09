func addDigits(num int) int {
    for num / 10 > 0 {
        sum, temp := 0, num
        for temp > 0 {
            sum += temp % 10
            temp /= 10
        }
        
        num = sum
    }
    
    return num
}
