// class NazwaTypu {}

// !!!!
// members data - fields
// members function - methods

// class TylkoDane {
//     // !!!!
//     // If fields has a basic type, it is zeroed by default.
//     int m_intVar1;
//     float m_floatVar1;
//     boolean m_boolVar1;
// }

// package, library

class StaticTest {
    static int m_intVar1 = 47;
}

class Main {
    // static int someFunction(String stringVar1) {
    // 	return 3 * stringVar1.length();
    // }

    public static void main(String[] args) {
	// byte byteVar1 = 0;

	// System.out.println("byteVar1: ");
	// System.out.println(byteVar1);

	// String stringVar1 = new String("łańcuch");

	// System.out.println("stringVar1: ");
	// System.out.println(stringVar1);

	// System.out.println("main scope.");

	// {
	//     String stringVar1 = new String("łańcuch");

	//     System.out.println("stringVar1: ");
	//     System.out.println(stringVar1);
	// }

	// System.out.println("stringVar1: ");
	// System.out.println(stringVar1);

	// NazwaTypu nazwaTypuVar1 = new NazwaTypu();

	// System.out.println("nazwaTypuVar1: ");
	// System.out.println(nazwaTypuVar1);

	// !!!!
	// float floatVar1 = 0.0f;

	// System.out.println("floatVar1: ");
	// System.out.println(floatVar1);

	// TylkoDane tylkoDaneVar1 = new TylkoDane();

	// System.out.println("tylkoDaneVar1: ");
	// !!!!
	// Basicaly useless data.
	// System.out.println(tylkoDaneVar1);

	// System.out.println("tylkoDaneVar1.m_intVar1: ");
	// System.out.println(tylkoDaneVar1.m_intVar1);

	// System.out.println("tylkoDaneVar1.m_floatVar1: ");
	// System.out.println(tylkoDaneVar1.m_floatVar1);

	// System.out.println("tylkoDaneVar1.m_boolVar1: ");
	// System.out.println(tylkoDaneVar1.m_boolVar1);

	// tylkoDaneVar1.m_intVar1 = 47;
	// tylkoDaneVar1.m_floatVar1 = 1.1f;
	// tylkoDaneVar1.m_boolVar1 = true;

	// System.out.println("tylkoDaneVar1.m_intVar1: ");
	// System.out.println(tylkoDaneVar1.m_intVar1);

	// System.out.println("tylkoDaneVar1.m_floatVar1: ");
	// System.out.println(tylkoDaneVar1.m_floatVar1);

	// System.out.println("tylkoDaneVar1.m_boolVar1: ");
	// System.out.println(tylkoDaneVar1.m_boolVar1);

	// String stringVar1 = new String("Łańcuch");

	// System.out.println("someFunction(): ");
	// System.out.println(Main.someFunction(stringVar1));

	StaticTest staticTestVar1 = new StaticTest();
	StaticTest staticTestVar2 = new StaticTest();

	System.out.println("staticTestVar1.m_intVar1: ");
	System.out.println(staticTestVar1.m_intVar1);

	System.out.println("staticTestVar1.m_intVar1: ");
	System.out.println(staticTestVar1.m_intVar1);

	System.out.println("StaticTest.m_intVar1++.");

	StaticTest.m_intVar1++;

	System.out.println("staticTestVar1.m_intVar1: ");
	System.out.println(staticTestVar1.m_intVar1);

	System.out.println("staticTestVar2.m_intVar1: ");
	System.out.println(staticTestVar2.m_intVar1);
    }
}

// Sending message to the object.
