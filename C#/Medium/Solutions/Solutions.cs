namespace Medium.Solutions;

public class Solutions{
    //1689. Partitioning Into Minimum Number Of Deci-Binary Numbers 77 ms Beats 68.8% Memory 43.5 MB Beats 32.86%
     public static int MinPartitions1(string n) {
        char max = ' ';
        foreach (char numeric in n){
            if (numeric > max){
                max = numeric;
            }
        }
        return max - '0';
    }

    //1689. Partitioning Into Minimum Number Of Deci-Binary Numbers 80 ms Beats 56.81% Memory 43.5 MB Beats 32.86%
    public static int MinPartitions2(string n) {
        return n.Max() - '0';
    }

    //1689. Partitioning Into Minimum Number Of Deci-Binary Numbers Runtime 76 ms Beats 72.30% Memory 43.4 MB Beats 32.86% WTF???
    public static int MinPartitions3(string n) => n.Max() - '0';
}