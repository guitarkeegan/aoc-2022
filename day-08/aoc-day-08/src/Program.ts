/* eslint-disable */
import { Effect } from "effect";
import { input } from "./input.js";

const log = (s: string) => console.log(s);

// Plain JavaScript for simple parsing
const parseMatrix = (input: string): number[][] =>
  input.split("\n").map((line) => line.split("").map((ch) => Number(ch)));

const matrix = parseMatrix(input);

function isVisible(i: number, j: number): number {
  const MAX_ROWS = matrix.length - 1;
  const MAX_COLS = matrix[0].length - 1;

  const cur = matrix[i][j];
  // log(`cur: ${cur}`);

  if (i === 0 || i === MAX_ROWS || j === 0 || j === MAX_COLS) {
    // log("outer");
    return 1;
  }

  let t = 0,
    r = 0,
    b = 0,
    l = 0;

  // look up
  let nxt = i - 1;
  while (nxt >= 0) {
    if (matrix[nxt][j] >= cur) {
      // log(`  up: ${matrix[nxt][j]}`);
      t = 1;
      break;
    }
    nxt--;
  }

  // look right
  nxt = j + 1;
  while (nxt <= MAX_COLS) {
    if (matrix[i][nxt] >= cur) {
      // log(`  right: ${matrix[i][nxt]}`);
      r = 1;
      break;
    }
    nxt++;
  }

  // look down
  nxt = i + 1;
  while (nxt <= MAX_ROWS) {
    if (matrix[nxt][j] >= cur) {
      // log(`  down: ${matrix[nxt][j]}`);
      b = 1;
      break;
    }
    nxt++;
  }

  // look left
  nxt = j - 1;
  while (nxt >= 0) {
    if (matrix[i][nxt] >= cur) {
      // log(`  left: ${matrix[i][nxt]}`);
      l = 1;
      break;
    }
    nxt--;
  }

  // log("res");
  // log(`${t + r + b + l === 4 ? 0 : 1}`);

  return t + r + b + l === 4 ? 0 : 1;
}

function scenicScore(i: number, j: number): number {
  const MAX_ROWS = matrix.length - 1;
  const MAX_COLS = matrix[0].length - 1;

  const cur = matrix[i][j];
  // log(`cur: ${cur}`);

  let t = 0,
    r = 0,
    b = 0,
    l = 0;

  // look up
  let nxt = i - 1;
  while (nxt >= 0) {
    if (matrix[nxt][j] >= cur) {
      // log(`  up: ${matrix[nxt][j]}`);
      t++;
      break;
    }
    t++;
    nxt--;
  }

  // look right
  nxt = j + 1;
  while (nxt <= MAX_COLS) {
    if (matrix[i][nxt] >= cur) {
      // log(`  right: ${matrix[i][nxt]}`);
      r++;
      break;
    }
    r++;
    nxt++;
  }

  // look down
  nxt = i + 1;
  while (nxt <= MAX_ROWS) {
    if (matrix[nxt][j] >= cur) {
      // log(`  down: ${matrix[nxt][j]}`);
      b++;
      break;
    }
    b++;
    nxt++;
  }

  // look left
  nxt = j - 1;
  while (nxt >= 0) {
    if (matrix[i][nxt] >= cur) {
      // log(`  left: ${matrix[i][nxt]}`);
      l++;
      break;
    }
    l++;
    nxt--;
  }

  // log("res");
  // log(`${t + r + b + l === 4 ? 0 : 1}`);

  return t * r * b * l;
}

// let visible = 0;
let topScenic = 0;
for (let i = 0; i < matrix.length; i++) {
  for (let j = 0; j < matrix[0].length; j++) {
    topScenic = Math.max(scenicScore(i, j), topScenic);
  }
}

log(`${topScenic}`);
