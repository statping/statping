import React from "react";
import { Box, Tooltip } from "@chakra-ui/react";
import Link from "../Link";
import Image from "../Image";
import Button from "../Button";

import RzpLogo from "../../static/razorpay-logo-white.svg";
import indiaFlagSvg from "./images/india-flag.svg";

const RazorpayLogoLink = () => (
  <Link
    to="https://razorpay.com/"
    isExternal
    py={{ xxs: "6", lg: "7" }}
    paddingRight={{ xxs: "0", lg: "22" }}
    paddingLeft={{ xxs: "2", lg: "0" }}
    display="inline-block"
  >
    <Image
      width="125px"
      htmlWidth="125px"
      height="auto"
      src={RzpLogo}
      alt="Razorpay Logo"
    />
  </Link>
);

const Navigation = () => {
  return (
    <Box
      // bg={props.backgroundColor}
      position="relative"
      as="nav"
      px={{ xxs: "2", md: "8", xl: "48" }}
    >
      <Box display="flex" maxWidth="1080px" margin="auto">
        <RazorpayLogoLink />
        <Box
          marginLeft="auto"
          display="flex"
          alignItems="center"
          justifyContent="center"
          my="auto"
          paddingRight={{ xxs: "16", lg: "0" }}
          py={{ xxs: "4", lg: "6" }}
        >
          <Tooltip
            width="56"
            textAlign="center"
            label="Razorpay is currently available only for Indian businesses"
            hasArrow
          >
            <Box
              role="group"
              as="span"
              marginRight="4"
              aria-label="Razorpay is currently available only for Indian businesses"
            >
              <Image
                display={{ base: "none", xxs: "none", lg: "inline-block" }}
                src={indiaFlagSvg}
                alt=""
                aria-hidden="true"
              />
            </Box>
          </Tooltip>
          <Button
            as={Link}
            size="sm"
            colorScheme="link"
            to="https://dashboard.razorpay.com/#/access/signin"
          >
            Log In
          </Button>
          <Button
            as={Link}
            to="https://dashboard.razorpay.com/signup"
            marginLeft={{ base: "1", xs: "4" }}
            size="sm"
            colorScheme="white"
            display={{ base: "none", xxs: "none", lg: "inline-block" }}
          >
            Sign Up
          </Button>
        </Box>
      </Box>
    </Box>
  );
};

export default Navigation;
