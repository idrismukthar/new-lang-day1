import * as fs from 'fs';
import * as http from 'http';

// The "Contract": Every transaction MUST have these exact types
interface Transaction {
    ref: string;
    merchant: string;
    amount: number;
    currency: string;
    date: string;
}

const LEDGER_FILE: string = 'ledger.csv';
const PORT: number = 3000;

function getTransactions(): Transaction[] | { error: string } {
    try {
        const data: string = fs.readFileSync(LEDGER_FILE, 'utf8');
        const lines: string[] = data.trim().split('\n');
        
        return lines.map(line => {
            const [ref, merchant, amount, currency, date] = line.split(',');
            return { 
                ref, 
                merchant, 
                amount: parseFloat(amount), 
                currency, 
                date 
            };
        });
    } catch (err) {
        return { error: "No transactions found yet." };
    }
}

const server = http.createServer((req, res) => {
    res.setHeader('Content-Type', 'application/json');
    res.setHeader('Access-Control-Allow-Origin', '*');

    if (req.url === '/api/transactions') {
        const transactions = getTransactions();
        res.writeHead(200);
        res.end(JSON.stringify(transactions, null, 2));
    } else {
        res.writeHead(404);
        res.end(JSON.stringify({ message: "Route not found. Try /api/transactions" }));
    }
});

server.listen(PORT, () => {
    console.log(`=== TS API DASHBOARD STARTED ===`);
    console.log(`Watching your Java/Go system at http://localhost:${PORT}/api/transactions`);
});