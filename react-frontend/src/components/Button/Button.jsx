import React from "react";
import { useStyleConfig, chakra, Box, forwardRef } from "@chakra-ui/react";

const Button = forwardRef((props, ref) => {
  const {
    variant = "solid",
    colorScheme = "blue",
    size = "md",
    leftIcon,
    rightIcon,
    children,
    onClick,
    ...rest
  } = props;

  const styles = useStyleConfig("Button", {
    variant,
    colorScheme,
    size,
  });

  return (
    /* using sx instead of __css so the styles cannot be overriden
    directly. To pass new styles, check the theme
     */
    <chakra.button
      display="inline-block"
      sx={styles}
      ref={ref}
      {...rest}
      onClick={onClick}
    >
      {leftIcon && (
        <Box as="span" marginRight={2} position="relative" top="-px">
          {leftIcon}
        </Box>
      )}
      {children}
      {rightIcon && (
        <Box as="span" marginLeft={2} position="relative" top="-px">
          {rightIcon}
        </Box>
      )}
    </chakra.button>
  );
});

export default Button;
