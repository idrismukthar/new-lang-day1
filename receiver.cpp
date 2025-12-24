#include <iostream>
#include <fstream>
#include <string>
#include <ctime>
#include <algorithm>
#include <iomanip>

using namespace std;

// Helper to get formatted time
string getCurrentTime()
{
    time_t now = time(0);
    char *dt = ctime(&now);
    string str(dt);
    if (!str.empty() && str[str.length() - 1] == '\n')
        str.erase(str.length() - 1);
    return str;
}

int main()
{
    cout << "========================================" << endl;
    cout << "    ELITEBANK C++ TERMINAL V3 (PRO)     " << endl;
    cout << "    (Ledger & Treasury Integration)     " << endl;
    cout << "========================================" << endl;

    string sender, currency, narration;
    double amount;

    while (true)
    {
        cout << "\nEnter Sender Name (or 'exit' to quit): ";
        getline(cin, sender);
        if (sender == "exit")
            break;

        // 1. Currency Selection with Validation
        while (true)
        {
            cout << "Enter Currency (NGN/USD): ";
            getline(cin, currency);
            for (auto &c : currency)
                c = toupper(c);
            if (currency == "NGN" || currency == "USD")
                break;
            cout << "INVALID: Please use NGN or USD." << endl;
        }

        // 2. Amount Entry
        cout << "Enter Amount: ";
        string amtStr;
        getline(cin, amtStr);

        // 3. Narration Entry
        cout << "Enter Narration (Description): ";
        getline(cin, narration);
        if (narration.empty())
            narration = "Cash Deposit";

        try
        {
            amount = stod(amtStr);

            // --- BILLING LOGIC (The Bank's Cut) ---
            double bankFee = 0.0;
            double govtTax = 0.0;

            // Revenue triggers: NGN >= 10,000 or USD >= 10
            bool shouldCharge = (currency == "NGN" && amount >= 10000) || (currency == "USD" && amount >= 10);

            if (shouldCharge)
            {
                bankFee = amount * 0.02; // 2% Bank Profit
                govtTax = amount * 0.01; // 1% Govt Tax
            }

            string ref = "P2P-" + to_string(time(0)).substr(4);
            string timestamp = getCurrentTime();

            // STEP 1: Update Customer Ledger
            ofstream ledger("ledger.csv", ios::app);
            if (ledger.is_open())
            {
                // Format: Ref,Merchant,Amount,Currency,Date,Desc,BankFee,GovtTax
                ledger << ref << "," << sender << "," << amount << "," << currency << ","
                       << timestamp << "," << narration << "," << fixed << setprecision(2)
                       << bankFee << "," << govtTax << endl;
                ledger.close();
            }

            // STEP 2: Update Bank Treasury (The Private Vault)
            if (bankFee > 0)
            {
                ofstream treasury("treasury.csv", ios::app);
                if (treasury.is_open())
                {
                    // Format: Date,Ref,Revenue,Tax,Currency
                    treasury << timestamp << "," << ref << "," << bankFee << ","
                             << govtTax << "," << currency << endl;
                    treasury.close();
                }
            }

            cout << "\n----------------------------------------" << endl;
            cout << " ✅ DEPOSIT SECURED" << endl;
            cout << " Ref: " << ref << endl;
            cout << " Bank Revenue: " << currency << " " << bankFee << endl;
            cout << "----------------------------------------" << endl;
        }
        catch (...)
        {
            cout << "ERROR: Invalid amount. Please enter numbers only." << endl;
        }
    }

    cout << "\nTerminal Closed. EliteBank Secure." << endl;
    return 0;
}