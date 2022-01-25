func validMountainArray(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    
    wasPeak := false
    for i := 0; i < len(arr) - 1; i++ {
        if i != 0 && !wasPeak && arr[i] > arr[i + 1] {
            wasPeak = true
        }
        if !wasPeak && arr[i] >= arr[i + 1] || wasPeak && arr[i] <= arr[i + 1] {
            return false
        }
    }
    return wasPeak
}
