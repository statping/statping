const baseStyle = {
  fontWeight: "bold",
  boxShadow: "md",
  border: "1px",
  _hover: {
    textDecoration: "none",
  },
};

const variantSolid = (props) => {
  const { colorScheme } = props;

  const blueVariant = {
    bg: "blue.400",
    color: "white.100",
    borderColor: "blue.400",
    _hover: {
      bg: "blue.500",
      color: "white.100",
      borderColor: "blue.500",
    },
    _focus: {
      boxShadow: "none",
    },
  };

  const whiteVariant = {
    bg: "white.100",
    color: "blue.300",
    borderColor: "white.100",
    _hover: {
      bg: "white.500",
      color: "blue.500",
      borderColor: "white.100",
    },
    _focus: {
      boxShadow: "none",
    },
  };

  const darkBlueVariant = {
    bg: "blue.900",
    color: "white.100",
    borderColor: "blue.900",
    _hover: {
      bg: "blue.900",
      color: "white.100",
      borderColor: "blue.900",
    },
    _focus: {
      boxShadow: "none",
    },
  };

  const linkVariant = {
    bg: "transparent",
    color: "white.100",
    borderColor: "transparent",
    boxShadow: "none",
    outline: "none",
    _hover: {
      bg: "transparent",
      color: "white.100",
      borderColor: "transparent",
    },
  };

  if (colorScheme === "blue") {
    return blueVariant;
  }

  if (colorScheme === "white") {
    return whiteVariant;
  }

  if (colorScheme === "darkBlue") {
    return darkBlueVariant;
  }

  if (colorScheme === "link") {
    return linkVariant;
  }

  return blueVariant;
};

const variantOutline = (props) => {
  const { colorScheme } = props;

  const darkBlueVariant = {
    bg: "blue.900",
    color: "white.100",
    borderColor: "blue.400",
    _hover: {
      bg: "blue.900",
      color: "white.100",
      borderColor: "blue.400",
    },
    _focus: {
      boxShadow: "none",
    },
  };

  if (colorScheme === "darkBlue") {
    return darkBlueVariant;
  }

  return darkBlueVariant;
};

const variants = {
  solid: variantSolid,
  outline: variantOutline,
};

const sizes = {
  md: {
    fontSize: "md",
    padding: "14px 18px",
    height: "auto",
    borderRadius: "base",
  },
  sm: {
    fontSize: "sm",
    padding: "12px 20px",
    height: "auto",
    borderRadius: "sm",
  },
};

const defaultProps = {
  variant: "outline",
  size: "md",
  colorScheme: "blue",
};

const buttonStyles = {
  baseStyle,
  variants,
  sizes,
  defaultProps,
};

export default buttonStyles;
