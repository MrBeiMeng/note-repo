import java.util.Arrays;
import java.util.HashMap;

class TwoNum{

    static class Solution {
        public int[] twoSum(int[] nums, int target) {
            HashMap<Integer, Integer> map = new HashMap<>();
            int tmpTarget;

            for (int i = 0; i < nums.length; i++) {
                tmpTarget = target - nums[i];
                if (map.containsKey(tmpTarget)){
                    return new int[]{map.get(tmpTarget),i};
                }

                map.put(nums[i],i);
            }

            return null;
        }
    }


    public static void main(String[] args) {
        System.out.println("new Solution().twoSum(new int[]{2,7,11,15},9) = " + Arrays.toString(new Solution().twoSum(new int[]{2, 7, 11, 15}, 9)));
        System.out.println("new Solution().twoSum(new int[]{3,2,4},6) = " + Arrays.toString(new Solution().twoSum(new int[]{3, 2, 4}, 6)));
        System.out.println("new Solution().twoSum(new int[]{3,3},6) = " + Arrays.toString(new Solution().twoSum(new int[]{3,3}, 6)));

    }
}