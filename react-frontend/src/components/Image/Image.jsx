import React from "react";
import { Image as ChakraImage } from "@chakra-ui/react";

const Image = (props) => (
  <ChakraImage loading="lazy" ignoreFallback={true} {...props} />
);

export default Image;
