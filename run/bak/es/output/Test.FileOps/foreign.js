import fs from "fs";

export const writeFileSync = (path) => (content) => () => {
  fs.writeFileSync(path, content, "utf8");
};

export const readFileSync = (path) => () => {
  return fs.readFileSync(path, "utf8");
};

export const loopE = (n) => (action) => () => {
  for (let i = 0; i < n; i++) {
    action();
  }
};
