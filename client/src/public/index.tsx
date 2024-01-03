import Gachi, { useState } from "../Gachi.js/src/core/framework";
import App from "../App";

const container = document.getElementById("root");

if (container) {
  Gachi.render(<App />, container);
} else {
  throw new Error("Problem with finding the root element");
}
