import React, { Fragment, useState, useEffect } from "react";
import { useLocalStorage } from '@rehooks/local-storage';

import logo from "./logo.svg";
import "./App.css";
import TransferForm from "./components/TransferForm";
import QRComponent from "./components/QRComponent";
import SwapForm from './components/SwapForm';
import ApproveForm from './components/ApproveForm';
import BalanceForm from "./components/BalanceForm";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <Fragment>
        <MainView />
      </Fragment>
    </div>
  );
}

function MainView() {
  const [qrURL, setQRUrl] = useLocalStorage("qrURL");
  const [authToken, setAuthToken] = useState("");
  const [pubKey] = useLocalStorage("pubkey");

  useEffect(() => {
    setQRUrl('');
    // get authToken from dApp backend
    fetch("/api/v0/genToken/123")
      .then((response) => response.json())
      .then((data) => {
        setAuthToken(data.token);
        return;
      })
      .catch((error) => console.error(error));
  }, []);

  useEffect(() => {

  }, [pubKey]);

  return (
    <div className="container">
      <Fragment>
        <TransferForm authToken={authToken}/>
        <ApproveForm authToken={authToken}/>
        <SwapForm authToken={authToken}/>
        <BalanceForm authToken={authToken}/>
        {qrURL && <QRComponent/>}
      </Fragment>
    </div>
  );
}

export default App;
