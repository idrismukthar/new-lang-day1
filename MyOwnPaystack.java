import java.util.*;
import java.io.*;

public class MyOwnPaystack {
    private static final String LEDGER_FILE = "ledger.csv";
    private static final String TREASURY_FILE = "treasury.csv";
    private static final String SECRET_PIN = "1234";

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        verifyUser(sc);

        double balanceNGN = 0.0;
        double balanceUSD = 0.0;

        // --- STEP 1: SYNC BALANCES FROM LEDGER ---
        try (BufferedReader br = new BufferedReader(new FileReader(LEDGER_FILE))) {
            String line;
            while ((line = br.readLine()) != null) {
                String[] data = line.split(",");
                if (data.length < 4)
                    continue;

                double amt = Double.parseDouble(data[2]);
                String curr = data[3];
                String ref = data[0];

                // P2P and CR entries are Credits (money in)
                boolean isCredit = ref.startsWith("P2P") || ref.startsWith("CR");

                if (curr.equalsIgnoreCase("NGN"))
                    balanceNGN += isCredit ? amt : -amt;
                else if (curr.equalsIgnoreCase("USD"))
                    balanceUSD += isCredit ? amt : -amt;
            }
            System.out.printf("\n[SYNC] Vault Loaded. NGN: ₦%,.2f | USD: $%,.2f\n", balanceNGN, balanceUSD);
        } catch (Exception e) {
            System.out.println("[INFO] No existing ledger found. Starting fresh.");
        }

        // --- STEP 2: MAIN MENU ---
        while (true) {
            try {
                System.out.println("\n1. Send Money | 2. View History | 3. Exit");
                System.out.print("Select: ");
                String choice = sc.nextLine();

                if (choice.equals("3"))
                    break;

                if (choice.equals("1")) {
                    handleTransaction(sc, balanceNGN, balanceUSD);
                    // Refresh balances after transaction
                    // (In a real app, you'd re-sync or update variables)
                } else if (choice.equals("2")) {
                    showHistory();
                }
            } catch (Exception e) {
                System.out.println("Error: " + e.getMessage());
            }
        }
        sc.close();
    }

    // Security Gate
    public static void verifyUser(Scanner sc) {
        System.out.println("\n--- ELITEBANK SECURE ACCESS ---");
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
        System.exit(0);
    }

    // Logic for sending money and taxing
    private static void handleTransaction(Scanner sc, double balNGN, double balUSD) {
        System.out.print("Currency (NGN/USD): ");
        String curr = sc.nextLine().toUpperCase().trim();
        System.out.print("Recipient: ");
        String recipient = sc.nextLine();
        System.out.print("Amount: ");
        double amt = Double.parseDouble(sc.nextLine());
        System.out.print("Narration: ");
        String desc = sc.nextLine();

        double currentBal = curr.equals("USD") ? balUSD : balNGN;

        if (amt <= currentBal) {
            simulateWebhook();
            String ref = "PAY-" + UUID.randomUUID().toString().substring(0, 6).toUpperCase();
            saveToVaults(ref, recipient, amt, curr, desc);
            System.out.println("TRANSACTION SUCCESSFUL. Ref: " + ref);
        } else {
            System.out.println("FAILED: Insufficient funds.");
        }
    }

    // Double-Entry Writing: Ledger and Treasury
    public static void saveToVaults(String ref, String m, double amt, String curr, String desc) {
        try {
            double bankCut = 0;
            double govtTax = 0;
            if ((curr.equals("NGN") && amt >= 10000) || (curr.equals("USD") && amt >= 10)) {
                bankCut = amt * 0.02;
                govtTax = amt * 0.01;
            }

            String date = new Date().toString().substring(0, 19);

            // 1. Write to Customer Ledger
            try (PrintWriter out = new PrintWriter(new FileWriter(LEDGER_FILE, true))) {
                out.println(ref + "," + m + "," + amt + "," + curr + "," + date + "," + desc + "," + bankCut + ","
                        + govtTax);
            }

            // 2. Write to Bank Treasury (The Profit Vault)
            if (bankCut > 0) {
                try (PrintWriter out = new PrintWriter(new FileWriter(TREASURY_FILE, true))) {
                    out.println(date + "," + ref + "," + bankCut + "," + govtTax + "," + curr);
                }
            }
        } catch (IOException e) {
            System.out.println("Vault Write Error!");
        }
    }

    public static void simulateWebhook() {
        System.out.print("Notifying Merchant ");
        try {
            for (int i = 0; i < 5; i++) {
                Thread.sleep(200);
                System.out.print(".");
            }
            System.out.println(" [DONE]");
        } catch (InterruptedException e) {
        }
    }

    public static void showHistory() throws IOException {
        System.out.println("\n--- TRANSACTION HISTORY ---");
        File f = new File(LEDGER_FILE);
        if (!f.exists())
            return;
        BufferedReader br = new BufferedReader(new FileReader(f));
        String line;
        while ((line = br.readLine()) != null)
            System.out.println(line.replace(",", " | "));
        br.close();
    }
}