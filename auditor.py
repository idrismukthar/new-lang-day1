import csv
import os

TREASURY_FILE = "treasury.csv"

def run_audit():
    print("=" * 40)
    print("      ELITEBANK WEALTH AUDITOR V1.0     ")
    print("=" * 40)

    if not os.path.exists(TREASURY_FILE):
        print(f"[ERROR] No vault data found at {TREASURY_FILE}")
        return

    # Profit Trackers
    stats = {
        "NGN": {"revenue": 0.0, "tax": 0.0, "count": 0},
        "USD": {"revenue": 0.0, "tax": 0.0, "count": 0}
    }

    try:
        with open(TREASURY_FILE, mode='r') as file:
            reader = csv.reader(file)
            for row in reader:
                if not row: continue # Skip empty lines
                
                # Format: Date, Ref, Revenue, Tax, Currency
                try:
                    revenue = float(row[2])
                    tax = float(row[3])
                    currency = row[4].strip().upper()

                    if currency in stats:
                        stats[currency]["revenue"] += revenue
                        stats[currency]["tax"] += tax
                        stats[currency]["count"] += 1
                except (ValueError, IndexError):
                    continue # Skip corrupted rows

        # --- DISPLAY RESULTS ---
        for curr, data in stats.items():
            symbol = "₦" if curr == "NGN" else "$"
            print(f"\n[{curr} ASSETS]")
            print(f"  Total Transactions: {data['count']}")
            print(f"  Bank Net Revenue:   {symbol}{data['revenue']:,.2f}")
            print(f"  Govt Tax Collected: {symbol}{data['tax']:,.2f}")
            print(f"  --------------------------")
            print(f"  TOTAL VAULT VALUE:  {symbol}{data['revenue'] + data['tax']:,.2f}")

    except Exception as e:
        print(f"Audit Failed: {e}")

if __name__ == "__main__":
    run_audit()