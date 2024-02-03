import Gachi, { useState } from "./Gachi.js/src/core/framework";
import MenuPage from "./pages/mainpage/mainpage";
import Lobby from "./pages/lobby/lobby";
import GamePage from "./pages/gamepage/gamepage";
import "./index.css";

export default function App() {
  const [currentPage, setCurrentPage] = useState("main");
  Gachi.createContext("switchPage", setCurrentPage);

  return (
    <>
      {currentPage === "main" ? <MenuPage /> : null}
      {currentPage === "lobby" ? <Lobby /> : null}
      {currentPage === "game" ? <GamePage /> : null}
    </>
  );
}
