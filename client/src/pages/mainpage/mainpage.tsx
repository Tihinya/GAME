import Gachi, {
  useContext,
  useEffect,
  useState,
} from "../../Gachi.js/src/core/framework";
import { subscribe, ws } from "../../additional-functions/websocket";

import { EventType, Page, PageState } from "../../types/Types";
import logo from "/sprites/img/logo.png";
import "./mainpage.css";

export default function MenuPage() {
  const [playerName, setPlayerName] = useState("");
  const [isNameShowen, setIsNameShowen] = useState(false);
  const navigate = useContext("switchPage");

  useEffect(() => {
    subscribe(EventType.GameEventGameState, ({ state }: Page) => {
      switch (state) {
        case PageState.Lobby:
          navigate("lobby");
      }
    });
  }, []);

  const acceptName = () => {
    ws.send(
      JSON.stringify({
        type: "register_user",
        payload: { username: playerName },
      })
    );
  };

  return (
    <div className="main-page">
      <div className="main-page_container">
        <img className="logo" src={logo} alt="Bomberman" />
        <div className="play-button-container">
          {!isNameShowen ? (
            <div
              id="play-button"
              className="play-button"
              onClick={() => setIsNameShowen(true)}
            >
              PLAY
            </div>
          ) : (
            <div id="name-container" className="name-container">
              <input
                className="write-name"
                placeholder="Write your name"
                maxLength="21"
                value={playerName}
                onChange={(e) => setPlayerName(e.target.value)}
              />
              <div className="accept-name" onClick={acceptName}>
                Accept
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
