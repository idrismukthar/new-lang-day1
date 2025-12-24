const express = require("express");
const fs = require("fs");
const cors = require("cors");
const app = express();

const LEDGER_FILE = "ledger.csv";
const TREASURY_FILE = "treasury.csv";
const PORT = 3000;

app.use(cors());
app.use(express.json());

// 1. GET: Read expanded ledger data for the React Dashboard
app.get("/api/transactions", (req, res) => {
  try {
    if (!fs.existsSync(LEDGER_FILE)) return res.json([]);

    const data = fs.readFileSync(LEDGER_FILE, "utf8");
    const lines = data
      .trim()
      .split("\n")
      .filter((line) => line.length > 5);

    const transactions = lines.map((line) => {
      // Matches the 8-column format: Ref, Merchant, Amount, Currency, Date, Desc, BankFee, GovtTax
      const [ref, merchant, amount, currency, date, desc, bankFee, govtTax] =
        line.split(",");
      return {
        ref,
        merchant,
        amount: parseFloat(amount),
        currency,
        date,
        description: desc || "No narration",
        bankFee: parseFloat(bankFee || 0),
        govtTax: parseFloat(govtTax || 0),
      };
    });
    res.json(transactions);
  } catch (err) {
    res.status(500).json({ error: "Failed to read ledger" });
  }
});

// 2. POST: Process transfer, Charge Fees, and Fund the Treasury
app.post("/api/transactions", (req, res) => {
  const { merchant, amount, currency, description } = req.body;
  const amt = parseFloat(amount);

  // --- REVENUE LOGIC ---
  let bankFee = 0;
  let govtTax = 0;
  const shouldCharge =
    (currency === "NGN" && amt >= 10000) || (currency === "USD" && amt >= 10);

  if (shouldCharge) {
    bankFee = amt * 0.02; // 2% Bank Profit
    govtTax = amt * 0.01; // 1% Government Tax
  }

  const ref = "DR-" + Math.random().toString(36).substring(7).toUpperCase();
  const date = new Date().toString().split(" ").slice(0, 5).join(" ");
  const narration = description || "N/A";

  // Line for Customer Ledger (8 columns)
  const ledgerLine = `${ref},${merchant},${amt},${currency},${date},${narration},${bankFee},${govtTax}\n`;

  // Line for Bank Treasury (5 columns)
  const treasuryLine = `${date},${ref},${bankFee},${govtTax},${currency}\n`;

  try {
    // Save to Customer Record
    fs.appendFileSync(LEDGER_FILE, ledgerLine);

    // Save to Bank Vault if there's profit
    if (bankFee > 0) {
      fs.appendFileSync(TREASURY_FILE, treasuryLine);
      console.log(`[TREASURY] Profit Secured: ${currency} ${bankFee}`);
    }

    res.status(200).json({ success: true, ref });
  } catch (err) {
    console.error("File Write Error:", err);
    res.status(500).json({ success: false });
  }
});

app.listen(PORT, () => {
  console.log(`\n=== ELITEBANK API V3 ACTIVE ===`);
  console.log(`Connected to Ledger: ${LEDGER_FILE}`);
  console.log(`Connected to Treasury: ${TREASURY_FILE}`);
});
