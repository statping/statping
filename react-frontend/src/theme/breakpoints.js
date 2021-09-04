// eslint-disable-next-line import/no-extraneous-dependencies
import { createBreakpoints } from '@chakra-ui/theme-tools';

/**
 * Breakpoints for responsive design
 * min-widths
 */
const breakpoints = createBreakpoints({
  xxs: '20em', // 320px
  xs: '22.5em', // 360px
  sm: '30em', // 480px
  md: '48em', // 768px
  lg: '64em', // 1024px
  xl: '80em', // 1280px
  '2xl': '85.375em', // 1366px (ipad pro)
  '3xl': '96em', // 1536px
});

export default breakpoints;
