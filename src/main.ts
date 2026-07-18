import "./styles.css";
import { route } from "./router";

function run() {
  const render = route();

  render();
  // setInterval(render, 1000);
}

run();
