import Gachi, {
  useContext,
  useEffect,
  useState,
} from "../../Gachi.js/src/core/framework";
import logo from "../../public/img/logo.png";
import { subscribe, ws } from "../../additional-functions/websocket";

import "./mainpage.css";

export default function MenuPage() {
  const [playerName, setPlayerName] = useState("");
  const [isNameShowen, setIsNameShowen] = useState(false);
  const navigate = useContext("switchPage");

  useEffect(() => {
    subscribe("game_state", ({ state }) => {
      console.log(state);

      if (state === "lobby") {
        navigate("lobby");
      }
    });
  }, []);

  const acceptName = () => {
    console.log(playerName);

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
