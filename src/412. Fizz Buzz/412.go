func fizzBuzz(n int) []string {
    result := make([]string, n)
    
    for i := 0; i < n; i++ {
        if (i + 1) % 15 == 0 {
            result[i] = "FizzBuzz"
        } else if (i + 1) % 3 == 0 {
            result[i] = "Fizz"
        } else if (i + 1) % 5 == 0 {
            result[i] = "Buzz"
        } else {
            result[i] = strconv.Itoa(i + 1)
        }
    }
    
    return result
}
