import { extendTheme } from '@chakra-ui/react';
import colors from './colors';
import { borders, radii } from './border';
import breakpoints from './breakpoints';
import typography from './typography';
import shadows from './shadows';
import buttonStyles from './button';
import textStyles from './text';
import headingStyles from './heading';
import zIndices from './z-indices';
import linkStyles from './link';
import listStyles from './list';

export default extendTheme({
  borders,
  breakpoints,
  radii,
  colors,
  ...typography,
  shadows,
  zIndices,
  components: {
    Button: buttonStyles,
    Text: textStyles,
    Heading: headingStyles,
    Link: linkStyles,
    List: listStyles,
  },
  styles: {
    global: {
      body: {
        color: 'gray.300',
      },
      // style override for freshdesk chatbot.
      '.mobile-chat-container': {
        backgroundColor: `${colors.blue[400]} !important`,
      },
    },
  },
});
