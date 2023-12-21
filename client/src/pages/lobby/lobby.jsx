import Gachi from "../../Gachi.js/src/core/framework";

import "./lobby.css";

export default function Lobby() {
  return (
    <div className="main-page">
      <div className="main-page-chat">
        <div className="chat-message">
          <div className="chat-message-name">Bibaboba</div>
          <div className="chat-message-message">oh yes sir!</div>
          <div className="chat-message-name">Bibaboba</div>
          <div className="chat-message-message">oh yes sir!</div>
          <div className="chat-message-name">Bibaboba</div>
          <div className="chat-message-message">oh yes sir!</div>
        </div>
        <div className="input-container">
          <input type="text" className="chat-input" />
        </div>
      </div>
      <div className="main-page-lobby-info"></div>
    </div>
  );
}
