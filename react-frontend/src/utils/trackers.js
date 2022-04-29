import analyticsService from  '@razorpay/universe-utils/analytics';

export const initLumberjack = () => {
  analyticsService.init({
    lumberjackAppName: 'test_source',
    lumberjackApiKey: 'test_key_1',
    lumberjackApiUrl: "https://lumberjack.stage.razorpay.in/v1/track",
  });
};