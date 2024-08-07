import { ChangeEvent, useState } from "react";
import useWebSocket from "react-use-websocket";

export default function Home() {
  const [messages, setMessages] = useState<string[]>([]);
  
  useWebSocket("ws://localhost:8080/ws?username=diogo", {
    onOpen: () => {
      console.log("Connection openned!");
    },
    onMessage: (e: WebSocketEventMap["message"]) => {
      const message = e.data;
      if(message == null) return;
      setMessages([...messages, message]);
    },
  });

  return (
    <main>
      <div>
        {
          messages.map(value => {
            return <div>{value}</div>
          })
        }
      </div>
    </main>
  );
}
