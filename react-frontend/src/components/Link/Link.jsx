import React from "react";
import { Link as ReactRouterLink } from "react-router-dom";
import { Link as ChakraLink } from "@chakra-ui/react";

const paths = {
  HOME: "/",
  ABOUT: "/about/",
};

const newInfraPaths = Object.values(paths);

const Link = (props) => {
  const isPathPresentInNewInfra = newInfraPaths.includes(props.to);
  if (
    !isPathPresentInNewInfra || // old infra paths can't be linked to with react-router
    props.to?.startsWith("http://") ||
    props.to?.startsWith("https://")
  ) {
    // if outer link, use normal ChakraLink without React Router
    const { to, ...rest } = props;
    return <ChakraLink href={to} {...rest} />;
  }

  // Link from React Router
  return <ChakraLink as={ReactRouterLink} {...props} />;
};

export default Link;
