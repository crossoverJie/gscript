
// [2, 7, 11, 15], target = 9 return [0,1]
int[] twoSum(int[] nums, int target){
    int[] ret = [2]{};

    for(int i =0; i<len(nums) ;i++){
        for (int j = len(nums) -1 ;j > 0 ;j--){
            if (nums[i] + nums[j] == target){
                ret[0]=i;
                ret[1]=j;
                println("ok");
                break;
            }
            println("内层循环");
        }
    }
    return ret;
}

int[] nums = {2,7,11,15};
int[] ret = twoSum(nums, 9);
println(ret);
int l = ret[0];
int r = ret[1];
assertEqual(l, 0);
assertEqual(r, 1);