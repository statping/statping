import React from "react";
import { Box, Divider, Flex, useBreakpointValue } from "@chakra-ui/react";
import {
  ChevronLeftIcon,
  ChevronRightIcon,
  CloseIcon,
  HamburgerIcon,
} from "@chakra-ui/icons";
import { AnimatePresence, motion } from "framer-motion";
import { isDescendant } from "../../utils/general";

import Link from "../Link";
import Text from "../Text";
import Image from "../Image";
import Badge from "../Badge";

export const NavContext = React.createContext({
  showContent: false,
  expandNavContent: () => {},
  collapseNavContent: () => {},
});

export function useNavBreakpointValue({ mobile, desktop }) {
  return (
    useBreakpointValue({
      xxs: mobile,
      sm: mobile,
      md: mobile,
      lg: desktop,
      xl: desktop,
      "2xl": desktop,
      "3xl": desktop,
      base: desktop, // apply to rest of the breakpoints
    }) ?? desktop
  );
}

const NewBadge = () => <Badge ml="1.5">NEW</Badge>;

export const NavMenu = React.forwardRef((props, ref) => {
  const { children } = props;

  const parentNavRef = React.useRef(null);

  const [showContent, setShowContent] = React.useState(false);
  const expandNavContent = () => {
    props.collapseAllNavSections();
    setShowContent(true);
    props.setActiveTab(parentNavRef);
  };
  const collapseNavContent = () => {
    setShowContent(false);
    props.setActiveTab(false);
  };

  React.useImperativeHandle(ref, () => ({
    // add properties that you may want to call from outside of Provider.
    collapseNavContent,
    expandNavContent,
    showContent,
  }));

  const navContextValue = {
    showContent,
    expandNavContent,
    collapseNavContent,
  };

  const blurHandler = (event) => {
    const child = event.relatedTarget;
    if (parentNavRef.current) {
      if (!isDescendant(parentNavRef.current, child)) {
        navContextValue.collapseNavContent();
      }
    }
  };

  const hoverListeners = useNavBreakpointValue({
    mobile: {},
    desktop: {
      onMouseEnter: navContextValue.expandNavContent,
      onMouseLeave: navContextValue.collapseNavContent,
      onFocus: navContextValue.expandNavContent,
      onBlur: blurHandler,
    },
  });

  return (
    <NavContext.Provider value={navContextValue}>
      <Box {...hoverListeners} ref={parentNavRef} position="relative">
        {children}
      </Box>
    </NavContext.Provider>
  );
});

export const NavTitle = (props) => {
  const navContextValue = React.useContext(NavContext);
  const ariaHasPopup = useNavBreakpointValue({ mobile: false, desktop: true });
  const renderAs = useNavBreakpointValue({
    mobile: "div",
    desktop: "button",
  });

  return (
    <Box
      as={renderAs}
      color={{
        xxs: "gray.200",
        lg: navContextValue.showContent ? "blue.400" : "gray.200",
      }}
      fontWeight={{ xxs: "bold", lg: "initial" }}
      fontSize={{ xxs: "sm", lg: "md" }}
      paddingTop={{ xxs: "5", lg: "8" }}
      paddingBottom="4"
      px={{ xxs: "5", lg: "3", xl: "4" }}
      textTransform={{ xxs: "uppercase", lg: "none" }}
      aria-haspopup={ariaHasPopup}
      aria-expanded={navContextValue.showContent}
      {...props}
    />
  );
};

export const NavTitleLink = (props) => {
  return (
    <Link
      paddingTop={{ xxs: "5", lg: "8" }}
      paddingBottom="4"
      px={{ xxs: "5", lg: "4" }}
      color="gray.200"
      _hover={{
        color: "blue.400",
        textDecoration: "none",
      }}
      {...props}
    />
  );
};

const MotionBox = motion(Box);

