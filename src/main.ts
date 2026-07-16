import "./styles.css";
import { render } from "./board/render";
import { getMockData } from "./board/mockData";
import { newRenderContext } from "./board/context";

const body = document.body;

function run() {
  const mockData = getMockData();

  function update() {
    const ctx = newRenderContext(mockData);
    render(body, ctx);
  }
  update();

  setInterval(update, 1000);
}

run();
