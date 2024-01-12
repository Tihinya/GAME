import Gachi, {
  useNavigate,
  useState,
} from "../../Gachi.js/src/core/framework";
import "./mainpage.css";
import logo from "../../public/img/logo.png";

export default function MenuPage() {
  const navigate = useNavigate();
  const [playerName, setPlayerName] = useState("");
  const [playersData, setPlayersData] = useState({
    players: [],
  });

  const showWriteName = () => {
    let playButton = document.getElementById("play-button");
    let nameContainer = document.getElementById("name-container");

    playButton.style.display = "none";

    nameContainer.style.display = "flex";
  };

  const acceptName = () => {
    // Update the JSON object with the entered name
    setPlayersData((prevData) => ({
      ...prevData,
      players: [...prevData.players, { name: playerName }],
    }));

    // Navigate to the lobby
    navigate("/lobby");
    console.log(playersData);
    console.log(playerName);
  };

  return (
    <div className="main-page">
      <div className="main-page_container">
        <img className="logo" src={logo} alt="Bomberman" />
        <div className="play-button-container">
          <div id="play-button" className="play-button" onClick={showWriteName}>
            PLAY
          </div>
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
        </div>
      </div>
    </div>
  );
}
