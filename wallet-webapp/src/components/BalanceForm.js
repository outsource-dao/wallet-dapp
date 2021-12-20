import React, { useState, useEffect } from "react";
import { useLocalStorage } from '@rehooks/local-storage';
export default function BalanceForm({ authToken }) {
  const [disabled, setDisabled] = useState(true);
  const [balance, setBalance] = useState(0);
  const [pubKey, setpubKey] = useLocalStorage("pubkey");
  
  useEffect(() => {
    checkDisabled();
  });
  const checkDisabled = () => {
    if (authToken && pubKey) {
      setDisabled(false);
      return;
    }
    setDisabled(true);
  };
  const handleTransfer = (evt) => {
    evt.preventDefault();
    const txParams = { pubkey: pubKey };
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json", Authorization: `Bearer ${authToken}` },
      body: JSON.stringify(txParams),
    };
    fetch("/api/v0/getTokenBalance", requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setBalance(data);
      });
  };
  return (
    <div>
      <form className="transfer-form" onSubmit={handleTransfer}>
        <label>
          Public address:
          <input type="text" value={pubKey} name="pubkey" onChange={(e) => setpubKey(e.target.value)} />
        </label>
        <input type="submit" value="Check balance" disabled={disabled} />
      </form>
      <p>{balance}</p>
    </div>
  );
}
