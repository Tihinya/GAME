import Gachi, { useEffect, useState } from "../../Gachi.js/src/core/framework";
import { subscribe, ws } from "../../additional-functions/websocket";

import "./lobby.css";

const maxUsers = 4;
const states = {
  awaiting_players_state: "Wait players to suck cock",
  closing_lobby_state: "Time to close lobby!",
  starting_game_state: "Gays starts in...",
};

export default function Lobby() {
  const [messages, setMessages] = useState([]);
  const [timer, setTimer] = useState({ state: "", currentTime: 0 });
  const [players, setPlayers] = useState([]);

  useEffect(() => {
    subscribe("online_users_list", (list) => {
      console.log(list);
      const onlineList = Object.keys(list).map((username) => ({
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

    subscribe("ama_boy_next_door", ({ state, currentTime }) => {
      console.log(currentTime, state);
      setTimer({ state, currentTime });
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
          <p>{states[timer.state]}</p>
          <p>
            {timer.state !== states["awaiting_players_state"]
              ? "0:" + timer.currentTime
              : ""}
          </p>
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
