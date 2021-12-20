import React, { useEffect, useState } from "react";
import { useLocalStorage } from '@rehooks/local-storage';
import QRCode from "qrcode.react";

export default function QRComponent() {
  const [bgColor, setBgColor] = useState('#FFFFFF');
  const [qrURL] = useLocalStorage("qrURL");
  useEffect(() => {
    setBgColor('#' + Math.floor(Math.random()*16777215).toString(16));
  }, [qrURL]);
  const containerStyle = {
    'border': 'solid 5px',
    'border-color': bgColor,
    'padding': '2px'
  }
  return (
    <div>
      <h2>Scan following QR code to load transaction in your ethential wallet</h2>
      <div style={containerStyle}>
        <QRCode
          value={qrURL}
          size={200}
          bgColor={"#ffffff"}
          fgColor={"#000000"}
          level={"L"}
          includeMargin={false}
          renderAs={"canvas"}
        />
      </div>
    </div>
  );
}
