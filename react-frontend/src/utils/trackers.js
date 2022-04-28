import analyticsService from  '@razorpay/universe-utils/analytics';

export const initLumberjack = () => {
  analyticsService.init({
    lumberjackAppName: 'test_source',
    lumberjackApiKey: window.LUMBERJACK_API_KEY,
    lumberjackApiUrl: "https://lumberjack.stage.razorpay.in/v1/track",
  });
};