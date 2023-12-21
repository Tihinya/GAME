import Gachi from "../../Gachi.js/src/core/framework";
import "../mainpage/mainpage.css";

export default function MenuPage() {
  const showWriteName = () => {
    let playButton = document.getElementById("play-button");
    let nameContainer = document.getElementById("name-container");

    // Hide the play button
    playButton.style.display = "none";

    // Show the write name div
    nameContainer.style.display = "flex";
  };

  return (
    <div className="main-page">
      <div className="main-page_container">
        <img className="logo" src="./pages/img/logo" alt="Bomberman" />
        <div className="play-button-container">
          <div id="play-button" className="play-button" onClick={showWriteName}>
            PLAY
          </div>
          <div id="name-container" className="name-container">
            <input
              className="write-name"
              placeholder="Write your name"
              maxLength="21"
            />
            <div className="accept-name">Accept</div>
          </div>
        </div>
      </div>
    </div>
  );
}
