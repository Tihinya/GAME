import Gachi, { useEffect, useState } from "../Gachi.js/src/core/framework";

const GameComponent = () => {
  const [keysPressed, setKeysPressed] = useState({
    KeyW: false,
    KeyA: false,
    KeyS: false,
    KeyD: false,
    Space: false,
  });

  const handleKeyDown = (event) => {
    const key = event.code;

    if (keysPressed.hasOwnProperty(key)) {
      if (event.code) {
        event.preventDefault();
      }

      setKeysPressed((prevKeys) => {
        const updatedKeys = { ...prevKeys, [key]: true };
        sendKeysToBackend(updatedKeys);
        return updatedKeys;
      });
    }
  };

  const handleKeyUp = (event) => {
    const key = event.code;
    if (keysPressed.hasOwnProperty(key)) {
      setKeysPressed((prevKeys) => {
        const updatedKeys = { ...prevKeys, [key]: false };
        sendKeysToBackend(updatedKeys);
        return updatedKeys;
      });
    }
  };

  const sendKeysToBackend = (updatedKeys) => {
    const jsonToSend = JSON.stringify(updatedKeys);
    console.log(jsonToSend);
    // Send jsonToSend to the backend
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
