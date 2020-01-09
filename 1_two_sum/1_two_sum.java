class Solution {
    public int[] twoSum(int[] nums, int target) {
        //int[] result =new int[2];
        for(int i=0; i<nums.length;i++){
               for(int j=i+1; j<nums.length;j++){
                   //int temp = nums[i]+nums[j];
                if(nums[j] == target - nums[i]){
                    //result[0] = nums[i];
                    //result[1] = nums[j];
                    return new int[] {i,j};
                }
            } 
        }
        throw new IllegalArgumentException("No two sum solution");
    }
}