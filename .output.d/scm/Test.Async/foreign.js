export const runAsyncTest = (n) => () => {
  return new Promise((resolve) => {
    let completed = 0;
    for (let i = 0; i < n; i++) {
      setTimeout(() => {
        completed++;
        if (completed === n) {
          resolve();
        }
      }, 0);
    }
  });
};
