import Gachi, { useEffect, useState } from "../Gachi.js/src/core/framework.ts";

const GameComponent = () => {
  const [keysPressed, setKeysPressed] = useState({
    KeyW: true,
    KeyA: true,
    KeyS: true,
    KeyD: true,
    Space: true,
  });

  const handleKeyDown = (event) => {
    const key = event.code;

    if (keysPressed.hasOwnProperty(key)) {
      if (event.code === "Space") {
        event.preventDefault();
      }

      setKeysPressed((prevKeys) => ({ ...prevKeys, [key]: false }));
      // Send JSON to backend
      sendKeysToBackend();
    }
  };

  const handleKeyUp = (event) => {
    const key = event.code;
    if (keysPressed.hasOwnProperty(key)) {
      setKeysPressed((prevKeys) => ({ ...prevKeys, [key]: true }));

      sendKeysToBackend();
    }
  };

  const sendKeysToBackend = () => {
    const jsonToSend = JSON.stringify(keysPressed);

    console.log(jsonToSend);
  };

  useEffect(() => {
    window.addEventListener("keydown", handleKeyDown);
    window.addEventListener("keyup", handleKeyUp);

    return () => {
      window.removeEventListener("keydown", handleKeyDown);
      window.removeEventListener("keyup", handleKeyUp);
    };
  }, [keysPressed]);
};

export default GameComponent;
