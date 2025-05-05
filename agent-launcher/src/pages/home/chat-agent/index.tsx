import { ipcRenderer } from "electron";
import { useEffect, useState } from "react";

const ChatAgent = () => {
  const [jsString, setJSString] = useState("");

  useEffect(() => {
    fetch("src/pages/home/chat-agent/bundle.js")
      .then((response) => response.text())
      .then(async (data) => {
        // console.log('data', data);
        // new Function(data)();
        // const result = await window.electronAPI.executeCode(data);
        // console.log('result', result);
      });
  }, []);

  console.log("jsString", jsString);

  return <div>ChatAgent</div>;
};

export default ChatAgent;
