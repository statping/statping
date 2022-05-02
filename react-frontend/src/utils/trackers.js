import analyticsService from  '@razorpay/universe-utils/analytics';

export const initLumberjack = () => {
  analyticsService.init({
    lumberjackAppName: 'website',
    lumberjackApiKey: '10pYUm55sa39zgTN1gzNwQzNyQjM54Cg',
    lumberjackApiUrl: "https://lumberjack.stage.razorpay.in/v1/track",
  });
};