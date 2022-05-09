import analyticsService from '@razorpay/universe-utils/analytics';

export const initLumberjack = () => {
	analyticsService.init({
		lumberjackAppName: 'website',
		lumberjackApiKey: window.LUMBERJACK_API_KEY,
		lumberjackApiUrl: window.LUMBERJACK_API_URL,
	});
};
