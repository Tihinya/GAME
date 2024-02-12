import Gachi from "../../Gachi.js/src/core/framework";
import { GameState } from "../../types/Types";

type GameComponentProps = {
  gameState: GameState;
};

const cellWidth = 50;

export function GameComponent({ gameState }: GameComponentProps) {
  if (gameState.map.length === 0) {
    return;
  }

  return (
    <div
      className="game-page-map"
      style={`grid-template-columns: repeat(${gameState.map[0].length}, ${cellWidth}px); grid-template-rows: repeat(${gameState.map.length}, ${cellWidth}px);`}
    >
      {gameState.map.flatMap((row, y) => {
        return row.map((block, x) => {
          return <Tile key={`cell${x}-${y}`} name={block?.name || ""} />;
        });
      })}

      {gameState.players.map((position) => {
        return (
          <div
            className={"cell player"}
            style={`top: ${position.y * 1.25}px; left: ${position.x * 1.25}px`}
          ></div>
        );
      })}

      {gameState.powerups.map((position) => {
        return (
          <div
            className={"powerup"}
            style={`top: ${position.y}px; left: ${position.x}px`}
          ></div>
        );
      })}
    </div>
  );
}

const tileClasses = {
  wall: "wall",
  explosion: "explosion",
  box: "box",
  bomb: "bomb",
  "": "",
};

function Tile({ key, name }: { key: string; name: string }) {
  return <div className={"cell " + tileClasses[name]} key={key} />;
}
