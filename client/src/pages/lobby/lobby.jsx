import Gachi, { useEffect, useState } from "../../Gachi.js/src/core/framework";
import { subscribe, ws } from "../../additional-functions/websocket";

import "./lobby.css";

const maxUsers = 4;

export default function Lobby() {
  const [messages, setMessages] = useState([]);
  const [timer, setTimer] = useState(20);
  const [players, setPlayers] = useState([]);

  useEffect(() => {
    subscribe("online_users_list", ({ list }) => {
      const onlineList = Object.entries(list).map(([id, username]) => ({
        id: Number(id),
        username: username,
      }));

      setPlayers(onlineList);
    });

    subscribe("receive_message", ({ name, message }) => {
      setMessages((prev) => {
        const temp = [...prev];
        temp.unshift({ username: name, content: message });

        return temp;
      });
    });
  }, []);

  const handleInput = (e) => {
    if (e.code !== "Enter") {
      return;
    }
    e.preventDefault();

    const content = e.target.value.trim();
    e.target.value = "";
    if (content === "") {
      return;
    }
    console.log(content);
    ws.send(
      JSON.stringify({
        type: "send_message",
        payload: { message: content },
      })
    );
  };

  return (
    <div className="main-page">
      <div className="main-page-chat">
        <div className="message-container">
          {messages.map((message) => {
            return (
              <div className="chat-message">
                <div className="chat-message-name">{message.username}</div>
                <div className="chat-message-message">{message.content}</div>
              </div>
            );
          })}
        </div>

        <div className="input-container">
          <input type="text" className="chat-input" onKeyDown={handleInput} />
        </div>
      </div>
      <div className="main-page-lobby-info">
        <div className="main-page-lobby-timer">
          <p>Time to close lobby:</p>
          <p>0:{timer}</p>
        </div>
        <div className="main-page-lobby-players">
          <p>
            Players:{players.length}/{maxUsers}
          </p>
          {players.map(({ username }) => {
            return <p>{username}</p>;
          })}
        </div>
      </div>
    </div>
  );
}
