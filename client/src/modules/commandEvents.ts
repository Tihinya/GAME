import Gachi, { useEffect, useState } from "../Gachi.js/src/core/framework";
import { ws } from "../additional-functions/websocket";
import { EventType } from "../types/Types";

const keysPressed = {
  KeyW: false,
  KeyA: false,
  KeyS: false,
  KeyD: false,
  Space: false,
};

const GameInput = () => {
  const handleKeyDown = (event: KeyboardEvent) => {
    const key = event.code;
    event.repeat;

    if (keysPressed.hasOwnProperty(key) && !event.repeat) {
      event.preventDefault();

      keysPressed[key] = true;
      sendKeysToBackend(keysPressed);
    }
  };

  const handleKeyUp = (event: KeyboardEvent) => {
    const key = event.code;
    if (keysPressed.hasOwnProperty(key)) {
      event.preventDefault();

      keysPressed[key] = false;
      sendKeysToBackend(keysPressed);
    }
  };

  useEffect(() => {
    window.addEventListener("keydown", handleKeyDown);
    window.addEventListener("keyup", handleKeyUp);

    return () => {
      window.removeEventListener("keydown", handleKeyDown);
      window.removeEventListener("keyup", handleKeyUp);
    };
  }, []);
};

export default GameInput;

function sendKeysToBackend(updatedKeys) {
  const jsonToSend = JSON.stringify({
    type: EventType.GameEventInput,
    payload: { keys: updatedKeys },
  });

  ws.send(jsonToSend);
}
