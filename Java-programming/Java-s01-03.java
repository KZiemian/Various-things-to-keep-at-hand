// class StaticTest {
//     static int m_intVar1 = 47;
// }

// class StaticFun {
//     static void incr() {
// 	StaticTest.m_intVar1++;
//     }
// }

// class Main {
//     public static void main(String[] args) {
// 	StaticFun staticFunVar1 = new StaticFun();

// 	System.out.println("StaticTest.m_intVar1: ");
// 	System.out.println(StaticTest.m_intVar1);

// 	System.out.println("\nStaticFun.incr().");

// 	StaticFun.incr();

// 	System.out.println("\nStaticTest.m_intVar1: ");
// 	System.out.println(StaticTest.m_intVar1);

// 	System.out.println("\nStaticFunVar1.incr().");

// 	staticFunVar1.incr();

// 	System.out.println("\nStaticTest.m_intVar1: ");
// 	System.out.println(StaticTest.m_intVar1);
//     }
// }

// import java.util.Date;

// class Main {
//     public static void main(String[] args) {
// 	System.out.println("Hello, World! Today is: ");
// 	System.out.println(new Date());
//     }
// }

// /** Komentarz opisujący klasę. */
// public class DocTest {
//     /** Komentarz do zmiennej. */
//     public int m_intVar1;
//     /** Komentarz do metody. */
//     public void funVoid() {}
// }

class Main {
    public static void main(String[] args) {
	// System.out.println("args[0] == " + args[0] + ".");
	// System.out.println("args[1] == " + args[1] + ".");
	// System.out.println("args[2] == " + args[2] + ".");

	// int intVar1;
	// int intVar2;
	// int intVar3;

	// System.out.println("Hello, World!");
	// System.out.println("intVar1 == " + intVar1 + ".");

	// int a = 0;
	// int x = 1;
	// int y = 2;
	// int z = 3;

	// System.out.println("a == " + a + ".");

	// a = x + y - 2/2 + z;

	// System.out.println("a == " + a + ".");

	// a = x + (y - 2)/(2 + z);

	// System.out.println("a == " + a + ".");

	int intVar1 = 0;
	int intVar2 = 2;

	intVar1 = intVar2;

	System.out.println("intVar1 == " + intVar1 + ".");
	System.out.println("intVar2 == " + intVar2 + ".");

	intVar1 = 3;

	System.out.println("intVar1 == " + intVar1 + ".");
	System.out.println("intVar2 == " + intVar2 + ".");
    }
}
