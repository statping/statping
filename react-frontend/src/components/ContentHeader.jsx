import React from "react";

const ContentHeader = () => {
  return (
    <div className="header">
      <h1 className="header-title mt-4 mb-3 font-24 fw-700">
        Razorpay Status Page
      </h1>
      <h5 className="header-description font-12">
        Razorpay status page publishes the most up-to-the-minute information on
        product availability. Check back here any time to get current
        status/information on individual products. If you are experiencing a
        real-time, operational issue with one of our products that is not
        described below, please reach out to{" "}
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://razorpay.com/support/"
        >
          our support team
        </a>{" "}
        and we will help you out.
      </h5>
    </div>
  );
};

export default ContentHeader;
