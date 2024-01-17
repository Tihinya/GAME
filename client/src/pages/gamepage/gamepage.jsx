import Gachi, { useNavigate } from "../../Gachi.js/src/core/framework";
import "./gamepage.css";
import GameComponent from "../../modules/commandEvents";

export default function GamePage() {
  return (
    <div className="game-page-container">
      <div className="game-page-info">
        <div className="game-page-panel">
          <div className="game-page-label">Players: 4/4</div>
          <div className="game-page-label-count-players">
            <div>HungryStepan</div>
            <div>Skibidick</div>
            <div>dabude dabudai</div>
            <div>bibabiba</div>
          </div>
        </div>
        <div className="game-page-panel">
          <div className="game-page-label">Time</div>
          <div className="game-page-label-count">4:19</div>
        </div>
        <div className="game-page-panel">
          <div className="game-page-label">Score</div>
          <div className="game-page-label-count">13000</div>
        </div>
        <div className="game-page-panel">
          <div className="game-page-label">Lives</div>
          <div className="game-page-label-count">2</div>
        </div>
      </div>
      <div className="game-page-map">
        <GameComponent />
      </div>
    </div>
  );
}
