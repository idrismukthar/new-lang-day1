import java.util.*;
import java.io.*;

public class MyOwnPaystack {
    private static final String LEDGER_FILE = "ledger.csv";
    private static final String SECRET_PIN = "1234";

    // 1. PIN Security
    public static void verifyUser(Scanner sc) {
        System.out.println("--- SECURE ACCESS REQUIRED ---");
        int attempts = 0;
        while (attempts < 3) {
            System.out.print("Enter 4-Digit PIN: ");
            if (sc.nextLine().equals(SECRET_PIN)) {
                System.out.println("Access Granted.");
                return;
            }
            attempts++;
            System.out.println("Wrong PIN. " + (3 - attempts) + " attempts left.");
        }
        System.out.println("Security Lockout. Program Terminated.");
        System.exit(0);
    }

    // 2. Webhook Simulator (Loading Bar)
    public static void simulateWebhook() {
        System.out.print("Processing Payment ");
        try {
            for (int i = 0; i <= 10; i++) {
                String bar = "[";
                for (int j = 0; j < 10; j++) {
                    bar += (j < i) ? "#" : ".";
                }
                bar += "]" + (i * 10) + "%";
                System.out.print("\r" + bar);
                Thread.sleep(200);
            }
            System.out.println("\nWebhook: Merchant Notification Sent Successfully!");
        } catch (InterruptedException e) {
            System.out.println("Error in simulation.");
        }
    }

    // 3. Analytics Engine
    public static void showAnalytics() {
        double totalNGN = 0, totalUSD = 0, maxSpend = 0;
        String topMerchant = "None";

        try (BufferedReader br = new BufferedReader(new FileReader(LEDGER_FILE))) {
            String line;
            while ((line = br.readLine()) != null) {
                String[] data = line.split(",");
                if (data.length < 4)
                    continue;

                double amt = Double.parseDouble(data[2]);
                String curr = data[3];

                if (curr.equals("NGN"))
                    totalNGN += amt;
                else
                    totalUSD += amt;

                if (amt > maxSpend) {
                    maxSpend = amt;
                    topMerchant = data[1];
                }
            }
            System.out.println("\n--- TRANSACTION ANALYTICS ---");
            System.out.println("Total Spent (NGN): " + totalNGN);
            System.out.println("Total Spent (USD): " + totalUSD);
            System.out.println("Biggest Purchase: " + maxSpend + " to " + topMerchant);
            System.out.println("-----------------------------");
        } catch (IOException e) {
            System.out.println("No data for analytics.");
        }
    }

    public static void saveToLedger(String ref, String m, double amt, String curr) {
        try (FileWriter fw = new FileWriter(LEDGER_FILE, true); PrintWriter out = new PrintWriter(fw)) {
            out.println(ref + "," + m + "," + amt + "," + curr + "," + new Date());
        } catch (IOException e) {
            System.out.println("Log Error.");
        }
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        verifyUser(sc);

        double balanceNGN = 125500.00;
        double rate = 1500.0;

        while (true) {
            try {
                System.out.println("\n1. Pay | 2. History | 3. Analytics | 4. Exit");
                System.out.print("Select: ");
                String choice = sc.nextLine();

                if (choice.equals("4"))
                    break;

                if (choice.equals("1")) {
                    System.out.print("Currency (NGN/USD): ");
                    String curr = sc.nextLine().toUpperCase();
                    System.out.print("Merchant: ");
                    String m = sc.nextLine();
                    System.out.print("Amount: ");
                    double amt = Double.parseDouble(sc.nextLine());

                    double cost = curr.equals("USD") ? amt * rate : amt;

                    if (cost <= balanceNGN) {
                        simulateWebhook();
                        balanceNGN -= cost;
                        String ref = "PAY-" + UUID.randomUUID().toString().substring(0, 6).toUpperCase();
                        saveToLedger(ref, m, amt, curr);
                        System.out.println("SUCCESS: Ref " + ref);
                        System.out.println("New Balance: NGN " + balanceNGN);
                    } else {
                        System.out.println("FAILED: Insufficient Funds.");
                    }
                } else if (choice.equals("2")) {
                    System.out.println("\n--- HISTORY ---");
                    try (BufferedReader br = new BufferedReader(new FileReader(LEDGER_FILE))) {
                        String line;
                        while ((line = br.readLine()) != null)
                            System.out.println(line.replace(",", " | "));
                    } catch (Exception e) {
                        System.out.println("No records.");
                    }
                } else if (choice.equals("3")) {
                    showAnalytics();
                }
            } catch (Exception e) {
                System.out.println("INPUT ERROR: Please check your entries.");
            }
        }
        sc.close();
    }
}