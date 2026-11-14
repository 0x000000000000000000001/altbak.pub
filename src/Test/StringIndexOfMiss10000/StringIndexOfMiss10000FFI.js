// FP-style native implementation; preparation is outside the timed action.
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "a").join("") + "needle";
  return count => {
    const text = count === size ? input : input.slice(0, count) + "needle";
    return text.indexOf("absent");
  };
};
