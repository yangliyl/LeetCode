func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    if m > n {
        return findMedianSortedArrays(nums2, nums1)
    }
    iMin, iMax := 0, m
    half := (n + m + 1) / 2
    for iMin <= iMax {
        i :=  (iMax + iMin ) / 2
        j := half - i
        if i < iMax && nums1[i] < nums2[j - 1] {
            iMin = i + 1
        } else if i > iMin && nums1[i - 1] > nums2[j] {
            iMax = i - 1
        } else {
            left := 0
            if i == 0 {
                left = nums2[j - 1]
            } else if j == 0 {
                left = nums1[i - 1]
            } else {
                if nums1[i - 1] > nums2[j - 1] {
                    left = nums1[i - 1]
                } else {
                    left = nums2[j - 1]
                }
            }
            if (n + m) % 2 == 1 {
                return float64(left)
            }

            right := 0
            if i == m {
                right = nums2[j]
            } else if j == n {
                right = nums1[i]
            } else {
                if nums1[i] < nums2[j] {
                    right = nums1[i]
                } else {
                    right = nums2[j]
                }
            }
            return float64(left + right) / 2
        }
    }
    return 0
}