export const NavContent = (props) => {
  const navContextValue = React.useContext(NavContext);
  const { children, ...rest } = props;

  const contentAnimationVariants = useNavBreakpointValue({
    mobile: {
      open: {
        x: 0,
        display: "block",
        transition: {
          type: "tween",
          duration: 0.3,
        },
      },
      closed: {
        x: "100vw",
        transition: {
          type: "tween",
          duration: 0.15,
        },
        transitionEnd: {
          display: "none",
        },
      },
    },
    desktop: {
      open: {
        opacity: 1,
        display: "block",
        transition: {
          duration: 0.25,
        },
      },
      closed: {
        opacity: 0,
        transition: {
          duration: 0.2,
        },
        transitionEnd: {
          display: "none",
        },
      },
    },
  });

  return (
    <MotionBox
      position={{ xxs: "fixed", lg: "absolute" }}
      top={{ xxs: "0", lg: "revert" }}
      height={{ xxs: "100%", lg: "revert" }}
      backgroundColor={{ xxs: "white.100", lg: "revert" }}
      overflowY="auto"
      display="none"
      variants={contentAnimationVariants}
      animate={navContextValue.showContent ? "open" : "closed"}
      zIndex="1"
      aria-hidden={!navContextValue.showContent}
      {...rest}
    >
      {children}
    </MotionBox>
  );
};

export const NavLinkTitle = (props) => {
  const { children, isNew, ...rest } = props;
  return (
    <Text
      className="nav-link-title"
      fontSize={{ xxs: "sm", lg: "md" }}
      color="blue.800"
      lineHeight="10"
      display="inline-block"
      fontWeight="bold"
      {...rest}
    >
      {children}
      {isNew && <NewBadge />}
    </Text>
  );
};

export const NavLinkThinTitle = (props) => (
  <NavLinkTitle color="revert" fontWeight="semibold" fontSize="sm" {...props} />
);

export const NavLinkDescription = (props) => (
  <Text
    className="nav-link-description"
    fontSize={{ xxs: "xs", lg: "sm" }}
    lineHeight={{ xxs: "6", lg: "9" }}
    color="blue.600"
    {...props}
  />
);

export const NavLinkMobileDescription = (props) => (
  <NavLinkDescription display={{ xxs: "revert", lg: "none" }} {...props} />
);

export const NavLink = (props) => {
  const { children, icon, iconVariant, isNew, ...rest } = props;

  const commonNavLinkStyles = {
    display: "block",
    px: { xxs: "5", lg: "6" },
    py: { xxs: "2.5", lg: "4" },
    borderRadius: "md",
    ".chakra-image": {
      background:
        iconVariant === "purple"
          ? "linear-gradient(188.57deg, #2095FF -16.05%, #6416FF 116.32%);"
          : "blue.400",
    },
    ".chakra-icon": {
      opacity: 0,
      position: "relative",
      left: "0",
      transition: "left .1s",
    },
    "&:hover, &:focus": {
      boxShadow: "sm",
      ".nav-tab-payments-lower &": {
        boxShadow: "none",
      },
      ".nav-link-title": {
        color: "blue.400",
        transition: "color .1s",
      },
      ".chakra-image": {
        background: "linear-gradient(30.55deg, #6AC7FF 11.05%, #148AFF 90.09%)",
        transition: "all .1s",
      },
      ".chakra-icon": {
        color: "blue.400",
        opacity: 1,
        left: "2",
        transition: "left .1s",
      },
    },
  };

  // if NavLink has icon property, render the nav link with icon and description
  if (Array.isArray(children) && icon) {
    // Render nav link with icon and description
    const [heading, paragraph] = children;

    return (
      <Link sx={commonNavLinkStyles} {...rest}>
        <Flex alignItems="center">
          <Box>
            <Image
              height="8"
              width="8"
              htmlHeight="32px"
              htmlWidth="32px"
              rounded="full"
              src={icon}
              alt=""
            />
          </Box>
          <Box flex="1" paddingLeft="3">
            <Box>
              {heading}
              {isNew && <NewBadge />}
              <ChevronRightIcon w="5" h="5" />
            </Box>
            {paragraph}
          </Box>
        </Flex>
      </Link>
    );
  } else if (icon) {
    return (
      <Link sx={{ ...commonNavLinkStyles, px: "5", py: "2.5" }} {...rest}>
        <Flex alignItems="center">
          <Box>
            <Image
              htmlHeight="24px"
              htmlWidth="24px"
              height="6"
              width="6"
              rounded="full"
              src={icon}
              alt=""
            />
          </Box>
          <Box fontSize="sm" color="gray.300" paddingLeft="2">
            {children}
            {isNew && <NewBadge />}
            <ChevronRightIcon w="5" h="5" />
          </Box>
        </Flex>
      </Link>
    );
  } else {
    // Render plain Nav Link without image
    const plainLinkStyles = {
      fontWeight: { xxs: "revert", lg: "normal" },
      letterSpacing: "tighter",
      position: "relative",
      fontSize: { xxs: "sm", lg: "md" },
      py: "1",
      my: "1.5",
      px: { xxs: "6", lg: "4" },
      lineHeight: "0",
      color: "blue.700",
      display: "flex",
      alignItems: "center",
      borderLeftWidth: "4px",
      borderLeftColor: "transparent",
      "&:before": {
        content: '""',
        height: "74px",
        width: "px",
        background:
          "linear-gradient(to bottom,rgba(0,0,0,0),rgba(0,0,0,0.267),rgba(0,0,0,0))",
        position: "absolute",
        left: "-2.1px",
        top: "-20px",
        transition: "opacity .2s ease",
        opacity: 0,
      },
      ".chakra-icon": {
        position: "relative",
        top: "px",
        opacity: 0,
        left: "0",
        transition: "left .1s",
      },
      "&:hover, &:focus": {
        color: "blue.400",
        ".chakra-icon": {
          opacity: 1,
          left: "2",
          transition: "left .1s",
        },
        borderLeftColor: "blue.400",
        "&:before": {
          opacity: 1,
        },
      },
    };

    return (
      <Link sx={plainLinkStyles} {...rest}>
        <Box display="inline-block" as="span">
          {children}
          {isNew && <NewBadge />}
        </Box>{" "}
        <ChevronRightIcon w="5" h="5" />
      </Link>
    );
  }
};

