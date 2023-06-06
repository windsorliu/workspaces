import java.util.Arrays;
import java.util.Collections;

public class Test {

    public static void main(String[] args) {
        Integer[] a = { 5, 7, 8, 9, 1, 3 };

        Arrays.sort(a); // 使用sort排序 小到大
        for (int x : a) {
            System.out.print(x);
        }

        System.out.println();

        Arrays.sort(a, Collections.reverseOrder()); // 使用sort排序 大到小
        for (int x : a) {
            System.out.print(x);
        }
    }
}