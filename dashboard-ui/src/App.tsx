import { useState, useEffect } from 'react'
import { jsPDF } from "jspdf"; 
import './App.css'

interface Transaction {
  ref: string;
  merchant: string;
  amount: number;
  currency: string;
  date: string;
  description: string;
  bankFee: number;
  govtTax: number;
}

function App() {
  const [txns, setTxns] = useState<Transaction[]>([]);
  const [showModal, setShowModal] = useState(false);
  
  // FIXED: Initialized amount as 0 (number) to satisfy TypeScript
  const [formData, setFormData] = useState({ 
    account: '', 
    amount: 0, 
    currency: 'NGN', 
    merchant: '', 
    description: '' 
  });

  const fetchData = () => {
    fetch('http://localhost:3000/api/transactions')
      .then(res => res.json())
      .then(data => { if (Array.isArray(data)) setTxns([...data].reverse()); })
      .catch(err => console.log("API Error:", err));
  };

  useEffect(() => {
    fetchData();
    const interval = setInterval(fetchData, 3000);
    return () => clearInterval(interval);
  }, []);

  const handleTransfer = async (e: React.FormEvent) => {
    e.preventDefault();
    const response = await fetch('http://localhost:3000/api/transactions', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        merchant: formData.merchant || formData.account,
        amount: formData.amount,
        currency: formData.currency,
        description: formData.description
      })
    });
    
    if (response.ok) {
      setShowModal(false);
      // Reset form
      setFormData({ account: '', amount: 0, currency: 'NGN', merchant: '', description: '' });
      fetchData();
    }
  };

  const generateReceipt = (t: Transaction) => {
    const doc = new jsPDF();
    doc.setFont("helvetica", "bold");
    doc.text("ELITEBANK OFFICIAL RECEIPT", 20, 20);
    doc.setFont("helvetica", "normal");
    doc.setFontSize(10);
    doc.text(`Ref: ${t.ref}`, 20, 35);
    doc.text(`Date: ${t.date}`, 20, 42);
    doc.line(20, 45, 100, 45);
    doc.text(`Recipient: ${t.merchant}`, 20, 55);
    doc.text(`Narration: ${t.description}`, 20, 62);
    doc.setFontSize(14);
    doc.text(`Amount: ${t.currency === 'USD' ? '$' : '₦'}${t.amount.toLocaleString()}`, 20, 75);
    doc.setFontSize(10);
    doc.text(`Bank Fee: ${t.bankFee}`, 20, 85);
    doc.text(`Govt Tax: ${t.govtTax}`, 20, 92);
    doc.save(`Receipt-${t.ref}.pdf`);
  };

  const calcTotal = (curr: string) => {
    return txns.reduce((acc, t) => {
      if (t.currency !== curr) return acc;
      // Recognized Credits (P2P from C++, CR from manual)
      const isCredit = t.ref.startsWith('CR') || t.ref.startsWith('P2P');
      // For debits, we subtract the amount PLUS the fees taken from the user
      return isCredit ? acc + t.amount : acc - (t.amount + t.bankFee + t.govtTax);
    }, 0);
  };

  const totalTax = (curr: string) => {
    return txns.reduce((acc, t) => (t.currency === curr ? acc + t.govtTax : acc), 0);
  };

  return (
    <div className="app-wrapper">
      <div className="container">
        <header className="bank-header">
          <h1>ELITEBANK<span>.OS</span></h1>
          <button className="transfer-btn" onClick={() => setShowModal(true)}>+ New Transfer</button>
        </header>

        <div className="dashboard-grid">
          <div className="balance-card naira">
            <p className="label">NGN BALANCE</p>
            <h2 className="amount">₦{calcTotal('NGN').toLocaleString()}</h2>
            <p className="tax-label">Tax Collected: ₦{totalTax('NGN').toLocaleString()}</p>
          </div>
          <div className="balance-card dollar">
            <p className="label">USD BALANCE</p>
            <h2 className="amount">${calcTotal('USD').toLocaleString()}</h2>
            <p className="tax-label">Tax Collected: ${totalTax('USD').toLocaleString()}</p>
          </div>
        </div>

        <div className="txn-container">
          <h3>Recent Activity</h3>
          <div className="txn-list">
            {txns.map((t, i) => {
              const isCredit = t.ref.startsWith('CR') || t.ref.startsWith('P2P');
              return (
                <div key={i} className="txn-item">
                  <div className="txn-left">
                    <div className={`avatar ${isCredit ? 'credit-icon' : 'debit-icon'}`}>
                      {isCredit ? '↓' : '↑'}
                    </div>
                    <div className="txn-details">
                      <span className="merchant-name">{t.merchant}</span>
                      <span className="txn-narration">{t.description}</span>
                      <span className="txn-ref">{t.ref} • {t.date}</span>
                    </div>
                  </div>
                  <div className="txn-right text-right">
                    <span className={`txn-value ${isCredit ? 'text-green' : 'text-red'}`}>
                      {isCredit ? '+' : '-'}{t.currency === 'USD' ? '$' : '₦'}{Math.abs(t.amount).toLocaleString()}
                    </span>
                    <button onClick={() => generateReceipt(t)} className="btn-receipt">Receipt</button>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </div>

      {showModal && (
        <div className="modal-overlay">
          <div className="modal-content">
            <h2>Send Money</h2>
            <form onSubmit={handleTransfer}>
              <input type="text" placeholder="Recipient / Account Number" required 
                onChange={e => setFormData({...formData, merchant: e.target.value})} />
              
              <input type="number" placeholder="Amount" required 
                value={formData.amount || ''} 
                onChange={e => setFormData({...formData, amount: parseFloat(e.target.value) || 0})} />
              
              <input type="text" placeholder="Narration (Description)" required 
                onChange={e => setFormData({...formData, description: e.target.value})} />
              
              <select onChange={e => setFormData({...formData, currency: e.target.value})}>
                <option value="NGN">NGN (Naira)</option>
                <option value="USD">USD (Dollar)</option>
              </select>
              <div className="modal-actions">
                <button type="button" onClick={() => setShowModal(false)}>Cancel</button>
                <button type="submit" className="confirm-btn">Confirm Transfer</button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  )
}

export default App;