export const NavColumnHeading = (props) => (
  <Text
    px="6"
    mt={{ xxs: "5", lg: "revert" }}
    fontSize="xs"
    fontWeight="normal"
    letterSpacing="wider"
    color="gray.200"
    paddingBottom="1"
    {...props}
  />
);

export const NavTab = (props) => {
  const { variant, ...rest } = props;
  if (variant === "payments-lower") {
    return (
      <Box
        className="nav-tab-payments-lower"
        mx={{ xxs: "0", lg: "6" }}
        py={{ xxs: "0", lg: "4" }}
        px={{ xxs: "0", lg: "3" }}
        bgColor={{ xxs: "white.100", lg: "white.200" }}
        borderRadius="0 0 0.375rem 0.375rem"
        {...rest}
      />
    );
  }

  if (variant === "without-link-description") {
    return (
      <Box
        className="nav-tab-without-link-description"
        px={{ xxs: "0", lg: "7" }}
        py={{ xxs: "0", lg: "10" }}
        borderRadius="md"
        bgColor="white.100"
        {...rest}
      />
    );
  }

  return (
    <Box
      className="nav-tab-default"
      px={{ xxs: "0", lg: "8" }}
      py={{ xxs: "0", lg: "10" }}
      borderRadius="lg"
      bgColor="white.100"
      {...rest}
    />
  );
};

export const Underline = ({ activeTab }) => {
  if (!activeTab) return null;
  const bounds = activeTab.current?.getBoundingClientRect();
  if (!bounds) return null;
  const left = bounds.left + 15; // 15px padding
  const width = bounds.width - 30;
  return (
    <MotionBox
      bgColor="blue.400"
      display={{ xxs: "none", lg: "revert" }}
      height="1"
      top="68px"
      position="absolute"
      animate={{ left: `${left}px` }}
      pointerEvents="none"
      width={width}
      initial={false}
    />
  );
};

export const NavDivider = () => (
  <Box display={{ xxs: "revert", lg: "none" }} px="5" py="2">
    <Divider borderColor="white.300" borderWidth="1px" />
  </Box>
);

// Mobile Components

