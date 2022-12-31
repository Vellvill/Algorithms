package main

func main() {
    topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
}

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for i := range nums {
        if _, ok := m[nums[i]]; ok {
            m[nums[i]]++
        } else {
            m[nums[i]] = 1
        }
    }
    temp := make([]int, len(m))
    var i int
    for key := range m {
        temp[i] = key
        i++
    }

    qs(temp, m)

    return temp[:k]
}

func qs(nums []int, m map[int]int) {
    if len(nums) < 2 {
        return
    }

    l, r := 0, len(nums)-1

    p := (l + r) >> 1

    nums[p], nums[r] = nums[r], nums[p]

    for i := range nums {
        if m[nums[i]] > m[nums[r]] {
            nums[i], nums[l] = nums[l], nums[i]
            l++
        }
    }

    nums[r], nums[l] = nums[l], nums[r]

    qs(nums[:l], m)
    qs(nums[l+1:], m)
}
