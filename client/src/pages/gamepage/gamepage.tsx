import Gachi, {
  useContext,
  useEffect,
  useState,
} from "../../Gachi.js/src/core/framework";
import GameInput from "../../modules/commandEvents";
import { subscribe } from "../../additional-functions/websocket";
import { GameComponent } from "./gameComponent";
import { EventType, GameState, Page, PageState } from "../../types/Types";

import "./gamepage.css";

const names = ["Billy Herrington", "Van Darkholme", "Steve Rambo", "Our Daddy"]
  .map((value) => ({ value, sort: Math.random() }))
  .sort((a, b) => a.sort - b.sort)
  .map(({ value }) => value);

export default function GamePage() {
  const [gameState, setGameState] = useState<GameState>({
    players: [],
    powerups: [],
    map: [],
  });
  const navigate = useContext("switchPage");

  useEffect(() => {
    subscribe(EventType.GameEvent, (state: GameState) => {
      setGameState(state);
    });

    subscribe(EventType.GameEventGameState, ({ state }: Page) => {
      switch (state) {
        case PageState.MainPage:
          navigate("main");
      }
    });
  }, []);

  return (
    <div className="game-page-container">
      <div className="game-page-info">
        <div className="game-page-panel">
          <div className="game-page-label">
            Players: {gameState.players.length}/4
          </div>
          <div className="game-page-label-count-players">
            {gameState.players.map((p, i) => (
              <div>{names[i]}</div>
            ))}
          </div>
        </div>
        <div className="game-page-panel">
          <div className="game-page-label">Time</div>
          {/* <div className="game-page-label-count">4:19</div> */}
        </div>
        <div className="game-page-panel">
          <div className="game-page-label">Bombs</div>
          {/* <div className="game-page-label-count">13000</div> */}
        </div>
        <div className="game-page-panel">
          <div className="game-page-label">Lives</div>
          {/* <div className="game-page-label-count">2</div> */}
        </div>
      </div>
      <GameInput />
      <GameComponent gameState={gameState} />
    </div>
  );
}
