package array

/**
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。
 *
 * 示例 1：
 * 输入：nums = [2,7,11,15], target = 9
 * 输出：[0,1]
 * 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 *
 * 示例 2：
 * 输入：nums = [3,2,4], target = 6
 * 输出：[1,2]
 *
 * 示例 3：
 * 输入：nums = [3,3], target = 6
 * 输出：[0,1]
 *
 * 提示：
 * 2 <= nums.length <= 104
 * -109 <= nums[i] <= 109
 * -109 <= target <= 109
 * 只会存在一个有效答案
*/

// TwoSum 查找数组中和为目标值的两个数的下标
func TwoSum(nums []int, target int) []int {
    // 创建哈希表存储数字和下标
    numMap := make(map[int]int)
    
    // 遍历数组
    for i, num := range nums {
        // 计算需要查找的补数
        complement := target - num
        
        // 如果补数存在于哈希表中，返回两个数的下标
        if j, exists := numMap[complement]; exists {
            return []int{j, i}
        }
        
        // 将当前数字和下标存入哈希表
        numMap[num] = i
    }
    
    // 如果没有找到结果，返回空数组
    return []int{}
}