export const NavMobileBackButton = (props) => {
  const navContextValue = React.useContext(NavContext);

  return (
    <Box
      as="button"
      display={{ xxs: "block", lg: "none" }}
      py="4"
      px="3"
      borderBottom="1px"
      borderBottomColor="white.300"
      bgColor="white.100"
      width="100%"
      textAlign="left"
      fontSize="md"
      color="blue.900"
      fontWeight="bold"
      onClick={navContextValue.collapseNavContent}
    >
      <ChevronLeftIcon h="5" w="5" marginRight="3" aria-label="Back Icon" />
      {props.children}
    </Box>
  );
};

/**
 * Use \<NavMobileExploreButton\> inside \<NavMenu\> or pass
 * `openMenuRef` with reference to NavMenu to open
 */
export const NavMobileExploreButton = (props) => {
  const navContextValue = React.useContext(NavContext);
  let { showContent, collapseNavContent, expandNavContent } = navContextValue;

  const toggleNavMenu = () => {
    if (props.openMenuRef?.current) {
      showContent = props.openMenuRef.current.showContent;
      collapseNavContent = props.openMenuRef.current.collapseNavContent;
      expandNavContent = props.openMenuRef.current.expandNavContent;
    }
    if (showContent) {
      collapseNavContent();
    } else {
      expandNavContent();
    }
  };
  return (
    <Box
      as="button"
      color={props.icon ? "gray.300" : "blue.400"}
      px="5"
      py={props.icon ? "2.5" : "4"}
      display={{ xxs: "flex", lg: "none" }}
      position="relative"
      alignItems="center"
      onClick={toggleNavMenu}
      width="100%"
    >
      {props.icon ? (
        <Image
          htmlHeight="32px"
          htmlWidth="32px"
          height="8"
          width="8"
          src={props.icon}
          alt=""
        />
      ) : null}
      <Text
        paddingLeft={props.icon ? "2" : "0"}
        as="span"
        fontWeight={props.icon ? "normal" : "bold"}
        fontSize="sm"
      >
        {props.children}
      </Text>
      <ChevronRightIcon
        marginLeft="auto"
        w="5"
        h="5"
        opacity={props.icon ? "0.4" : "1"}
      />
    </Box>
  );
};

export const NavMobileIconLink = (props) => {
  return (
    <Link
      display={{ xxs: "block", lg: "none" }}
      px="5"
      py="2.5"
      fontSize="sm"
      color="gray.300"
      {...props}
    >
      <Flex alignItems="center">
        <Box>
          <Image
            htmlHeight="32px"
            htmlWidth="32px"
            height="8"
            width="8"
            rounded="full"
            src={props.icon}
            alt=""
          />
        </Box>
        <Box paddingLeft="2">{props.children}</Box>
      </Flex>
    </Link>
  );
};

const MotionCloseIcon = motion(CloseIcon);
const MotionHamburgerIcon = motion(HamburgerIcon);

export const NavHamMenuButton = ({ isMobileNavOpen, toggleMobileNavMenu }) => {
  const navHamIconVariants = {
    initial: {
      rotate: 0,
    },
    rotated: {
      rotate: -180,
      transition: {
        duration: 0.3,
      },
    },
  };

  return (
    <Box
      as="button"
      position={isMobileNavOpen ? "fixed" : "absolute"}
      top="0"
      right="0"
      zIndex="navbutton"
      padding={isMobileNavOpen ? "4" : "6"}
      onClick={toggleMobileNavMenu}
      display={{ xxs: "inline-block", lg: "none" }}
      aria-expanded={isMobileNavOpen}
      aria-label={isMobileNavOpen ? "Close Nav Menu" : "Open Nav Menu"}
    >
      <AnimatePresence>
        {isMobileNavOpen ? (
          <MotionCloseIcon
            color="gray.900"
            fontSize="sm"
            zIndex="1"
            initial="initial"
            top="6"
            right="6"
            animate="rotated"
            exit="initial"
            variants={navHamIconVariants}
          />
        ) : (
          <MotionHamburgerIcon
            color="gray.900"
            fontSize="xl"
            initial="initial"
            animate="rotated"
            exit="initial"
            variants={navHamIconVariants}
          />
        )}
      </AnimatePresence>
    </Box>
  );
};
