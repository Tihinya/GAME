import Gachi from "../../Gachi.js/src/core/framework";
import { GameState } from "../../types/Types";

type GameComponentProps = {
  gameState: GameState;
};

const cellWidth = 50;

const powerupClasses = {
  speedPowerUp: "powerupSpeed",
  exposionPowerup: "powerupBlastRadius",
  bombPowerUp: "powerupBomb",
};
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

      {gameState.players
        .sort(({ id: idA }, { id: idB }) => idA - idB)
        .map((position) => {
          return (
            <div
              className={"cell player"}
              style={`top: ${position.y * 1.25}px; left: ${
                position.x * 1.25
              }px`}
            ></div>
          );
        })}

      {gameState.powerups.map((position, index) => {
        const powerUpName = gameState.powerups[index]?.name;
        const powerupClass = powerupClasses[powerUpName];

        return (
          <div
            className={`powerup ${powerupClass}`}
            style={`top: ${position.y * 1.25}px; left: ${position.x * 1.25}px`}
            key={`powerup-${index}`}
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
