import React from "react";
import { Box, Grid, Tooltip } from "@chakra-ui/react";
// import { ArrowForwardIcon } from "@chakra-ui/icons";
import { motion } from "framer-motion";
import { SkipNavLink } from "@chakra-ui/skip-nav";
import Link from "../Link";
import Image from "../Image";
import Button from "../Button";

import RzpLogo from "../../static/razorpay-logo.svg";
import * as productIcons from "../../static/product-icons-blue";
import indiaFlagSvg from "./images/india-flag.svg";

import {
  NavMenu,
  NavTitle,
  NavContent,
  NavTab,
  NavColumnHeading,
  NavLink,
  NavLinkTitle,
  NavLinkDescription,
  NavMobileExploreButton,
  NavMobileBackButton,
  Underline,
  useNavBreakpointValue,
  NavTitleLink,
  NavMobileIconLink,
  NavLinkMobileDescription,
  NavLinkThinTitle,
  NavDivider,
  // NavHamMenuButton,
} from "./NavHelpers";

const MotionBox = motion(Box);

const RazorpayLogoLink = () => (
  <Link
    to="/"
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

const Navbar = (props) => {
  const [isMobileNavOpen, setIsMobileNavOpen] = React.useState(false);
  const [activeTab, setActiveTab] = React.useState(false);

  const paymentsMenuRef = React.createRef();
  const bankingMenuRef = React.createRef();
  const resourcesMenuRef = React.createRef();
  const supportMenuRef = React.createRef();

  const collapseAllNavSections = () => {
    const navMenuRefs = [
      paymentsMenuRef,
      bankingMenuRef,
      resourcesMenuRef,
      supportMenuRef,
    ];
    for (const navMenuRef of navMenuRefs) {
      navMenuRef.current?.collapseNavContent();
    }
  };

  const toggleMobileNavMenu = () => {
    setIsMobileNavOpen(!isMobileNavOpen);
    if (isMobileNavOpen) {
      // when closing navbar, collapse all tabs
      collapseAllNavSections();
    }
  };

  const navStatusVariants = useNavBreakpointValue({
    mobile: {
      open: {
        x: "-100vw",
        display: "block",
        transition: {
          type: "tween",
          duration: 0.3,
        },
      },
      closed: {
        x: 0,
        transition: {
          duration: 0.15,
        },
        transitionEnd: {
          display: "none",
        },
      },
    },
    desktop: {},
  });

  return (
    <Box
      // bg={props.backgroundColor}
      position="relative"
      as="nav"
      px={{ xxs: "2", md: "8", xl: "48" }}
    >
      <Box display="flex" maxWidth="1080px" margin="auto">
        <SkipNavLink>Skip to content</SkipNavLink>
        <RazorpayLogoLink />
        {/* Ham Menu Button */}
        {/* <NavHamMenuButton
          isMobileNavOpen={isMobileNavOpen}
          toggleMobileNavMenu={toggleMobileNavMenu}
        /> */}
        <MotionBox
          className="nav-container"
          display={{ xxs: "inline-block", lg: "flex" }}
          position={{ xxs: "fixed", lg: "revert" }}
          zIndex="navbar"
          right={{ xxs: "-100%", lg: "revert" }}
          width={{ xxs: "300px", lg: "revert" }}
          bg={{ xxs: "white.100", lg: "revert" }}
          height={{ xxs: "100%", lg: "revert" }}
          overflowY={{ xxs: activeTab ? "revert" : "auto", lg: "revert" }}
          variants={navStatusVariants}
          animate={isMobileNavOpen ? "open" : "closed"}
          aria-hidden={!isMobileNavOpen}
        >
          <NavMenu
            ref={paymentsMenuRef}
            setActiveTab={setActiveTab}
            collapseAllNavSections={collapseAllNavSections}
          >
            <NavTitle>Payments</NavTitle>
            <Box display={{ xxs: "revert", lg: "none" }}>
              <NavLink
                isExternal
                to="https://razorpay.com/payment-gateway"
                icon={productIcons.paymentGatewaySvg}
              >
                Payment Gateway
              </NavLink>
              <NavLink
                isExternal
                to="https://razorpay.com/payment-links"
                icon={productIcons.paymentLinksSvg}
              >
                Payment Links
              </NavLink>
              <NavLink
                isExternal
                to="https://razorpay.com/payment-buttons"
                icon={productIcons.paymentButtonsSvg}
                isNew
              >
                Payment Buttons
              </NavLink>
            </Box>
            <NavMobileExploreButton>
              Explore Payment Suite
            </NavMobileExploreButton>
            <NavDivider />
            <NavContent
              left={{ xxs: "0", md: "-40", xl: "-56" }}
              width={{ xxs: "100%", md: "1000px", xl: "1225px" }}
              maxWidth={{ xxs: "100%", md: "1000px", xl: "1225px" }}
            >
              <NavMobileBackButton>Razorpay Payment Suite</NavMobileBackButton>
              <NavTab>
                <Grid
                  templateColumns={{
                    xxs: "repeat(1, 1fr)",
                    lg: "repeat(3, 1fr)",
                  }}
                  gap="0"
                >
                  <Box>
                    <NavColumnHeading>ACCEPT PAYMENTS</NavColumnHeading>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/payment-gateway"
                      icon={productIcons.paymentGatewaySvg}
                    >
                      <NavLinkTitle>Payment Gateway</NavLinkTitle>
                      <NavLinkDescription>
                        Payments on your Website & App
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/payment-links"
                      icon={productIcons.paymentLinksSvg}
                    >
                      <NavLinkTitle>Payment Links</NavLinkTitle>
                      <NavLinkDescription>
                        Create & send links to collect money
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/payment-pages"
                      icon={productIcons.paymentPagesSvg}
                    >
                      <NavLinkTitle>Payment Pages</NavLinkTitle>
                      <NavLinkDescription>
                        Get paid with personalized pages
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/payment-buttons"
                      icon={productIcons.paymentButtonsSvg}
                    >
                      <NavLinkTitle isNew>Payment Buttons</NavLinkTitle>
                      <NavLinkDescription>
                        Create, Copy and Collect in 5 mins
                      </NavLinkDescription>
                    </NavLink>
                  </Box>
                  <Box>
                    {/* Gap in column 2 */}
                    <Box
                      display={{ xxs: "none", lg: "block" }}
                      height="8"
                      marginBottom="0.5"
                    />
                    <NavLink
                      isExternal
                      to="https://razorpay.com/capital/instant-settlements/"
                      icon={productIcons.instantSettlementSvg}
                    >
                      <NavLinkTitle>Instant Settlement</NavLinkTitle>
                      <NavLinkDescription>
                        Customer payments settled faster
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/invoices"
                      icon={productIcons.invoicesSvg}
                    >
                      <NavLinkTitle>Invoices</NavLinkTitle>
                      <NavLinkDescription>
                        Create & send GST compliant invoices
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/smart-collect"
                      icon={productIcons.smartCollectSvg}
                    >
                      <NavLinkTitle>Smart Collect</NavLinkTitle>
                      <NavLinkDescription>
                        Automate NEFT, RTGS, IMPS payments
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/invoices"
                      icon={productIcons.subscriptionsSvg}
                    >
                      <NavLinkTitle>Subscriptions</NavLinkTitle>
                      <NavLinkDescription>
                        Collect recurring subscription payments
                      </NavLinkDescription>
                    </NavLink>
                  </Box>
                  <Box>
                    <NavColumnHeading>DISBURSE PAYMENTS</NavColumnHeading>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/route"
                      icon={productIcons.routeSvg}
                    >
                      <NavLinkTitle>Route</NavLinkTitle>
                      <NavLinkDescription>
                        Split & manage market payments
                      </NavLinkDescription>
                    </NavLink>

                    <NavColumnHeading mt={{ xxs: "5", lg: "17px" }}>
                      RISK AND FRAUD
                    </NavColumnHeading>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/thirdwatch"
                      icon={productIcons.thirdwatchSvg}
                    >
                      <NavLinkTitle>Thirdwatch</NavLinkTitle>
                      <NavLinkDescription>
                        Fight fraud with Artificial Intelligence
                      </NavLinkDescription>
                    </NavLink>

                    <NavColumnHeading mt={{ xxs: "5", lg: "17px" }}>
                      PARTNER APPS
                    </NavColumnHeading>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/app-store"
                      icon={productIcons.appStoreSvg}
                    >
                      <NavLinkTitle isNew>App Store</NavLinkTitle>
                      <NavLinkDescription>
                        Find right app for your business
                      </NavLinkDescription>
                    </NavLink>
                  </Box>
                </Grid>
              </NavTab>
              <NavTab variant="payments-lower">
                <Grid
                  templateColumns={{
                    xxs: "repeat(1, 1fr)",
                    lg: "repeat(3, 1fr)",
                  }}
                  gap="0"
                >
                  <NavColumnHeading display={{ xxs: "revert", lg: "none" }}>
                    MORE
                  </NavColumnHeading>
                  <NavLink
                    isExternal
                    to="https://razorpay.com/payments-app"
                    icon={productIcons.paymentsMobileAppSvg}
                  >
                    <NavLinkTitle isNew>Payments Mobile App</NavLinkTitle>
                    <NavLinkDescription>
                      Track and Accept payments
                    </NavLinkDescription>
                  </NavLink>
                  <NavLink
                    isExternal
                    to="https://razorpay.com/cred-pay"
                    icon={productIcons.credPaySvg}
                  >
                    <NavLinkTitle isNew>CRED Pay</NavLinkTitle>
                    <NavLinkDescription>
                      Make payments using CRED coins
                    </NavLinkDescription>
                  </NavLink>
                  <NavLink
                    isExternal
                    to="https://razorpay.com/upi-autopay"
                    icon={productIcons.upiAutopaySvg}
                  >
                    <NavLinkTitle isNew>UPI AutoPay</NavLinkTitle>
                    <NavLinkDescription>
                      Recurring payments using UPI App
                    </NavLinkDescription>
                  </NavLink>
                </Grid>
              </NavTab>
              <MobileBottomLinks
                paymentsMenuRef={paymentsMenuRef}
                bankingMenuRef={bankingMenuRef}
                resourcesMenuRef={resourcesMenuRef}
                supportMenuRef={supportMenuRef}
                hideTab="payments"
              />
            </NavContent>
          </NavMenu>
          <NavMenu
            ref={bankingMenuRef}
            setActiveTab={setActiveTab}
            collapseAllNavSections={collapseAllNavSections}
          >
            <NavTitle display={{ xxs: "none", lg: "revert" }}>Banking</NavTitle>
            <NavTitle display={{ xxs: "revert", lg: "none" }}>
              Razorpay X - Banking Suite
            </NavTitle>
            <Box display={{ xxs: "revert", lg: "none" }}>
              <NavLink
                isExternal
                to="https://razorpay.com/x/current-accounts"
                icon={productIcons.currentAcountSvg}
              >
                Current Accounts
              </NavLink>
              <NavLink
                isExternal
                to="https://razorpay.com/x/vendor-payments"
                icon={productIcons.vendorPaymentsSvg}
              >
                Vendor Payments
              </NavLink>
            </Box>
            <NavMobileExploreButton>
              Explore Banking Suite
            </NavMobileExploreButton>
            <NavDivider />
            <NavContent
              left={{ xxs: "0", md: "-64", xl: "-20.7rem" }}
              width={{ xxs: "100%", md: "1000px", xl: "1225px" }}
              maxWidth={{ xxs: "100%", md: "1000px", xl: "1225px" }}
            >
              <NavMobileBackButton>Razorpay Banking Suite</NavMobileBackButton>
              <NavTab>
                <Grid
                  templateColumns={{
                    xxs: "repeat(1, 1fr)",
                    lg: "repeat(3, 1fr)",
                  }}
                  gap="0"
                >
                  <Box>
                    <NavColumnHeading>BUSINESS BANKING</NavColumnHeading>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/x/"
                      icon={productIcons.razorpayXSvg}
                      iconVariant="purple"
                    >
                      <NavLinkTitle>RazorpayX</NavLinkTitle>
                      <NavLinkDescription>
                        Business Banking built for disruptors
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/x/vendor-payments/"
                      icon={productIcons.vendorPaymentsSvg}
                    >
                      <NavLinkTitle>Vendor Payments</NavLinkTitle>
                      <NavLinkDescription>
                        Automate vendor invoice and TDS payments
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/x/payout-links/"
                      icon={productIcons.payoutLinkSvg}
                    >
                      <NavLinkTitle>Payout Links</NavLinkTitle>
                      <NavLinkDescription>
                        Send money without recipient account details
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/x/payouts/"
                      icon={productIcons.payoutsSvg}
                    >
                      <NavLinkTitle>Payouts</NavLinkTitle>
                      <NavLinkDescription>
                        24x7, Instant & Automated Payouts
                      </NavLinkDescription>
                    </NavLink>
                  </Box>
                  <Box>
                    <Box
                      display={{ xxs: "none", lg: "block" }}
                      height="8"
                      marginBottom="0.5"
                    />
                    <NavLink
                      isExternal
                      to="https://razorpay.com/x/current-accounts/"
                      icon={productIcons.currentAcountSvg}
                    >
                      <NavLinkTitle>Current Account</NavLinkTitle>
                      <NavLinkDescription>
                        Current Accounts for fast growing businesses
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/x/tax-payments/"
                      icon={productIcons.taxPaymentsSvg}
                    >
                      <NavLinkTitle>Tax Payments</NavLinkTitle>
                      <NavLinkDescription>
                        Pay your business taxes in under 30 seconds
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/payroll/"
                      icon={productIcons.payrollSvg}
                    >
                      <NavLinkTitle>Payroll</NavLinkTitle>
                      <NavLinkDescription>
                        Automate and execute payroll
                      </NavLinkDescription>
                    </NavLink>
                  </Box>
                  <Box>
                    <NavColumnHeading>CREDIT</NavColumnHeading>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/capital/"
                      icon={productIcons.capitalSvg}
                      iconVariant="purple"
                    >
                      <NavLinkTitle>Razorpay Capital</NavLinkTitle>
                      <NavLinkDescription>
                        Get money for your business needs
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/capital/cash-advance/"
                      icon={productIcons.cashAdvanceSvg}
                    >
                      <NavLinkTitle>Cash Advance</NavLinkTitle>
                      <NavLinkDescription>
                        Instant additional cash for business use
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/x/corporate-cards/"
                      icon={productIcons.corporateCardsSvg}
                    >
                      <NavLinkTitle>Corporate Cards</NavLinkTitle>
                      <NavLinkDescription>
                        Credit Card for growing businesses
                      </NavLinkDescription>
                    </NavLink>
                    <NavLink
                      isExternal
                      to="https://razorpay.com/capital/working-capital-loans/"
                      icon={productIcons.workingCapitalLoansSvg}
                    >
                      <NavLinkTitle>Working Capital Loans</NavLinkTitle>
                      <NavLinkDescription>
                        Avail collateral-free business loans
                      </NavLinkDescription>
                    </NavLink>
                  </Box>
                </Grid>
              </NavTab>
              <MobileBottomLinks
                paymentsMenuRef={paymentsMenuRef}
                bankingMenuRef={bankingMenuRef}
                resourcesMenuRef={resourcesMenuRef}
                supportMenuRef={supportMenuRef}
                hideTab="banking"
              />
            </NavContent>
          </NavMenu>
          <NavMenu
            ref={resourcesMenuRef}
            setActiveTab={setActiveTab}
            collapseAllNavSections={collapseAllNavSections}
          >
            <NavTitle display={{ xxs: "none", lg: "revert" }}>
              Resources
            </NavTitle>
            <NavContent
              width={{ xxs: "100%", lg: "840px" }}
              maxWidth={{ xxs: "100%", lg: "840px" }}
              left={{ xxs: "0", lg: "-56" }}
            >
              <NavMobileBackButton>Resources</NavMobileBackButton>
              <NavTab variant="without-link-description">
                <Grid
                  templateColumns={{
                    xxs: "repeat(1, 1fr)",
                    lg: "repeat(4, 1fr)",
                  }}
                  gap="0"
                >
                  <Box>
                    <NavColumnHeading>AWARENESS</NavColumnHeading>
                    <NavLink to="https://razorpay.com/blog" isExternal>
                      <NavLinkThinTitle>Blog</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Know about the nitty gritty of Payments, Banking & more!
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/learn" isExternal>
                      <NavLinkThinTitle>Learn</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Learn about Business Management, Freelance & more!
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/events" isExternal>
                      <NavLinkThinTitle>Events</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Learn more about Startups, Products, Sales and Funding
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/white-papers">
                      <NavLinkThinTitle>White papers</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        From data-driven fintech insights to best practices
                        around handling payments
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/customer-stories">
                      <NavLinkThinTitle>Customer Stories</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        50,00,000+ businesses powering payments with Razorpay
                      </NavLinkMobileDescription>
                    </NavLink>
                  </Box>
                  <Box>
                    <NavColumnHeading>DEVELOPERS</NavColumnHeading>
                    <NavLink to="https://razorpay.com/docs">
                      <NavLinkThinTitle>Developer Docs</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Get started with SDKs here
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/integrations">
                      <NavLinkThinTitle>Integrations</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        30+ platforms that Razorpay supports
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/docs/api" isExternal>
                      <NavLinkThinTitle>API Reference</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Official references for the Razorpay APIs
                      </NavLinkMobileDescription>
                    </NavLink>
                  </Box>
                  <Box>
                    <NavColumnHeading>SOLUTIONS</NavColumnHeading>
                    <NavLink to="https://razorpay.com/solutions/saas">
                      <NavLinkThinTitle>SaaS</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Multi-channel, Multi-mode Payments Experience
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/solutions/e-commerce">
                      <NavLinkThinTitle>E-commerce</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Seamlessly accept, manage and disburse money!
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/solutions/education">
                      <NavLinkThinTitle>Education</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Facilitate learning & growth for your students &
                        customers
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/solutions/bfsi">
                      <NavLinkThinTitle>BFSI</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Solve unique challenges across lending, wealth
                        management, and insurance sectors
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/freelancer-unregistered-business">
                      <NavLinkThinTitle>Freelance</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Accept both domestic & international payments from your
                        clients and cutomers!
                      </NavLinkMobileDescription>
                    </NavLink>
                  </Box>
                  <Box minWidth={{ base: "revert", lg: "250px" }}>
                    <NavColumnHeading>FREE TOOLS</NavColumnHeading>
                    <NavLink to="https://razorpay.com/gst-calculator">
                      <NavLinkThinTitle isNew>GST Calculator</NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        The easiest way for businesses to calculate their GST
                      </NavLinkMobileDescription>
                    </NavLink>
                    <NavLink to="https://razorpay.com/x/tds-online-payment/?ref=topnav">
                      <NavLinkThinTitle isNew>
                        Online TDS Payment
                      </NavLinkThinTitle>
                      <NavLinkMobileDescription>
                        Pay TDS for your business in 30 seconds
                      </NavLinkMobileDescription>
                    </NavLink>
                  </Box>
                </Grid>
              </NavTab>
              <MobileBottomLinks
                paymentsMenuRef={paymentsMenuRef}
                bankingMenuRef={bankingMenuRef}
                resourcesMenuRef={resourcesMenuRef}
                supportMenuRef={supportMenuRef}
                hideTab="resources"
              />
            </NavContent>
          </NavMenu>
          <NavMenu
            ref={supportMenuRef}
            setActiveTab={setActiveTab}
            collapseAllNavSections={collapseAllNavSections}
          >
            <NavTitle display={{ xxs: "none", lg: "revert" }}>Support</NavTitle>
            <NavContent
              width={{ xxs: "100%", lg: "272px" }}
              maxWidth={{ xxs: "100%", lg: "272px" }}
              left={{ xxs: "0", lg: "-20" }}
            >
              <NavMobileBackButton>Support</NavMobileBackButton>
              <NavTab variant="without-link-description">
                <Box py={{ xxs: "4", lg: "0" }}>
                  <NavColumnHeading>GET SUPPORT</NavColumnHeading>
                  <NavLink to="https://razorpay.com/support#request" isExternal>
                    <NavLinkThinTitle>Raise a request</NavLinkThinTitle>
                  </NavLink>
                  <NavLink to="https://razorpay.com/knowledgebase" isExternal>
                    <NavLinkThinTitle>Knowledgebase</NavLinkThinTitle>
                  </NavLink>
                  <NavLink to="https://razorpay.com/chargeback" isExternal>
                    <NavLinkThinTitle>Chargeback Guides</NavLinkThinTitle>
                  </NavLink>
                  <NavLink to="https://razorpay.com/settlement" isExternal>
                    <NavLinkThinTitle>Settlement Guides</NavLinkThinTitle>
                  </NavLink>
                </Box>
              </NavTab>
              <MobileBottomLinks
                paymentsMenuRef={paymentsMenuRef}
                bankingMenuRef={bankingMenuRef}
                resourcesMenuRef={resourcesMenuRef}
                supportMenuRef={supportMenuRef}
                hideTab="support"
              />
            </NavContent>
          </NavMenu>

          {/* Desktop Links */}
          <NavTitleLink
            display={{ xxs: "none", lg: "inline-block" }}
            to="https://razorpay.com/partners"
          >
            Partners
          </NavTitleLink>
          <NavTitleLink
            display={{ xxs: "none", lg: "inline-block" }}
            to="https://razorpay.com/pricing"
          >
            Pricing
          </NavTitleLink>
          {/* Mobile Links */}
          <MobileBottomLinks
            paymentsMenuRef={paymentsMenuRef}
            bankingMenuRef={bankingMenuRef}
            resourcesMenuRef={resourcesMenuRef}
            supportMenuRef={supportMenuRef}
            backgroundColor="white.100"
            hideTab={["payments", "banking"]}
            mt="0"
            pt="2"
          />
          <Underline activeTab={activeTab} />
        </MotionBox>

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
            colorScheme="white"
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

const MobileBottomLinks = (props) => {
  const {
    paymentsMenuRef,
    bankingMenuRef,
    resourcesMenuRef,
    supportMenuRef,
    hideTab,
    ...rest
  } = props;

  return (
    <>
      <Box
        display={{ xxs: "revert", lg: "none" }}
        backgroundColor="gray.100"
        py="6"
        mt="8"
        {...rest}
      >
        {!hideTab?.includes("payments") && (
          <NavMobileExploreButton
            openMenuRef={paymentsMenuRef}
            icon={productIcons.explorePaymentsSvg}
          >
            Explore Payments Suite
          </NavMobileExploreButton>
        )}
        {!hideTab?.includes("banking") && (
          <NavMobileExploreButton
            openMenuRef={bankingMenuRef}
            icon={productIcons.exploreBankingSvg}
          >
            Explore Banking Suite
          </NavMobileExploreButton>
        )}
        <NavMobileIconLink
          to="https://razorpay.com/pricing"
          isExternal
          icon={productIcons.pricingSvg}
          color="yellow.200"
        >
          Pricing
        </NavMobileIconLink>
        {!hideTab?.includes("resources") && (
          <NavMobileExploreButton
            openMenuRef={resourcesMenuRef}
            icon={productIcons.resourcesSvg}
          >
            Resources
          </NavMobileExploreButton>
        )}
        <NavMobileIconLink
          to="https://razorpay.com/partners"
          isExternal
          icon={productIcons.partnersSvg}
        >
          Partners (Refer & Earn)
        </NavMobileIconLink>
        {!hideTab?.includes("support") && (
          <NavMobileExploreButton
            openMenuRef={supportMenuRef}
            icon={productIcons.supportSvg}
          >
            Support
          </NavMobileExploreButton>
        )}
      </Box>
      <Box
        display={{ xxs: "revert", lg: "none" }}
        px="4"
        py="2"
        position="sticky"
        bottom="0px"
        width="100%"
        backgroundColor="white.100"
      >
        <Button
          as={Link}
          display="block"
          textAlign="center"
          size="sm"
          to="https://dashboard.razorpay.com/#/access/signin"
        >
          Log In
        </Button>
      </Box>
    </>
  );
};

export default Navbar;
