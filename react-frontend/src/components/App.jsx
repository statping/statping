import React, { useEffect } from "react";
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import { ChakraProvider } from "@chakra-ui/react";
import theme from "../theme";
import ServicesPage from "./ServicesPage";
import Navigation from "./Navbar";
import { initLumberjack } from "../utils/trackers";
import { generateUUID, getUserId, setUerId } from "../utils/helper";

const App = () => {
  console.log(`Application running on ${process.env.NODE_ENV} mode.`);
  useEffect(() => {
    if(!getUserId()) {
      setUerId(generateUUID(14));
    }

    initLumberjack();
  }, [])

  return (
    <ChakraProvider theme={theme}>
      <BrowserRouter>
        <div className="app-layout">
          <Navigation />
          <Switch>
            <Route exact path="/" component={ServicesPage} />
            <Redirect from="*" to="/" />
          </Switch>
        </div>
      </BrowserRouter>
    </ChakraProvider>
  );
};

export default App;
