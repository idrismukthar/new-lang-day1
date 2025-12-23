import java.util.Scanner;

public class ATM {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        double balance = 2500.50;

        System.out.println("--- JAVA BANKING SYSTEM ---");
        System.out.println("Welcome, Idris Mukthar.");
        System.out.println("Current Balance: $" + balance);

        System.out.print("\nEnter amount to withdraw: ");

        if (input.hasNextDouble()) {
            double withdrawal = input.nextDouble();

            if (withdrawal <= 0) {
                System.out.println("Error: Please enter a positive amount.");
            } else if (withdrawal <= balance) {
                balance -= withdrawal;
                System.out.println("Success: Withdrawal Successful!");
                System.out.println("New Balance: $" + balance);
            } else {
                System.out.println("Error: Insufficient funds!");
            }
        } else {
            System.out.println("Error: Invalid input! Please enter a number.");
        }

        System.out.println("\nThank you for using Java Bank.");
        input.close();
    }
}