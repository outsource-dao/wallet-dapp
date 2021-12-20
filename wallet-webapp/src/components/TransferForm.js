import React, { useState, useEffect } from "react";
import { useLocalStorage } from '@rehooks/local-storage';

export default function TransferForm({ authToken }) {
  const [toAddr, setToAddr] = useState("");
  const [disabled, setDisabled] = useState(true);
  const [tokenAmount, setTokenAmount] = useState(0);
  const [pubKey, setpubKey] = useLocalStorage("pubkey");
  const [qrURL, setQRUrl] = useLocalStorage("qrURL");
  
  useEffect(() => {
    checkDisabled();
  });
  const checkDisabled = () => {
    if (authToken && pubKey && toAddr && tokenAmount > 0) {
      setDisabled(false);
      return;
    }
    setDisabled(true);
  };
  const handleTransfer = (evt) => {
    evt.preventDefault();
    const txParams = { from: pubKey, toAddr, tokenAmount };
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${authToken}` },
      // TODO: better we use base64 encoding to transfer values
      body: JSON.stringify(txParams)
    };
    // create unsigned transaction from dApp-backend
    fetch('/api/v0/transferToken', requestOptions)
      .then(response => response.json())
      .then(data => setQRUrl(data))
  }
  return (
    <form className="transfer-form" onSubmit={handleTransfer}>
      <label>
        Your Public address that you will use to sign the transaction:
        <input type="text" value={pubKey} name="pubkey" onChange={(e) => setpubKey(e.target.value)} />
      </label>
      <label>
        To address:
        <input type="text" value={toAddr} name="toAddress" onChange={(e) => setToAddr(e.target.value)} />
      </label>
      <label>
        Token amount:
        <input
          type="number"
          step="1"
          value={tokenAmount}
          name="token"
          onChange={(e) => setTokenAmount(e.target.value)}
        />
      </label>
      <input type="submit" value="Transfer token" disabled={disabled} />
    </form>
  );
}
