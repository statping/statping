/**
 * Defines typography
 * fonts, font size, font weight
 * line height,
 * letter spacing
 */

const CUSTOM_FONT = "Lato";
const FALLBACK_FONTS = `-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol"`;
const FONTS = `${CUSTOM_FONT},${FALLBACK_FONTS}`;

const fonts = {
  heading: FONTS,
  body: FONTS,
};

const fontSizes = {
  xxxs: ".5rem", // 8px
  xxs: ".625rem", // 10px
  xs: ".75rem", // 12px
  sm: ".875rem", //14px
  md: "1rem", // 16px
  lg: "1.125rem", // 18px
  xl: "1.25rem", // 20px
  "2xl": "1.375rem", //22px
  "3xl": "1.5rem", // 24px
  "4xl": "1.625rem", // 26px
  "5xl": "1.75rem", // 28px
  "6xl": "1.875rem", //30px
  "7xl": "2rem", // 32px
  "8xl": "2.125rem", // 34px
  "9xl": "2.25rem", // 36px
  "10xl": "2.375rem", // 38px
  "11xl": "2.5rem", // 40px
};

const fontWeights = {
  thin: 200,
  normal: 400,
  semibold: 600,
  bold: 700,
  extrabold: 800,
  black: 900,
};

const lineHeights = {
  normal: "normal",
  none: 1,
  shorter: 1.25,
  short: 1.375,
  base: 1.5,
  tall: 1.625,
  taller: 2,
  3: ".65rem", // 8px
  4: ".75rem",
  5: ".875rem",
  6: "1rem",
  7: "1.125rem",
  8: "1.25rem",
  9: "1.375rem",
  10: "1.5rem",
  11: "1.625rem",
  12: "1.75rem",
  13: "1.875rem",
  14: "2rem",
  15: "2.125rem",
  16: "2.25rem",
  17: "2.375rem",
  18: "2.5", // 40px,
  25: "3.375rem", // 54px
};

const letterSpacings = {
  tighter: "-0.008em",
  normal: "0",
  wider: "0.05em",
  heading: "0.002em",
};

const typography = {
  fonts,
  fontSizes,
  fontWeights,
  lineHeights,
  letterSpacings,
};

export default typography;
