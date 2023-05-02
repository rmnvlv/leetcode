namespace Easy.Solutions;
using System.Text;
public class Solutions{
    
    //1929. Concatenation of Array Runtime 144 ms Beats 70.86% Memory 48.1 MB Beats 16.3%
    public static int[] GetConcatenation1(int[] nums) {
            int [] ans = new int [nums.Length*2];

            for (int i=0; i < nums.Length; ++i){
                ans[i] = nums[i];
                ans[i + nums.Length] = nums[i];
            }

            return ans;
        }

    //1929. Concatenation of Array Runtime 137 ms Beats 93.72%% Memory 48.3 MB Beats 8.34%
    public static int[] GetConcatenation2(int[] nums) {
            int [] ans = new int [nums.Length*2];

            Array.Copy(nums, ans, nums.Length);
            Array.Copy(nums, 0, ans, nums.Length, nums.Length);

            return ans;
        }

    //1920. Build Array from Permutation Runtime 149 ms Beats 72.66% Memory 50.4 MB Beats 6.4%
    public static int[] BuildArray1(int[] nums){
        int[] ans = new int[nums.Length];
        for (int i = 0; i< nums.Length; ++i){
            ans[i] = nums[nums[i]];
        }
        return ans;
    }

    //1920. Build Array from Permutation Runtime 140 ms Beats 94.92% Memory 50.1 MB Beats 7.97%
    public static int[] BuildArray2(int[] nums){
        for (int i = 0; i<nums.Length; ++i){
            nums[i] += (nums[nums[i]] % nums.Length) * nums.Length;
        }
        for (int i = 0; i<nums.Length; ++i){
            nums[i] = nums[i] / nums.Length;
        }
        return nums;
    }

    //1108. Defanging an IP Address Runtime 76 ms Beats 75.66% Memory 36.4 MB Beats 90.23%
    public static string DefangIPaddr1(string address) {
        return address.Replace(".", "[.]");
    }

    //1108. Defanging an IP Address Runtime 77 ms Beats 72.19% Memory 36.2 MB Beats 95.20%
    public static string DefangIPaddr2(string address) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i<address.Length; ++i){
            if (address[i] == '.'){
                sb.Append("[.]");
            } else {
                sb.Append(address[i]);
            }
        }

        return sb.ToString();
    }

    //1470. Shuffle the Array Runtime 146 ms Beats 44.78% Memory 45.1 MB Beats 6.40%
    public static int[] Shuffle(int[] nums, int n) {
        int[] ans = new int[n*2];

        for (int i = 0, j = 0; i <nums.Length - n; ++i, ++j) {
            ans[i] = nums[i];
            ans[j++] = nums[i+n];
        } 
        return nums;
    }

    //2469. Convert the Temperature Runtime 81 ms Beats 95.89% Memory 36.1 MB Beats 45.24%
    public static double[] ConvertTemperature1(double celsius) {
        double[] ans = new double[2];

        ans[0] = celsius + 273.15;
        ans[1] = celsius * 1.80 + 32.0;

        return ans;
    }

      //2469. Convert the Temperature Runtime 81 ms Beats 95.89% Memory 36.1 MB Beats 45.24%
    public static double[] ConvertTemperature2(double celsius) {
        return new double[2] {celsius + 273.15, celsius * 1.80 + 32.0};
    }
}
