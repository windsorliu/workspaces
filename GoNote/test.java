import java.util.ArrayList;
import java.util.List;

class test {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        list.add("0");
        list.add("1");
        list.add("2");
        list.add("3");

        // while (true) {
        // for (int i = 0; i < list.size(); i++) {
        // System.out.print(list.get(i) + "\t");
        // }
        // System.out.println("size:" + list.size());
        // list.remove(0);

        // if (list.isEmpty()) {
        // break;
        // }
        // }

        // System.out.println(list.size());
        // System.out.println(list.isEmpty());
        System.out.println("before:" + list.size());
        list.remove(2);
        System.out.println("after:" + list.size());
        System.out.println("index:0 --> " + list.get(2));
    }
}