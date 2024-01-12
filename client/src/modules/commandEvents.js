import Gachi, { useEffect, useState } from "../Gachi.js/src/core/framework.ts";

const GameComponent = () => {
  const [keysPressed, setKeysPressed] = useState({
    W: false,
    A: false,
    S: false,
    D: false,
    Space: false,
  });

  const handleKeyDown = (event) => {
    const key = event.key.toUpperCase(); // Convert to uppercase for consistency
    if (keysPressed.hasOwnProperty(key)) {
      setKeysPressed((prevKeys) => ({ ...prevKeys, [key]: false }));
      // Send JSON to backend
      sendKeysToBackend();
    }
  };

  const handleKeyUp = (event) => {
    const key = event.key.toUpperCase();
    if (keysPressed.hasOwnProperty(key)) {
      setKeysPressed((prevKeys) => ({ ...prevKeys, [key]: true }));
      // Send JSON to backend
      sendKeysToBackend();
    }
  };

  const sendKeysToBackend = () => {
    // Assuming you have some function to send JSON to the backend
    const jsonToSend = JSON.stringify(keysPressed);
    // Send jsonToSend to the backend
    console.log(jsonToSend);
  };

  useEffect(() => {
    // Add event listeners when the component mounts
    window.addEventListener("keydown", handleKeyDown);
    window.addEventListener("keyup", handleKeyUp);

    // Remove event listeners when the component unmounts
    return () => {
      window.removeEventListener("keydown", handleKeyDown);
      window.removeEventListener("keyup", handleKeyUp);
    };
  }, [keysPressed]);

  return <div>{/* Your game component goes here */}</div>;
};

export default GameComponent;
