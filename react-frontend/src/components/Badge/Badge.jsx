import React from "react";
import { Box } from "@chakra-ui/react";

function Badge(props) {
  const { children, ...rest } = props;
  return (
    <Box
      as="span"
      backgroundColor="green.200"
      color="white.100"
      fontSize="xs"
      py="px"
      px="1"
      ml="px"
      borderRadius="sm"
      fontWeight="bold"
      {...rest}
    >
      {children}
    </Box>
  );
}

export default Badge;
